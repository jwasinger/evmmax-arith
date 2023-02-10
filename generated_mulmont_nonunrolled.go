package evmmax_arith

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/bits"
)

func MulMontNonUnrolled64(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [1]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[0] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[0] = binary.BigEndian.Uint64(y_bytes[0:8])

	mod := ctx.ModulusLimbs
	var t [2]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

	C, t[0] = bits.Mul64(x[0], y[0])

	t[1], D = bits.Add64(t[1], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		t[0], C = bits.Add64(t[1], C, 0)
		t[1], _ = bits.Add64(0, D, C)
	}
	z[0], D = bits.Sub64(t[0], mod[0], 0)

	var src []uint64
	if D != 0 && t[1] == 0 {
		src = t[:1]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[0])

	return nil
}

func MulMontNonUnrolled128(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [2]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[1] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[1] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[0] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[0] = binary.BigEndian.Uint64(y_bytes[8:16])

	mod := ctx.ModulusLimbs
	var t [3]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)

	t[2], D = bits.Add64(t[2], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		t[1], C = bits.Add64(t[2], C, 0)
		t[2], _ = bits.Add64(0, D, C)
	}
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)

	var src []uint64
	if D != 0 && t[2] == 0 {
		src = t[:2]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[1])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[0])

	return nil
}

func MulMontNonUnrolled192(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [3]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[2] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[2] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[1] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[1] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[0] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[0] = binary.BigEndian.Uint64(y_bytes[16:24])

	mod := ctx.ModulusLimbs
	var t [4]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)

	t[3], D = bits.Add64(t[3], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		t[2], C = bits.Add64(t[3], C, 0)
		t[3], _ = bits.Add64(0, D, C)
	}
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)

	var src []uint64
	if D != 0 && t[3] == 0 {
		src = t[:3]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[2])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[1])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[0])

	return nil
}

func MulMontNonUnrolled256(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [4]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[3] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[3] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[2] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[2] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[1] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[1] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[0] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[0] = binary.BigEndian.Uint64(y_bytes[24:32])

	mod := ctx.ModulusLimbs
	var t [5]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)

	t[4], D = bits.Add64(t[4], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

		// -----------------------------------
		// Second inner loop: reduce 1 limb at a time (B**1, B**2, ...)
		C = madd0(m, mod[0], t[0])
		C, t[0] = madd2(m, mod[1], t[1], C)
		C, t[1] = madd2(m, mod[2], t[2], C)
		C, t[2] = madd2(m, mod[3], t[3], C)
		t[3], C = bits.Add64(t[4], C, 0)
		t[4], _ = bits.Add64(0, D, C)
	}
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)

	var src []uint64
	if D != 0 && t[4] == 0 {
		src = t[:4]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[3])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[2])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[1])
	binary.BigEndian.PutUint64(z_bytes[24:32], src[0])

	return nil
}

func MulMontNonUnrolled320(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [5]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[4] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[4] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[3] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[3] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[2] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[2] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[1] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[1] = binary.BigEndian.Uint64(y_bytes[24:32])
	x[0] = binary.BigEndian.Uint64(x_bytes[32:40])
	y[0] = binary.BigEndian.Uint64(y_bytes[32:40])

	mod := ctx.ModulusLimbs
	var t [6]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], y[4], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)

	t[5], D = bits.Add64(t[5], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)

	var src []uint64
	if D != 0 && t[5] == 0 {
		src = t[:5]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[4])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[3])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[2])
	binary.BigEndian.PutUint64(z_bytes[24:32], src[1])
	binary.BigEndian.PutUint64(z_bytes[32:40], src[0])

	return nil
}

func MulMontNonUnrolled384(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [6]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[5] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[5] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[4] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[4] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[3] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[3] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[2] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[2] = binary.BigEndian.Uint64(y_bytes[24:32])
	x[1] = binary.BigEndian.Uint64(x_bytes[32:40])
	y[1] = binary.BigEndian.Uint64(y_bytes[32:40])
	x[0] = binary.BigEndian.Uint64(x_bytes[40:48])
	y[0] = binary.BigEndian.Uint64(y_bytes[40:48])

	mod := ctx.ModulusLimbs
	var t [7]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], y[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], y[5], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)
	C, t[5] = madd1(x[0], y[5], C)

	t[6], D = bits.Add64(t[6], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)

	var src []uint64
	if D != 0 && t[6] == 0 {
		src = t[:6]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[5])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[4])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[3])
	binary.BigEndian.PutUint64(z_bytes[24:32], src[2])
	binary.BigEndian.PutUint64(z_bytes[32:40], src[1])
	binary.BigEndian.PutUint64(z_bytes[40:48], src[0])

	return nil
}

func MulMontNonUnrolled448(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [7]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[6] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[6] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[5] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[5] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[4] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[4] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[3] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[3] = binary.BigEndian.Uint64(y_bytes[24:32])
	x[2] = binary.BigEndian.Uint64(x_bytes[32:40])
	y[2] = binary.BigEndian.Uint64(y_bytes[32:40])
	x[1] = binary.BigEndian.Uint64(x_bytes[40:48])
	y[1] = binary.BigEndian.Uint64(y_bytes[40:48])
	x[0] = binary.BigEndian.Uint64(x_bytes[48:56])
	y[0] = binary.BigEndian.Uint64(y_bytes[48:56])

	mod := ctx.ModulusLimbs
	var t [8]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC1 = bits.Sub64(mod[6], x[6], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], y[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], y[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], y[6], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)
	C, t[5] = madd1(x[0], y[5], C)
	C, t[6] = madd1(x[0], y[6], C)

	t[7], D = bits.Add64(t[7], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)
	z[6], D = bits.Sub64(t[6], mod[6], D)

	var src []uint64
	if D != 0 && t[7] == 0 {
		src = t[:7]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[6])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[5])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[4])
	binary.BigEndian.PutUint64(z_bytes[24:32], src[3])
	binary.BigEndian.PutUint64(z_bytes[32:40], src[2])
	binary.BigEndian.PutUint64(z_bytes[40:48], src[1])
	binary.BigEndian.PutUint64(z_bytes[48:56], src[0])

	return nil
}

func MulMontNonUnrolled512(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [8]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[7] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[7] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[6] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[6] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[5] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[5] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[4] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[4] = binary.BigEndian.Uint64(y_bytes[24:32])
	x[3] = binary.BigEndian.Uint64(x_bytes[32:40])
	y[3] = binary.BigEndian.Uint64(y_bytes[32:40])
	x[2] = binary.BigEndian.Uint64(x_bytes[40:48])
	y[2] = binary.BigEndian.Uint64(y_bytes[40:48])
	x[1] = binary.BigEndian.Uint64(x_bytes[48:56])
	y[1] = binary.BigEndian.Uint64(y_bytes[48:56])
	x[0] = binary.BigEndian.Uint64(x_bytes[56:64])
	y[0] = binary.BigEndian.Uint64(y_bytes[56:64])

	mod := ctx.ModulusLimbs
	var t [9]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC1 = bits.Sub64(mod[6], x[6], gteC1)
	_, gteC1 = bits.Sub64(mod[7], x[7], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], y[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], y[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], y[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], y[7], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

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
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)
	z[6], D = bits.Sub64(t[6], mod[6], D)
	z[7], D = bits.Sub64(t[7], mod[7], D)

	var src []uint64
	if D != 0 && t[8] == 0 {
		src = t[:8]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[7])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[6])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[5])
	binary.BigEndian.PutUint64(z_bytes[24:32], src[4])
	binary.BigEndian.PutUint64(z_bytes[32:40], src[3])
	binary.BigEndian.PutUint64(z_bytes[40:48], src[2])
	binary.BigEndian.PutUint64(z_bytes[48:56], src[1])
	binary.BigEndian.PutUint64(z_bytes[56:64], src[0])

	return nil
}

func MulMontNonUnrolled576(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [9]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[8] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[8] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[7] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[7] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[6] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[6] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[5] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[5] = binary.BigEndian.Uint64(y_bytes[24:32])
	x[4] = binary.BigEndian.Uint64(x_bytes[32:40])
	y[4] = binary.BigEndian.Uint64(y_bytes[32:40])
	x[3] = binary.BigEndian.Uint64(x_bytes[40:48])
	y[3] = binary.BigEndian.Uint64(y_bytes[40:48])
	x[2] = binary.BigEndian.Uint64(x_bytes[48:56])
	y[2] = binary.BigEndian.Uint64(y_bytes[48:56])
	x[1] = binary.BigEndian.Uint64(x_bytes[56:64])
	y[1] = binary.BigEndian.Uint64(y_bytes[56:64])
	x[0] = binary.BigEndian.Uint64(x_bytes[64:72])
	y[0] = binary.BigEndian.Uint64(y_bytes[64:72])

	mod := ctx.ModulusLimbs
	var t [10]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC1 = bits.Sub64(mod[6], x[6], gteC1)
	_, gteC1 = bits.Sub64(mod[7], x[7], gteC1)
	_, gteC1 = bits.Sub64(mod[8], x[8], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], y[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], y[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], y[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], y[7], gteC2)
	_, gteC2 = bits.Sub64(mod[8], y[8], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

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
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)
	z[6], D = bits.Sub64(t[6], mod[6], D)
	z[7], D = bits.Sub64(t[7], mod[7], D)
	z[8], D = bits.Sub64(t[8], mod[8], D)

	var src []uint64
	if D != 0 && t[9] == 0 {
		src = t[:9]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[8])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[7])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[6])
	binary.BigEndian.PutUint64(z_bytes[24:32], src[5])
	binary.BigEndian.PutUint64(z_bytes[32:40], src[4])
	binary.BigEndian.PutUint64(z_bytes[40:48], src[3])
	binary.BigEndian.PutUint64(z_bytes[48:56], src[2])
	binary.BigEndian.PutUint64(z_bytes[56:64], src[1])
	binary.BigEndian.PutUint64(z_bytes[64:72], src[0])

	return nil
}

func MulMontNonUnrolled640(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [10]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[9] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[9] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[8] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[8] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[7] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[7] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[6] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[6] = binary.BigEndian.Uint64(y_bytes[24:32])
	x[5] = binary.BigEndian.Uint64(x_bytes[32:40])
	y[5] = binary.BigEndian.Uint64(y_bytes[32:40])
	x[4] = binary.BigEndian.Uint64(x_bytes[40:48])
	y[4] = binary.BigEndian.Uint64(y_bytes[40:48])
	x[3] = binary.BigEndian.Uint64(x_bytes[48:56])
	y[3] = binary.BigEndian.Uint64(y_bytes[48:56])
	x[2] = binary.BigEndian.Uint64(x_bytes[56:64])
	y[2] = binary.BigEndian.Uint64(y_bytes[56:64])
	x[1] = binary.BigEndian.Uint64(x_bytes[64:72])
	y[1] = binary.BigEndian.Uint64(y_bytes[64:72])
	x[0] = binary.BigEndian.Uint64(x_bytes[72:80])
	y[0] = binary.BigEndian.Uint64(y_bytes[72:80])

	mod := ctx.ModulusLimbs
	var t [11]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC1 = bits.Sub64(mod[6], x[6], gteC1)
	_, gteC1 = bits.Sub64(mod[7], x[7], gteC1)
	_, gteC1 = bits.Sub64(mod[8], x[8], gteC1)
	_, gteC1 = bits.Sub64(mod[9], x[9], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], y[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], y[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], y[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], y[7], gteC2)
	_, gteC2 = bits.Sub64(mod[8], y[8], gteC2)
	_, gteC2 = bits.Sub64(mod[9], y[9], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

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
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)
	z[6], D = bits.Sub64(t[6], mod[6], D)
	z[7], D = bits.Sub64(t[7], mod[7], D)
	z[8], D = bits.Sub64(t[8], mod[8], D)
	z[9], D = bits.Sub64(t[9], mod[9], D)

	var src []uint64
	if D != 0 && t[10] == 0 {
		src = t[:10]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[9])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[8])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[7])
	binary.BigEndian.PutUint64(z_bytes[24:32], src[6])
	binary.BigEndian.PutUint64(z_bytes[32:40], src[5])
	binary.BigEndian.PutUint64(z_bytes[40:48], src[4])
	binary.BigEndian.PutUint64(z_bytes[48:56], src[3])
	binary.BigEndian.PutUint64(z_bytes[56:64], src[2])
	binary.BigEndian.PutUint64(z_bytes[64:72], src[1])
	binary.BigEndian.PutUint64(z_bytes[72:80], src[0])

	return nil
}

func MulMontNonUnrolled704(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [11]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[10] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[10] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[9] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[9] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[8] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[8] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[7] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[7] = binary.BigEndian.Uint64(y_bytes[24:32])
	x[6] = binary.BigEndian.Uint64(x_bytes[32:40])
	y[6] = binary.BigEndian.Uint64(y_bytes[32:40])
	x[5] = binary.BigEndian.Uint64(x_bytes[40:48])
	y[5] = binary.BigEndian.Uint64(y_bytes[40:48])
	x[4] = binary.BigEndian.Uint64(x_bytes[48:56])
	y[4] = binary.BigEndian.Uint64(y_bytes[48:56])
	x[3] = binary.BigEndian.Uint64(x_bytes[56:64])
	y[3] = binary.BigEndian.Uint64(y_bytes[56:64])
	x[2] = binary.BigEndian.Uint64(x_bytes[64:72])
	y[2] = binary.BigEndian.Uint64(y_bytes[64:72])
	x[1] = binary.BigEndian.Uint64(x_bytes[72:80])
	y[1] = binary.BigEndian.Uint64(y_bytes[72:80])
	x[0] = binary.BigEndian.Uint64(x_bytes[80:88])
	y[0] = binary.BigEndian.Uint64(y_bytes[80:88])

	mod := ctx.ModulusLimbs
	var t [12]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC1 = bits.Sub64(mod[6], x[6], gteC1)
	_, gteC1 = bits.Sub64(mod[7], x[7], gteC1)
	_, gteC1 = bits.Sub64(mod[8], x[8], gteC1)
	_, gteC1 = bits.Sub64(mod[9], x[9], gteC1)
	_, gteC1 = bits.Sub64(mod[10], x[10], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], y[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], y[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], y[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], y[7], gteC2)
	_, gteC2 = bits.Sub64(mod[8], y[8], gteC2)
	_, gteC2 = bits.Sub64(mod[9], y[9], gteC2)
	_, gteC2 = bits.Sub64(mod[10], y[10], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

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
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)
	z[6], D = bits.Sub64(t[6], mod[6], D)
	z[7], D = bits.Sub64(t[7], mod[7], D)
	z[8], D = bits.Sub64(t[8], mod[8], D)
	z[9], D = bits.Sub64(t[9], mod[9], D)
	z[10], D = bits.Sub64(t[10], mod[10], D)

	var src []uint64
	if D != 0 && t[11] == 0 {
		src = t[:11]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[10])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[9])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[8])
	binary.BigEndian.PutUint64(z_bytes[24:32], src[7])
	binary.BigEndian.PutUint64(z_bytes[32:40], src[6])
	binary.BigEndian.PutUint64(z_bytes[40:48], src[5])
	binary.BigEndian.PutUint64(z_bytes[48:56], src[4])
	binary.BigEndian.PutUint64(z_bytes[56:64], src[3])
	binary.BigEndian.PutUint64(z_bytes[64:72], src[2])
	binary.BigEndian.PutUint64(z_bytes[72:80], src[1])
	binary.BigEndian.PutUint64(z_bytes[80:88], src[0])

	return nil
}

func MulMontNonUnrolled768(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [12]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[11] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[11] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[10] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[10] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[9] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[9] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[8] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[8] = binary.BigEndian.Uint64(y_bytes[24:32])
	x[7] = binary.BigEndian.Uint64(x_bytes[32:40])
	y[7] = binary.BigEndian.Uint64(y_bytes[32:40])
	x[6] = binary.BigEndian.Uint64(x_bytes[40:48])
	y[6] = binary.BigEndian.Uint64(y_bytes[40:48])
	x[5] = binary.BigEndian.Uint64(x_bytes[48:56])
	y[5] = binary.BigEndian.Uint64(y_bytes[48:56])
	x[4] = binary.BigEndian.Uint64(x_bytes[56:64])
	y[4] = binary.BigEndian.Uint64(y_bytes[56:64])
	x[3] = binary.BigEndian.Uint64(x_bytes[64:72])
	y[3] = binary.BigEndian.Uint64(y_bytes[64:72])
	x[2] = binary.BigEndian.Uint64(x_bytes[72:80])
	y[2] = binary.BigEndian.Uint64(y_bytes[72:80])
	x[1] = binary.BigEndian.Uint64(x_bytes[80:88])
	y[1] = binary.BigEndian.Uint64(y_bytes[80:88])
	x[0] = binary.BigEndian.Uint64(x_bytes[88:96])
	y[0] = binary.BigEndian.Uint64(y_bytes[88:96])

	mod := ctx.ModulusLimbs
	var t [13]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC1 = bits.Sub64(mod[6], x[6], gteC1)
	_, gteC1 = bits.Sub64(mod[7], x[7], gteC1)
	_, gteC1 = bits.Sub64(mod[8], x[8], gteC1)
	_, gteC1 = bits.Sub64(mod[9], x[9], gteC1)
	_, gteC1 = bits.Sub64(mod[10], x[10], gteC1)
	_, gteC1 = bits.Sub64(mod[11], x[11], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], y[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], y[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], y[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], y[7], gteC2)
	_, gteC2 = bits.Sub64(mod[8], y[8], gteC2)
	_, gteC2 = bits.Sub64(mod[9], y[9], gteC2)
	_, gteC2 = bits.Sub64(mod[10], y[10], gteC2)
	_, gteC2 = bits.Sub64(mod[11], y[11], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

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
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)
	z[6], D = bits.Sub64(t[6], mod[6], D)
	z[7], D = bits.Sub64(t[7], mod[7], D)
	z[8], D = bits.Sub64(t[8], mod[8], D)
	z[9], D = bits.Sub64(t[9], mod[9], D)
	z[10], D = bits.Sub64(t[10], mod[10], D)
	z[11], D = bits.Sub64(t[11], mod[11], D)

	var src []uint64
	if D != 0 && t[12] == 0 {
		src = t[:12]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[11])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[10])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[9])
	binary.BigEndian.PutUint64(z_bytes[24:32], src[8])
	binary.BigEndian.PutUint64(z_bytes[32:40], src[7])
	binary.BigEndian.PutUint64(z_bytes[40:48], src[6])
	binary.BigEndian.PutUint64(z_bytes[48:56], src[5])
	binary.BigEndian.PutUint64(z_bytes[56:64], src[4])
	binary.BigEndian.PutUint64(z_bytes[64:72], src[3])
	binary.BigEndian.PutUint64(z_bytes[72:80], src[2])
	binary.BigEndian.PutUint64(z_bytes[80:88], src[1])
	binary.BigEndian.PutUint64(z_bytes[88:96], src[0])

	return nil
}

func MulMontNonUnrolled832(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [13]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[12] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[12] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[11] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[11] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[10] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[10] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[9] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[9] = binary.BigEndian.Uint64(y_bytes[24:32])
	x[8] = binary.BigEndian.Uint64(x_bytes[32:40])
	y[8] = binary.BigEndian.Uint64(y_bytes[32:40])
	x[7] = binary.BigEndian.Uint64(x_bytes[40:48])
	y[7] = binary.BigEndian.Uint64(y_bytes[40:48])
	x[6] = binary.BigEndian.Uint64(x_bytes[48:56])
	y[6] = binary.BigEndian.Uint64(y_bytes[48:56])
	x[5] = binary.BigEndian.Uint64(x_bytes[56:64])
	y[5] = binary.BigEndian.Uint64(y_bytes[56:64])
	x[4] = binary.BigEndian.Uint64(x_bytes[64:72])
	y[4] = binary.BigEndian.Uint64(y_bytes[64:72])
	x[3] = binary.BigEndian.Uint64(x_bytes[72:80])
	y[3] = binary.BigEndian.Uint64(y_bytes[72:80])
	x[2] = binary.BigEndian.Uint64(x_bytes[80:88])
	y[2] = binary.BigEndian.Uint64(y_bytes[80:88])
	x[1] = binary.BigEndian.Uint64(x_bytes[88:96])
	y[1] = binary.BigEndian.Uint64(y_bytes[88:96])
	x[0] = binary.BigEndian.Uint64(x_bytes[96:104])
	y[0] = binary.BigEndian.Uint64(y_bytes[96:104])

	mod := ctx.ModulusLimbs
	var t [14]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC1 = bits.Sub64(mod[6], x[6], gteC1)
	_, gteC1 = bits.Sub64(mod[7], x[7], gteC1)
	_, gteC1 = bits.Sub64(mod[8], x[8], gteC1)
	_, gteC1 = bits.Sub64(mod[9], x[9], gteC1)
	_, gteC1 = bits.Sub64(mod[10], x[10], gteC1)
	_, gteC1 = bits.Sub64(mod[11], x[11], gteC1)
	_, gteC1 = bits.Sub64(mod[12], x[12], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], y[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], y[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], y[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], y[7], gteC2)
	_, gteC2 = bits.Sub64(mod[8], y[8], gteC2)
	_, gteC2 = bits.Sub64(mod[9], y[9], gteC2)
	_, gteC2 = bits.Sub64(mod[10], y[10], gteC2)
	_, gteC2 = bits.Sub64(mod[11], y[11], gteC2)
	_, gteC2 = bits.Sub64(mod[12], y[12], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

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
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)
	z[6], D = bits.Sub64(t[6], mod[6], D)
	z[7], D = bits.Sub64(t[7], mod[7], D)
	z[8], D = bits.Sub64(t[8], mod[8], D)
	z[9], D = bits.Sub64(t[9], mod[9], D)
	z[10], D = bits.Sub64(t[10], mod[10], D)
	z[11], D = bits.Sub64(t[11], mod[11], D)
	z[12], D = bits.Sub64(t[12], mod[12], D)

	var src []uint64
	if D != 0 && t[13] == 0 {
		src = t[:13]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[12])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[11])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[10])
	binary.BigEndian.PutUint64(z_bytes[24:32], src[9])
	binary.BigEndian.PutUint64(z_bytes[32:40], src[8])
	binary.BigEndian.PutUint64(z_bytes[40:48], src[7])
	binary.BigEndian.PutUint64(z_bytes[48:56], src[6])
	binary.BigEndian.PutUint64(z_bytes[56:64], src[5])
	binary.BigEndian.PutUint64(z_bytes[64:72], src[4])
	binary.BigEndian.PutUint64(z_bytes[72:80], src[3])
	binary.BigEndian.PutUint64(z_bytes[80:88], src[2])
	binary.BigEndian.PutUint64(z_bytes[88:96], src[1])
	binary.BigEndian.PutUint64(z_bytes[96:104], src[0])

	return nil
}

func MulMontNonUnrolled896(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [14]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[13] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[13] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[12] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[12] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[11] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[11] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[10] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[10] = binary.BigEndian.Uint64(y_bytes[24:32])
	x[9] = binary.BigEndian.Uint64(x_bytes[32:40])
	y[9] = binary.BigEndian.Uint64(y_bytes[32:40])
	x[8] = binary.BigEndian.Uint64(x_bytes[40:48])
	y[8] = binary.BigEndian.Uint64(y_bytes[40:48])
	x[7] = binary.BigEndian.Uint64(x_bytes[48:56])
	y[7] = binary.BigEndian.Uint64(y_bytes[48:56])
	x[6] = binary.BigEndian.Uint64(x_bytes[56:64])
	y[6] = binary.BigEndian.Uint64(y_bytes[56:64])
	x[5] = binary.BigEndian.Uint64(x_bytes[64:72])
	y[5] = binary.BigEndian.Uint64(y_bytes[64:72])
	x[4] = binary.BigEndian.Uint64(x_bytes[72:80])
	y[4] = binary.BigEndian.Uint64(y_bytes[72:80])
	x[3] = binary.BigEndian.Uint64(x_bytes[80:88])
	y[3] = binary.BigEndian.Uint64(y_bytes[80:88])
	x[2] = binary.BigEndian.Uint64(x_bytes[88:96])
	y[2] = binary.BigEndian.Uint64(y_bytes[88:96])
	x[1] = binary.BigEndian.Uint64(x_bytes[96:104])
	y[1] = binary.BigEndian.Uint64(y_bytes[96:104])
	x[0] = binary.BigEndian.Uint64(x_bytes[104:112])
	y[0] = binary.BigEndian.Uint64(y_bytes[104:112])

	mod := ctx.ModulusLimbs
	var t [15]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC1 = bits.Sub64(mod[6], x[6], gteC1)
	_, gteC1 = bits.Sub64(mod[7], x[7], gteC1)
	_, gteC1 = bits.Sub64(mod[8], x[8], gteC1)
	_, gteC1 = bits.Sub64(mod[9], x[9], gteC1)
	_, gteC1 = bits.Sub64(mod[10], x[10], gteC1)
	_, gteC1 = bits.Sub64(mod[11], x[11], gteC1)
	_, gteC1 = bits.Sub64(mod[12], x[12], gteC1)
	_, gteC1 = bits.Sub64(mod[13], x[13], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], y[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], y[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], y[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], y[7], gteC2)
	_, gteC2 = bits.Sub64(mod[8], y[8], gteC2)
	_, gteC2 = bits.Sub64(mod[9], y[9], gteC2)
	_, gteC2 = bits.Sub64(mod[10], y[10], gteC2)
	_, gteC2 = bits.Sub64(mod[11], y[11], gteC2)
	_, gteC2 = bits.Sub64(mod[12], y[12], gteC2)
	_, gteC2 = bits.Sub64(mod[13], y[13], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

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
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)
	z[6], D = bits.Sub64(t[6], mod[6], D)
	z[7], D = bits.Sub64(t[7], mod[7], D)
	z[8], D = bits.Sub64(t[8], mod[8], D)
	z[9], D = bits.Sub64(t[9], mod[9], D)
	z[10], D = bits.Sub64(t[10], mod[10], D)
	z[11], D = bits.Sub64(t[11], mod[11], D)
	z[12], D = bits.Sub64(t[12], mod[12], D)
	z[13], D = bits.Sub64(t[13], mod[13], D)

	var src []uint64
	if D != 0 && t[14] == 0 {
		src = t[:14]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[13])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[12])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[11])
	binary.BigEndian.PutUint64(z_bytes[24:32], src[10])
	binary.BigEndian.PutUint64(z_bytes[32:40], src[9])
	binary.BigEndian.PutUint64(z_bytes[40:48], src[8])
	binary.BigEndian.PutUint64(z_bytes[48:56], src[7])
	binary.BigEndian.PutUint64(z_bytes[56:64], src[6])
	binary.BigEndian.PutUint64(z_bytes[64:72], src[5])
	binary.BigEndian.PutUint64(z_bytes[72:80], src[4])
	binary.BigEndian.PutUint64(z_bytes[80:88], src[3])
	binary.BigEndian.PutUint64(z_bytes[88:96], src[2])
	binary.BigEndian.PutUint64(z_bytes[96:104], src[1])
	binary.BigEndian.PutUint64(z_bytes[104:112], src[0])

	return nil
}

func MulMontNonUnrolled960(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [15]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[14] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[14] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[13] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[13] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[12] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[12] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[11] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[11] = binary.BigEndian.Uint64(y_bytes[24:32])
	x[10] = binary.BigEndian.Uint64(x_bytes[32:40])
	y[10] = binary.BigEndian.Uint64(y_bytes[32:40])
	x[9] = binary.BigEndian.Uint64(x_bytes[40:48])
	y[9] = binary.BigEndian.Uint64(y_bytes[40:48])
	x[8] = binary.BigEndian.Uint64(x_bytes[48:56])
	y[8] = binary.BigEndian.Uint64(y_bytes[48:56])
	x[7] = binary.BigEndian.Uint64(x_bytes[56:64])
	y[7] = binary.BigEndian.Uint64(y_bytes[56:64])
	x[6] = binary.BigEndian.Uint64(x_bytes[64:72])
	y[6] = binary.BigEndian.Uint64(y_bytes[64:72])
	x[5] = binary.BigEndian.Uint64(x_bytes[72:80])
	y[5] = binary.BigEndian.Uint64(y_bytes[72:80])
	x[4] = binary.BigEndian.Uint64(x_bytes[80:88])
	y[4] = binary.BigEndian.Uint64(y_bytes[80:88])
	x[3] = binary.BigEndian.Uint64(x_bytes[88:96])
	y[3] = binary.BigEndian.Uint64(y_bytes[88:96])
	x[2] = binary.BigEndian.Uint64(x_bytes[96:104])
	y[2] = binary.BigEndian.Uint64(y_bytes[96:104])
	x[1] = binary.BigEndian.Uint64(x_bytes[104:112])
	y[1] = binary.BigEndian.Uint64(y_bytes[104:112])
	x[0] = binary.BigEndian.Uint64(x_bytes[112:120])
	y[0] = binary.BigEndian.Uint64(y_bytes[112:120])

	mod := ctx.ModulusLimbs
	var t [16]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC1 = bits.Sub64(mod[6], x[6], gteC1)
	_, gteC1 = bits.Sub64(mod[7], x[7], gteC1)
	_, gteC1 = bits.Sub64(mod[8], x[8], gteC1)
	_, gteC1 = bits.Sub64(mod[9], x[9], gteC1)
	_, gteC1 = bits.Sub64(mod[10], x[10], gteC1)
	_, gteC1 = bits.Sub64(mod[11], x[11], gteC1)
	_, gteC1 = bits.Sub64(mod[12], x[12], gteC1)
	_, gteC1 = bits.Sub64(mod[13], x[13], gteC1)
	_, gteC1 = bits.Sub64(mod[14], x[14], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], y[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], y[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], y[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], y[7], gteC2)
	_, gteC2 = bits.Sub64(mod[8], y[8], gteC2)
	_, gteC2 = bits.Sub64(mod[9], y[9], gteC2)
	_, gteC2 = bits.Sub64(mod[10], y[10], gteC2)
	_, gteC2 = bits.Sub64(mod[11], y[11], gteC2)
	_, gteC2 = bits.Sub64(mod[12], y[12], gteC2)
	_, gteC2 = bits.Sub64(mod[13], y[13], gteC2)
	_, gteC2 = bits.Sub64(mod[14], y[14], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

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
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)
	z[6], D = bits.Sub64(t[6], mod[6], D)
	z[7], D = bits.Sub64(t[7], mod[7], D)
	z[8], D = bits.Sub64(t[8], mod[8], D)
	z[9], D = bits.Sub64(t[9], mod[9], D)
	z[10], D = bits.Sub64(t[10], mod[10], D)
	z[11], D = bits.Sub64(t[11], mod[11], D)
	z[12], D = bits.Sub64(t[12], mod[12], D)
	z[13], D = bits.Sub64(t[13], mod[13], D)
	z[14], D = bits.Sub64(t[14], mod[14], D)

	var src []uint64
	if D != 0 && t[15] == 0 {
		src = t[:15]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[14])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[13])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[12])
	binary.BigEndian.PutUint64(z_bytes[24:32], src[11])
	binary.BigEndian.PutUint64(z_bytes[32:40], src[10])
	binary.BigEndian.PutUint64(z_bytes[40:48], src[9])
	binary.BigEndian.PutUint64(z_bytes[48:56], src[8])
	binary.BigEndian.PutUint64(z_bytes[56:64], src[7])
	binary.BigEndian.PutUint64(z_bytes[64:72], src[6])
	binary.BigEndian.PutUint64(z_bytes[72:80], src[5])
	binary.BigEndian.PutUint64(z_bytes[80:88], src[4])
	binary.BigEndian.PutUint64(z_bytes[88:96], src[3])
	binary.BigEndian.PutUint64(z_bytes[96:104], src[2])
	binary.BigEndian.PutUint64(z_bytes[104:112], src[1])
	binary.BigEndian.PutUint64(z_bytes[112:120], src[0])

	return nil
}

func MulMontNonUnrolled1024(ctx *Field, z_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [16]uint64

	// conversion to little-endian limb-order, system limb-endianess
	x[15] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[15] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[14] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[14] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[13] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[13] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[12] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[12] = binary.BigEndian.Uint64(y_bytes[24:32])
	x[11] = binary.BigEndian.Uint64(x_bytes[32:40])
	y[11] = binary.BigEndian.Uint64(y_bytes[32:40])
	x[10] = binary.BigEndian.Uint64(x_bytes[40:48])
	y[10] = binary.BigEndian.Uint64(y_bytes[40:48])
	x[9] = binary.BigEndian.Uint64(x_bytes[48:56])
	y[9] = binary.BigEndian.Uint64(y_bytes[48:56])
	x[8] = binary.BigEndian.Uint64(x_bytes[56:64])
	y[8] = binary.BigEndian.Uint64(y_bytes[56:64])
	x[7] = binary.BigEndian.Uint64(x_bytes[64:72])
	y[7] = binary.BigEndian.Uint64(y_bytes[64:72])
	x[6] = binary.BigEndian.Uint64(x_bytes[72:80])
	y[6] = binary.BigEndian.Uint64(y_bytes[72:80])
	x[5] = binary.BigEndian.Uint64(x_bytes[80:88])
	y[5] = binary.BigEndian.Uint64(y_bytes[80:88])
	x[4] = binary.BigEndian.Uint64(x_bytes[88:96])
	y[4] = binary.BigEndian.Uint64(y_bytes[88:96])
	x[3] = binary.BigEndian.Uint64(x_bytes[96:104])
	y[3] = binary.BigEndian.Uint64(y_bytes[96:104])
	x[2] = binary.BigEndian.Uint64(x_bytes[104:112])
	y[2] = binary.BigEndian.Uint64(y_bytes[104:112])
	x[1] = binary.BigEndian.Uint64(x_bytes[112:120])
	y[1] = binary.BigEndian.Uint64(y_bytes[112:120])
	x[0] = binary.BigEndian.Uint64(x_bytes[120:128])
	y[0] = binary.BigEndian.Uint64(y_bytes[120:128])

	mod := ctx.ModulusLimbs
	var t [17]uint64
	var D uint64
	var m, C uint64

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC1 = bits.Sub64(mod[6], x[6], gteC1)
	_, gteC1 = bits.Sub64(mod[7], x[7], gteC1)
	_, gteC1 = bits.Sub64(mod[8], x[8], gteC1)
	_, gteC1 = bits.Sub64(mod[9], x[9], gteC1)
	_, gteC1 = bits.Sub64(mod[10], x[10], gteC1)
	_, gteC1 = bits.Sub64(mod[11], x[11], gteC1)
	_, gteC1 = bits.Sub64(mod[12], x[12], gteC1)
	_, gteC1 = bits.Sub64(mod[13], x[13], gteC1)
	_, gteC1 = bits.Sub64(mod[14], x[14], gteC1)
	_, gteC1 = bits.Sub64(mod[15], x[15], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], y[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], y[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], y[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], y[7], gteC2)
	_, gteC2 = bits.Sub64(mod[8], y[8], gteC2)
	_, gteC2 = bits.Sub64(mod[9], y[9], gteC2)
	_, gteC2 = bits.Sub64(mod[10], y[10], gteC2)
	_, gteC2 = bits.Sub64(mod[11], y[11], gteC2)
	_, gteC2 = bits.Sub64(mod[12], y[12], gteC2)
	_, gteC2 = bits.Sub64(mod[13], y[13], gteC2)
	_, gteC2 = bits.Sub64(mod[14], y[14], gteC2)
	_, gteC2 = bits.Sub64(mod[15], y[15], gteC2)

	/*
	   fmt.Println()
	   fmt.Println()
	   fmt.Println("foo")
	   fmt.Println(x)
	   fmt.Println(y)
	   fmt.Println(mod)
	*/

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus: x=%x,y=%x,mod=%x", x, y, z))
	}

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
	m = t[0] * ctx.MontParamInterleaved

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
		m = t[0] * ctx.MontParamInterleaved

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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)
	z[6], D = bits.Sub64(t[6], mod[6], D)
	z[7], D = bits.Sub64(t[7], mod[7], D)
	z[8], D = bits.Sub64(t[8], mod[8], D)
	z[9], D = bits.Sub64(t[9], mod[9], D)
	z[10], D = bits.Sub64(t[10], mod[10], D)
	z[11], D = bits.Sub64(t[11], mod[11], D)
	z[12], D = bits.Sub64(t[12], mod[12], D)
	z[13], D = bits.Sub64(t[13], mod[13], D)
	z[14], D = bits.Sub64(t[14], mod[14], D)
	z[15], D = bits.Sub64(t[15], mod[15], D)

	var src []uint64
	if D != 0 && t[16] == 0 {
		src = t[:16]
	} else {
		src = z[:]
	}
	binary.BigEndian.PutUint64(z_bytes[0:8], src[15])
	binary.BigEndian.PutUint64(z_bytes[8:16], src[14])
	binary.BigEndian.PutUint64(z_bytes[16:24], src[13])
	binary.BigEndian.PutUint64(z_bytes[24:32], src[12])
	binary.BigEndian.PutUint64(z_bytes[32:40], src[11])
	binary.BigEndian.PutUint64(z_bytes[40:48], src[10])
	binary.BigEndian.PutUint64(z_bytes[48:56], src[9])
	binary.BigEndian.PutUint64(z_bytes[56:64], src[8])
	binary.BigEndian.PutUint64(z_bytes[64:72], src[7])
	binary.BigEndian.PutUint64(z_bytes[72:80], src[6])
	binary.BigEndian.PutUint64(z_bytes[80:88], src[5])
	binary.BigEndian.PutUint64(z_bytes[88:96], src[4])
	binary.BigEndian.PutUint64(z_bytes[96:104], src[3])
	binary.BigEndian.PutUint64(z_bytes[104:112], src[2])
	binary.BigEndian.PutUint64(z_bytes[112:120], src[1])
	binary.BigEndian.PutUint64(z_bytes[120:128], src[0])

	return nil
}
