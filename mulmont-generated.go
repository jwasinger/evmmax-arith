package evmmax_arith

import (
	"math/bits"
)

var mulmodPreset = []mulFunc{
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
}

var addmodPreset = []addOrSubFunc{
	AddMod64,
	AddMod128,
	AddMod192,
	AddMod256,
	AddMod320,
	AddMod384,
	AddMod448,
	AddMod512,
	AddMod576,
	AddMod640,
	AddMod704,
	AddMod768,
}

var submodPreset = []addOrSubFunc{
	SubMod64,
	SubMod128,
	SubMod192,
	SubMod256,
	SubMod320,
	SubMod384,
	SubMod448,
	SubMod512,
	SubMod576,
	SubMod640,
	SubMod704,
	SubMod768,
}

func MontMul64(out, x, y, mod []uint64, modInv uint64) {
	var t [2]uint64
	var D uint64
	var m, C uint64

	var res [1]uint64

	// signal to compiler to avoid subsequent bounds checks
	_ = x[0]
	_ = y[0]
	_ = out[0]
	_ = mod[0]

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

	copy(out[:], src)
}

func MontMul128(out, x, y, mod []uint64, modInv uint64) {
	var t [3]uint64
	var D uint64
	var m, C uint64

	var res [2]uint64

	// signal to compiler to avoid subsequent bounds checks
	_ = x[1]
	_ = y[1]
	_ = out[1]
	_ = mod[1]

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

	copy(out[:], src)
}

func MontMul192(out, x, y, mod []uint64, modInv uint64) {
	var t [4]uint64
	var D uint64
	var m, C uint64

	var res [3]uint64

	// signal to compiler to avoid subsequent bounds checks
	_ = x[2]
	_ = y[2]
	_ = out[2]
	_ = mod[2]

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

	copy(out[:], src)
}

func MontMul256(out, x, y, mod []uint64, modInv uint64) {
	var t [5]uint64
	var D uint64
	var m, C uint64

	var res [4]uint64

	// signal to compiler to avoid subsequent bounds checks
	_ = x[3]
	_ = y[3]
	_ = out[3]
	_ = mod[3]

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

	copy(out[:], src)
}

func MontMul320(out, x, y, mod []uint64, modInv uint64) {
	var t [6]uint64
	var D uint64
	var m, C uint64

	var res [5]uint64

	// signal to compiler to avoid subsequent bounds checks
	_ = x[4]
	_ = y[4]
	_ = out[4]
	_ = mod[4]

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

	copy(out[:], src)
}

func MontMul384(out, x, y, mod []uint64, modInv uint64) {
	var t [7]uint64
	var D uint64
	var m, C uint64

	var res [6]uint64

	// signal to compiler to avoid subsequent bounds checks
	_ = x[5]
	_ = y[5]
	_ = out[5]
	_ = mod[5]

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

	copy(out[:], src)
}

func MontMul448(out, x, y, mod []uint64, modInv uint64) {
	var t [8]uint64
	var D uint64
	var m, C uint64

	var res [7]uint64

	// signal to compiler to avoid subsequent bounds checks
	_ = x[6]
	_ = y[6]
	_ = out[6]
	_ = mod[6]

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

	copy(out[:], src)
}

func MontMul512(out, x, y, mod []uint64, modInv uint64) {
	var t [9]uint64
	var D uint64
	var m, C uint64

	var res [8]uint64

	// signal to compiler to avoid subsequent bounds checks
	_ = x[7]
	_ = y[7]
	_ = out[7]
	_ = mod[7]

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

	copy(out[:], src)
}

func MontMul576(out, x, y, mod []uint64, modInv uint64) {
	var t [10]uint64
	var D uint64
	var m, C uint64

	var res [9]uint64

	// signal to compiler to avoid subsequent bounds checks
	_ = x[8]
	_ = y[8]
	_ = out[8]
	_ = mod[8]

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

	copy(out[:], src)
}

func MontMul640(out, x, y, mod []uint64, modInv uint64) {
	var t [11]uint64
	var D uint64
	var m, C uint64

	var res [10]uint64

	// signal to compiler to avoid subsequent bounds checks
	_ = x[9]
	_ = y[9]
	_ = out[9]
	_ = mod[9]

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

	copy(out[:], src)
}

func MontMul704(out, x, y, mod []uint64, modInv uint64) {
	var t [12]uint64
	var D uint64
	var m, C uint64

	var res [11]uint64

	// signal to compiler to avoid subsequent bounds checks
	_ = x[10]
	_ = y[10]
	_ = out[10]
	_ = mod[10]

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

	copy(out[:], src)
}

func MontMul768(out, x, y, mod []uint64, modInv uint64) {
	var t [13]uint64
	var D uint64
	var m, C uint64

	var res [12]uint64

	// signal to compiler to avoid subsequent bounds checks
	_ = x[11]
	_ = y[11]
	_ = out[11]
	_ = mod[11]

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

	copy(out[:], src)
}
