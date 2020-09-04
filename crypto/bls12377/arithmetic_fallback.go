// +build !amd64 generic

// Native go field arithmetic code is generated with 'goff'
// https://github.com/ConsenSys/goff
// Many function signature of field operations are renamed.

// Copyright 2020 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// field modulus q =
//
// 258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458177
// Code generated by goff DO NOT EDIT
// goff version: v0.1.0 - build: 790f1f56eac432441e043abff8819eacddd1d668
// fe are assumed to be in Montgomery form in all methods

// /!\ WARNING /!\
// this code has not been audited and is provided as-is. In particular,
// there is no security guarantees such as constant time implementation
// or side-channel attack resistance
// /!\ WARNING /!\

// Package bls (generated by goff) contains field arithmetics operations

package bls12377

import (
	"math/bits"
)

func add(z, x, y *fe) {
	var carry uint64

	z[0], carry = bits.Add64(x[0], y[0], 0)
	z[1], carry = bits.Add64(x[1], y[1], carry)
	z[2], carry = bits.Add64(x[2], y[2], carry)
	z[3], carry = bits.Add64(x[3], y[3], carry)
	z[4], carry = bits.Add64(x[4], y[4], carry)
	z[5], _ = bits.Add64(x[5], y[5], carry)

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[5] < 0x01ae3a4617c510ea || (z[5] == 0x01ae3a4617c510ea && (z[4] < 0xc63b05c06ca1493b || (z[4] == 0xc63b05c06ca1493b && (z[3] < 0x1a22d9f300f5138f || (z[3] == 0x1a22d9f300f5138f && (z[2] < 0x1ef3622fba094800 || (z[2] == 0x1ef3622fba094800 && (z[1] < 0x170b5d4430000000 || (z[1] == 0x170b5d4430000000 && (z[0] < 0x8508c00000000001))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 0x8508c00000000001, 0)
		z[1], b = bits.Sub64(z[1], 0x170b5d4430000000, b)
		z[2], b = bits.Sub64(z[2], 0x1ef3622fba094800, b)
		z[3], b = bits.Sub64(z[3], 0x1a22d9f300f5138f, b)
		z[4], b = bits.Sub64(z[4], 0xc63b05c06ca1493b, b)
		z[5], _ = bits.Sub64(z[5], 0x01ae3a4617c510ea, b)
	}
}

func addAssign(x, y *fe) {
	var carry uint64

	x[0], carry = bits.Add64(x[0], y[0], 0)
	x[1], carry = bits.Add64(x[1], y[1], carry)
	x[2], carry = bits.Add64(x[2], y[2], carry)
	x[3], carry = bits.Add64(x[3], y[3], carry)
	x[4], carry = bits.Add64(x[4], y[4], carry)
	x[5], _ = bits.Add64(x[5], y[5], carry)

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(x[5] < 0x01ae3a4617c510ea || (x[5] == 0x01ae3a4617c510ea && (x[4] < 0xc63b05c06ca1493b || (x[4] == 0xc63b05c06ca1493b && (x[3] < 0x1a22d9f300f5138f || (x[3] == 0x1a22d9f300f5138f && (x[2] < 0x1ef3622fba094800 || (x[2] == 0x1ef3622fba094800 && (x[1] < 0x170b5d4430000000 || (x[1] == 0x170b5d4430000000 && (x[0] < 0x8508c00000000001))))))))))) {
		var b uint64
		x[0], b = bits.Sub64(x[0], 0x8508c00000000001, 0)
		x[1], b = bits.Sub64(x[1], 0x170b5d4430000000, b)
		x[2], b = bits.Sub64(x[2], 0x1ef3622fba094800, b)
		x[3], b = bits.Sub64(x[3], 0x1a22d9f300f5138f, b)
		x[4], b = bits.Sub64(x[4], 0xc63b05c06ca1493b, b)
		x[5], _ = bits.Sub64(x[5], 0x01ae3a4617c510ea, b)
	}
}

func ladd(z, x, y *fe) {
	var carry uint64
	z[0], carry = bits.Add64(x[0], y[0], 0)
	z[1], carry = bits.Add64(x[1], y[1], carry)
	z[2], carry = bits.Add64(x[2], y[2], carry)
	z[3], carry = bits.Add64(x[3], y[3], carry)
	z[4], carry = bits.Add64(x[4], y[4], carry)
	z[5], _ = bits.Add64(x[5], y[5], carry)
}

func laddAssign(x, y *fe) {
	var carry uint64
	x[0], carry = bits.Add64(x[0], y[0], 0)
	x[1], carry = bits.Add64(x[1], y[1], carry)
	x[2], carry = bits.Add64(x[2], y[2], carry)
	x[3], carry = bits.Add64(x[3], y[3], carry)
	x[4], carry = bits.Add64(x[4], y[4], carry)
	x[5], _ = bits.Add64(x[5], y[5], carry)
}

func double(z, x *fe) {
	var carry uint64

	z[0], carry = bits.Add64(x[0], x[0], 0)
	z[1], carry = bits.Add64(x[1], x[1], carry)
	z[2], carry = bits.Add64(x[2], x[2], carry)
	z[3], carry = bits.Add64(x[3], x[3], carry)
	z[4], carry = bits.Add64(x[4], x[4], carry)
	z[5], _ = bits.Add64(x[5], x[5], carry)

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[5] < 0x01ae3a4617c510ea || (z[5] == 0x01ae3a4617c510ea && (z[4] < 0xc63b05c06ca1493b || (z[4] == 0xc63b05c06ca1493b && (z[3] < 0x1a22d9f300f5138f || (z[3] == 0x1a22d9f300f5138f && (z[2] < 0x1ef3622fba094800 || (z[2] == 0x1ef3622fba094800 && (z[1] < 0x170b5d4430000000 || (z[1] == 0x170b5d4430000000 && (z[0] < 0x8508c00000000001))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 0x8508c00000000001, 0)
		z[1], b = bits.Sub64(z[1], 0x170b5d4430000000, b)
		z[2], b = bits.Sub64(z[2], 0x1ef3622fba094800, b)
		z[3], b = bits.Sub64(z[3], 0x1a22d9f300f5138f, b)
		z[4], b = bits.Sub64(z[4], 0xc63b05c06ca1493b, b)
		z[5], _ = bits.Sub64(z[5], 0x01ae3a4617c510ea, b)
	}
}

func doubleAssign(z *fe) {
	var carry uint64

	z[0], carry = bits.Add64(z[0], z[0], 0)
	z[1], carry = bits.Add64(z[1], z[1], carry)
	z[2], carry = bits.Add64(z[2], z[2], carry)
	z[3], carry = bits.Add64(z[3], z[3], carry)
	z[4], carry = bits.Add64(z[4], z[4], carry)
	z[5], _ = bits.Add64(z[5], z[5], carry)

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[5] < 0x01ae3a4617c510ea || (z[5] == 0x01ae3a4617c510ea && (z[4] < 0xc63b05c06ca1493b || (z[4] == 0xc63b05c06ca1493b && (z[3] < 0x1a22d9f300f5138f || (z[3] == 0x1a22d9f300f5138f && (z[2] < 0x1ef3622fba094800 || (z[2] == 0x1ef3622fba094800 && (z[1] < 0x170b5d4430000000 || (z[1] == 0x170b5d4430000000 && (z[0] < 0x8508c00000000001))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 0x8508c00000000001, 0)
		z[1], b = bits.Sub64(z[1], 0x170b5d4430000000, b)
		z[2], b = bits.Sub64(z[2], 0x1ef3622fba094800, b)
		z[3], b = bits.Sub64(z[3], 0x1a22d9f300f5138f, b)
		z[4], b = bits.Sub64(z[4], 0xc63b05c06ca1493b, b)
		z[5], _ = bits.Sub64(z[5], 0x01ae3a4617c510ea, b)
	}
}

func ldouble(z, x *fe) {
	var carry uint64

	z[0], carry = bits.Add64(x[0], x[0], 0)
	z[1], carry = bits.Add64(x[1], x[1], carry)
	z[2], carry = bits.Add64(x[2], x[2], carry)
	z[3], carry = bits.Add64(x[3], x[3], carry)
	z[4], carry = bits.Add64(x[4], x[4], carry)
	z[5], _ = bits.Add64(x[5], x[5], carry)
}

func sub(z, x, y *fe) {
	var b uint64
	z[0], b = bits.Sub64(x[0], y[0], 0)
	z[1], b = bits.Sub64(x[1], y[1], b)
	z[2], b = bits.Sub64(x[2], y[2], b)
	z[3], b = bits.Sub64(x[3], y[3], b)
	z[4], b = bits.Sub64(x[4], y[4], b)
	z[5], b = bits.Sub64(x[5], y[5], b)
	if b != 0 {
		var c uint64
		z[0], c = bits.Add64(z[0], 0x8508c00000000001, 0)
		z[1], c = bits.Add64(z[1], 0x170b5d4430000000, c)
		z[2], c = bits.Add64(z[2], 0x1ef3622fba094800, c)
		z[3], c = bits.Add64(z[3], 0x1a22d9f300f5138f, c)
		z[4], c = bits.Add64(z[4], 0xc63b05c06ca1493b, c)
		z[5], _ = bits.Add64(z[5], 0x01ae3a4617c510ea, c)
	}
}

func subAssign(z, x *fe) {
	var b uint64
	z[0], b = bits.Sub64(z[0], x[0], 0)
	z[1], b = bits.Sub64(z[1], x[1], b)
	z[2], b = bits.Sub64(z[2], x[2], b)
	z[3], b = bits.Sub64(z[3], x[3], b)
	z[4], b = bits.Sub64(z[4], x[4], b)
	z[5], b = bits.Sub64(z[5], x[5], b)
	if b != 0 {
		var c uint64
		z[0], c = bits.Add64(z[0], 0x8508c00000000001, 0)
		z[1], c = bits.Add64(z[1], 0x170b5d4430000000, c)
		z[2], c = bits.Add64(z[2], 0x1ef3622fba094800, c)
		z[3], c = bits.Add64(z[3], 0x1a22d9f300f5138f, c)
		z[4], c = bits.Add64(z[4], 0xc63b05c06ca1493b, c)
		z[5], _ = bits.Add64(z[5], 0x01ae3a4617c510ea, c)
	}
}

func lsubAssign(z, x *fe) {
	var b uint64
	z[0], b = bits.Sub64(z[0], x[0], 0)
	z[1], b = bits.Sub64(z[1], x[1], b)
	z[2], b = bits.Sub64(z[2], x[2], b)
	z[3], b = bits.Sub64(z[3], x[3], b)
	z[4], b = bits.Sub64(z[4], x[4], b)
	z[5], _ = bits.Sub64(z[5], x[5], b)
}

func neg(z *fe, x *fe) {
	if x.isZero() {
		z.zero()
		return
	}
	var borrow uint64
	z[0], borrow = bits.Sub64(0x8508c00000000001, x[0], 0)
	z[1], borrow = bits.Sub64(0x170b5d4430000000, x[1], borrow)
	z[2], borrow = bits.Sub64(0x1ef3622fba094800, x[2], borrow)
	z[3], borrow = bits.Sub64(0x1a22d9f300f5138f, x[3], borrow)
	z[4], borrow = bits.Sub64(0xc63b05c06ca1493b, x[4], borrow)
	z[5], _ = bits.Sub64(0x01ae3a4617c510ea, x[5], borrow)
}

func mul(z, x, y *fe) {
	var t [6]uint64
	var c [3]uint64
	{
		// round 0
		v := x[0]
		c[1], c[0] = bits.Mul64(v, y[0])
		m := c[0] * 0x8508bfffffffffff
		c[2] = madd0(m, 0x8508c00000000001, c[0])
		c[1], c[0] = madd1(v, y[1], c[1])
		c[2], t[0] = madd2(m, 0x170b5d4430000000, c[2], c[0])
		c[1], c[0] = madd1(v, y[2], c[1])
		c[2], t[1] = madd2(m, 0x1ef3622fba094800, c[2], c[0])
		c[1], c[0] = madd1(v, y[3], c[1])
		c[2], t[2] = madd2(m, 0x1a22d9f300f5138f, c[2], c[0])
		c[1], c[0] = madd1(v, y[4], c[1])
		c[2], t[3] = madd2(m, 0xc63b05c06ca1493b, c[2], c[0])
		c[1], c[0] = madd1(v, y[5], c[1])
		t[5], t[4] = madd3(m, 0x01ae3a4617c510ea, c[0], c[2], c[1])
	}
	{
		// round 1
		v := x[1]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 0x8508bfffffffffff
		c[2] = madd0(m, 0x8508c00000000001, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 0x170b5d4430000000, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 0x1ef3622fba094800, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], t[2] = madd2(m, 0x1a22d9f300f5138f, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], t[3] = madd2(m, 0xc63b05c06ca1493b, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		t[5], t[4] = madd3(m, 0x01ae3a4617c510ea, c[0], c[2], c[1])
	}
	{
		// round 2
		v := x[2]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 0x8508bfffffffffff
		c[2] = madd0(m, 0x8508c00000000001, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 0x170b5d4430000000, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 0x1ef3622fba094800, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], t[2] = madd2(m, 0x1a22d9f300f5138f, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], t[3] = madd2(m, 0xc63b05c06ca1493b, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		t[5], t[4] = madd3(m, 0x01ae3a4617c510ea, c[0], c[2], c[1])
	}
	{
		// round 3
		v := x[3]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 0x8508bfffffffffff
		c[2] = madd0(m, 0x8508c00000000001, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 0x170b5d4430000000, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 0x1ef3622fba094800, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], t[2] = madd2(m, 0x1a22d9f300f5138f, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], t[3] = madd2(m, 0xc63b05c06ca1493b, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		t[5], t[4] = madd3(m, 0x01ae3a4617c510ea, c[0], c[2], c[1])
	}
	{
		// round 4
		v := x[4]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 0x8508bfffffffffff
		c[2] = madd0(m, 0x8508c00000000001, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 0x170b5d4430000000, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 0x1ef3622fba094800, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], t[2] = madd2(m, 0x1a22d9f300f5138f, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], t[3] = madd2(m, 0xc63b05c06ca1493b, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		t[5], t[4] = madd3(m, 0x01ae3a4617c510ea, c[0], c[2], c[1])
	}
	{
		// round 5
		v := x[5]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 0x8508bfffffffffff
		c[2] = madd0(m, 0x8508c00000000001, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], z[0] = madd2(m, 0x170b5d4430000000, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], z[1] = madd2(m, 0x1ef3622fba094800, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], z[2] = madd2(m, 0x1a22d9f300f5138f, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], z[3] = madd2(m, 0xc63b05c06ca1493b, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		z[5], z[4] = madd3(m, 0x01ae3a4617c510ea, c[0], c[2], c[1])
	}

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[5] < 0x01ae3a4617c510ea || (z[5] == 0x01ae3a4617c510ea && (z[4] < 0xc63b05c06ca1493b || (z[4] == 0xc63b05c06ca1493b && (z[3] < 0x1a22d9f300f5138f || (z[3] == 0x1a22d9f300f5138f && (z[2] < 0x1ef3622fba094800 || (z[2] == 0x1ef3622fba094800 && (z[1] < 0x170b5d4430000000 || (z[1] == 0x170b5d4430000000 && (z[0] < 0x8508c00000000001))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 0x8508c00000000001, 0)
		z[1], b = bits.Sub64(z[1], 0x170b5d4430000000, b)
		z[2], b = bits.Sub64(z[2], 0x1ef3622fba094800, b)
		z[3], b = bits.Sub64(z[3], 0x1a22d9f300f5138f, b)
		z[4], b = bits.Sub64(z[4], 0xc63b05c06ca1493b, b)
		z[5], _ = bits.Sub64(z[5], 0x01ae3a4617c510ea, b)
	}
}

func square(z, x *fe) {

	var p [6]uint64

	var u, v uint64
	{
		// round 0
		u, p[0] = bits.Mul64(x[0], x[0])
		m := p[0] * 0x8508bfffffffffff
		C := madd0(m, 0x8508c00000000001, p[0])
		var t uint64
		t, u, v = madd1sb(x[0], x[1], u)
		C, p[0] = madd2(m, 0x170b5d4430000000, v, C)
		t, u, v = madd1s(x[0], x[2], t, u)
		C, p[1] = madd2(m, 0x1ef3622fba094800, v, C)
		t, u, v = madd1s(x[0], x[3], t, u)
		C, p[2] = madd2(m, 0x1a22d9f300f5138f, v, C)
		t, u, v = madd1s(x[0], x[4], t, u)
		C, p[3] = madd2(m, 0xc63b05c06ca1493b, v, C)
		_, u, v = madd1s(x[0], x[5], t, u)
		p[5], p[4] = madd3(m, 0x01ae3a4617c510ea, v, C, u)
	}
	{
		// round 1
		m := p[0] * 0x8508bfffffffffff
		C := madd0(m, 0x8508c00000000001, p[0])
		u, v = madd1(x[1], x[1], p[1])
		C, p[0] = madd2(m, 0x170b5d4430000000, v, C)
		var t uint64
		t, u, v = madd2sb(x[1], x[2], p[2], u)
		C, p[1] = madd2(m, 0x1ef3622fba094800, v, C)
		t, u, v = madd2s(x[1], x[3], p[3], t, u)
		C, p[2] = madd2(m, 0x1a22d9f300f5138f, v, C)
		t, u, v = madd2s(x[1], x[4], p[4], t, u)
		C, p[3] = madd2(m, 0xc63b05c06ca1493b, v, C)
		_, u, v = madd2s(x[1], x[5], p[5], t, u)
		p[5], p[4] = madd3(m, 0x01ae3a4617c510ea, v, C, u)
	}
	{
		// round 2
		m := p[0] * 0x8508bfffffffffff
		C := madd0(m, 0x8508c00000000001, p[0])
		C, p[0] = madd2(m, 0x170b5d4430000000, p[1], C)
		u, v = madd1(x[2], x[2], p[2])
		C, p[1] = madd2(m, 0x1ef3622fba094800, v, C)
		var t uint64
		t, u, v = madd2sb(x[2], x[3], p[3], u)
		C, p[2] = madd2(m, 0x1a22d9f300f5138f, v, C)
		t, u, v = madd2s(x[2], x[4], p[4], t, u)
		C, p[3] = madd2(m, 0xc63b05c06ca1493b, v, C)
		_, u, v = madd2s(x[2], x[5], p[5], t, u)
		p[5], p[4] = madd3(m, 0x01ae3a4617c510ea, v, C, u)
	}
	{
		// round 3
		m := p[0] * 0x8508bfffffffffff
		C := madd0(m, 0x8508c00000000001, p[0])
		C, p[0] = madd2(m, 0x170b5d4430000000, p[1], C)
		C, p[1] = madd2(m, 0x1ef3622fba094800, p[2], C)
		u, v = madd1(x[3], x[3], p[3])
		C, p[2] = madd2(m, 0x1a22d9f300f5138f, v, C)
		var t uint64
		t, u, v = madd2sb(x[3], x[4], p[4], u)
		C, p[3] = madd2(m, 0xc63b05c06ca1493b, v, C)
		_, u, v = madd2s(x[3], x[5], p[5], t, u)
		p[5], p[4] = madd3(m, 0x01ae3a4617c510ea, v, C, u)
	}
	{
		// round 4
		m := p[0] * 0x8508bfffffffffff
		C := madd0(m, 0x8508c00000000001, p[0])
		C, p[0] = madd2(m, 0x170b5d4430000000, p[1], C)
		C, p[1] = madd2(m, 0x1ef3622fba094800, p[2], C)
		C, p[2] = madd2(m, 0x1a22d9f300f5138f, p[3], C)
		u, v = madd1(x[4], x[4], p[4])
		C, p[3] = madd2(m, 0xc63b05c06ca1493b, v, C)
		_, u, v = madd2sb(x[4], x[5], p[5], u)
		p[5], p[4] = madd3(m, 0x01ae3a4617c510ea, v, C, u)
	}
	{
		// round 5
		m := p[0] * 0x8508bfffffffffff
		C := madd0(m, 0x8508c00000000001, p[0])
		C, z[0] = madd2(m, 0x170b5d4430000000, p[1], C)
		C, z[1] = madd2(m, 0x1ef3622fba094800, p[2], C)
		C, z[2] = madd2(m, 0x1a22d9f300f5138f, p[3], C)
		C, z[3] = madd2(m, 0xc63b05c06ca1493b, p[4], C)
		u, v = madd1(x[5], x[5], p[5])
		z[5], z[4] = madd3(m, 0x01ae3a4617c510ea, v, C, u)
	}

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[5] < 0x01ae3a4617c510ea || (z[5] == 0x01ae3a4617c510ea && (z[4] < 0xc63b05c06ca1493b || (z[4] == 0xc63b05c06ca1493b && (z[3] < 0x1a22d9f300f5138f || (z[3] == 0x1a22d9f300f5138f && (z[2] < 0x1ef3622fba094800 || (z[2] == 0x1ef3622fba094800 && (z[1] < 0x170b5d4430000000 || (z[1] == 0x170b5d4430000000 && (z[0] < 0x8508c00000000001))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 0x8508c00000000001, 0)
		z[1], b = bits.Sub64(z[1], 0x170b5d4430000000, b)
		z[2], b = bits.Sub64(z[2], 0x1ef3622fba094800, b)
		z[3], b = bits.Sub64(z[3], 0x1a22d9f300f5138f, b)
		z[4], b = bits.Sub64(z[4], 0xc63b05c06ca1493b, b)
		z[5], _ = bits.Sub64(z[5], 0x01ae3a4617c510ea, b)
	}
}

// arith.go
// Copyright 2020 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by goff DO NOT EDIT

func madd(a, b, t, u, v uint64) (uint64, uint64, uint64) {
	var carry uint64
	hi, lo := bits.Mul64(a, b)
	v, carry = bits.Add64(lo, v, 0)
	u, carry = bits.Add64(hi, u, carry)
	t, _ = bits.Add64(t, 0, carry)
	return t, u, v
}

// madd0 hi = a*b + c (discards lo bits)
func madd0(a, b, c uint64) (hi uint64) {
	var carry, lo uint64
	hi, lo = bits.Mul64(a, b)
	_, carry = bits.Add64(lo, c, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return
}

// madd1 hi, lo = a*b + c
func madd1(a, b, c uint64) (hi uint64, lo uint64) {
	var carry uint64
	hi, lo = bits.Mul64(a, b)
	lo, carry = bits.Add64(lo, c, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return
}

// madd2 hi, lo = a*b + c + d
func madd2(a, b, c, d uint64) (hi uint64, lo uint64) {
	var carry uint64
	hi, lo = bits.Mul64(a, b)
	c, carry = bits.Add64(c, d, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	lo, carry = bits.Add64(lo, c, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return
}

// madd2s superhi, hi, lo = 2*a*b + c + d + e
func madd2s(a, b, c, d, e uint64) (superhi, hi, lo uint64) {
	var carry, sum uint64

	hi, lo = bits.Mul64(a, b)
	lo, carry = bits.Add64(lo, lo, 0)
	hi, superhi = bits.Add64(hi, hi, carry)

	sum, carry = bits.Add64(c, e, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	lo, carry = bits.Add64(lo, sum, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	hi, _ = bits.Add64(hi, 0, d)
	return
}

func madd1s(a, b, d, e uint64) (superhi, hi, lo uint64) {
	var carry uint64

	hi, lo = bits.Mul64(a, b)
	lo, carry = bits.Add64(lo, lo, 0)
	hi, superhi = bits.Add64(hi, hi, carry)
	lo, carry = bits.Add64(lo, e, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	hi, _ = bits.Add64(hi, 0, d)
	return
}

func madd2sb(a, b, c, e uint64) (superhi, hi, lo uint64) {
	var carry, sum uint64

	hi, lo = bits.Mul64(a, b)
	lo, carry = bits.Add64(lo, lo, 0)
	hi, superhi = bits.Add64(hi, hi, carry)

	sum, carry = bits.Add64(c, e, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	lo, carry = bits.Add64(lo, sum, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return
}

func madd1sb(a, b, e uint64) (superhi, hi, lo uint64) {
	var carry uint64

	hi, lo = bits.Mul64(a, b)
	lo, carry = bits.Add64(lo, lo, 0)
	hi, superhi = bits.Add64(hi, hi, carry)
	lo, carry = bits.Add64(lo, e, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return
}

func madd3(a, b, c, d, e uint64) (hi uint64, lo uint64) {
	var carry uint64
	hi, lo = bits.Mul64(a, b)
	c, carry = bits.Add64(c, d, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	lo, carry = bits.Add64(lo, c, 0)
	hi, _ = bits.Add64(hi, e, carry)
	return
}
