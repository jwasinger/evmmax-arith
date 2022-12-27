package evmmax_arith

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/bits"
)

func SubModNonUnrolled64(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [1]uint64
	x[0] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[0] = binary.BigEndian.Uint64(y_bytes[0:8])

	mod := f.ModulusLimbs
	_ = mod[0]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [1]uint64{0}

	for i := 0; i < 1; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 1; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[0]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[0])
	return nil
}

func SubModNonUnrolled128(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [2]uint64
	x[1] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[1] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[0] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[0] = binary.BigEndian.Uint64(y_bytes[8:16])

	mod := f.ModulusLimbs
	_ = mod[1]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [2]uint64{0, 0}

	for i := 0; i < 2; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 2; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[1]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[1])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[0])
	return nil
}

func SubModNonUnrolled192(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [3]uint64
	x[2] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[2] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[1] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[1] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[0] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[0] = binary.BigEndian.Uint64(y_bytes[16:24])

	mod := f.ModulusLimbs
	_ = mod[2]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [3]uint64{0, 0, 0}

	for i := 0; i < 3; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 3; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[2]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[2])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[1])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[0])
	return nil
}

func SubModNonUnrolled256(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [4]uint64
	x[3] = binary.BigEndian.Uint64(x_bytes[0:8])
	y[3] = binary.BigEndian.Uint64(y_bytes[0:8])
	x[2] = binary.BigEndian.Uint64(x_bytes[8:16])
	y[2] = binary.BigEndian.Uint64(y_bytes[8:16])
	x[1] = binary.BigEndian.Uint64(x_bytes[16:24])
	y[1] = binary.BigEndian.Uint64(y_bytes[16:24])
	x[0] = binary.BigEndian.Uint64(x_bytes[24:32])
	y[0] = binary.BigEndian.Uint64(y_bytes[24:32])

	mod := f.ModulusLimbs
	_ = mod[3]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [4]uint64{0, 0, 0, 0}

	for i := 0; i < 4; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 4; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[3]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[3])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[2])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[1])
	binary.BigEndian.PutUint64(out_bytes[24:32], src[0])
	return nil
}

func SubModNonUnrolled320(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [5]uint64
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

	mod := f.ModulusLimbs
	_ = mod[4]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [5]uint64{0, 0, 0, 0, 0}

	for i := 0; i < 5; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 5; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[4]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[4])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[3])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[2])
	binary.BigEndian.PutUint64(out_bytes[24:32], src[1])
	binary.BigEndian.PutUint64(out_bytes[32:40], src[0])
	return nil
}

func SubModNonUnrolled384(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [6]uint64
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

	mod := f.ModulusLimbs
	_ = mod[5]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [6]uint64{0, 0, 0, 0, 0, 0}

	for i := 0; i < 6; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 6; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[5]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[5])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[4])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[3])
	binary.BigEndian.PutUint64(out_bytes[24:32], src[2])
	binary.BigEndian.PutUint64(out_bytes[32:40], src[1])
	binary.BigEndian.PutUint64(out_bytes[40:48], src[0])
	return nil
}

func SubModNonUnrolled448(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [7]uint64
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

	mod := f.ModulusLimbs
	_ = mod[6]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [7]uint64{0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 7; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 7; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[6]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[6])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[5])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[4])
	binary.BigEndian.PutUint64(out_bytes[24:32], src[3])
	binary.BigEndian.PutUint64(out_bytes[32:40], src[2])
	binary.BigEndian.PutUint64(out_bytes[40:48], src[1])
	binary.BigEndian.PutUint64(out_bytes[48:56], src[0])
	return nil
}

func SubModNonUnrolled512(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [8]uint64
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

	mod := f.ModulusLimbs
	_ = mod[7]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [8]uint64{0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 8; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 8; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[7]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[7])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[6])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[5])
	binary.BigEndian.PutUint64(out_bytes[24:32], src[4])
	binary.BigEndian.PutUint64(out_bytes[32:40], src[3])
	binary.BigEndian.PutUint64(out_bytes[40:48], src[2])
	binary.BigEndian.PutUint64(out_bytes[48:56], src[1])
	binary.BigEndian.PutUint64(out_bytes[56:64], src[0])
	return nil
}

func SubModNonUnrolled576(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [9]uint64
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

	mod := f.ModulusLimbs
	_ = mod[8]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [9]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 9; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 9; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[8]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[8])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[7])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[6])
	binary.BigEndian.PutUint64(out_bytes[24:32], src[5])
	binary.BigEndian.PutUint64(out_bytes[32:40], src[4])
	binary.BigEndian.PutUint64(out_bytes[40:48], src[3])
	binary.BigEndian.PutUint64(out_bytes[48:56], src[2])
	binary.BigEndian.PutUint64(out_bytes[56:64], src[1])
	binary.BigEndian.PutUint64(out_bytes[64:72], src[0])
	return nil
}

func SubModNonUnrolled640(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [10]uint64
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

	mod := f.ModulusLimbs
	_ = mod[9]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [10]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 10; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 10; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[9]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[9])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[8])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[7])
	binary.BigEndian.PutUint64(out_bytes[24:32], src[6])
	binary.BigEndian.PutUint64(out_bytes[32:40], src[5])
	binary.BigEndian.PutUint64(out_bytes[40:48], src[4])
	binary.BigEndian.PutUint64(out_bytes[48:56], src[3])
	binary.BigEndian.PutUint64(out_bytes[56:64], src[2])
	binary.BigEndian.PutUint64(out_bytes[64:72], src[1])
	binary.BigEndian.PutUint64(out_bytes[72:80], src[0])
	return nil
}

func SubModNonUnrolled704(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [11]uint64
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

	mod := f.ModulusLimbs
	_ = mod[10]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [11]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 11; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 11; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[10]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[10])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[9])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[8])
	binary.BigEndian.PutUint64(out_bytes[24:32], src[7])
	binary.BigEndian.PutUint64(out_bytes[32:40], src[6])
	binary.BigEndian.PutUint64(out_bytes[40:48], src[5])
	binary.BigEndian.PutUint64(out_bytes[48:56], src[4])
	binary.BigEndian.PutUint64(out_bytes[56:64], src[3])
	binary.BigEndian.PutUint64(out_bytes[64:72], src[2])
	binary.BigEndian.PutUint64(out_bytes[72:80], src[1])
	binary.BigEndian.PutUint64(out_bytes[80:88], src[0])
	return nil
}

func SubModNonUnrolled768(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [12]uint64
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

	mod := f.ModulusLimbs
	_ = mod[11]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [12]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 12; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 12; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[11]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[11])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[10])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[9])
	binary.BigEndian.PutUint64(out_bytes[24:32], src[8])
	binary.BigEndian.PutUint64(out_bytes[32:40], src[7])
	binary.BigEndian.PutUint64(out_bytes[40:48], src[6])
	binary.BigEndian.PutUint64(out_bytes[48:56], src[5])
	binary.BigEndian.PutUint64(out_bytes[56:64], src[4])
	binary.BigEndian.PutUint64(out_bytes[64:72], src[3])
	binary.BigEndian.PutUint64(out_bytes[72:80], src[2])
	binary.BigEndian.PutUint64(out_bytes[80:88], src[1])
	binary.BigEndian.PutUint64(out_bytes[88:96], src[0])
	return nil
}

func SubModNonUnrolled832(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [13]uint64
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

	mod := f.ModulusLimbs
	_ = mod[12]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [13]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 13; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 13; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[12]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[12])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[11])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[10])
	binary.BigEndian.PutUint64(out_bytes[24:32], src[9])
	binary.BigEndian.PutUint64(out_bytes[32:40], src[8])
	binary.BigEndian.PutUint64(out_bytes[40:48], src[7])
	binary.BigEndian.PutUint64(out_bytes[48:56], src[6])
	binary.BigEndian.PutUint64(out_bytes[56:64], src[5])
	binary.BigEndian.PutUint64(out_bytes[64:72], src[4])
	binary.BigEndian.PutUint64(out_bytes[72:80], src[3])
	binary.BigEndian.PutUint64(out_bytes[80:88], src[2])
	binary.BigEndian.PutUint64(out_bytes[88:96], src[1])
	binary.BigEndian.PutUint64(out_bytes[96:104], src[0])
	return nil
}

func SubModNonUnrolled896(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [14]uint64
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

	mod := f.ModulusLimbs
	_ = mod[13]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [14]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 14; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 14; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[13]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[13])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[12])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[11])
	binary.BigEndian.PutUint64(out_bytes[24:32], src[10])
	binary.BigEndian.PutUint64(out_bytes[32:40], src[9])
	binary.BigEndian.PutUint64(out_bytes[40:48], src[8])
	binary.BigEndian.PutUint64(out_bytes[48:56], src[7])
	binary.BigEndian.PutUint64(out_bytes[56:64], src[6])
	binary.BigEndian.PutUint64(out_bytes[64:72], src[5])
	binary.BigEndian.PutUint64(out_bytes[72:80], src[4])
	binary.BigEndian.PutUint64(out_bytes[80:88], src[3])
	binary.BigEndian.PutUint64(out_bytes[88:96], src[2])
	binary.BigEndian.PutUint64(out_bytes[96:104], src[1])
	binary.BigEndian.PutUint64(out_bytes[104:112], src[0])
	return nil
}

func SubModNonUnrolled960(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [15]uint64
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

	mod := f.ModulusLimbs
	_ = mod[14]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [15]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 15; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 15; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[14]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[14])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[13])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[12])
	binary.BigEndian.PutUint64(out_bytes[24:32], src[11])
	binary.BigEndian.PutUint64(out_bytes[32:40], src[10])
	binary.BigEndian.PutUint64(out_bytes[40:48], src[9])
	binary.BigEndian.PutUint64(out_bytes[48:56], src[8])
	binary.BigEndian.PutUint64(out_bytes[56:64], src[7])
	binary.BigEndian.PutUint64(out_bytes[64:72], src[6])
	binary.BigEndian.PutUint64(out_bytes[72:80], src[5])
	binary.BigEndian.PutUint64(out_bytes[80:88], src[4])
	binary.BigEndian.PutUint64(out_bytes[88:96], src[3])
	binary.BigEndian.PutUint64(out_bytes[96:104], src[2])
	binary.BigEndian.PutUint64(out_bytes[104:112], src[1])
	binary.BigEndian.PutUint64(out_bytes[112:120], src[0])
	return nil
}

func SubModNonUnrolled1024(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	var x, y, z [16]uint64
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

	mod := f.ModulusLimbs
	_ = mod[15]

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
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [16]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 16; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 16; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	var src []uint64
	// final sub was unnecessary
	if c == 0 {
		src = tmp[:]
	} else {
		src = z[:]
	}

	// pre-hint to compiler: TODO check asm to make sure this actually does something.
	_ = src[15]
	binary.BigEndian.PutUint64(out_bytes[0:8], src[15])
	binary.BigEndian.PutUint64(out_bytes[8:16], src[14])
	binary.BigEndian.PutUint64(out_bytes[16:24], src[13])
	binary.BigEndian.PutUint64(out_bytes[24:32], src[12])
	binary.BigEndian.PutUint64(out_bytes[32:40], src[11])
	binary.BigEndian.PutUint64(out_bytes[40:48], src[10])
	binary.BigEndian.PutUint64(out_bytes[48:56], src[9])
	binary.BigEndian.PutUint64(out_bytes[56:64], src[8])
	binary.BigEndian.PutUint64(out_bytes[64:72], src[7])
	binary.BigEndian.PutUint64(out_bytes[72:80], src[6])
	binary.BigEndian.PutUint64(out_bytes[80:88], src[5])
	binary.BigEndian.PutUint64(out_bytes[88:96], src[4])
	binary.BigEndian.PutUint64(out_bytes[96:104], src[3])
	binary.BigEndian.PutUint64(out_bytes[104:112], src[2])
	binary.BigEndian.PutUint64(out_bytes[112:120], src[1])
	binary.BigEndian.PutUint64(out_bytes[120:128], src[0])
	return nil
}
