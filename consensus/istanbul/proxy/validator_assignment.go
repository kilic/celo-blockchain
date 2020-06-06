// Copyright 2017 The celo Authors
// This file is part of the celo library.
//
// The celo library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The celo library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the celo library. If not, see <http://www.gnu.org/licenses/>.

package proxy

import (
	"github.com/buraksezer/consistent"
	"github.com/cespare/xxhash"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

// This type stores the assignment of remote validators to proxies, as well as the
// reverse assignment.
// If a validator is assigned to a nil proxy, then that means that it's
// not assigned yet.
// WARNING:  None of this object's functions are threadsafe, so it's
//           the user's responsibility to ensure that.
type valAssignments struct {
	valToProxy  map[common.Address]*enode.ID             // map of validator address -> proxy assignment ID
	proxyToVals map[enode.ID]map[common.Address]struct{} // map of proxy ID to set of validator addresses
}

func newValAssignments() *valAssignments {
	return &valAssignments{
		valToProxy:  make(map[common.Address]*enode.ID),
		proxyToVals: make(map[enode.ID]map[common.Address]struct{}),
	}
}

// addValidators adds validators to valToProxy without an assigned proxy
func (va *valAssignments) addValidators(vals []common.Address) {
	for _, val := range vals {
		va.valToProxy[val] = nil
	}
}

// removeValidators removes validators from any proxy assignments and deletes
// them from valToProxy
func (va *valAssignments) removeValidators(vals []common.Address) {
	for _, val := range vals {
		va.unassignValidator(val)
		delete(va.valToProxy, val)
	}
}

// assignValidator assigns a validator with address valAddress to the proxy
// with ID proxyID
func (va *valAssignments) assignValidator(valAddress common.Address, proxyID enode.ID) {
	va.valToProxy[valAddress] = &proxyID

	if _, ok := va.proxyToVals[proxyID]; !ok {
		va.proxyToVals[proxyID] = make(map[common.Address]struct{})
	}

	va.proxyToVals[proxyID][valAddress] = struct{}{}
}

// unassignValidator unassigns a validator with address valAddress from
// its proxy. If it was never assigned, this does nothing
func (va *valAssignments) unassignValidator(valAddress common.Address) {
	proxyID := va.valToProxy[valAddress]

	if proxyID != nil {
		va.valToProxy[valAddress] = nil
		delete(va.proxyToVals[*proxyID], valAddress)

		if len(va.proxyToVals[*proxyID]) == 0 {
			delete(va.proxyToVals, *proxyID)
		}
	}
}

// getValidators returns all validator addresses that are found in valToProxy.
// Note that it will also return both assigned and unassigne validators.
func (va *valAssignments) getValidators() []common.Address {
	vals := make([]common.Address, len(va.valToProxy), len(va.valToProxy))

	for val := range va.valToProxy {
		vals = append(vals, val)
	}
	return vals
}

// assignmentPolicy is intended to allow multiple implementations of validator assignment
// policies
type assignmentPolicy interface {
	assignProxy(proxy *proxy, valAssignments *valAssignments) bool
	removeProxy(proxy *proxy, valAssignments *valAssignments) bool
	assignRemoteValidators(validators []common.Address, valAssignments *valAssignments) bool
	removeRemoteValidators(validators []common.Address, valAssignments *valAssignments) bool
}

// ==============================================
//
// define the consistent hashing assignment policy implementation

type hasher struct{}

func (h hasher) Sum64(data []byte) uint64 {
	return xxhash.Sum64(data)
}

// consistentHashingPolicy uses consistent hashing to assign validators to proxies.
// Validator <-> proxy pairings are recalculated every time a proxy or validator
// is added/removed
// WARNING:  None of this object's functions are threadsafe, so it's
//           the user's responsibility to ensure that.
type consistentHashingPolicy struct {
	c *consistent.Consistent // used for consistent hashing
}

func newConsistentHashingPolicy() *consistentHashingPolicy {
	// This sets up a consistent hasher with bounded loads:
	// https://ai.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html
	// Partitions are assigned to members (proxies in this case)
	// using a hash ring.
	// When locating a key's member using `LocateKey`, the key is assigned
	// to a partition using hash(key) % PartitionCount in constant time.
	cfg := consistent.Config{
		// Prime to distribute keys more uniformly.
		// Higher partition count generally gives a more even distribution
		PartitionCount: 271,
		// The number of replications of a member (proxy) on the hash ring
		ReplicationFactor: 40,
		// Used to enforce a max # of partitions assigned per member, which is
		// (PartitionCount / len(members)) * Load. A load closer to 1 gives
		// more uniformity in the # of partitions assigned to specific members,
		// but a higher load results in less relocations when members are added/removed
		Load:   1.2,
		Hasher: hasher{},
	}

	return &consistentHashingPolicy{
		c: consistent.New(nil, cfg),
	}
}

// assignProxy adds a proxy to the consistent hasher and recalculates all validator assignments
func (ch *consistentHashingPolicy) assignProxy(proxy *proxy, valAssignments *valAssignments) bool {
	ch.c.Add(proxy.ID())
	return ch.reassignValidators(valAssignments)
}

// removeProxy removes a proxy from the consistent hasher and recalculates all validator assignments
func (ch *consistentHashingPolicy) removeProxy(proxy *proxy, valAssignments *valAssignments) bool {
	ch.c.Remove(proxy.ID().String())
	return ch.reassignValidators(valAssignments)
}

// assignRemoteValidators adds remote validators to the valAssignments struct and recalculates
// all validator assignments
func (ch *consistentHashingPolicy) assignRemoteValidators(vals []common.Address, valAssignments *valAssignments) bool {
	valAssignments.addValidators(vals)
	return ch.reassignValidators(valAssignments)
}

// removeRemoteValidators removes remote validators from the valAssignments struct and recalculates
// all validator assignments
func (ch *consistentHashingPolicy) removeRemoteValidators(vals []common.Address, valAssignments *valAssignments) bool {
	valAssignments.removeValidators(vals)
	return ch.reassignValidators(valAssignments)
}

// reassignValidators recalculates all validator <-> proxy pairings
func (ch *consistentHashingPolicy) reassignValidators(valAssignments *valAssignments) bool {
        anyAssignmentsChanged := false
	for val, proxyID := range valAssignments.valToProxy {
		newProxyID := ch.c.LocateKey(val.Bytes())

		valReassigned := false
		if newProxyID == nil {
			valAssignments.unassignValidator(val)
			valReassigned = true
		} else if proxyID == nil || newProxyID.String() != proxyID.String() {
			valAssignments.unassignValidator(val)
			valAssignments.assignValidator(val, enode.HexID(newProxyID.String()))
			valReassigned = true
		}

		if !anyAssignmentsChanged && valReassigned {
		   anyAssignmentsChanged = true
		}
	}

	return anyAssignmentsChanged
}