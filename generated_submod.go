package mont_arith

import (
	"errors"
	"fmt"
	"math/bits"
	"unsafe"
)

func SubModNonUnrolled64(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[1]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[1]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[1]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c, c1 uint64
	tmp := [1]uint64{0}
	tmp[0], c = bits.Sub64(x[0], y[0], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], c1)

	if c == 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func SubModNonUnrolled128(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var gteC1, gteC2 uint64
	_, gteC1 = bits.Sub64(mod[0], x[0], gteC1)
	_, gteC1 = bits.Sub64(mod[1], x[1], gteC1)
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c, c1 uint64
	tmp := [2]uint64{0, 0}
	tmp[0], c = bits.Sub64(x[0], y[0], c)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)

	if c == 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func SubModNonUnrolled192(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[3]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	var c, c1 uint64
	tmp := [3]uint64{0, 0, 0}
	tmp[0], c = bits.Sub64(x[0], y[0], c)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	tmp[2], c = bits.Sub64(x[2], y[2], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Add64(tmp[2], mod[2], c1)

	if c == 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func SubModNonUnrolled256(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[4]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

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

	var c, c1 uint64
	tmp := [4]uint64{0, 0, 0, 0}
	tmp[0], c = bits.Sub64(x[0], y[0], c)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	tmp[2], c = bits.Sub64(x[2], y[2], c)
	tmp[3], c = bits.Sub64(x[3], y[3], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Add64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Add64(tmp[3], mod[3], c1)

	if c == 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func SubModNonUnrolled320(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
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
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], x[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], x[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], x[4], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c, c1 uint64
	tmp := [5]uint64{0, 0, 0, 0, 0}
	tmp[0], c = bits.Sub64(x[0], y[0], c)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	tmp[2], c = bits.Sub64(x[2], y[2], c)
	tmp[3], c = bits.Sub64(x[3], y[3], c)
	tmp[4], c = bits.Sub64(x[4], y[4], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Add64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Add64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Add64(tmp[4], mod[4], c1)

	if c == 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func SubModNonUnrolled384(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
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
	_, gteC2 = bits.Sub64(mod[0], x[0], gteC2)
	_, gteC2 = bits.Sub64(mod[1], x[1], gteC2)
	_, gteC2 = bits.Sub64(mod[2], x[2], gteC2)
	_, gteC2 = bits.Sub64(mod[3], x[3], gteC2)
	_, gteC2 = bits.Sub64(mod[4], x[4], gteC2)
	_, gteC2 = bits.Sub64(mod[5], x[5], gteC2)

	if gteC1 != 0 || gteC2 != 0 {
		return errors.New(fmt.Sprintf("input gte modulus"))
	}

	var c, c1 uint64
	tmp := [6]uint64{0, 0, 0, 0, 0, 0}
	tmp[0], c = bits.Sub64(x[0], y[0], c)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	tmp[2], c = bits.Sub64(x[2], y[2], c)
	tmp[3], c = bits.Sub64(x[3], y[3], c)
	tmp[4], c = bits.Sub64(x[4], y[4], c)
	tmp[5], c = bits.Sub64(x[5], y[5], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Add64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Add64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Add64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Add64(tmp[5], mod[5], c1)

	if c == 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func SubModNonUnrolled448(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
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

	var c, c1 uint64
	tmp := [7]uint64{0, 0, 0, 0, 0, 0, 0}
	tmp[0], c = bits.Sub64(x[0], y[0], c)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	tmp[2], c = bits.Sub64(x[2], y[2], c)
	tmp[3], c = bits.Sub64(x[3], y[3], c)
	tmp[4], c = bits.Sub64(x[4], y[4], c)
	tmp[5], c = bits.Sub64(x[5], y[5], c)
	tmp[6], c = bits.Sub64(x[6], y[6], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Add64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Add64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Add64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Add64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Add64(tmp[6], mod[6], c1)

	if c == 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func SubModNonUnrolled512(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
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

	var c, c1 uint64
	tmp := [8]uint64{0, 0, 0, 0, 0, 0, 0, 0}
	tmp[0], c = bits.Sub64(x[0], y[0], c)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	tmp[2], c = bits.Sub64(x[2], y[2], c)
	tmp[3], c = bits.Sub64(x[3], y[3], c)
	tmp[4], c = bits.Sub64(x[4], y[4], c)
	tmp[5], c = bits.Sub64(x[5], y[5], c)
	tmp[6], c = bits.Sub64(x[6], y[6], c)
	tmp[7], c = bits.Sub64(x[7], y[7], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Add64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Add64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Add64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Add64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Add64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Add64(tmp[7], mod[7], c1)

	if c == 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func SubModNonUnrolled576(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
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

	var c, c1 uint64
	tmp := [9]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	tmp[0], c = bits.Sub64(x[0], y[0], c)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	tmp[2], c = bits.Sub64(x[2], y[2], c)
	tmp[3], c = bits.Sub64(x[3], y[3], c)
	tmp[4], c = bits.Sub64(x[4], y[4], c)
	tmp[5], c = bits.Sub64(x[5], y[5], c)
	tmp[6], c = bits.Sub64(x[6], y[6], c)
	tmp[7], c = bits.Sub64(x[7], y[7], c)
	tmp[8], c = bits.Sub64(x[8], y[8], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Add64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Add64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Add64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Add64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Add64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Add64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Add64(tmp[8], mod[8], c1)

	if c == 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func SubModNonUnrolled640(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
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

	var c, c1 uint64
	tmp := [10]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	tmp[0], c = bits.Sub64(x[0], y[0], c)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	tmp[2], c = bits.Sub64(x[2], y[2], c)
	tmp[3], c = bits.Sub64(x[3], y[3], c)
	tmp[4], c = bits.Sub64(x[4], y[4], c)
	tmp[5], c = bits.Sub64(x[5], y[5], c)
	tmp[6], c = bits.Sub64(x[6], y[6], c)
	tmp[7], c = bits.Sub64(x[7], y[7], c)
	tmp[8], c = bits.Sub64(x[8], y[8], c)
	tmp[9], c = bits.Sub64(x[9], y[9], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Add64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Add64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Add64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Add64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Add64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Add64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Add64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Add64(tmp[9], mod[9], c1)

	if c == 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func SubModNonUnrolled704(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
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

	var c, c1 uint64
	tmp := [11]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	tmp[0], c = bits.Sub64(x[0], y[0], c)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	tmp[2], c = bits.Sub64(x[2], y[2], c)
	tmp[3], c = bits.Sub64(x[3], y[3], c)
	tmp[4], c = bits.Sub64(x[4], y[4], c)
	tmp[5], c = bits.Sub64(x[5], y[5], c)
	tmp[6], c = bits.Sub64(x[6], y[6], c)
	tmp[7], c = bits.Sub64(x[7], y[7], c)
	tmp[8], c = bits.Sub64(x[8], y[8], c)
	tmp[9], c = bits.Sub64(x[9], y[9], c)
	tmp[10], c = bits.Sub64(x[10], y[10], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Add64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Add64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Add64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Add64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Add64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Add64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Add64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Add64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Add64(tmp[10], mod[10], c1)

	if c == 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func SubModNonUnrolled768(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
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

	var c, c1 uint64
	tmp := [12]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	tmp[0], c = bits.Sub64(x[0], y[0], c)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	tmp[2], c = bits.Sub64(x[2], y[2], c)
	tmp[3], c = bits.Sub64(x[3], y[3], c)
	tmp[4], c = bits.Sub64(x[4], y[4], c)
	tmp[5], c = bits.Sub64(x[5], y[5], c)
	tmp[6], c = bits.Sub64(x[6], y[6], c)
	tmp[7], c = bits.Sub64(x[7], y[7], c)
	tmp[8], c = bits.Sub64(x[8], y[8], c)
	tmp[9], c = bits.Sub64(x[9], y[9], c)
	tmp[10], c = bits.Sub64(x[10], y[10], c)
	tmp[11], c = bits.Sub64(x[11], y[11], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Add64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Add64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Add64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Add64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Add64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Add64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Add64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Add64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Add64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Add64(tmp[11], mod[11], c1)

	if c == 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}
