package evmmax_arith

import (
	"encoding/binary"
	"math/bits"
)

var Preset = []arithFunc{
	MontMul64,
	MontMul128,
	MontMul192,
	MontMul256,
	MontMul320,
	MontMul384,
	MontMul448,
	MontMul512,
	MontMul576,
	MontMul640,
	MontMul704,
	MontMul768,
	MontMul832,
	MontMul896,
}

func MontMul64(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [2]uint64
	var D uint64
	var m, C uint64

	var (
		mod [1]uint64
		x   [1]uint64
		y   [1]uint64
		res [1]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[7]
	_ = yBytes[7]
	_ = outBytes[7]
	_ = modBytes[7]

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])

	t[1], D = bits.Add64(t[1], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	t[0], C = bits.Add64(t[1], C, 0)
	t[1], _ = bits.Add64(0, D, C)

	for j := 1; j < 1; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		t[1], D = bits.Add64(t[1], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		t[0], C = bits.Add64(t[1], C, 0)
		t[1], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)

	var src []uint64
	if D != 0 && t[1] == 0 {
		src = t[:1]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[0])
}

func MontMul128(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [3]uint64
	var D uint64
	var m, C uint64

	var (
		mod [2]uint64
		x   [2]uint64
		y   [2]uint64
		res [2]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[15]
	_ = yBytes[15]
	_ = outBytes[15]
	_ = modBytes[15]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)

	t[2], D = bits.Add64(t[2], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	t[1], C = bits.Add64(t[2], C, 0)
	t[2], _ = bits.Add64(0, D, C)

	for j := 1; j < 2; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		t[2], D = bits.Add64(t[2], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		t[1], C = bits.Add64(t[2], C, 0)
		t[2], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)

	var src []uint64
	if D != 0 && t[2] == 0 {
		src = t[:2]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[1])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[0])
}

func MontMul192(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [4]uint64
	var D uint64
	var m, C uint64

	var (
		mod [3]uint64
		x   [3]uint64
		y   [3]uint64
		res [3]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[23]
	_ = yBytes[23]
	_ = outBytes[23]
	_ = modBytes[23]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)

	t[3], D = bits.Add64(t[3], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	t[2], C = bits.Add64(t[3], C, 0)
	t[3], _ = bits.Add64(0, D, C)

	for j := 1; j < 3; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		t[3], D = bits.Add64(t[3], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		t[2], C = bits.Add64(t[3], C, 0)
		t[3], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)

	var src []uint64
	if D != 0 && t[3] == 0 {
		src = t[:3]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[2])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[1])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[0])
}

func MontMul256(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [5]uint64
	var D uint64
	var m, C uint64

	var (
		mod [4]uint64
		x   [4]uint64
		y   [4]uint64
		res [4]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[31]
	_ = yBytes[31]
	_ = outBytes[31]
	_ = modBytes[31]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	x[2] = binary.LittleEndian.Uint64(xBytes[16:24])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	y[2] = binary.LittleEndian.Uint64(yBytes[16:24])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])
	mod[2] = binary.LittleEndian.Uint64(modBytes[16:24])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)

	t[4], D = bits.Add64(t[4], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	t[3], C = bits.Add64(t[4], C, 0)
	t[4], _ = bits.Add64(0, D, C)

	for j := 1; j < 4; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		C, t[3] = madd2(x[j], y[3], t[3], C)
		t[4], D = bits.Add64(t[4], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		t[3], C = bits.Add64(t[4], C, 0)
		t[4], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)
	res[3], D = bits.Sub64(t[3], mod[3], D)

	var src []uint64
	if D != 0 && t[4] == 0 {
		src = t[:4]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[3])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[2])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[1])
	binary.LittleEndian.PutUint64(outBytes[24:32], src[0])
}

func MontMul320(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [6]uint64
	var D uint64
	var m, C uint64

	var (
		mod [5]uint64
		x   [5]uint64
		y   [5]uint64
		res [5]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[39]
	_ = yBytes[39]
	_ = outBytes[39]
	_ = modBytes[39]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	x[2] = binary.LittleEndian.Uint64(xBytes[16:24])
	x[3] = binary.LittleEndian.Uint64(xBytes[24:32])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	y[2] = binary.LittleEndian.Uint64(yBytes[16:24])
	y[3] = binary.LittleEndian.Uint64(yBytes[24:32])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])
	mod[2] = binary.LittleEndian.Uint64(modBytes[16:24])
	mod[3] = binary.LittleEndian.Uint64(modBytes[24:32])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)

	t[5], D = bits.Add64(t[5], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	t[4], C = bits.Add64(t[5], C, 0)
	t[5], _ = bits.Add64(0, D, C)

	for j := 1; j < 5; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		C, t[3] = madd2(x[j], y[3], t[3], C)
		C, t[4] = madd2(x[j], y[4], t[4], C)
		t[5], D = bits.Add64(t[5], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		C, t[3] = madd2(m, mod[4], t[4], C)
		t[4], C = bits.Add64(t[5], C, 0)
		t[5], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)
	res[3], D = bits.Sub64(t[3], mod[3], D)
	res[4], D = bits.Sub64(t[4], mod[4], D)

	var src []uint64
	if D != 0 && t[5] == 0 {
		src = t[:5]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[4])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[3])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[2])
	binary.LittleEndian.PutUint64(outBytes[24:32], src[1])
	binary.LittleEndian.PutUint64(outBytes[32:40], src[0])
}

func MontMul384(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [7]uint64
	var D uint64
	var m, C uint64

	var (
		mod [6]uint64
		x   [6]uint64
		y   [6]uint64
		res [6]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[47]
	_ = yBytes[47]
	_ = outBytes[47]
	_ = modBytes[47]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	x[2] = binary.LittleEndian.Uint64(xBytes[16:24])
	x[3] = binary.LittleEndian.Uint64(xBytes[24:32])
	x[4] = binary.LittleEndian.Uint64(xBytes[32:40])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	y[2] = binary.LittleEndian.Uint64(yBytes[16:24])
	y[3] = binary.LittleEndian.Uint64(yBytes[24:32])
	y[4] = binary.LittleEndian.Uint64(yBytes[32:40])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])
	mod[2] = binary.LittleEndian.Uint64(modBytes[16:24])
	mod[3] = binary.LittleEndian.Uint64(modBytes[24:32])
	mod[4] = binary.LittleEndian.Uint64(modBytes[32:40])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)
	C, t[5] = madd1(x[0], y[5], C)

	t[6], D = bits.Add64(t[6], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	t[5], C = bits.Add64(t[6], C, 0)
	t[6], _ = bits.Add64(0, D, C)

	for j := 1; j < 6; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		C, t[3] = madd2(x[j], y[3], t[3], C)
		C, t[4] = madd2(x[j], y[4], t[4], C)
		C, t[5] = madd2(x[j], y[5], t[5], C)
		t[6], D = bits.Add64(t[6], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		C, t[3] = madd2(m, mod[4], t[4], C)
		C, t[4] = madd2(m, mod[5], t[5], C)
		t[5], C = bits.Add64(t[6], C, 0)
		t[6], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)
	res[3], D = bits.Sub64(t[3], mod[3], D)
	res[4], D = bits.Sub64(t[4], mod[4], D)
	res[5], D = bits.Sub64(t[5], mod[5], D)

	var src []uint64
	if D != 0 && t[6] == 0 {
		src = t[:6]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[5])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[4])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[3])
	binary.LittleEndian.PutUint64(outBytes[24:32], src[2])
	binary.LittleEndian.PutUint64(outBytes[32:40], src[1])
	binary.LittleEndian.PutUint64(outBytes[40:48], src[0])
}

func MontMul448(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [8]uint64
	var D uint64
	var m, C uint64

	var (
		mod [7]uint64
		x   [7]uint64
		y   [7]uint64
		res [7]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[55]
	_ = yBytes[55]
	_ = outBytes[55]
	_ = modBytes[55]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	x[2] = binary.LittleEndian.Uint64(xBytes[16:24])
	x[3] = binary.LittleEndian.Uint64(xBytes[24:32])
	x[4] = binary.LittleEndian.Uint64(xBytes[32:40])
	x[5] = binary.LittleEndian.Uint64(xBytes[40:48])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	y[2] = binary.LittleEndian.Uint64(yBytes[16:24])
	y[3] = binary.LittleEndian.Uint64(yBytes[24:32])
	y[4] = binary.LittleEndian.Uint64(yBytes[32:40])
	y[5] = binary.LittleEndian.Uint64(yBytes[40:48])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])
	mod[2] = binary.LittleEndian.Uint64(modBytes[16:24])
	mod[3] = binary.LittleEndian.Uint64(modBytes[24:32])
	mod[4] = binary.LittleEndian.Uint64(modBytes[32:40])
	mod[5] = binary.LittleEndian.Uint64(modBytes[40:48])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)
	C, t[5] = madd1(x[0], y[5], C)
	C, t[6] = madd1(x[0], y[6], C)

	t[7], D = bits.Add64(t[7], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	t[6], C = bits.Add64(t[7], C, 0)
	t[7], _ = bits.Add64(0, D, C)

	for j := 1; j < 7; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		C, t[3] = madd2(x[j], y[3], t[3], C)
		C, t[4] = madd2(x[j], y[4], t[4], C)
		C, t[5] = madd2(x[j], y[5], t[5], C)
		C, t[6] = madd2(x[j], y[6], t[6], C)
		t[7], D = bits.Add64(t[7], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		C, t[3] = madd2(m, mod[4], t[4], C)
		C, t[4] = madd2(m, mod[5], t[5], C)
		C, t[5] = madd2(m, mod[6], t[6], C)
		t[6], C = bits.Add64(t[7], C, 0)
		t[7], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)
	res[3], D = bits.Sub64(t[3], mod[3], D)
	res[4], D = bits.Sub64(t[4], mod[4], D)
	res[5], D = bits.Sub64(t[5], mod[5], D)
	res[6], D = bits.Sub64(t[6], mod[6], D)

	var src []uint64
	if D != 0 && t[7] == 0 {
		src = t[:7]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[6])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[5])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[4])
	binary.LittleEndian.PutUint64(outBytes[24:32], src[3])
	binary.LittleEndian.PutUint64(outBytes[32:40], src[2])
	binary.LittleEndian.PutUint64(outBytes[40:48], src[1])
	binary.LittleEndian.PutUint64(outBytes[48:56], src[0])
}

func MontMul512(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [9]uint64
	var D uint64
	var m, C uint64

	var (
		mod [8]uint64
		x   [8]uint64
		y   [8]uint64
		res [8]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[63]
	_ = yBytes[63]
	_ = outBytes[63]
	_ = modBytes[63]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	x[2] = binary.LittleEndian.Uint64(xBytes[16:24])
	x[3] = binary.LittleEndian.Uint64(xBytes[24:32])
	x[4] = binary.LittleEndian.Uint64(xBytes[32:40])
	x[5] = binary.LittleEndian.Uint64(xBytes[40:48])
	x[6] = binary.LittleEndian.Uint64(xBytes[48:56])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	y[2] = binary.LittleEndian.Uint64(yBytes[16:24])
	y[3] = binary.LittleEndian.Uint64(yBytes[24:32])
	y[4] = binary.LittleEndian.Uint64(yBytes[32:40])
	y[5] = binary.LittleEndian.Uint64(yBytes[40:48])
	y[6] = binary.LittleEndian.Uint64(yBytes[48:56])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])
	mod[2] = binary.LittleEndian.Uint64(modBytes[16:24])
	mod[3] = binary.LittleEndian.Uint64(modBytes[24:32])
	mod[4] = binary.LittleEndian.Uint64(modBytes[32:40])
	mod[5] = binary.LittleEndian.Uint64(modBytes[40:48])
	mod[6] = binary.LittleEndian.Uint64(modBytes[48:56])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)
	C, t[5] = madd1(x[0], y[5], C)
	C, t[6] = madd1(x[0], y[6], C)
	C, t[7] = madd1(x[0], y[7], C)

	t[8], D = bits.Add64(t[8], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	C, t[6] = madd2(m, mod[7], t[7], C)
	t[7], C = bits.Add64(t[8], C, 0)
	t[8], _ = bits.Add64(0, D, C)

	for j := 1; j < 8; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		C, t[3] = madd2(x[j], y[3], t[3], C)
		C, t[4] = madd2(x[j], y[4], t[4], C)
		C, t[5] = madd2(x[j], y[5], t[5], C)
		C, t[6] = madd2(x[j], y[6], t[6], C)
		C, t[7] = madd2(x[j], y[7], t[7], C)
		t[8], D = bits.Add64(t[8], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		C, t[3] = madd2(m, mod[4], t[4], C)
		C, t[4] = madd2(m, mod[5], t[5], C)
		C, t[5] = madd2(m, mod[6], t[6], C)
		C, t[6] = madd2(m, mod[7], t[7], C)
		t[7], C = bits.Add64(t[8], C, 0)
		t[8], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)
	res[3], D = bits.Sub64(t[3], mod[3], D)
	res[4], D = bits.Sub64(t[4], mod[4], D)
	res[5], D = bits.Sub64(t[5], mod[5], D)
	res[6], D = bits.Sub64(t[6], mod[6], D)
	res[7], D = bits.Sub64(t[7], mod[7], D)

	var src []uint64
	if D != 0 && t[8] == 0 {
		src = t[:8]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[7])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[6])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[5])
	binary.LittleEndian.PutUint64(outBytes[24:32], src[4])
	binary.LittleEndian.PutUint64(outBytes[32:40], src[3])
	binary.LittleEndian.PutUint64(outBytes[40:48], src[2])
	binary.LittleEndian.PutUint64(outBytes[48:56], src[1])
	binary.LittleEndian.PutUint64(outBytes[56:64], src[0])
}

func MontMul576(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [10]uint64
	var D uint64
	var m, C uint64

	var (
		mod [9]uint64
		x   [9]uint64
		y   [9]uint64
		res [9]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[71]
	_ = yBytes[71]
	_ = outBytes[71]
	_ = modBytes[71]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	x[2] = binary.LittleEndian.Uint64(xBytes[16:24])
	x[3] = binary.LittleEndian.Uint64(xBytes[24:32])
	x[4] = binary.LittleEndian.Uint64(xBytes[32:40])
	x[5] = binary.LittleEndian.Uint64(xBytes[40:48])
	x[6] = binary.LittleEndian.Uint64(xBytes[48:56])
	x[7] = binary.LittleEndian.Uint64(xBytes[56:64])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	y[2] = binary.LittleEndian.Uint64(yBytes[16:24])
	y[3] = binary.LittleEndian.Uint64(yBytes[24:32])
	y[4] = binary.LittleEndian.Uint64(yBytes[32:40])
	y[5] = binary.LittleEndian.Uint64(yBytes[40:48])
	y[6] = binary.LittleEndian.Uint64(yBytes[48:56])
	y[7] = binary.LittleEndian.Uint64(yBytes[56:64])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])
	mod[2] = binary.LittleEndian.Uint64(modBytes[16:24])
	mod[3] = binary.LittleEndian.Uint64(modBytes[24:32])
	mod[4] = binary.LittleEndian.Uint64(modBytes[32:40])
	mod[5] = binary.LittleEndian.Uint64(modBytes[40:48])
	mod[6] = binary.LittleEndian.Uint64(modBytes[48:56])
	mod[7] = binary.LittleEndian.Uint64(modBytes[56:64])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)
	C, t[5] = madd1(x[0], y[5], C)
	C, t[6] = madd1(x[0], y[6], C)
	C, t[7] = madd1(x[0], y[7], C)
	C, t[8] = madd1(x[0], y[8], C)

	t[9], D = bits.Add64(t[9], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	C, t[6] = madd2(m, mod[7], t[7], C)
	C, t[7] = madd2(m, mod[8], t[8], C)
	t[8], C = bits.Add64(t[9], C, 0)
	t[9], _ = bits.Add64(0, D, C)

	for j := 1; j < 9; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		C, t[3] = madd2(x[j], y[3], t[3], C)
		C, t[4] = madd2(x[j], y[4], t[4], C)
		C, t[5] = madd2(x[j], y[5], t[5], C)
		C, t[6] = madd2(x[j], y[6], t[6], C)
		C, t[7] = madd2(x[j], y[7], t[7], C)
		C, t[8] = madd2(x[j], y[8], t[8], C)
		t[9], D = bits.Add64(t[9], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		C, t[3] = madd2(m, mod[4], t[4], C)
		C, t[4] = madd2(m, mod[5], t[5], C)
		C, t[5] = madd2(m, mod[6], t[6], C)
		C, t[6] = madd2(m, mod[7], t[7], C)
		C, t[7] = madd2(m, mod[8], t[8], C)
		t[8], C = bits.Add64(t[9], C, 0)
		t[9], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)
	res[3], D = bits.Sub64(t[3], mod[3], D)
	res[4], D = bits.Sub64(t[4], mod[4], D)
	res[5], D = bits.Sub64(t[5], mod[5], D)
	res[6], D = bits.Sub64(t[6], mod[6], D)
	res[7], D = bits.Sub64(t[7], mod[7], D)
	res[8], D = bits.Sub64(t[8], mod[8], D)

	var src []uint64
	if D != 0 && t[9] == 0 {
		src = t[:9]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[8])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[7])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[6])
	binary.LittleEndian.PutUint64(outBytes[24:32], src[5])
	binary.LittleEndian.PutUint64(outBytes[32:40], src[4])
	binary.LittleEndian.PutUint64(outBytes[40:48], src[3])
	binary.LittleEndian.PutUint64(outBytes[48:56], src[2])
	binary.LittleEndian.PutUint64(outBytes[56:64], src[1])
	binary.LittleEndian.PutUint64(outBytes[64:72], src[0])
}

func MontMul640(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [11]uint64
	var D uint64
	var m, C uint64

	var (
		mod [10]uint64
		x   [10]uint64
		y   [10]uint64
		res [10]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[79]
	_ = yBytes[79]
	_ = outBytes[79]
	_ = modBytes[79]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	x[2] = binary.LittleEndian.Uint64(xBytes[16:24])
	x[3] = binary.LittleEndian.Uint64(xBytes[24:32])
	x[4] = binary.LittleEndian.Uint64(xBytes[32:40])
	x[5] = binary.LittleEndian.Uint64(xBytes[40:48])
	x[6] = binary.LittleEndian.Uint64(xBytes[48:56])
	x[7] = binary.LittleEndian.Uint64(xBytes[56:64])
	x[8] = binary.LittleEndian.Uint64(xBytes[64:72])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	y[2] = binary.LittleEndian.Uint64(yBytes[16:24])
	y[3] = binary.LittleEndian.Uint64(yBytes[24:32])
	y[4] = binary.LittleEndian.Uint64(yBytes[32:40])
	y[5] = binary.LittleEndian.Uint64(yBytes[40:48])
	y[6] = binary.LittleEndian.Uint64(yBytes[48:56])
	y[7] = binary.LittleEndian.Uint64(yBytes[56:64])
	y[8] = binary.LittleEndian.Uint64(yBytes[64:72])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])
	mod[2] = binary.LittleEndian.Uint64(modBytes[16:24])
	mod[3] = binary.LittleEndian.Uint64(modBytes[24:32])
	mod[4] = binary.LittleEndian.Uint64(modBytes[32:40])
	mod[5] = binary.LittleEndian.Uint64(modBytes[40:48])
	mod[6] = binary.LittleEndian.Uint64(modBytes[48:56])
	mod[7] = binary.LittleEndian.Uint64(modBytes[56:64])
	mod[8] = binary.LittleEndian.Uint64(modBytes[64:72])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)
	C, t[5] = madd1(x[0], y[5], C)
	C, t[6] = madd1(x[0], y[6], C)
	C, t[7] = madd1(x[0], y[7], C)
	C, t[8] = madd1(x[0], y[8], C)
	C, t[9] = madd1(x[0], y[9], C)

	t[10], D = bits.Add64(t[10], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	C, t[6] = madd2(m, mod[7], t[7], C)
	C, t[7] = madd2(m, mod[8], t[8], C)
	C, t[8] = madd2(m, mod[9], t[9], C)
	t[9], C = bits.Add64(t[10], C, 0)
	t[10], _ = bits.Add64(0, D, C)

	for j := 1; j < 10; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		C, t[3] = madd2(x[j], y[3], t[3], C)
		C, t[4] = madd2(x[j], y[4], t[4], C)
		C, t[5] = madd2(x[j], y[5], t[5], C)
		C, t[6] = madd2(x[j], y[6], t[6], C)
		C, t[7] = madd2(x[j], y[7], t[7], C)
		C, t[8] = madd2(x[j], y[8], t[8], C)
		C, t[9] = madd2(x[j], y[9], t[9], C)
		t[10], D = bits.Add64(t[10], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		C, t[3] = madd2(m, mod[4], t[4], C)
		C, t[4] = madd2(m, mod[5], t[5], C)
		C, t[5] = madd2(m, mod[6], t[6], C)
		C, t[6] = madd2(m, mod[7], t[7], C)
		C, t[7] = madd2(m, mod[8], t[8], C)
		C, t[8] = madd2(m, mod[9], t[9], C)
		t[9], C = bits.Add64(t[10], C, 0)
		t[10], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)
	res[3], D = bits.Sub64(t[3], mod[3], D)
	res[4], D = bits.Sub64(t[4], mod[4], D)
	res[5], D = bits.Sub64(t[5], mod[5], D)
	res[6], D = bits.Sub64(t[6], mod[6], D)
	res[7], D = bits.Sub64(t[7], mod[7], D)
	res[8], D = bits.Sub64(t[8], mod[8], D)
	res[9], D = bits.Sub64(t[9], mod[9], D)

	var src []uint64
	if D != 0 && t[10] == 0 {
		src = t[:10]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[9])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[8])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[7])
	binary.LittleEndian.PutUint64(outBytes[24:32], src[6])
	binary.LittleEndian.PutUint64(outBytes[32:40], src[5])
	binary.LittleEndian.PutUint64(outBytes[40:48], src[4])
	binary.LittleEndian.PutUint64(outBytes[48:56], src[3])
	binary.LittleEndian.PutUint64(outBytes[56:64], src[2])
	binary.LittleEndian.PutUint64(outBytes[64:72], src[1])
	binary.LittleEndian.PutUint64(outBytes[72:80], src[0])
}

func MontMul704(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [12]uint64
	var D uint64
	var m, C uint64

	var (
		mod [11]uint64
		x   [11]uint64
		y   [11]uint64
		res [11]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[87]
	_ = yBytes[87]
	_ = outBytes[87]
	_ = modBytes[87]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	x[2] = binary.LittleEndian.Uint64(xBytes[16:24])
	x[3] = binary.LittleEndian.Uint64(xBytes[24:32])
	x[4] = binary.LittleEndian.Uint64(xBytes[32:40])
	x[5] = binary.LittleEndian.Uint64(xBytes[40:48])
	x[6] = binary.LittleEndian.Uint64(xBytes[48:56])
	x[7] = binary.LittleEndian.Uint64(xBytes[56:64])
	x[8] = binary.LittleEndian.Uint64(xBytes[64:72])
	x[9] = binary.LittleEndian.Uint64(xBytes[72:80])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	y[2] = binary.LittleEndian.Uint64(yBytes[16:24])
	y[3] = binary.LittleEndian.Uint64(yBytes[24:32])
	y[4] = binary.LittleEndian.Uint64(yBytes[32:40])
	y[5] = binary.LittleEndian.Uint64(yBytes[40:48])
	y[6] = binary.LittleEndian.Uint64(yBytes[48:56])
	y[7] = binary.LittleEndian.Uint64(yBytes[56:64])
	y[8] = binary.LittleEndian.Uint64(yBytes[64:72])
	y[9] = binary.LittleEndian.Uint64(yBytes[72:80])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])
	mod[2] = binary.LittleEndian.Uint64(modBytes[16:24])
	mod[3] = binary.LittleEndian.Uint64(modBytes[24:32])
	mod[4] = binary.LittleEndian.Uint64(modBytes[32:40])
	mod[5] = binary.LittleEndian.Uint64(modBytes[40:48])
	mod[6] = binary.LittleEndian.Uint64(modBytes[48:56])
	mod[7] = binary.LittleEndian.Uint64(modBytes[56:64])
	mod[8] = binary.LittleEndian.Uint64(modBytes[64:72])
	mod[9] = binary.LittleEndian.Uint64(modBytes[72:80])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)
	C, t[5] = madd1(x[0], y[5], C)
	C, t[6] = madd1(x[0], y[6], C)
	C, t[7] = madd1(x[0], y[7], C)
	C, t[8] = madd1(x[0], y[8], C)
	C, t[9] = madd1(x[0], y[9], C)
	C, t[10] = madd1(x[0], y[10], C)

	t[11], D = bits.Add64(t[11], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	C, t[6] = madd2(m, mod[7], t[7], C)
	C, t[7] = madd2(m, mod[8], t[8], C)
	C, t[8] = madd2(m, mod[9], t[9], C)
	C, t[9] = madd2(m, mod[10], t[10], C)
	t[10], C = bits.Add64(t[11], C, 0)
	t[11], _ = bits.Add64(0, D, C)

	for j := 1; j < 11; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		C, t[3] = madd2(x[j], y[3], t[3], C)
		C, t[4] = madd2(x[j], y[4], t[4], C)
		C, t[5] = madd2(x[j], y[5], t[5], C)
		C, t[6] = madd2(x[j], y[6], t[6], C)
		C, t[7] = madd2(x[j], y[7], t[7], C)
		C, t[8] = madd2(x[j], y[8], t[8], C)
		C, t[9] = madd2(x[j], y[9], t[9], C)
		C, t[10] = madd2(x[j], y[10], t[10], C)
		t[11], D = bits.Add64(t[11], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		C, t[3] = madd2(m, mod[4], t[4], C)
		C, t[4] = madd2(m, mod[5], t[5], C)
		C, t[5] = madd2(m, mod[6], t[6], C)
		C, t[6] = madd2(m, mod[7], t[7], C)
		C, t[7] = madd2(m, mod[8], t[8], C)
		C, t[8] = madd2(m, mod[9], t[9], C)
		C, t[9] = madd2(m, mod[10], t[10], C)
		t[10], C = bits.Add64(t[11], C, 0)
		t[11], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)
	res[3], D = bits.Sub64(t[3], mod[3], D)
	res[4], D = bits.Sub64(t[4], mod[4], D)
	res[5], D = bits.Sub64(t[5], mod[5], D)
	res[6], D = bits.Sub64(t[6], mod[6], D)
	res[7], D = bits.Sub64(t[7], mod[7], D)
	res[8], D = bits.Sub64(t[8], mod[8], D)
	res[9], D = bits.Sub64(t[9], mod[9], D)
	res[10], D = bits.Sub64(t[10], mod[10], D)

	var src []uint64
	if D != 0 && t[11] == 0 {
		src = t[:11]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[10])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[9])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[8])
	binary.LittleEndian.PutUint64(outBytes[24:32], src[7])
	binary.LittleEndian.PutUint64(outBytes[32:40], src[6])
	binary.LittleEndian.PutUint64(outBytes[40:48], src[5])
	binary.LittleEndian.PutUint64(outBytes[48:56], src[4])
	binary.LittleEndian.PutUint64(outBytes[56:64], src[3])
	binary.LittleEndian.PutUint64(outBytes[64:72], src[2])
	binary.LittleEndian.PutUint64(outBytes[72:80], src[1])
	binary.LittleEndian.PutUint64(outBytes[80:88], src[0])
}

func MontMul768(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [13]uint64
	var D uint64
	var m, C uint64

	var (
		mod [12]uint64
		x   [12]uint64
		y   [12]uint64
		res [12]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[95]
	_ = yBytes[95]
	_ = outBytes[95]
	_ = modBytes[95]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	x[2] = binary.LittleEndian.Uint64(xBytes[16:24])
	x[3] = binary.LittleEndian.Uint64(xBytes[24:32])
	x[4] = binary.LittleEndian.Uint64(xBytes[32:40])
	x[5] = binary.LittleEndian.Uint64(xBytes[40:48])
	x[6] = binary.LittleEndian.Uint64(xBytes[48:56])
	x[7] = binary.LittleEndian.Uint64(xBytes[56:64])
	x[8] = binary.LittleEndian.Uint64(xBytes[64:72])
	x[9] = binary.LittleEndian.Uint64(xBytes[72:80])
	x[10] = binary.LittleEndian.Uint64(xBytes[80:88])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	y[2] = binary.LittleEndian.Uint64(yBytes[16:24])
	y[3] = binary.LittleEndian.Uint64(yBytes[24:32])
	y[4] = binary.LittleEndian.Uint64(yBytes[32:40])
	y[5] = binary.LittleEndian.Uint64(yBytes[40:48])
	y[6] = binary.LittleEndian.Uint64(yBytes[48:56])
	y[7] = binary.LittleEndian.Uint64(yBytes[56:64])
	y[8] = binary.LittleEndian.Uint64(yBytes[64:72])
	y[9] = binary.LittleEndian.Uint64(yBytes[72:80])
	y[10] = binary.LittleEndian.Uint64(yBytes[80:88])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])
	mod[2] = binary.LittleEndian.Uint64(modBytes[16:24])
	mod[3] = binary.LittleEndian.Uint64(modBytes[24:32])
	mod[4] = binary.LittleEndian.Uint64(modBytes[32:40])
	mod[5] = binary.LittleEndian.Uint64(modBytes[40:48])
	mod[6] = binary.LittleEndian.Uint64(modBytes[48:56])
	mod[7] = binary.LittleEndian.Uint64(modBytes[56:64])
	mod[8] = binary.LittleEndian.Uint64(modBytes[64:72])
	mod[9] = binary.LittleEndian.Uint64(modBytes[72:80])
	mod[10] = binary.LittleEndian.Uint64(modBytes[80:88])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)
	C, t[5] = madd1(x[0], y[5], C)
	C, t[6] = madd1(x[0], y[6], C)
	C, t[7] = madd1(x[0], y[7], C)
	C, t[8] = madd1(x[0], y[8], C)
	C, t[9] = madd1(x[0], y[9], C)
	C, t[10] = madd1(x[0], y[10], C)
	C, t[11] = madd1(x[0], y[11], C)

	t[12], D = bits.Add64(t[12], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	C, t[6] = madd2(m, mod[7], t[7], C)
	C, t[7] = madd2(m, mod[8], t[8], C)
	C, t[8] = madd2(m, mod[9], t[9], C)
	C, t[9] = madd2(m, mod[10], t[10], C)
	C, t[10] = madd2(m, mod[11], t[11], C)
	t[11], C = bits.Add64(t[12], C, 0)
	t[12], _ = bits.Add64(0, D, C)

	for j := 1; j < 12; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		C, t[3] = madd2(x[j], y[3], t[3], C)
		C, t[4] = madd2(x[j], y[4], t[4], C)
		C, t[5] = madd2(x[j], y[5], t[5], C)
		C, t[6] = madd2(x[j], y[6], t[6], C)
		C, t[7] = madd2(x[j], y[7], t[7], C)
		C, t[8] = madd2(x[j], y[8], t[8], C)
		C, t[9] = madd2(x[j], y[9], t[9], C)
		C, t[10] = madd2(x[j], y[10], t[10], C)
		C, t[11] = madd2(x[j], y[11], t[11], C)
		t[12], D = bits.Add64(t[12], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		C, t[3] = madd2(m, mod[4], t[4], C)
		C, t[4] = madd2(m, mod[5], t[5], C)
		C, t[5] = madd2(m, mod[6], t[6], C)
		C, t[6] = madd2(m, mod[7], t[7], C)
		C, t[7] = madd2(m, mod[8], t[8], C)
		C, t[8] = madd2(m, mod[9], t[9], C)
		C, t[9] = madd2(m, mod[10], t[10], C)
		C, t[10] = madd2(m, mod[11], t[11], C)
		t[11], C = bits.Add64(t[12], C, 0)
		t[12], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)
	res[3], D = bits.Sub64(t[3], mod[3], D)
	res[4], D = bits.Sub64(t[4], mod[4], D)
	res[5], D = bits.Sub64(t[5], mod[5], D)
	res[6], D = bits.Sub64(t[6], mod[6], D)
	res[7], D = bits.Sub64(t[7], mod[7], D)
	res[8], D = bits.Sub64(t[8], mod[8], D)
	res[9], D = bits.Sub64(t[9], mod[9], D)
	res[10], D = bits.Sub64(t[10], mod[10], D)
	res[11], D = bits.Sub64(t[11], mod[11], D)

	var src []uint64
	if D != 0 && t[12] == 0 {
		src = t[:12]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[11])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[10])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[9])
	binary.LittleEndian.PutUint64(outBytes[24:32], src[8])
	binary.LittleEndian.PutUint64(outBytes[32:40], src[7])
	binary.LittleEndian.PutUint64(outBytes[40:48], src[6])
	binary.LittleEndian.PutUint64(outBytes[48:56], src[5])
	binary.LittleEndian.PutUint64(outBytes[56:64], src[4])
	binary.LittleEndian.PutUint64(outBytes[64:72], src[3])
	binary.LittleEndian.PutUint64(outBytes[72:80], src[2])
	binary.LittleEndian.PutUint64(outBytes[80:88], src[1])
	binary.LittleEndian.PutUint64(outBytes[88:96], src[0])
}

func MontMul832(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [14]uint64
	var D uint64
	var m, C uint64

	var (
		mod [13]uint64
		x   [13]uint64
		y   [13]uint64
		res [13]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[103]
	_ = yBytes[103]
	_ = outBytes[103]
	_ = modBytes[103]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	x[2] = binary.LittleEndian.Uint64(xBytes[16:24])
	x[3] = binary.LittleEndian.Uint64(xBytes[24:32])
	x[4] = binary.LittleEndian.Uint64(xBytes[32:40])
	x[5] = binary.LittleEndian.Uint64(xBytes[40:48])
	x[6] = binary.LittleEndian.Uint64(xBytes[48:56])
	x[7] = binary.LittleEndian.Uint64(xBytes[56:64])
	x[8] = binary.LittleEndian.Uint64(xBytes[64:72])
	x[9] = binary.LittleEndian.Uint64(xBytes[72:80])
	x[10] = binary.LittleEndian.Uint64(xBytes[80:88])
	x[11] = binary.LittleEndian.Uint64(xBytes[88:96])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	y[2] = binary.LittleEndian.Uint64(yBytes[16:24])
	y[3] = binary.LittleEndian.Uint64(yBytes[24:32])
	y[4] = binary.LittleEndian.Uint64(yBytes[32:40])
	y[5] = binary.LittleEndian.Uint64(yBytes[40:48])
	y[6] = binary.LittleEndian.Uint64(yBytes[48:56])
	y[7] = binary.LittleEndian.Uint64(yBytes[56:64])
	y[8] = binary.LittleEndian.Uint64(yBytes[64:72])
	y[9] = binary.LittleEndian.Uint64(yBytes[72:80])
	y[10] = binary.LittleEndian.Uint64(yBytes[80:88])
	y[11] = binary.LittleEndian.Uint64(yBytes[88:96])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])
	mod[2] = binary.LittleEndian.Uint64(modBytes[16:24])
	mod[3] = binary.LittleEndian.Uint64(modBytes[24:32])
	mod[4] = binary.LittleEndian.Uint64(modBytes[32:40])
	mod[5] = binary.LittleEndian.Uint64(modBytes[40:48])
	mod[6] = binary.LittleEndian.Uint64(modBytes[48:56])
	mod[7] = binary.LittleEndian.Uint64(modBytes[56:64])
	mod[8] = binary.LittleEndian.Uint64(modBytes[64:72])
	mod[9] = binary.LittleEndian.Uint64(modBytes[72:80])
	mod[10] = binary.LittleEndian.Uint64(modBytes[80:88])
	mod[11] = binary.LittleEndian.Uint64(modBytes[88:96])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)
	C, t[5] = madd1(x[0], y[5], C)
	C, t[6] = madd1(x[0], y[6], C)
	C, t[7] = madd1(x[0], y[7], C)
	C, t[8] = madd1(x[0], y[8], C)
	C, t[9] = madd1(x[0], y[9], C)
	C, t[10] = madd1(x[0], y[10], C)
	C, t[11] = madd1(x[0], y[11], C)
	C, t[12] = madd1(x[0], y[12], C)

	t[13], D = bits.Add64(t[13], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	C, t[6] = madd2(m, mod[7], t[7], C)
	C, t[7] = madd2(m, mod[8], t[8], C)
	C, t[8] = madd2(m, mod[9], t[9], C)
	C, t[9] = madd2(m, mod[10], t[10], C)
	C, t[10] = madd2(m, mod[11], t[11], C)
	C, t[11] = madd2(m, mod[12], t[12], C)
	t[12], C = bits.Add64(t[13], C, 0)
	t[13], _ = bits.Add64(0, D, C)

	for j := 1; j < 13; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		C, t[3] = madd2(x[j], y[3], t[3], C)
		C, t[4] = madd2(x[j], y[4], t[4], C)
		C, t[5] = madd2(x[j], y[5], t[5], C)
		C, t[6] = madd2(x[j], y[6], t[6], C)
		C, t[7] = madd2(x[j], y[7], t[7], C)
		C, t[8] = madd2(x[j], y[8], t[8], C)
		C, t[9] = madd2(x[j], y[9], t[9], C)
		C, t[10] = madd2(x[j], y[10], t[10], C)
		C, t[11] = madd2(x[j], y[11], t[11], C)
		C, t[12] = madd2(x[j], y[12], t[12], C)
		t[13], D = bits.Add64(t[13], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		C, t[3] = madd2(m, mod[4], t[4], C)
		C, t[4] = madd2(m, mod[5], t[5], C)
		C, t[5] = madd2(m, mod[6], t[6], C)
		C, t[6] = madd2(m, mod[7], t[7], C)
		C, t[7] = madd2(m, mod[8], t[8], C)
		C, t[8] = madd2(m, mod[9], t[9], C)
		C, t[9] = madd2(m, mod[10], t[10], C)
		C, t[10] = madd2(m, mod[11], t[11], C)
		C, t[11] = madd2(m, mod[12], t[12], C)
		t[12], C = bits.Add64(t[13], C, 0)
		t[13], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)
	res[3], D = bits.Sub64(t[3], mod[3], D)
	res[4], D = bits.Sub64(t[4], mod[4], D)
	res[5], D = bits.Sub64(t[5], mod[5], D)
	res[6], D = bits.Sub64(t[6], mod[6], D)
	res[7], D = bits.Sub64(t[7], mod[7], D)
	res[8], D = bits.Sub64(t[8], mod[8], D)
	res[9], D = bits.Sub64(t[9], mod[9], D)
	res[10], D = bits.Sub64(t[10], mod[10], D)
	res[11], D = bits.Sub64(t[11], mod[11], D)
	res[12], D = bits.Sub64(t[12], mod[12], D)

	var src []uint64
	if D != 0 && t[13] == 0 {
		src = t[:13]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[12])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[11])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[10])
	binary.LittleEndian.PutUint64(outBytes[24:32], src[9])
	binary.LittleEndian.PutUint64(outBytes[32:40], src[8])
	binary.LittleEndian.PutUint64(outBytes[40:48], src[7])
	binary.LittleEndian.PutUint64(outBytes[48:56], src[6])
	binary.LittleEndian.PutUint64(outBytes[56:64], src[5])
	binary.LittleEndian.PutUint64(outBytes[64:72], src[4])
	binary.LittleEndian.PutUint64(outBytes[72:80], src[3])
	binary.LittleEndian.PutUint64(outBytes[80:88], src[2])
	binary.LittleEndian.PutUint64(outBytes[88:96], src[1])
	binary.LittleEndian.PutUint64(outBytes[96:104], src[0])
}

func MontMul896(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [15]uint64
	var D uint64
	var m, C uint64

	var (
		mod [14]uint64
		x   [14]uint64
		y   [14]uint64
		res [14]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[111]
	_ = yBytes[111]
	_ = outBytes[111]
	_ = modBytes[111]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	x[2] = binary.LittleEndian.Uint64(xBytes[16:24])
	x[3] = binary.LittleEndian.Uint64(xBytes[24:32])
	x[4] = binary.LittleEndian.Uint64(xBytes[32:40])
	x[5] = binary.LittleEndian.Uint64(xBytes[40:48])
	x[6] = binary.LittleEndian.Uint64(xBytes[48:56])
	x[7] = binary.LittleEndian.Uint64(xBytes[56:64])
	x[8] = binary.LittleEndian.Uint64(xBytes[64:72])
	x[9] = binary.LittleEndian.Uint64(xBytes[72:80])
	x[10] = binary.LittleEndian.Uint64(xBytes[80:88])
	x[11] = binary.LittleEndian.Uint64(xBytes[88:96])
	x[12] = binary.LittleEndian.Uint64(xBytes[96:104])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	y[2] = binary.LittleEndian.Uint64(yBytes[16:24])
	y[3] = binary.LittleEndian.Uint64(yBytes[24:32])
	y[4] = binary.LittleEndian.Uint64(yBytes[32:40])
	y[5] = binary.LittleEndian.Uint64(yBytes[40:48])
	y[6] = binary.LittleEndian.Uint64(yBytes[48:56])
	y[7] = binary.LittleEndian.Uint64(yBytes[56:64])
	y[8] = binary.LittleEndian.Uint64(yBytes[64:72])
	y[9] = binary.LittleEndian.Uint64(yBytes[72:80])
	y[10] = binary.LittleEndian.Uint64(yBytes[80:88])
	y[11] = binary.LittleEndian.Uint64(yBytes[88:96])
	y[12] = binary.LittleEndian.Uint64(yBytes[96:104])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])
	mod[2] = binary.LittleEndian.Uint64(modBytes[16:24])
	mod[3] = binary.LittleEndian.Uint64(modBytes[24:32])
	mod[4] = binary.LittleEndian.Uint64(modBytes[32:40])
	mod[5] = binary.LittleEndian.Uint64(modBytes[40:48])
	mod[6] = binary.LittleEndian.Uint64(modBytes[48:56])
	mod[7] = binary.LittleEndian.Uint64(modBytes[56:64])
	mod[8] = binary.LittleEndian.Uint64(modBytes[64:72])
	mod[9] = binary.LittleEndian.Uint64(modBytes[72:80])
	mod[10] = binary.LittleEndian.Uint64(modBytes[80:88])
	mod[11] = binary.LittleEndian.Uint64(modBytes[88:96])
	mod[12] = binary.LittleEndian.Uint64(modBytes[96:104])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)
	C, t[5] = madd1(x[0], y[5], C)
	C, t[6] = madd1(x[0], y[6], C)
	C, t[7] = madd1(x[0], y[7], C)
	C, t[8] = madd1(x[0], y[8], C)
	C, t[9] = madd1(x[0], y[9], C)
	C, t[10] = madd1(x[0], y[10], C)
	C, t[11] = madd1(x[0], y[11], C)
	C, t[12] = madd1(x[0], y[12], C)
	C, t[13] = madd1(x[0], y[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	C, t[6] = madd2(m, mod[7], t[7], C)
	C, t[7] = madd2(m, mod[8], t[8], C)
	C, t[8] = madd2(m, mod[9], t[9], C)
	C, t[9] = madd2(m, mod[10], t[10], C)
	C, t[10] = madd2(m, mod[11], t[11], C)
	C, t[11] = madd2(m, mod[12], t[12], C)
	C, t[12] = madd2(m, mod[13], t[13], C)
	t[13], C = bits.Add64(t[14], C, 0)
	t[14], _ = bits.Add64(0, D, C)

	for j := 1; j < 14; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		C, t[3] = madd2(x[j], y[3], t[3], C)
		C, t[4] = madd2(x[j], y[4], t[4], C)
		C, t[5] = madd2(x[j], y[5], t[5], C)
		C, t[6] = madd2(x[j], y[6], t[6], C)
		C, t[7] = madd2(x[j], y[7], t[7], C)
		C, t[8] = madd2(x[j], y[8], t[8], C)
		C, t[9] = madd2(x[j], y[9], t[9], C)
		C, t[10] = madd2(x[j], y[10], t[10], C)
		C, t[11] = madd2(x[j], y[11], t[11], C)
		C, t[12] = madd2(x[j], y[12], t[12], C)
		C, t[13] = madd2(x[j], y[13], t[13], C)
		t[14], D = bits.Add64(t[14], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		C, t[3] = madd2(m, mod[4], t[4], C)
		C, t[4] = madd2(m, mod[5], t[5], C)
		C, t[5] = madd2(m, mod[6], t[6], C)
		C, t[6] = madd2(m, mod[7], t[7], C)
		C, t[7] = madd2(m, mod[8], t[8], C)
		C, t[8] = madd2(m, mod[9], t[9], C)
		C, t[9] = madd2(m, mod[10], t[10], C)
		C, t[10] = madd2(m, mod[11], t[11], C)
		C, t[11] = madd2(m, mod[12], t[12], C)
		C, t[12] = madd2(m, mod[13], t[13], C)
		t[13], C = bits.Add64(t[14], C, 0)
		t[14], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)
	res[3], D = bits.Sub64(t[3], mod[3], D)
	res[4], D = bits.Sub64(t[4], mod[4], D)
	res[5], D = bits.Sub64(t[5], mod[5], D)
	res[6], D = bits.Sub64(t[6], mod[6], D)
	res[7], D = bits.Sub64(t[7], mod[7], D)
	res[8], D = bits.Sub64(t[8], mod[8], D)
	res[9], D = bits.Sub64(t[9], mod[9], D)
	res[10], D = bits.Sub64(t[10], mod[10], D)
	res[11], D = bits.Sub64(t[11], mod[11], D)
	res[12], D = bits.Sub64(t[12], mod[12], D)
	res[13], D = bits.Sub64(t[13], mod[13], D)

	var src []uint64
	if D != 0 && t[14] == 0 {
		src = t[:14]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[13])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[12])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[11])
	binary.LittleEndian.PutUint64(outBytes[24:32], src[10])
	binary.LittleEndian.PutUint64(outBytes[32:40], src[9])
	binary.LittleEndian.PutUint64(outBytes[40:48], src[8])
	binary.LittleEndian.PutUint64(outBytes[48:56], src[7])
	binary.LittleEndian.PutUint64(outBytes[56:64], src[6])
	binary.LittleEndian.PutUint64(outBytes[64:72], src[5])
	binary.LittleEndian.PutUint64(outBytes[72:80], src[4])
	binary.LittleEndian.PutUint64(outBytes[80:88], src[3])
	binary.LittleEndian.PutUint64(outBytes[88:96], src[2])
	binary.LittleEndian.PutUint64(outBytes[96:104], src[1])
	binary.LittleEndian.PutUint64(outBytes[104:112], src[0])
}

func MontMul960(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [16]uint64
	var D uint64
	var m, C uint64

	var (
		mod [15]uint64
		x   [15]uint64
		y   [15]uint64
		res [15]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[119]
	_ = yBytes[119]
	_ = outBytes[119]
	_ = modBytes[119]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	x[2] = binary.LittleEndian.Uint64(xBytes[16:24])
	x[3] = binary.LittleEndian.Uint64(xBytes[24:32])
	x[4] = binary.LittleEndian.Uint64(xBytes[32:40])
	x[5] = binary.LittleEndian.Uint64(xBytes[40:48])
	x[6] = binary.LittleEndian.Uint64(xBytes[48:56])
	x[7] = binary.LittleEndian.Uint64(xBytes[56:64])
	x[8] = binary.LittleEndian.Uint64(xBytes[64:72])
	x[9] = binary.LittleEndian.Uint64(xBytes[72:80])
	x[10] = binary.LittleEndian.Uint64(xBytes[80:88])
	x[11] = binary.LittleEndian.Uint64(xBytes[88:96])
	x[12] = binary.LittleEndian.Uint64(xBytes[96:104])
	x[13] = binary.LittleEndian.Uint64(xBytes[104:112])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	y[2] = binary.LittleEndian.Uint64(yBytes[16:24])
	y[3] = binary.LittleEndian.Uint64(yBytes[24:32])
	y[4] = binary.LittleEndian.Uint64(yBytes[32:40])
	y[5] = binary.LittleEndian.Uint64(yBytes[40:48])
	y[6] = binary.LittleEndian.Uint64(yBytes[48:56])
	y[7] = binary.LittleEndian.Uint64(yBytes[56:64])
	y[8] = binary.LittleEndian.Uint64(yBytes[64:72])
	y[9] = binary.LittleEndian.Uint64(yBytes[72:80])
	y[10] = binary.LittleEndian.Uint64(yBytes[80:88])
	y[11] = binary.LittleEndian.Uint64(yBytes[88:96])
	y[12] = binary.LittleEndian.Uint64(yBytes[96:104])
	y[13] = binary.LittleEndian.Uint64(yBytes[104:112])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])
	mod[2] = binary.LittleEndian.Uint64(modBytes[16:24])
	mod[3] = binary.LittleEndian.Uint64(modBytes[24:32])
	mod[4] = binary.LittleEndian.Uint64(modBytes[32:40])
	mod[5] = binary.LittleEndian.Uint64(modBytes[40:48])
	mod[6] = binary.LittleEndian.Uint64(modBytes[48:56])
	mod[7] = binary.LittleEndian.Uint64(modBytes[56:64])
	mod[8] = binary.LittleEndian.Uint64(modBytes[64:72])
	mod[9] = binary.LittleEndian.Uint64(modBytes[72:80])
	mod[10] = binary.LittleEndian.Uint64(modBytes[80:88])
	mod[11] = binary.LittleEndian.Uint64(modBytes[88:96])
	mod[12] = binary.LittleEndian.Uint64(modBytes[96:104])
	mod[13] = binary.LittleEndian.Uint64(modBytes[104:112])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)
	C, t[5] = madd1(x[0], y[5], C)
	C, t[6] = madd1(x[0], y[6], C)
	C, t[7] = madd1(x[0], y[7], C)
	C, t[8] = madd1(x[0], y[8], C)
	C, t[9] = madd1(x[0], y[9], C)
	C, t[10] = madd1(x[0], y[10], C)
	C, t[11] = madd1(x[0], y[11], C)
	C, t[12] = madd1(x[0], y[12], C)
	C, t[13] = madd1(x[0], y[13], C)
	C, t[14] = madd1(x[0], y[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	C, t[6] = madd2(m, mod[7], t[7], C)
	C, t[7] = madd2(m, mod[8], t[8], C)
	C, t[8] = madd2(m, mod[9], t[9], C)
	C, t[9] = madd2(m, mod[10], t[10], C)
	C, t[10] = madd2(m, mod[11], t[11], C)
	C, t[11] = madd2(m, mod[12], t[12], C)
	C, t[12] = madd2(m, mod[13], t[13], C)
	C, t[13] = madd2(m, mod[14], t[14], C)
	t[14], C = bits.Add64(t[15], C, 0)
	t[15], _ = bits.Add64(0, D, C)

	for j := 1; j < 15; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		C, t[3] = madd2(x[j], y[3], t[3], C)
		C, t[4] = madd2(x[j], y[4], t[4], C)
		C, t[5] = madd2(x[j], y[5], t[5], C)
		C, t[6] = madd2(x[j], y[6], t[6], C)
		C, t[7] = madd2(x[j], y[7], t[7], C)
		C, t[8] = madd2(x[j], y[8], t[8], C)
		C, t[9] = madd2(x[j], y[9], t[9], C)
		C, t[10] = madd2(x[j], y[10], t[10], C)
		C, t[11] = madd2(x[j], y[11], t[11], C)
		C, t[12] = madd2(x[j], y[12], t[12], C)
		C, t[13] = madd2(x[j], y[13], t[13], C)
		C, t[14] = madd2(x[j], y[14], t[14], C)
		t[15], D = bits.Add64(t[15], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		C, t[3] = madd2(m, mod[4], t[4], C)
		C, t[4] = madd2(m, mod[5], t[5], C)
		C, t[5] = madd2(m, mod[6], t[6], C)
		C, t[6] = madd2(m, mod[7], t[7], C)
		C, t[7] = madd2(m, mod[8], t[8], C)
		C, t[8] = madd2(m, mod[9], t[9], C)
		C, t[9] = madd2(m, mod[10], t[10], C)
		C, t[10] = madd2(m, mod[11], t[11], C)
		C, t[11] = madd2(m, mod[12], t[12], C)
		C, t[12] = madd2(m, mod[13], t[13], C)
		C, t[13] = madd2(m, mod[14], t[14], C)
		t[14], C = bits.Add64(t[15], C, 0)
		t[15], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)
	res[3], D = bits.Sub64(t[3], mod[3], D)
	res[4], D = bits.Sub64(t[4], mod[4], D)
	res[5], D = bits.Sub64(t[5], mod[5], D)
	res[6], D = bits.Sub64(t[6], mod[6], D)
	res[7], D = bits.Sub64(t[7], mod[7], D)
	res[8], D = bits.Sub64(t[8], mod[8], D)
	res[9], D = bits.Sub64(t[9], mod[9], D)
	res[10], D = bits.Sub64(t[10], mod[10], D)
	res[11], D = bits.Sub64(t[11], mod[11], D)
	res[12], D = bits.Sub64(t[12], mod[12], D)
	res[13], D = bits.Sub64(t[13], mod[13], D)
	res[14], D = bits.Sub64(t[14], mod[14], D)

	var src []uint64
	if D != 0 && t[15] == 0 {
		src = t[:15]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[14])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[13])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[12])
	binary.LittleEndian.PutUint64(outBytes[24:32], src[11])
	binary.LittleEndian.PutUint64(outBytes[32:40], src[10])
	binary.LittleEndian.PutUint64(outBytes[40:48], src[9])
	binary.LittleEndian.PutUint64(outBytes[48:56], src[8])
	binary.LittleEndian.PutUint64(outBytes[56:64], src[7])
	binary.LittleEndian.PutUint64(outBytes[64:72], src[6])
	binary.LittleEndian.PutUint64(outBytes[72:80], src[5])
	binary.LittleEndian.PutUint64(outBytes[80:88], src[4])
	binary.LittleEndian.PutUint64(outBytes[88:96], src[3])
	binary.LittleEndian.PutUint64(outBytes[96:104], src[2])
	binary.LittleEndian.PutUint64(outBytes[104:112], src[1])
	binary.LittleEndian.PutUint64(outBytes[112:120], src[0])
}

func MontMul1024(modInv uint64, modBytes, outBytes, xBytes, yBytes []byte) {
	var t [17]uint64
	var D uint64
	var m, C uint64

	var (
		mod [16]uint64
		x   [16]uint64
		y   [16]uint64
		res [16]uint64
	)

	// signal to compiler to avoid subsequent bounds checks
	_ = xBytes[127]
	_ = yBytes[127]
	_ = outBytes[127]
	_ = modBytes[127]
	x[0] = binary.LittleEndian.Uint64(xBytes[0:8])
	x[1] = binary.LittleEndian.Uint64(xBytes[8:16])
	x[2] = binary.LittleEndian.Uint64(xBytes[16:24])
	x[3] = binary.LittleEndian.Uint64(xBytes[24:32])
	x[4] = binary.LittleEndian.Uint64(xBytes[32:40])
	x[5] = binary.LittleEndian.Uint64(xBytes[40:48])
	x[6] = binary.LittleEndian.Uint64(xBytes[48:56])
	x[7] = binary.LittleEndian.Uint64(xBytes[56:64])
	x[8] = binary.LittleEndian.Uint64(xBytes[64:72])
	x[9] = binary.LittleEndian.Uint64(xBytes[72:80])
	x[10] = binary.LittleEndian.Uint64(xBytes[80:88])
	x[11] = binary.LittleEndian.Uint64(xBytes[88:96])
	x[12] = binary.LittleEndian.Uint64(xBytes[96:104])
	x[13] = binary.LittleEndian.Uint64(xBytes[104:112])
	x[14] = binary.LittleEndian.Uint64(xBytes[112:120])
	y[0] = binary.LittleEndian.Uint64(yBytes[0:8])
	y[1] = binary.LittleEndian.Uint64(yBytes[8:16])
	y[2] = binary.LittleEndian.Uint64(yBytes[16:24])
	y[3] = binary.LittleEndian.Uint64(yBytes[24:32])
	y[4] = binary.LittleEndian.Uint64(yBytes[32:40])
	y[5] = binary.LittleEndian.Uint64(yBytes[40:48])
	y[6] = binary.LittleEndian.Uint64(yBytes[48:56])
	y[7] = binary.LittleEndian.Uint64(yBytes[56:64])
	y[8] = binary.LittleEndian.Uint64(yBytes[64:72])
	y[9] = binary.LittleEndian.Uint64(yBytes[72:80])
	y[10] = binary.LittleEndian.Uint64(yBytes[80:88])
	y[11] = binary.LittleEndian.Uint64(yBytes[88:96])
	y[12] = binary.LittleEndian.Uint64(yBytes[96:104])
	y[13] = binary.LittleEndian.Uint64(yBytes[104:112])
	y[14] = binary.LittleEndian.Uint64(yBytes[112:120])
	mod[0] = binary.LittleEndian.Uint64(modBytes[0:8])
	mod[1] = binary.LittleEndian.Uint64(modBytes[8:16])
	mod[2] = binary.LittleEndian.Uint64(modBytes[16:24])
	mod[3] = binary.LittleEndian.Uint64(modBytes[24:32])
	mod[4] = binary.LittleEndian.Uint64(modBytes[32:40])
	mod[5] = binary.LittleEndian.Uint64(modBytes[40:48])
	mod[6] = binary.LittleEndian.Uint64(modBytes[48:56])
	mod[7] = binary.LittleEndian.Uint64(modBytes[56:64])
	mod[8] = binary.LittleEndian.Uint64(modBytes[64:72])
	mod[9] = binary.LittleEndian.Uint64(modBytes[72:80])
	mod[10] = binary.LittleEndian.Uint64(modBytes[80:88])
	mod[11] = binary.LittleEndian.Uint64(modBytes[88:96])
	mod[12] = binary.LittleEndian.Uint64(modBytes[96:104])
	mod[13] = binary.LittleEndian.Uint64(modBytes[104:112])
	mod[14] = binary.LittleEndian.Uint64(modBytes[112:120])

	// 1st outer loop:
	// 1st inner loop: t <- x[0] * y
	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)
	C, t[5] = madd1(x[0], y[5], C)
	C, t[6] = madd1(x[0], y[6], C)
	C, t[7] = madd1(x[0], y[7], C)
	C, t[8] = madd1(x[0], y[8], C)
	C, t[9] = madd1(x[0], y[9], C)
	C, t[10] = madd1(x[0], y[10], C)
	C, t[11] = madd1(x[0], y[11], C)
	C, t[12] = madd1(x[0], y[12], C)
	C, t[13] = madd1(x[0], y[13], C)
	C, t[14] = madd1(x[0], y[14], C)
	C, t[15] = madd1(x[0], y[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * modInv

	// -----------------------------------
	// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	C, t[6] = madd2(m, mod[7], t[7], C)
	C, t[7] = madd2(m, mod[8], t[8], C)
	C, t[8] = madd2(m, mod[9], t[9], C)
	C, t[9] = madd2(m, mod[10], t[10], C)
	C, t[10] = madd2(m, mod[11], t[11], C)
	C, t[11] = madd2(m, mod[12], t[12], C)
	C, t[12] = madd2(m, mod[13], t[13], C)
	C, t[13] = madd2(m, mod[14], t[14], C)
	C, t[14] = madd2(m, mod[15], t[15], C)
	t[15], C = bits.Add64(t[16], C, 0)
	t[16], _ = bits.Add64(0, D, C)

	for j := 1; j < 16; j++ {
		//  first inner loop (second iteration)
		C, t[0] = madd1(x[j], y[0], t[0])
		C, t[1] = madd2(x[j], y[1], t[1], C)
		C, t[2] = madd2(x[j], y[2], t[2], C)
		C, t[3] = madd2(x[j], y[3], t[3], C)
		C, t[4] = madd2(x[j], y[4], t[4], C)
		C, t[5] = madd2(x[j], y[5], t[5], C)
		C, t[6] = madd2(x[j], y[6], t[6], C)
		C, t[7] = madd2(x[j], y[7], t[7], C)
		C, t[8] = madd2(x[j], y[8], t[8], C)
		C, t[9] = madd2(x[j], y[9], t[9], C)
		C, t[10] = madd2(x[j], y[10], t[10], C)
		C, t[11] = madd2(x[j], y[11], t[11], C)
		C, t[12] = madd2(x[j], y[12], t[12], C)
		C, t[13] = madd2(x[j], y[13], t[13], C)
		C, t[14] = madd2(x[j], y[14], t[14], C)
		C, t[15] = madd2(x[j], y[15], t[15], C)
		t[16], D = bits.Add64(t[16], C, 0)
		// m = t[0]n'[0] mod W
		m = t[0] * modInv

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		C, t[3] = madd2(m, mod[4], t[4], C)
		C, t[4] = madd2(m, mod[5], t[5], C)
		C, t[5] = madd2(m, mod[6], t[6], C)
		C, t[6] = madd2(m, mod[7], t[7], C)
		C, t[7] = madd2(m, mod[8], t[8], C)
		C, t[8] = madd2(m, mod[9], t[9], C)
		C, t[9] = madd2(m, mod[10], t[10], C)
		C, t[10] = madd2(m, mod[11], t[11], C)
		C, t[11] = madd2(m, mod[12], t[12], C)
		C, t[12] = madd2(m, mod[13], t[13], C)
		C, t[13] = madd2(m, mod[14], t[14], C)
		C, t[14] = madd2(m, mod[15], t[15], C)
		t[15], C = bits.Add64(t[16], C, 0)
		t[16], _ = bits.Add64(0, D, C)
	}
	res[0], D = bits.Sub64(t[0], mod[0], 0)
	res[1], D = bits.Sub64(t[1], mod[1], D)
	res[2], D = bits.Sub64(t[2], mod[2], D)
	res[3], D = bits.Sub64(t[3], mod[3], D)
	res[4], D = bits.Sub64(t[4], mod[4], D)
	res[5], D = bits.Sub64(t[5], mod[5], D)
	res[6], D = bits.Sub64(t[6], mod[6], D)
	res[7], D = bits.Sub64(t[7], mod[7], D)
	res[8], D = bits.Sub64(t[8], mod[8], D)
	res[9], D = bits.Sub64(t[9], mod[9], D)
	res[10], D = bits.Sub64(t[10], mod[10], D)
	res[11], D = bits.Sub64(t[11], mod[11], D)
	res[12], D = bits.Sub64(t[12], mod[12], D)
	res[13], D = bits.Sub64(t[13], mod[13], D)
	res[14], D = bits.Sub64(t[14], mod[14], D)
	res[15], D = bits.Sub64(t[15], mod[15], D)

	var src []uint64
	if D != 0 && t[16] == 0 {
		src = t[:16]
	} else {
		src = res[:]
	}
	binary.LittleEndian.PutUint64(outBytes[0:8], src[15])
	binary.LittleEndian.PutUint64(outBytes[8:16], src[14])
	binary.LittleEndian.PutUint64(outBytes[16:24], src[13])
	binary.LittleEndian.PutUint64(outBytes[24:32], src[12])
	binary.LittleEndian.PutUint64(outBytes[32:40], src[11])
	binary.LittleEndian.PutUint64(outBytes[40:48], src[10])
	binary.LittleEndian.PutUint64(outBytes[48:56], src[9])
	binary.LittleEndian.PutUint64(outBytes[56:64], src[8])
	binary.LittleEndian.PutUint64(outBytes[64:72], src[7])
	binary.LittleEndian.PutUint64(outBytes[72:80], src[6])
	binary.LittleEndian.PutUint64(outBytes[80:88], src[5])
	binary.LittleEndian.PutUint64(outBytes[88:96], src[4])
	binary.LittleEndian.PutUint64(outBytes[96:104], src[3])
	binary.LittleEndian.PutUint64(outBytes[104:112], src[2])
	binary.LittleEndian.PutUint64(outBytes[112:120], src[1])
	binary.LittleEndian.PutUint64(outBytes[120:128], src[0])
}
