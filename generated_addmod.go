package mont_arith

import (
	"errors"
	"fmt"
	"math/bits"
	"unsafe"
)

// TODO check unrolled speed
func AddModNonUnrolled64(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[1]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[1]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := f.Modulus

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	tmp := make([]uint64, 1)
	tmp[0], c = bits.Add64(x[0], y[0], c)

	c = 0
	z[0], c = bits.Sub64(tmp[0], mod[0], c)

	// final sub was unnecessary
	if c != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModNonUnrolled128(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := f.Modulus

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	tmp := make([]uint64, 2)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)

	c = 0
	z[0], c = bits.Sub64(tmp[0], mod[0], c)
	z[1], c = bits.Sub64(tmp[1], mod[1], c)

	// final sub was unnecessary
	if c != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModNonUnrolled192(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := f.Modulus

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], x[2], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	tmp := make([]uint64, 3)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)

	c = 0
	z[0], c = bits.Sub64(tmp[0], mod[0], c)
	z[1], c = bits.Sub64(tmp[1], mod[1], c)
	z[2], c = bits.Sub64(tmp[2], mod[2], c)

	// final sub was unnecessary
	if c != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModNonUnrolled256(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := f.Modulus

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], x[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], x[3], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	tmp := make([]uint64, 4)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)

	c = 0
	z[0], c = bits.Sub64(tmp[0], mod[0], c)
	z[1], c = bits.Sub64(tmp[1], mod[1], c)
	z[2], c = bits.Sub64(tmp[2], mod[2], c)
	z[3], c = bits.Sub64(tmp[3], mod[3], c)

	// final sub was unnecessary
	if c != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModNonUnrolled320(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[5]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[5]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[5]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := f.Modulus

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], x[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], x[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], x[4], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	tmp := make([]uint64, 5)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)

	c = 0
	z[0], c = bits.Sub64(tmp[0], mod[0], c)
	z[1], c = bits.Sub64(tmp[1], mod[1], c)
	z[2], c = bits.Sub64(tmp[2], mod[2], c)
	z[3], c = bits.Sub64(tmp[3], mod[3], c)
	z[4], c = bits.Sub64(tmp[4], mod[4], c)

	// final sub was unnecessary
	if c != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModNonUnrolled384(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := f.Modulus

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], x[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], x[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], x[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], x[5], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	tmp := make([]uint64, 6)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)

	c = 0
	z[0], c = bits.Sub64(tmp[0], mod[0], c)
	z[1], c = bits.Sub64(tmp[1], mod[1], c)
	z[2], c = bits.Sub64(tmp[2], mod[2], c)
	z[3], c = bits.Sub64(tmp[3], mod[3], c)
	z[4], c = bits.Sub64(tmp[4], mod[4], c)
	z[5], c = bits.Sub64(tmp[5], mod[5], c)

	// final sub was unnecessary
	if c != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModNonUnrolled448(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := f.Modulus

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC1 = bits.Sub64(mod[6], x[6], gteC1)
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], x[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], x[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], x[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], x[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], x[6], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	tmp := make([]uint64, 7)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)

	c = 0
	z[0], c = bits.Sub64(tmp[0], mod[0], c)
	z[1], c = bits.Sub64(tmp[1], mod[1], c)
	z[2], c = bits.Sub64(tmp[2], mod[2], c)
	z[3], c = bits.Sub64(tmp[3], mod[3], c)
	z[4], c = bits.Sub64(tmp[4], mod[4], c)
	z[5], c = bits.Sub64(tmp[5], mod[5], c)
	z[6], c = bits.Sub64(tmp[6], mod[6], c)

	// final sub was unnecessary
	if c != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModNonUnrolled512(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[8]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[8]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[8]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := f.Modulus

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC1 = bits.Sub64(mod[2], x[2], gteC1)
	_, gteC1 = bits.Sub64(mod[3], x[3], gteC1)
	_, gteC1 = bits.Sub64(mod[4], x[4], gteC1)
	_, gteC1 = bits.Sub64(mod[5], x[5], gteC1)
	_, gteC1 = bits.Sub64(mod[6], x[6], gteC1)
	_, gteC1 = bits.Sub64(mod[7], x[7], gteC1)
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], x[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], x[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], x[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], x[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], x[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], x[7], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	tmp := make([]uint64, 8)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)

	c = 0
	z[0], c = bits.Sub64(tmp[0], mod[0], c)
	z[1], c = bits.Sub64(tmp[1], mod[1], c)
	z[2], c = bits.Sub64(tmp[2], mod[2], c)
	z[3], c = bits.Sub64(tmp[3], mod[3], c)
	z[4], c = bits.Sub64(tmp[4], mod[4], c)
	z[5], c = bits.Sub64(tmp[5], mod[5], c)
	z[6], c = bits.Sub64(tmp[6], mod[6], c)
	z[7], c = bits.Sub64(tmp[7], mod[7], c)

	// final sub was unnecessary
	if c != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModNonUnrolled576(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := f.Modulus

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
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], x[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], x[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], x[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], x[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], x[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], x[7], gteC2)
	_, gteC2 = bits.Sub64(mod[8], x[8], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	tmp := make([]uint64, 9)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)

	c = 0
	z[0], c = bits.Sub64(tmp[0], mod[0], c)
	z[1], c = bits.Sub64(tmp[1], mod[1], c)
	z[2], c = bits.Sub64(tmp[2], mod[2], c)
	z[3], c = bits.Sub64(tmp[3], mod[3], c)
	z[4], c = bits.Sub64(tmp[4], mod[4], c)
	z[5], c = bits.Sub64(tmp[5], mod[5], c)
	z[6], c = bits.Sub64(tmp[6], mod[6], c)
	z[7], c = bits.Sub64(tmp[7], mod[7], c)
	z[8], c = bits.Sub64(tmp[8], mod[8], c)

	// final sub was unnecessary
	if c != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModNonUnrolled640(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := f.Modulus

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
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], x[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], x[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], x[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], x[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], x[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], x[7], gteC2)
	_, gteC2 = bits.Sub64(mod[8], x[8], gteC2)
	_, gteC2 = bits.Sub64(mod[9], x[9], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	tmp := make([]uint64, 10)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)

	c = 0
	z[0], c = bits.Sub64(tmp[0], mod[0], c)
	z[1], c = bits.Sub64(tmp[1], mod[1], c)
	z[2], c = bits.Sub64(tmp[2], mod[2], c)
	z[3], c = bits.Sub64(tmp[3], mod[3], c)
	z[4], c = bits.Sub64(tmp[4], mod[4], c)
	z[5], c = bits.Sub64(tmp[5], mod[5], c)
	z[6], c = bits.Sub64(tmp[6], mod[6], c)
	z[7], c = bits.Sub64(tmp[7], mod[7], c)
	z[8], c = bits.Sub64(tmp[8], mod[8], c)
	z[9], c = bits.Sub64(tmp[9], mod[9], c)

	// final sub was unnecessary
	if c != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModNonUnrolled704(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[11]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[11]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[11]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := f.Modulus

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
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], x[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], x[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], x[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], x[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], x[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], x[7], gteC2)
	_, gteC2 = bits.Sub64(mod[8], x[8], gteC2)
	_, gteC2 = bits.Sub64(mod[9], x[9], gteC2)
	_, gteC2 = bits.Sub64(mod[10], x[10], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	tmp := make([]uint64, 11)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)
	tmp[10], c = bits.Add64(x[10], y[10], c)

	c = 0
	z[0], c = bits.Sub64(tmp[0], mod[0], c)
	z[1], c = bits.Sub64(tmp[1], mod[1], c)
	z[2], c = bits.Sub64(tmp[2], mod[2], c)
	z[3], c = bits.Sub64(tmp[3], mod[3], c)
	z[4], c = bits.Sub64(tmp[4], mod[4], c)
	z[5], c = bits.Sub64(tmp[5], mod[5], c)
	z[6], c = bits.Sub64(tmp[6], mod[6], c)
	z[7], c = bits.Sub64(tmp[7], mod[7], c)
	z[8], c = bits.Sub64(tmp[8], mod[8], c)
	z[9], c = bits.Sub64(tmp[9], mod[9], c)
	z[10], c = bits.Sub64(tmp[10], mod[10], c)

	// final sub was unnecessary
	if c != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModNonUnrolled768(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[12]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[12]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := f.Modulus

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
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], x[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], x[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], x[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], x[5], gteC2)
	_, gteC2 = bits.Sub64(mod[6], x[6], gteC2)
	_, gteC2 = bits.Sub64(mod[7], x[7], gteC2)
	_, gteC2 = bits.Sub64(mod[8], x[8], gteC2)
	_, gteC2 = bits.Sub64(mod[9], x[9], gteC2)
	_, gteC2 = bits.Sub64(mod[10], x[10], gteC2)
	_, gteC2 = bits.Sub64(mod[11], x[11], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c uint64 = 0
	tmp := make([]uint64, 12)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)
	tmp[10], c = bits.Add64(x[10], y[10], c)
	tmp[11], c = bits.Add64(x[11], y[11], c)

	c = 0
	z[0], c = bits.Sub64(tmp[0], mod[0], c)
	z[1], c = bits.Sub64(tmp[1], mod[1], c)
	z[2], c = bits.Sub64(tmp[2], mod[2], c)
	z[3], c = bits.Sub64(tmp[3], mod[3], c)
	z[4], c = bits.Sub64(tmp[4], mod[4], c)
	z[5], c = bits.Sub64(tmp[5], mod[5], c)
	z[6], c = bits.Sub64(tmp[6], mod[6], c)
	z[7], c = bits.Sub64(tmp[7], mod[7], c)
	z[8], c = bits.Sub64(tmp[8], mod[8], c)
	z[9], c = bits.Sub64(tmp[9], mod[9], c)
	z[10], c = bits.Sub64(tmp[10], mod[10], c)
	z[11], c = bits.Sub64(tmp[11], mod[11], c)

	// final sub was unnecessary
	if c != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}
