package mont_arith

import (
	"errors"
	"fmt"
	"math/bits"
	"unsafe"
)

func AddModNonUnrolled64(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[1]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[1]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[1]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [1]uint64{0}

	for i := 0; i < 1; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 1; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled128(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [2]uint64{0, 0}

	for i := 0; i < 2; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 2; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled192(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[3]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [3]uint64{0, 0, 0}

	for i := 0; i < 3; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 3; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled256(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[4]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC2 = bits.Sub64(mod[0], y[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], y[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], y[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], y[3], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [4]uint64{0, 0, 0, 0}

	for i := 0; i < 4; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 4; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled320(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[5]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[5]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[5]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[5]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [5]uint64{0, 0, 0, 0, 0}

	for i := 0; i < 5; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 5; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled384(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [6]uint64{0, 0, 0, 0, 0, 0}

	for i := 0; i < 6; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 6; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled448(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[7]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [7]uint64{0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 7; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 7; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled512(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[8]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[8]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[8]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[8]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [8]uint64{0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 8; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 8; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled576(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [9]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 9; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 9; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled640(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[10]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [10]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 10; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 10; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled704(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[11]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[11]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[11]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[11]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [11]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 11; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 11; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled768(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[12]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[12]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[12]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [12]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 12; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 12; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled832(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[13]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[13]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[13]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[13]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [13]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 13; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 13; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled896(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[14]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[14]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[14]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[14]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [14]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 14; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 14; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled960(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[15]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[15]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[15]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[15]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [15]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 15; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 15; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1024(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[16]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[16]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[16]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[16]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [16]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 16; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 16; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1088(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[17]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[17]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[17]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[17]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [17]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 17; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 17; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1152(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[18]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[18]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[18]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[18]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [18]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 18; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 18; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1216(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[19]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[19]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[19]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[19]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [19]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 19; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 19; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1280(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[20]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[20]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[20]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[20]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [20]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 20; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 20; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1344(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[21]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[21]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[21]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[21]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [21]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 21; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 21; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1408(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[22]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[22]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[22]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[22]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [22]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 22; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 22; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1472(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[23]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[23]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[23]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[23]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [23]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 23; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 23; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1536(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[24]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[24]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[24]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[24]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [24]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 24; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 24; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1600(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[25]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[25]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[25]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[25]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [25]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 25; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 25; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1664(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[26]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[26]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[26]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[26]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [26]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 26; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 26; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1728(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[27]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[27]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[27]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[27]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [27]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 27; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 27; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1792(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[28]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[28]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[28]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[28]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [28]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 28; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 28; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1856(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[29]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[29]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[29]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[29]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [29]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 29; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 29; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1920(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[30]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[30]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[30]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[30]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [30]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 30; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 30; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled1984(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[31]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[31]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[31]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[31]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [31]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 31; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 31; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2048(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[32]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[32]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[32]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[32]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [32]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 32; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 32; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2112(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[33]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[33]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[33]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[33]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [33]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 33; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 33; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2176(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[34]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[34]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[34]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[34]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [34]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 34; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 34; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2240(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[35]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[35]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[35]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[35]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [35]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 35; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 35; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2304(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[36]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[36]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[36]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[36]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [36]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 36; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 36; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2368(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[37]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[37]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[37]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[37]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [37]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 37; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 37; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2432(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[38]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[38]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[38]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[38]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [38]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 38; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 38; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2496(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[39]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[39]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[39]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[39]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [39]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 39; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 39; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2560(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[40]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[40]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[40]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[40]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [40]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 40; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 40; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2624(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[41]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[41]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[41]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[41]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [41]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 41; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 41; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2688(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[42]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[42]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[42]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[42]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [42]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 42; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 42; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2752(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[43]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[43]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[43]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[43]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [43]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 43; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 43; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2816(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[44]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[44]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[44]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[44]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [44]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 44; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 44; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2880(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[45]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[45]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[45]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[45]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [45]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 45; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 45; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled2944(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[46]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[46]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[46]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[46]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [46]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 46; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 46; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3008(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[47]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[47]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[47]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[47]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [47]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 47; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 47; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3072(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[48]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[48]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[48]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[48]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [48]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 48; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 48; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3136(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[49]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[49]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[49]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[49]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [49]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 49; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 49; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3200(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[50]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[50]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[50]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[50]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [50]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 50; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 50; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3264(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[51]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[51]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[51]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[51]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [51]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 51; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 51; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3328(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[52]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[52]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[52]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[52]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
	_, gteC1 = bits.Sub64(mod[51], x[51], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)
	_, gteC2 = bits.Sub64(mod[51], y[51], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [52]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 52; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 52; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3392(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[53]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[53]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[53]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[53]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
	_, gteC1 = bits.Sub64(mod[51], x[51], gteC1)
	_, gteC1 = bits.Sub64(mod[52], x[52], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)
	_, gteC2 = bits.Sub64(mod[51], y[51], gteC2)
	_, gteC2 = bits.Sub64(mod[52], y[52], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [53]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 53; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 53; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3456(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[54]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[54]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[54]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[54]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
	_, gteC1 = bits.Sub64(mod[51], x[51], gteC1)
	_, gteC1 = bits.Sub64(mod[52], x[52], gteC1)
	_, gteC1 = bits.Sub64(mod[53], x[53], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)
	_, gteC2 = bits.Sub64(mod[51], y[51], gteC2)
	_, gteC2 = bits.Sub64(mod[52], y[52], gteC2)
	_, gteC2 = bits.Sub64(mod[53], y[53], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [54]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 54; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 54; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3520(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[55]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[55]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[55]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[55]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
	_, gteC1 = bits.Sub64(mod[51], x[51], gteC1)
	_, gteC1 = bits.Sub64(mod[52], x[52], gteC1)
	_, gteC1 = bits.Sub64(mod[53], x[53], gteC1)
	_, gteC1 = bits.Sub64(mod[54], x[54], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)
	_, gteC2 = bits.Sub64(mod[51], y[51], gteC2)
	_, gteC2 = bits.Sub64(mod[52], y[52], gteC2)
	_, gteC2 = bits.Sub64(mod[53], y[53], gteC2)
	_, gteC2 = bits.Sub64(mod[54], y[54], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [55]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 55; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 55; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3584(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[56]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[56]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[56]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[56]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
	_, gteC1 = bits.Sub64(mod[51], x[51], gteC1)
	_, gteC1 = bits.Sub64(mod[52], x[52], gteC1)
	_, gteC1 = bits.Sub64(mod[53], x[53], gteC1)
	_, gteC1 = bits.Sub64(mod[54], x[54], gteC1)
	_, gteC1 = bits.Sub64(mod[55], x[55], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)
	_, gteC2 = bits.Sub64(mod[51], y[51], gteC2)
	_, gteC2 = bits.Sub64(mod[52], y[52], gteC2)
	_, gteC2 = bits.Sub64(mod[53], y[53], gteC2)
	_, gteC2 = bits.Sub64(mod[54], y[54], gteC2)
	_, gteC2 = bits.Sub64(mod[55], y[55], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [56]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 56; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 56; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3648(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[57]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[57]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[57]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[57]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
	_, gteC1 = bits.Sub64(mod[51], x[51], gteC1)
	_, gteC1 = bits.Sub64(mod[52], x[52], gteC1)
	_, gteC1 = bits.Sub64(mod[53], x[53], gteC1)
	_, gteC1 = bits.Sub64(mod[54], x[54], gteC1)
	_, gteC1 = bits.Sub64(mod[55], x[55], gteC1)
	_, gteC1 = bits.Sub64(mod[56], x[56], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)
	_, gteC2 = bits.Sub64(mod[51], y[51], gteC2)
	_, gteC2 = bits.Sub64(mod[52], y[52], gteC2)
	_, gteC2 = bits.Sub64(mod[53], y[53], gteC2)
	_, gteC2 = bits.Sub64(mod[54], y[54], gteC2)
	_, gteC2 = bits.Sub64(mod[55], y[55], gteC2)
	_, gteC2 = bits.Sub64(mod[56], y[56], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [57]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 57; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 57; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3712(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[58]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[58]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[58]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[58]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
	_, gteC1 = bits.Sub64(mod[51], x[51], gteC1)
	_, gteC1 = bits.Sub64(mod[52], x[52], gteC1)
	_, gteC1 = bits.Sub64(mod[53], x[53], gteC1)
	_, gteC1 = bits.Sub64(mod[54], x[54], gteC1)
	_, gteC1 = bits.Sub64(mod[55], x[55], gteC1)
	_, gteC1 = bits.Sub64(mod[56], x[56], gteC1)
	_, gteC1 = bits.Sub64(mod[57], x[57], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)
	_, gteC2 = bits.Sub64(mod[51], y[51], gteC2)
	_, gteC2 = bits.Sub64(mod[52], y[52], gteC2)
	_, gteC2 = bits.Sub64(mod[53], y[53], gteC2)
	_, gteC2 = bits.Sub64(mod[54], y[54], gteC2)
	_, gteC2 = bits.Sub64(mod[55], y[55], gteC2)
	_, gteC2 = bits.Sub64(mod[56], y[56], gteC2)
	_, gteC2 = bits.Sub64(mod[57], y[57], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [58]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 58; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 58; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3776(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[59]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[59]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[59]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[59]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
	_, gteC1 = bits.Sub64(mod[51], x[51], gteC1)
	_, gteC1 = bits.Sub64(mod[52], x[52], gteC1)
	_, gteC1 = bits.Sub64(mod[53], x[53], gteC1)
	_, gteC1 = bits.Sub64(mod[54], x[54], gteC1)
	_, gteC1 = bits.Sub64(mod[55], x[55], gteC1)
	_, gteC1 = bits.Sub64(mod[56], x[56], gteC1)
	_, gteC1 = bits.Sub64(mod[57], x[57], gteC1)
	_, gteC1 = bits.Sub64(mod[58], x[58], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)
	_, gteC2 = bits.Sub64(mod[51], y[51], gteC2)
	_, gteC2 = bits.Sub64(mod[52], y[52], gteC2)
	_, gteC2 = bits.Sub64(mod[53], y[53], gteC2)
	_, gteC2 = bits.Sub64(mod[54], y[54], gteC2)
	_, gteC2 = bits.Sub64(mod[55], y[55], gteC2)
	_, gteC2 = bits.Sub64(mod[56], y[56], gteC2)
	_, gteC2 = bits.Sub64(mod[57], y[57], gteC2)
	_, gteC2 = bits.Sub64(mod[58], y[58], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [59]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 59; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 59; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3840(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[60]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[60]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[60]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[60]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
	_, gteC1 = bits.Sub64(mod[51], x[51], gteC1)
	_, gteC1 = bits.Sub64(mod[52], x[52], gteC1)
	_, gteC1 = bits.Sub64(mod[53], x[53], gteC1)
	_, gteC1 = bits.Sub64(mod[54], x[54], gteC1)
	_, gteC1 = bits.Sub64(mod[55], x[55], gteC1)
	_, gteC1 = bits.Sub64(mod[56], x[56], gteC1)
	_, gteC1 = bits.Sub64(mod[57], x[57], gteC1)
	_, gteC1 = bits.Sub64(mod[58], x[58], gteC1)
	_, gteC1 = bits.Sub64(mod[59], x[59], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)
	_, gteC2 = bits.Sub64(mod[51], y[51], gteC2)
	_, gteC2 = bits.Sub64(mod[52], y[52], gteC2)
	_, gteC2 = bits.Sub64(mod[53], y[53], gteC2)
	_, gteC2 = bits.Sub64(mod[54], y[54], gteC2)
	_, gteC2 = bits.Sub64(mod[55], y[55], gteC2)
	_, gteC2 = bits.Sub64(mod[56], y[56], gteC2)
	_, gteC2 = bits.Sub64(mod[57], y[57], gteC2)
	_, gteC2 = bits.Sub64(mod[58], y[58], gteC2)
	_, gteC2 = bits.Sub64(mod[59], y[59], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [60]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 60; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 60; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3904(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[61]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[61]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[61]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[61]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
	_, gteC1 = bits.Sub64(mod[51], x[51], gteC1)
	_, gteC1 = bits.Sub64(mod[52], x[52], gteC1)
	_, gteC1 = bits.Sub64(mod[53], x[53], gteC1)
	_, gteC1 = bits.Sub64(mod[54], x[54], gteC1)
	_, gteC1 = bits.Sub64(mod[55], x[55], gteC1)
	_, gteC1 = bits.Sub64(mod[56], x[56], gteC1)
	_, gteC1 = bits.Sub64(mod[57], x[57], gteC1)
	_, gteC1 = bits.Sub64(mod[58], x[58], gteC1)
	_, gteC1 = bits.Sub64(mod[59], x[59], gteC1)
	_, gteC1 = bits.Sub64(mod[60], x[60], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)
	_, gteC2 = bits.Sub64(mod[51], y[51], gteC2)
	_, gteC2 = bits.Sub64(mod[52], y[52], gteC2)
	_, gteC2 = bits.Sub64(mod[53], y[53], gteC2)
	_, gteC2 = bits.Sub64(mod[54], y[54], gteC2)
	_, gteC2 = bits.Sub64(mod[55], y[55], gteC2)
	_, gteC2 = bits.Sub64(mod[56], y[56], gteC2)
	_, gteC2 = bits.Sub64(mod[57], y[57], gteC2)
	_, gteC2 = bits.Sub64(mod[58], y[58], gteC2)
	_, gteC2 = bits.Sub64(mod[59], y[59], gteC2)
	_, gteC2 = bits.Sub64(mod[60], y[60], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [61]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 61; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 61; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled3968(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[62]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[62]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[62]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[62]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
	_, gteC1 = bits.Sub64(mod[51], x[51], gteC1)
	_, gteC1 = bits.Sub64(mod[52], x[52], gteC1)
	_, gteC1 = bits.Sub64(mod[53], x[53], gteC1)
	_, gteC1 = bits.Sub64(mod[54], x[54], gteC1)
	_, gteC1 = bits.Sub64(mod[55], x[55], gteC1)
	_, gteC1 = bits.Sub64(mod[56], x[56], gteC1)
	_, gteC1 = bits.Sub64(mod[57], x[57], gteC1)
	_, gteC1 = bits.Sub64(mod[58], x[58], gteC1)
	_, gteC1 = bits.Sub64(mod[59], x[59], gteC1)
	_, gteC1 = bits.Sub64(mod[60], x[60], gteC1)
	_, gteC1 = bits.Sub64(mod[61], x[61], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)
	_, gteC2 = bits.Sub64(mod[51], y[51], gteC2)
	_, gteC2 = bits.Sub64(mod[52], y[52], gteC2)
	_, gteC2 = bits.Sub64(mod[53], y[53], gteC2)
	_, gteC2 = bits.Sub64(mod[54], y[54], gteC2)
	_, gteC2 = bits.Sub64(mod[55], y[55], gteC2)
	_, gteC2 = bits.Sub64(mod[56], y[56], gteC2)
	_, gteC2 = bits.Sub64(mod[57], y[57], gteC2)
	_, gteC2 = bits.Sub64(mod[58], y[58], gteC2)
	_, gteC2 = bits.Sub64(mod[59], y[59], gteC2)
	_, gteC2 = bits.Sub64(mod[60], y[60], gteC2)
	_, gteC2 = bits.Sub64(mod[61], y[61], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [62]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 62; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 62; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled4032(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[63]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[63]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[63]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[63]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
	_, gteC1 = bits.Sub64(mod[51], x[51], gteC1)
	_, gteC1 = bits.Sub64(mod[52], x[52], gteC1)
	_, gteC1 = bits.Sub64(mod[53], x[53], gteC1)
	_, gteC1 = bits.Sub64(mod[54], x[54], gteC1)
	_, gteC1 = bits.Sub64(mod[55], x[55], gteC1)
	_, gteC1 = bits.Sub64(mod[56], x[56], gteC1)
	_, gteC1 = bits.Sub64(mod[57], x[57], gteC1)
	_, gteC1 = bits.Sub64(mod[58], x[58], gteC1)
	_, gteC1 = bits.Sub64(mod[59], x[59], gteC1)
	_, gteC1 = bits.Sub64(mod[60], x[60], gteC1)
	_, gteC1 = bits.Sub64(mod[61], x[61], gteC1)
	_, gteC1 = bits.Sub64(mod[62], x[62], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)
	_, gteC2 = bits.Sub64(mod[51], y[51], gteC2)
	_, gteC2 = bits.Sub64(mod[52], y[52], gteC2)
	_, gteC2 = bits.Sub64(mod[53], y[53], gteC2)
	_, gteC2 = bits.Sub64(mod[54], y[54], gteC2)
	_, gteC2 = bits.Sub64(mod[55], y[55], gteC2)
	_, gteC2 = bits.Sub64(mod[56], y[56], gteC2)
	_, gteC2 = bits.Sub64(mod[57], y[57], gteC2)
	_, gteC2 = bits.Sub64(mod[58], y[58], gteC2)
	_, gteC2 = bits.Sub64(mod[59], y[59], gteC2)
	_, gteC2 = bits.Sub64(mod[60], y[60], gteC2)
	_, gteC2 = bits.Sub64(mod[61], y[61], gteC2)
	_, gteC2 = bits.Sub64(mod[62], y[62], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [63]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 63; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 63; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}

func AddModNonUnrolled4096(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[64]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[64]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[64]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[64]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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
	_, gteC1 = bits.Sub64(mod[16], x[16], gteC1)
	_, gteC1 = bits.Sub64(mod[17], x[17], gteC1)
	_, gteC1 = bits.Sub64(mod[18], x[18], gteC1)
	_, gteC1 = bits.Sub64(mod[19], x[19], gteC1)
	_, gteC1 = bits.Sub64(mod[20], x[20], gteC1)
	_, gteC1 = bits.Sub64(mod[21], x[21], gteC1)
	_, gteC1 = bits.Sub64(mod[22], x[22], gteC1)
	_, gteC1 = bits.Sub64(mod[23], x[23], gteC1)
	_, gteC1 = bits.Sub64(mod[24], x[24], gteC1)
	_, gteC1 = bits.Sub64(mod[25], x[25], gteC1)
	_, gteC1 = bits.Sub64(mod[26], x[26], gteC1)
	_, gteC1 = bits.Sub64(mod[27], x[27], gteC1)
	_, gteC1 = bits.Sub64(mod[28], x[28], gteC1)
	_, gteC1 = bits.Sub64(mod[29], x[29], gteC1)
	_, gteC1 = bits.Sub64(mod[30], x[30], gteC1)
	_, gteC1 = bits.Sub64(mod[31], x[31], gteC1)
	_, gteC1 = bits.Sub64(mod[32], x[32], gteC1)
	_, gteC1 = bits.Sub64(mod[33], x[33], gteC1)
	_, gteC1 = bits.Sub64(mod[34], x[34], gteC1)
	_, gteC1 = bits.Sub64(mod[35], x[35], gteC1)
	_, gteC1 = bits.Sub64(mod[36], x[36], gteC1)
	_, gteC1 = bits.Sub64(mod[37], x[37], gteC1)
	_, gteC1 = bits.Sub64(mod[38], x[38], gteC1)
	_, gteC1 = bits.Sub64(mod[39], x[39], gteC1)
	_, gteC1 = bits.Sub64(mod[40], x[40], gteC1)
	_, gteC1 = bits.Sub64(mod[41], x[41], gteC1)
	_, gteC1 = bits.Sub64(mod[42], x[42], gteC1)
	_, gteC1 = bits.Sub64(mod[43], x[43], gteC1)
	_, gteC1 = bits.Sub64(mod[44], x[44], gteC1)
	_, gteC1 = bits.Sub64(mod[45], x[45], gteC1)
	_, gteC1 = bits.Sub64(mod[46], x[46], gteC1)
	_, gteC1 = bits.Sub64(mod[47], x[47], gteC1)
	_, gteC1 = bits.Sub64(mod[48], x[48], gteC1)
	_, gteC1 = bits.Sub64(mod[49], x[49], gteC1)
	_, gteC1 = bits.Sub64(mod[50], x[50], gteC1)
	_, gteC1 = bits.Sub64(mod[51], x[51], gteC1)
	_, gteC1 = bits.Sub64(mod[52], x[52], gteC1)
	_, gteC1 = bits.Sub64(mod[53], x[53], gteC1)
	_, gteC1 = bits.Sub64(mod[54], x[54], gteC1)
	_, gteC1 = bits.Sub64(mod[55], x[55], gteC1)
	_, gteC1 = bits.Sub64(mod[56], x[56], gteC1)
	_, gteC1 = bits.Sub64(mod[57], x[57], gteC1)
	_, gteC1 = bits.Sub64(mod[58], x[58], gteC1)
	_, gteC1 = bits.Sub64(mod[59], x[59], gteC1)
	_, gteC1 = bits.Sub64(mod[60], x[60], gteC1)
	_, gteC1 = bits.Sub64(mod[61], x[61], gteC1)
	_, gteC1 = bits.Sub64(mod[62], x[62], gteC1)
	_, gteC1 = bits.Sub64(mod[63], x[63], gteC1)
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
	_, gteC2 = bits.Sub64(mod[16], y[16], gteC2)
	_, gteC2 = bits.Sub64(mod[17], y[17], gteC2)
	_, gteC2 = bits.Sub64(mod[18], y[18], gteC2)
	_, gteC2 = bits.Sub64(mod[19], y[19], gteC2)
	_, gteC2 = bits.Sub64(mod[20], y[20], gteC2)
	_, gteC2 = bits.Sub64(mod[21], y[21], gteC2)
	_, gteC2 = bits.Sub64(mod[22], y[22], gteC2)
	_, gteC2 = bits.Sub64(mod[23], y[23], gteC2)
	_, gteC2 = bits.Sub64(mod[24], y[24], gteC2)
	_, gteC2 = bits.Sub64(mod[25], y[25], gteC2)
	_, gteC2 = bits.Sub64(mod[26], y[26], gteC2)
	_, gteC2 = bits.Sub64(mod[27], y[27], gteC2)
	_, gteC2 = bits.Sub64(mod[28], y[28], gteC2)
	_, gteC2 = bits.Sub64(mod[29], y[29], gteC2)
	_, gteC2 = bits.Sub64(mod[30], y[30], gteC2)
	_, gteC2 = bits.Sub64(mod[31], y[31], gteC2)
	_, gteC2 = bits.Sub64(mod[32], y[32], gteC2)
	_, gteC2 = bits.Sub64(mod[33], y[33], gteC2)
	_, gteC2 = bits.Sub64(mod[34], y[34], gteC2)
	_, gteC2 = bits.Sub64(mod[35], y[35], gteC2)
	_, gteC2 = bits.Sub64(mod[36], y[36], gteC2)
	_, gteC2 = bits.Sub64(mod[37], y[37], gteC2)
	_, gteC2 = bits.Sub64(mod[38], y[38], gteC2)
	_, gteC2 = bits.Sub64(mod[39], y[39], gteC2)
	_, gteC2 = bits.Sub64(mod[40], y[40], gteC2)
	_, gteC2 = bits.Sub64(mod[41], y[41], gteC2)
	_, gteC2 = bits.Sub64(mod[42], y[42], gteC2)
	_, gteC2 = bits.Sub64(mod[43], y[43], gteC2)
	_, gteC2 = bits.Sub64(mod[44], y[44], gteC2)
	_, gteC2 = bits.Sub64(mod[45], y[45], gteC2)
	_, gteC2 = bits.Sub64(mod[46], y[46], gteC2)
	_, gteC2 = bits.Sub64(mod[47], y[47], gteC2)
	_, gteC2 = bits.Sub64(mod[48], y[48], gteC2)
	_, gteC2 = bits.Sub64(mod[49], y[49], gteC2)
	_, gteC2 = bits.Sub64(mod[50], y[50], gteC2)
	_, gteC2 = bits.Sub64(mod[51], y[51], gteC2)
	_, gteC2 = bits.Sub64(mod[52], y[52], gteC2)
	_, gteC2 = bits.Sub64(mod[53], y[53], gteC2)
	_, gteC2 = bits.Sub64(mod[54], y[54], gteC2)
	_, gteC2 = bits.Sub64(mod[55], y[55], gteC2)
	_, gteC2 = bits.Sub64(mod[56], y[56], gteC2)
	_, gteC2 = bits.Sub64(mod[57], y[57], gteC2)
	_, gteC2 = bits.Sub64(mod[58], y[58], gteC2)
	_, gteC2 = bits.Sub64(mod[59], y[59], gteC2)
	_, gteC2 = bits.Sub64(mod[60], y[60], gteC2)
	_, gteC2 = bits.Sub64(mod[61], y[61], gteC2)
	_, gteC2 = bits.Sub64(mod[62], y[62], gteC2)
	_, gteC2 = bits.Sub64(mod[63], y[63], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := [64]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 64; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 64; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} else {
		panic("not worst case performance")
	}
	return nil
}
