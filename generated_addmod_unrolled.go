package mont_arith

import (
	"math/bits"
	"unsafe"
)

// TODO check unrolled speed
func AddModUnrolled64(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[1]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[1]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[1]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 1)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled128(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 2)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled192(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[3]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 3)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled256(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[4]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 4)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled320(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[5]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[5]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[5]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[5]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 5)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled384(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 6)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled448(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[7]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 7)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled512(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[8]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[8]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[8]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[8]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 8)
	tmp[0], c = bits.Add64(x[0], y[0], c)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled576(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
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
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled640(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[10]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
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
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled704(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[11]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[11]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[11]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[11]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
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
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled768(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[12]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[12]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[12]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
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
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled832(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[13]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[13]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[13]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[13]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 13)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled896(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[14]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[14]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[14]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[14]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 14)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled960(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[15]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[15]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[15]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[15]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 15)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1024(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[16]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[16]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[16]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[16]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 16)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1088(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[17]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[17]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[17]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[17]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 17)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1152(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[18]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[18]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[18]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[18]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 18)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1216(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[19]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[19]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[19]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[19]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 19)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1280(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[20]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[20]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[20]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[20]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 20)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1344(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[21]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[21]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[21]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[21]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 21)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1408(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[22]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[22]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[22]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[22]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 22)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1472(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[23]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[23]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[23]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[23]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 23)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1536(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[24]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[24]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[24]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[24]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 24)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1600(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[25]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[25]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[25]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[25]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 25)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1664(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[26]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[26]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[26]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[26]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 26)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1728(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[27]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[27]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[27]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[27]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 27)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1792(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[28]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[28]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[28]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[28]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 28)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1856(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[29]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[29]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[29]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[29]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 29)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1920(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[30]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[30]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[30]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[30]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 30)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled1984(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[31]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[31]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[31]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[31]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 31)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2048(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[32]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[32]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[32]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[32]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 32)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2112(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[33]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[33]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[33]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[33]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 33)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2176(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[34]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[34]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[34]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[34]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 34)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2240(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[35]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[35]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[35]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[35]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 35)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2304(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[36]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[36]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[36]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[36]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 36)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2368(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[37]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[37]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[37]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[37]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 37)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2432(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[38]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[38]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[38]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[38]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 38)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2496(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[39]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[39]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[39]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[39]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 39)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2560(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[40]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[40]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[40]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[40]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 40)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2624(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[41]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[41]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[41]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[41]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 41)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2688(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[42]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[42]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[42]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[42]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 42)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2752(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[43]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[43]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[43]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[43]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 43)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2816(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[44]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[44]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[44]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[44]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 44)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2880(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[45]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[45]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[45]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[45]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 45)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled2944(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[46]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[46]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[46]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[46]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 46)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3008(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[47]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[47]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[47]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[47]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 47)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3072(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[48]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[48]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[48]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[48]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 48)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3136(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[49]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[49]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[49]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[49]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 49)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3200(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[50]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[50]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[50]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[50]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 50)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3264(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[51]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[51]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[51]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[51]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 51)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3328(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[52]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[52]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[52]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[52]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 52)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	tmp[51], c = bits.Add64(x[51], y[51], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)
	z[51], c1 = bits.Sub64(tmp[51], mod[51], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3392(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[53]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[53]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[53]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[53]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 53)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	tmp[51], c = bits.Add64(x[51], y[51], c)
	tmp[52], c = bits.Add64(x[52], y[52], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)
	z[51], c1 = bits.Sub64(tmp[51], mod[51], c1)
	z[52], c1 = bits.Sub64(tmp[52], mod[52], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3456(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[54]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[54]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[54]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[54]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 54)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	tmp[51], c = bits.Add64(x[51], y[51], c)
	tmp[52], c = bits.Add64(x[52], y[52], c)
	tmp[53], c = bits.Add64(x[53], y[53], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)
	z[51], c1 = bits.Sub64(tmp[51], mod[51], c1)
	z[52], c1 = bits.Sub64(tmp[52], mod[52], c1)
	z[53], c1 = bits.Sub64(tmp[53], mod[53], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3520(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[55]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[55]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[55]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[55]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 55)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	tmp[51], c = bits.Add64(x[51], y[51], c)
	tmp[52], c = bits.Add64(x[52], y[52], c)
	tmp[53], c = bits.Add64(x[53], y[53], c)
	tmp[54], c = bits.Add64(x[54], y[54], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)
	z[51], c1 = bits.Sub64(tmp[51], mod[51], c1)
	z[52], c1 = bits.Sub64(tmp[52], mod[52], c1)
	z[53], c1 = bits.Sub64(tmp[53], mod[53], c1)
	z[54], c1 = bits.Sub64(tmp[54], mod[54], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3584(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[56]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[56]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[56]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[56]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 56)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	tmp[51], c = bits.Add64(x[51], y[51], c)
	tmp[52], c = bits.Add64(x[52], y[52], c)
	tmp[53], c = bits.Add64(x[53], y[53], c)
	tmp[54], c = bits.Add64(x[54], y[54], c)
	tmp[55], c = bits.Add64(x[55], y[55], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)
	z[51], c1 = bits.Sub64(tmp[51], mod[51], c1)
	z[52], c1 = bits.Sub64(tmp[52], mod[52], c1)
	z[53], c1 = bits.Sub64(tmp[53], mod[53], c1)
	z[54], c1 = bits.Sub64(tmp[54], mod[54], c1)
	z[55], c1 = bits.Sub64(tmp[55], mod[55], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3648(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[57]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[57]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[57]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[57]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 57)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	tmp[51], c = bits.Add64(x[51], y[51], c)
	tmp[52], c = bits.Add64(x[52], y[52], c)
	tmp[53], c = bits.Add64(x[53], y[53], c)
	tmp[54], c = bits.Add64(x[54], y[54], c)
	tmp[55], c = bits.Add64(x[55], y[55], c)
	tmp[56], c = bits.Add64(x[56], y[56], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)
	z[51], c1 = bits.Sub64(tmp[51], mod[51], c1)
	z[52], c1 = bits.Sub64(tmp[52], mod[52], c1)
	z[53], c1 = bits.Sub64(tmp[53], mod[53], c1)
	z[54], c1 = bits.Sub64(tmp[54], mod[54], c1)
	z[55], c1 = bits.Sub64(tmp[55], mod[55], c1)
	z[56], c1 = bits.Sub64(tmp[56], mod[56], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3712(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[58]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[58]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[58]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[58]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 58)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	tmp[51], c = bits.Add64(x[51], y[51], c)
	tmp[52], c = bits.Add64(x[52], y[52], c)
	tmp[53], c = bits.Add64(x[53], y[53], c)
	tmp[54], c = bits.Add64(x[54], y[54], c)
	tmp[55], c = bits.Add64(x[55], y[55], c)
	tmp[56], c = bits.Add64(x[56], y[56], c)
	tmp[57], c = bits.Add64(x[57], y[57], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)
	z[51], c1 = bits.Sub64(tmp[51], mod[51], c1)
	z[52], c1 = bits.Sub64(tmp[52], mod[52], c1)
	z[53], c1 = bits.Sub64(tmp[53], mod[53], c1)
	z[54], c1 = bits.Sub64(tmp[54], mod[54], c1)
	z[55], c1 = bits.Sub64(tmp[55], mod[55], c1)
	z[56], c1 = bits.Sub64(tmp[56], mod[56], c1)
	z[57], c1 = bits.Sub64(tmp[57], mod[57], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3776(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[59]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[59]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[59]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[59]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 59)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	tmp[51], c = bits.Add64(x[51], y[51], c)
	tmp[52], c = bits.Add64(x[52], y[52], c)
	tmp[53], c = bits.Add64(x[53], y[53], c)
	tmp[54], c = bits.Add64(x[54], y[54], c)
	tmp[55], c = bits.Add64(x[55], y[55], c)
	tmp[56], c = bits.Add64(x[56], y[56], c)
	tmp[57], c = bits.Add64(x[57], y[57], c)
	tmp[58], c = bits.Add64(x[58], y[58], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)
	z[51], c1 = bits.Sub64(tmp[51], mod[51], c1)
	z[52], c1 = bits.Sub64(tmp[52], mod[52], c1)
	z[53], c1 = bits.Sub64(tmp[53], mod[53], c1)
	z[54], c1 = bits.Sub64(tmp[54], mod[54], c1)
	z[55], c1 = bits.Sub64(tmp[55], mod[55], c1)
	z[56], c1 = bits.Sub64(tmp[56], mod[56], c1)
	z[57], c1 = bits.Sub64(tmp[57], mod[57], c1)
	z[58], c1 = bits.Sub64(tmp[58], mod[58], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3840(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[60]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[60]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[60]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[60]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 60)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	tmp[51], c = bits.Add64(x[51], y[51], c)
	tmp[52], c = bits.Add64(x[52], y[52], c)
	tmp[53], c = bits.Add64(x[53], y[53], c)
	tmp[54], c = bits.Add64(x[54], y[54], c)
	tmp[55], c = bits.Add64(x[55], y[55], c)
	tmp[56], c = bits.Add64(x[56], y[56], c)
	tmp[57], c = bits.Add64(x[57], y[57], c)
	tmp[58], c = bits.Add64(x[58], y[58], c)
	tmp[59], c = bits.Add64(x[59], y[59], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)
	z[51], c1 = bits.Sub64(tmp[51], mod[51], c1)
	z[52], c1 = bits.Sub64(tmp[52], mod[52], c1)
	z[53], c1 = bits.Sub64(tmp[53], mod[53], c1)
	z[54], c1 = bits.Sub64(tmp[54], mod[54], c1)
	z[55], c1 = bits.Sub64(tmp[55], mod[55], c1)
	z[56], c1 = bits.Sub64(tmp[56], mod[56], c1)
	z[57], c1 = bits.Sub64(tmp[57], mod[57], c1)
	z[58], c1 = bits.Sub64(tmp[58], mod[58], c1)
	z[59], c1 = bits.Sub64(tmp[59], mod[59], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3904(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[61]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[61]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[61]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[61]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 61)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	tmp[51], c = bits.Add64(x[51], y[51], c)
	tmp[52], c = bits.Add64(x[52], y[52], c)
	tmp[53], c = bits.Add64(x[53], y[53], c)
	tmp[54], c = bits.Add64(x[54], y[54], c)
	tmp[55], c = bits.Add64(x[55], y[55], c)
	tmp[56], c = bits.Add64(x[56], y[56], c)
	tmp[57], c = bits.Add64(x[57], y[57], c)
	tmp[58], c = bits.Add64(x[58], y[58], c)
	tmp[59], c = bits.Add64(x[59], y[59], c)
	tmp[60], c = bits.Add64(x[60], y[60], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)
	z[51], c1 = bits.Sub64(tmp[51], mod[51], c1)
	z[52], c1 = bits.Sub64(tmp[52], mod[52], c1)
	z[53], c1 = bits.Sub64(tmp[53], mod[53], c1)
	z[54], c1 = bits.Sub64(tmp[54], mod[54], c1)
	z[55], c1 = bits.Sub64(tmp[55], mod[55], c1)
	z[56], c1 = bits.Sub64(tmp[56], mod[56], c1)
	z[57], c1 = bits.Sub64(tmp[57], mod[57], c1)
	z[58], c1 = bits.Sub64(tmp[58], mod[58], c1)
	z[59], c1 = bits.Sub64(tmp[59], mod[59], c1)
	z[60], c1 = bits.Sub64(tmp[60], mod[60], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled3968(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[62]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[62]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[62]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[62]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 62)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	tmp[51], c = bits.Add64(x[51], y[51], c)
	tmp[52], c = bits.Add64(x[52], y[52], c)
	tmp[53], c = bits.Add64(x[53], y[53], c)
	tmp[54], c = bits.Add64(x[54], y[54], c)
	tmp[55], c = bits.Add64(x[55], y[55], c)
	tmp[56], c = bits.Add64(x[56], y[56], c)
	tmp[57], c = bits.Add64(x[57], y[57], c)
	tmp[58], c = bits.Add64(x[58], y[58], c)
	tmp[59], c = bits.Add64(x[59], y[59], c)
	tmp[60], c = bits.Add64(x[60], y[60], c)
	tmp[61], c = bits.Add64(x[61], y[61], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)
	z[51], c1 = bits.Sub64(tmp[51], mod[51], c1)
	z[52], c1 = bits.Sub64(tmp[52], mod[52], c1)
	z[53], c1 = bits.Sub64(tmp[53], mod[53], c1)
	z[54], c1 = bits.Sub64(tmp[54], mod[54], c1)
	z[55], c1 = bits.Sub64(tmp[55], mod[55], c1)
	z[56], c1 = bits.Sub64(tmp[56], mod[56], c1)
	z[57], c1 = bits.Sub64(tmp[57], mod[57], c1)
	z[58], c1 = bits.Sub64(tmp[58], mod[58], c1)
	z[59], c1 = bits.Sub64(tmp[59], mod[59], c1)
	z[60], c1 = bits.Sub64(tmp[60], mod[60], c1)
	z[61], c1 = bits.Sub64(tmp[61], mod[61], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled4032(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[63]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[63]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[63]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[63]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 63)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	tmp[51], c = bits.Add64(x[51], y[51], c)
	tmp[52], c = bits.Add64(x[52], y[52], c)
	tmp[53], c = bits.Add64(x[53], y[53], c)
	tmp[54], c = bits.Add64(x[54], y[54], c)
	tmp[55], c = bits.Add64(x[55], y[55], c)
	tmp[56], c = bits.Add64(x[56], y[56], c)
	tmp[57], c = bits.Add64(x[57], y[57], c)
	tmp[58], c = bits.Add64(x[58], y[58], c)
	tmp[59], c = bits.Add64(x[59], y[59], c)
	tmp[60], c = bits.Add64(x[60], y[60], c)
	tmp[61], c = bits.Add64(x[61], y[61], c)
	tmp[62], c = bits.Add64(x[62], y[62], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)
	z[51], c1 = bits.Sub64(tmp[51], mod[51], c1)
	z[52], c1 = bits.Sub64(tmp[52], mod[52], c1)
	z[53], c1 = bits.Sub64(tmp[53], mod[53], c1)
	z[54], c1 = bits.Sub64(tmp[54], mod[54], c1)
	z[55], c1 = bits.Sub64(tmp[55], mod[55], c1)
	z[56], c1 = bits.Sub64(tmp[56], mod[56], c1)
	z[57], c1 = bits.Sub64(tmp[57], mod[57], c1)
	z[58], c1 = bits.Sub64(tmp[58], mod[58], c1)
	z[59], c1 = bits.Sub64(tmp[59], mod[59], c1)
	z[60], c1 = bits.Sub64(tmp[60], mod[60], c1)
	z[61], c1 = bits.Sub64(tmp[61], mod[61], c1)
	z[62], c1 = bits.Sub64(tmp[62], mod[62], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}

// TODO check unrolled speed
func AddModUnrolled4096(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[64]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[64]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[64]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[64]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0
	tmp := make([]uint64, 64)
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
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	tmp[20], c = bits.Add64(x[20], y[20], c)
	tmp[21], c = bits.Add64(x[21], y[21], c)
	tmp[22], c = bits.Add64(x[22], y[22], c)
	tmp[23], c = bits.Add64(x[23], y[23], c)
	tmp[24], c = bits.Add64(x[24], y[24], c)
	tmp[25], c = bits.Add64(x[25], y[25], c)
	tmp[26], c = bits.Add64(x[26], y[26], c)
	tmp[27], c = bits.Add64(x[27], y[27], c)
	tmp[28], c = bits.Add64(x[28], y[28], c)
	tmp[29], c = bits.Add64(x[29], y[29], c)
	tmp[30], c = bits.Add64(x[30], y[30], c)
	tmp[31], c = bits.Add64(x[31], y[31], c)
	tmp[32], c = bits.Add64(x[32], y[32], c)
	tmp[33], c = bits.Add64(x[33], y[33], c)
	tmp[34], c = bits.Add64(x[34], y[34], c)
	tmp[35], c = bits.Add64(x[35], y[35], c)
	tmp[36], c = bits.Add64(x[36], y[36], c)
	tmp[37], c = bits.Add64(x[37], y[37], c)
	tmp[38], c = bits.Add64(x[38], y[38], c)
	tmp[39], c = bits.Add64(x[39], y[39], c)
	tmp[40], c = bits.Add64(x[40], y[40], c)
	tmp[41], c = bits.Add64(x[41], y[41], c)
	tmp[42], c = bits.Add64(x[42], y[42], c)
	tmp[43], c = bits.Add64(x[43], y[43], c)
	tmp[44], c = bits.Add64(x[44], y[44], c)
	tmp[45], c = bits.Add64(x[45], y[45], c)
	tmp[46], c = bits.Add64(x[46], y[46], c)
	tmp[47], c = bits.Add64(x[47], y[47], c)
	tmp[48], c = bits.Add64(x[48], y[48], c)
	tmp[49], c = bits.Add64(x[49], y[49], c)
	tmp[50], c = bits.Add64(x[50], y[50], c)
	tmp[51], c = bits.Add64(x[51], y[51], c)
	tmp[52], c = bits.Add64(x[52], y[52], c)
	tmp[53], c = bits.Add64(x[53], y[53], c)
	tmp[54], c = bits.Add64(x[54], y[54], c)
	tmp[55], c = bits.Add64(x[55], y[55], c)
	tmp[56], c = bits.Add64(x[56], y[56], c)
	tmp[57], c = bits.Add64(x[57], y[57], c)
	tmp[58], c = bits.Add64(x[58], y[58], c)
	tmp[59], c = bits.Add64(x[59], y[59], c)
	tmp[60], c = bits.Add64(x[60], y[60], c)
	tmp[61], c = bits.Add64(x[61], y[61], c)
	tmp[62], c = bits.Add64(x[62], y[62], c)
	tmp[63], c = bits.Add64(x[63], y[63], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], c1)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Sub64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Sub64(tmp[21], mod[21], c1)
	z[22], c1 = bits.Sub64(tmp[22], mod[22], c1)
	z[23], c1 = bits.Sub64(tmp[23], mod[23], c1)
	z[24], c1 = bits.Sub64(tmp[24], mod[24], c1)
	z[25], c1 = bits.Sub64(tmp[25], mod[25], c1)
	z[26], c1 = bits.Sub64(tmp[26], mod[26], c1)
	z[27], c1 = bits.Sub64(tmp[27], mod[27], c1)
	z[28], c1 = bits.Sub64(tmp[28], mod[28], c1)
	z[29], c1 = bits.Sub64(tmp[29], mod[29], c1)
	z[30], c1 = bits.Sub64(tmp[30], mod[30], c1)
	z[31], c1 = bits.Sub64(tmp[31], mod[31], c1)
	z[32], c1 = bits.Sub64(tmp[32], mod[32], c1)
	z[33], c1 = bits.Sub64(tmp[33], mod[33], c1)
	z[34], c1 = bits.Sub64(tmp[34], mod[34], c1)
	z[35], c1 = bits.Sub64(tmp[35], mod[35], c1)
	z[36], c1 = bits.Sub64(tmp[36], mod[36], c1)
	z[37], c1 = bits.Sub64(tmp[37], mod[37], c1)
	z[38], c1 = bits.Sub64(tmp[38], mod[38], c1)
	z[39], c1 = bits.Sub64(tmp[39], mod[39], c1)
	z[40], c1 = bits.Sub64(tmp[40], mod[40], c1)
	z[41], c1 = bits.Sub64(tmp[41], mod[41], c1)
	z[42], c1 = bits.Sub64(tmp[42], mod[42], c1)
	z[43], c1 = bits.Sub64(tmp[43], mod[43], c1)
	z[44], c1 = bits.Sub64(tmp[44], mod[44], c1)
	z[45], c1 = bits.Sub64(tmp[45], mod[45], c1)
	z[46], c1 = bits.Sub64(tmp[46], mod[46], c1)
	z[47], c1 = bits.Sub64(tmp[47], mod[47], c1)
	z[48], c1 = bits.Sub64(tmp[48], mod[48], c1)
	z[49], c1 = bits.Sub64(tmp[49], mod[49], c1)
	z[50], c1 = bits.Sub64(tmp[50], mod[50], c1)
	z[51], c1 = bits.Sub64(tmp[51], mod[51], c1)
	z[52], c1 = bits.Sub64(tmp[52], mod[52], c1)
	z[53], c1 = bits.Sub64(tmp[53], mod[53], c1)
	z[54], c1 = bits.Sub64(tmp[54], mod[54], c1)
	z[55], c1 = bits.Sub64(tmp[55], mod[55], c1)
	z[56], c1 = bits.Sub64(tmp[56], mod[56], c1)
	z[57], c1 = bits.Sub64(tmp[57], mod[57], c1)
	z[58], c1 = bits.Sub64(tmp[58], mod[58], c1)
	z[59], c1 = bits.Sub64(tmp[59], mod[59], c1)
	z[60], c1 = bits.Sub64(tmp[60], mod[60], c1)
	z[61], c1 = bits.Sub64(tmp[61], mod[61], c1)
	z[62], c1 = bits.Sub64(tmp[62], mod[62], c1)
	z[63], c1 = bits.Sub64(tmp[63], mod[63], c1)

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	} /* else {
	    panic("not worst case performance")
	}*/
	return nil
}
