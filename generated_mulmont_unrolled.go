package mont_arith

import (
	"math/bits"
	"unsafe"
)

func MulMontUnrolled64(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[1]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[1]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[1]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [2]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

	C, t[0] = bits.Mul64(x[0], y[0])

	t[1], D = bits.Add64(t[1], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	t[0], C = bits.Add64(t[1], C, 0)
	t[1], _ = bits.Add64(0, D, C)
	z[0], D = bits.Sub64(t[0], mod[0], 0)

	if D != 0 && t[1] == 0 {
		// reduction was not necessary
		copy(z[:], t[:1])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled128(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [3]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)

	t[2], D = bits.Add64(t[2], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	t[1], C = bits.Add64(t[2], C, 0)
	t[2], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)

	t[2], D = bits.Add64(t[2], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	t[1], C = bits.Add64(t[2], C, 0)
	t[2], _ = bits.Add64(0, D, C)
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)

	if D != 0 && t[2] == 0 {
		// reduction was not necessary
		copy(z[:], t[:2])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled192(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[3]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [4]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)

	t[3], D = bits.Add64(t[3], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	t[2], C = bits.Add64(t[3], C, 0)
	t[3], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)

	t[3], D = bits.Add64(t[3], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	t[2], C = bits.Add64(t[3], C, 0)
	t[3], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)

	t[3], D = bits.Add64(t[3], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	t[2], C = bits.Add64(t[3], C, 0)
	t[3], _ = bits.Add64(0, D, C)
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)

	if D != 0 && t[3] == 0 {
		// reduction was not necessary
		copy(z[:], t[:3])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled256(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[4]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [5]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)

	t[4], D = bits.Add64(t[4], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	t[3], C = bits.Add64(t[4], C, 0)
	t[4], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)
	C, t[3] = madd2(x[1], y[3], t[3], C)

	t[4], D = bits.Add64(t[4], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	t[3], C = bits.Add64(t[4], C, 0)
	t[4], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)
	C, t[3] = madd2(x[2], y[3], t[3], C)

	t[4], D = bits.Add64(t[4], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	t[3], C = bits.Add64(t[4], C, 0)
	t[4], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], y[0], t[0])
	C, t[1] = madd2(x[3], y[1], t[1], C)
	C, t[2] = madd2(x[3], y[2], t[2], C)
	C, t[3] = madd2(x[3], y[3], t[3], C)

	t[4], D = bits.Add64(t[4], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	t[3], C = bits.Add64(t[4], C, 0)
	t[4], _ = bits.Add64(0, D, C)
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)

	if D != 0 && t[4] == 0 {
		// reduction was not necessary
		copy(z[:], t[:4])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled320(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[5]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[5]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[5]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[5]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [6]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

	C, t[0] = bits.Mul64(x[0], y[0])
	C, t[1] = madd1(x[0], y[1], C)
	C, t[2] = madd1(x[0], y[2], C)
	C, t[3] = madd1(x[0], y[3], C)
	C, t[4] = madd1(x[0], y[4], C)

	t[5], D = bits.Add64(t[5], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	t[4], C = bits.Add64(t[5], C, 0)
	t[5], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)
	C, t[3] = madd2(x[1], y[3], t[3], C)
	C, t[4] = madd2(x[1], y[4], t[4], C)

	t[5], D = bits.Add64(t[5], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	t[4], C = bits.Add64(t[5], C, 0)
	t[5], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)
	C, t[3] = madd2(x[2], y[3], t[3], C)
	C, t[4] = madd2(x[2], y[4], t[4], C)

	t[5], D = bits.Add64(t[5], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	t[4], C = bits.Add64(t[5], C, 0)
	t[5], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], y[0], t[0])
	C, t[1] = madd2(x[3], y[1], t[1], C)
	C, t[2] = madd2(x[3], y[2], t[2], C)
	C, t[3] = madd2(x[3], y[3], t[3], C)
	C, t[4] = madd2(x[3], y[4], t[4], C)

	t[5], D = bits.Add64(t[5], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	t[4], C = bits.Add64(t[5], C, 0)
	t[5], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[4], y[0], t[0])
	C, t[1] = madd2(x[4], y[1], t[1], C)
	C, t[2] = madd2(x[4], y[2], t[2], C)
	C, t[3] = madd2(x[4], y[3], t[3], C)
	C, t[4] = madd2(x[4], y[4], t[4], C)

	t[5], D = bits.Add64(t[5], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	t[4], C = bits.Add64(t[5], C, 0)
	t[5], _ = bits.Add64(0, D, C)
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)

	if D != 0 && t[5] == 0 {
		// reduction was not necessary
		copy(z[:], t[:5])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled384(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [7]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

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
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	t[5], C = bits.Add64(t[6], C, 0)
	t[6], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)
	C, t[3] = madd2(x[1], y[3], t[3], C)
	C, t[4] = madd2(x[1], y[4], t[4], C)
	C, t[5] = madd2(x[1], y[5], t[5], C)

	t[6], D = bits.Add64(t[6], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	t[5], C = bits.Add64(t[6], C, 0)
	t[6], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)
	C, t[3] = madd2(x[2], y[3], t[3], C)
	C, t[4] = madd2(x[2], y[4], t[4], C)
	C, t[5] = madd2(x[2], y[5], t[5], C)

	t[6], D = bits.Add64(t[6], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	t[5], C = bits.Add64(t[6], C, 0)
	t[6], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], y[0], t[0])
	C, t[1] = madd2(x[3], y[1], t[1], C)
	C, t[2] = madd2(x[3], y[2], t[2], C)
	C, t[3] = madd2(x[3], y[3], t[3], C)
	C, t[4] = madd2(x[3], y[4], t[4], C)
	C, t[5] = madd2(x[3], y[5], t[5], C)

	t[6], D = bits.Add64(t[6], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	t[5], C = bits.Add64(t[6], C, 0)
	t[6], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[4], y[0], t[0])
	C, t[1] = madd2(x[4], y[1], t[1], C)
	C, t[2] = madd2(x[4], y[2], t[2], C)
	C, t[3] = madd2(x[4], y[3], t[3], C)
	C, t[4] = madd2(x[4], y[4], t[4], C)
	C, t[5] = madd2(x[4], y[5], t[5], C)

	t[6], D = bits.Add64(t[6], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	t[5], C = bits.Add64(t[6], C, 0)
	t[6], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[5], y[0], t[0])
	C, t[1] = madd2(x[5], y[1], t[1], C)
	C, t[2] = madd2(x[5], y[2], t[2], C)
	C, t[3] = madd2(x[5], y[3], t[3], C)
	C, t[4] = madd2(x[5], y[4], t[4], C)
	C, t[5] = madd2(x[5], y[5], t[5], C)

	t[6], D = bits.Add64(t[6], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	t[5], C = bits.Add64(t[6], C, 0)
	t[6], _ = bits.Add64(0, D, C)
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)

	if D != 0 && t[6] == 0 {
		// reduction was not necessary
		copy(z[:], t[:6])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled448(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[7]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [8]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

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
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	t[6], C = bits.Add64(t[7], C, 0)
	t[7], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)
	C, t[3] = madd2(x[1], y[3], t[3], C)
	C, t[4] = madd2(x[1], y[4], t[4], C)
	C, t[5] = madd2(x[1], y[5], t[5], C)
	C, t[6] = madd2(x[1], y[6], t[6], C)

	t[7], D = bits.Add64(t[7], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	t[6], C = bits.Add64(t[7], C, 0)
	t[7], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)
	C, t[3] = madd2(x[2], y[3], t[3], C)
	C, t[4] = madd2(x[2], y[4], t[4], C)
	C, t[5] = madd2(x[2], y[5], t[5], C)
	C, t[6] = madd2(x[2], y[6], t[6], C)

	t[7], D = bits.Add64(t[7], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	t[6], C = bits.Add64(t[7], C, 0)
	t[7], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], y[0], t[0])
	C, t[1] = madd2(x[3], y[1], t[1], C)
	C, t[2] = madd2(x[3], y[2], t[2], C)
	C, t[3] = madd2(x[3], y[3], t[3], C)
	C, t[4] = madd2(x[3], y[4], t[4], C)
	C, t[5] = madd2(x[3], y[5], t[5], C)
	C, t[6] = madd2(x[3], y[6], t[6], C)

	t[7], D = bits.Add64(t[7], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	t[6], C = bits.Add64(t[7], C, 0)
	t[7], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[4], y[0], t[0])
	C, t[1] = madd2(x[4], y[1], t[1], C)
	C, t[2] = madd2(x[4], y[2], t[2], C)
	C, t[3] = madd2(x[4], y[3], t[3], C)
	C, t[4] = madd2(x[4], y[4], t[4], C)
	C, t[5] = madd2(x[4], y[5], t[5], C)
	C, t[6] = madd2(x[4], y[6], t[6], C)

	t[7], D = bits.Add64(t[7], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	t[6], C = bits.Add64(t[7], C, 0)
	t[7], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[5], y[0], t[0])
	C, t[1] = madd2(x[5], y[1], t[1], C)
	C, t[2] = madd2(x[5], y[2], t[2], C)
	C, t[3] = madd2(x[5], y[3], t[3], C)
	C, t[4] = madd2(x[5], y[4], t[4], C)
	C, t[5] = madd2(x[5], y[5], t[5], C)
	C, t[6] = madd2(x[5], y[6], t[6], C)

	t[7], D = bits.Add64(t[7], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	t[6], C = bits.Add64(t[7], C, 0)
	t[7], _ = bits.Add64(0, D, C)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[6], y[0], t[0])
	C, t[1] = madd2(x[6], y[1], t[1], C)
	C, t[2] = madd2(x[6], y[2], t[2], C)
	C, t[3] = madd2(x[6], y[3], t[3], C)
	C, t[4] = madd2(x[6], y[4], t[4], C)
	C, t[5] = madd2(x[6], y[5], t[5], C)
	C, t[6] = madd2(x[6], y[6], t[6], C)

	t[7], D = bits.Add64(t[7], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
	C = madd0(m, mod[0], t[0])
	C, t[0] = madd2(m, mod[1], t[1], C)
	C, t[1] = madd2(m, mod[2], t[2], C)
	C, t[2] = madd2(m, mod[3], t[3], C)
	C, t[3] = madd2(m, mod[4], t[4], C)
	C, t[4] = madd2(m, mod[5], t[5], C)
	C, t[5] = madd2(m, mod[6], t[6], C)
	t[6], C = bits.Add64(t[7], C, 0)
	t[7], _ = bits.Add64(0, D, C)
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)
	z[6], D = bits.Sub64(t[6], mod[6], D)

	if D != 0 && t[7] == 0 {
		// reduction was not necessary
		copy(z[:], t[:7])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled512(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[8]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[8]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[8]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[8]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [9]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

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
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)
	C, t[3] = madd2(x[1], y[3], t[3], C)
	C, t[4] = madd2(x[1], y[4], t[4], C)
	C, t[5] = madd2(x[1], y[5], t[5], C)
	C, t[6] = madd2(x[1], y[6], t[6], C)
	C, t[7] = madd2(x[1], y[7], t[7], C)

	t[8], D = bits.Add64(t[8], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)
	C, t[3] = madd2(x[2], y[3], t[3], C)
	C, t[4] = madd2(x[2], y[4], t[4], C)
	C, t[5] = madd2(x[2], y[5], t[5], C)
	C, t[6] = madd2(x[2], y[6], t[6], C)
	C, t[7] = madd2(x[2], y[7], t[7], C)

	t[8], D = bits.Add64(t[8], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], y[0], t[0])
	C, t[1] = madd2(x[3], y[1], t[1], C)
	C, t[2] = madd2(x[3], y[2], t[2], C)
	C, t[3] = madd2(x[3], y[3], t[3], C)
	C, t[4] = madd2(x[3], y[4], t[4], C)
	C, t[5] = madd2(x[3], y[5], t[5], C)
	C, t[6] = madd2(x[3], y[6], t[6], C)
	C, t[7] = madd2(x[3], y[7], t[7], C)

	t[8], D = bits.Add64(t[8], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[4], y[0], t[0])
	C, t[1] = madd2(x[4], y[1], t[1], C)
	C, t[2] = madd2(x[4], y[2], t[2], C)
	C, t[3] = madd2(x[4], y[3], t[3], C)
	C, t[4] = madd2(x[4], y[4], t[4], C)
	C, t[5] = madd2(x[4], y[5], t[5], C)
	C, t[6] = madd2(x[4], y[6], t[6], C)
	C, t[7] = madd2(x[4], y[7], t[7], C)

	t[8], D = bits.Add64(t[8], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[5], y[0], t[0])
	C, t[1] = madd2(x[5], y[1], t[1], C)
	C, t[2] = madd2(x[5], y[2], t[2], C)
	C, t[3] = madd2(x[5], y[3], t[3], C)
	C, t[4] = madd2(x[5], y[4], t[4], C)
	C, t[5] = madd2(x[5], y[5], t[5], C)
	C, t[6] = madd2(x[5], y[6], t[6], C)
	C, t[7] = madd2(x[5], y[7], t[7], C)

	t[8], D = bits.Add64(t[8], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[6], y[0], t[0])
	C, t[1] = madd2(x[6], y[1], t[1], C)
	C, t[2] = madd2(x[6], y[2], t[2], C)
	C, t[3] = madd2(x[6], y[3], t[3], C)
	C, t[4] = madd2(x[6], y[4], t[4], C)
	C, t[5] = madd2(x[6], y[5], t[5], C)
	C, t[6] = madd2(x[6], y[6], t[6], C)
	C, t[7] = madd2(x[6], y[7], t[7], C)

	t[8], D = bits.Add64(t[8], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[7], y[0], t[0])
	C, t[1] = madd2(x[7], y[1], t[1], C)
	C, t[2] = madd2(x[7], y[2], t[2], C)
	C, t[3] = madd2(x[7], y[3], t[3], C)
	C, t[4] = madd2(x[7], y[4], t[4], C)
	C, t[5] = madd2(x[7], y[5], t[5], C)
	C, t[6] = madd2(x[7], y[6], t[6], C)
	C, t[7] = madd2(x[7], y[7], t[7], C)

	t[8], D = bits.Add64(t[8], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)
	z[6], D = bits.Sub64(t[6], mod[6], D)
	z[7], D = bits.Sub64(t[7], mod[7], D)

	if D != 0 && t[8] == 0 {
		// reduction was not necessary
		copy(z[:], t[:8])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled576(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [10]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

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
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)
	C, t[3] = madd2(x[1], y[3], t[3], C)
	C, t[4] = madd2(x[1], y[4], t[4], C)
	C, t[5] = madd2(x[1], y[5], t[5], C)
	C, t[6] = madd2(x[1], y[6], t[6], C)
	C, t[7] = madd2(x[1], y[7], t[7], C)
	C, t[8] = madd2(x[1], y[8], t[8], C)

	t[9], D = bits.Add64(t[9], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)
	C, t[3] = madd2(x[2], y[3], t[3], C)
	C, t[4] = madd2(x[2], y[4], t[4], C)
	C, t[5] = madd2(x[2], y[5], t[5], C)
	C, t[6] = madd2(x[2], y[6], t[6], C)
	C, t[7] = madd2(x[2], y[7], t[7], C)
	C, t[8] = madd2(x[2], y[8], t[8], C)

	t[9], D = bits.Add64(t[9], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], y[0], t[0])
	C, t[1] = madd2(x[3], y[1], t[1], C)
	C, t[2] = madd2(x[3], y[2], t[2], C)
	C, t[3] = madd2(x[3], y[3], t[3], C)
	C, t[4] = madd2(x[3], y[4], t[4], C)
	C, t[5] = madd2(x[3], y[5], t[5], C)
	C, t[6] = madd2(x[3], y[6], t[6], C)
	C, t[7] = madd2(x[3], y[7], t[7], C)
	C, t[8] = madd2(x[3], y[8], t[8], C)

	t[9], D = bits.Add64(t[9], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[4], y[0], t[0])
	C, t[1] = madd2(x[4], y[1], t[1], C)
	C, t[2] = madd2(x[4], y[2], t[2], C)
	C, t[3] = madd2(x[4], y[3], t[3], C)
	C, t[4] = madd2(x[4], y[4], t[4], C)
	C, t[5] = madd2(x[4], y[5], t[5], C)
	C, t[6] = madd2(x[4], y[6], t[6], C)
	C, t[7] = madd2(x[4], y[7], t[7], C)
	C, t[8] = madd2(x[4], y[8], t[8], C)

	t[9], D = bits.Add64(t[9], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[5], y[0], t[0])
	C, t[1] = madd2(x[5], y[1], t[1], C)
	C, t[2] = madd2(x[5], y[2], t[2], C)
	C, t[3] = madd2(x[5], y[3], t[3], C)
	C, t[4] = madd2(x[5], y[4], t[4], C)
	C, t[5] = madd2(x[5], y[5], t[5], C)
	C, t[6] = madd2(x[5], y[6], t[6], C)
	C, t[7] = madd2(x[5], y[7], t[7], C)
	C, t[8] = madd2(x[5], y[8], t[8], C)

	t[9], D = bits.Add64(t[9], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[6], y[0], t[0])
	C, t[1] = madd2(x[6], y[1], t[1], C)
	C, t[2] = madd2(x[6], y[2], t[2], C)
	C, t[3] = madd2(x[6], y[3], t[3], C)
	C, t[4] = madd2(x[6], y[4], t[4], C)
	C, t[5] = madd2(x[6], y[5], t[5], C)
	C, t[6] = madd2(x[6], y[6], t[6], C)
	C, t[7] = madd2(x[6], y[7], t[7], C)
	C, t[8] = madd2(x[6], y[8], t[8], C)

	t[9], D = bits.Add64(t[9], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[7], y[0], t[0])
	C, t[1] = madd2(x[7], y[1], t[1], C)
	C, t[2] = madd2(x[7], y[2], t[2], C)
	C, t[3] = madd2(x[7], y[3], t[3], C)
	C, t[4] = madd2(x[7], y[4], t[4], C)
	C, t[5] = madd2(x[7], y[5], t[5], C)
	C, t[6] = madd2(x[7], y[6], t[6], C)
	C, t[7] = madd2(x[7], y[7], t[7], C)
	C, t[8] = madd2(x[7], y[8], t[8], C)

	t[9], D = bits.Add64(t[9], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[8], y[0], t[0])
	C, t[1] = madd2(x[8], y[1], t[1], C)
	C, t[2] = madd2(x[8], y[2], t[2], C)
	C, t[3] = madd2(x[8], y[3], t[3], C)
	C, t[4] = madd2(x[8], y[4], t[4], C)
	C, t[5] = madd2(x[8], y[5], t[5], C)
	C, t[6] = madd2(x[8], y[6], t[6], C)
	C, t[7] = madd2(x[8], y[7], t[7], C)
	C, t[8] = madd2(x[8], y[8], t[8], C)

	t[9], D = bits.Add64(t[9], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	z[0], D = bits.Sub64(t[0], mod[0], 0)
	z[1], D = bits.Sub64(t[1], mod[1], D)
	z[2], D = bits.Sub64(t[2], mod[2], D)
	z[3], D = bits.Sub64(t[3], mod[3], D)
	z[4], D = bits.Sub64(t[4], mod[4], D)
	z[5], D = bits.Sub64(t[5], mod[5], D)
	z[6], D = bits.Sub64(t[6], mod[6], D)
	z[7], D = bits.Sub64(t[7], mod[7], D)
	z[8], D = bits.Sub64(t[8], mod[8], D)

	if D != 0 && t[9] == 0 {
		// reduction was not necessary
		copy(z[:], t[:9])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled640(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[10]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [11]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

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
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)
	C, t[3] = madd2(x[1], y[3], t[3], C)
	C, t[4] = madd2(x[1], y[4], t[4], C)
	C, t[5] = madd2(x[1], y[5], t[5], C)
	C, t[6] = madd2(x[1], y[6], t[6], C)
	C, t[7] = madd2(x[1], y[7], t[7], C)
	C, t[8] = madd2(x[1], y[8], t[8], C)
	C, t[9] = madd2(x[1], y[9], t[9], C)

	t[10], D = bits.Add64(t[10], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)
	C, t[3] = madd2(x[2], y[3], t[3], C)
	C, t[4] = madd2(x[2], y[4], t[4], C)
	C, t[5] = madd2(x[2], y[5], t[5], C)
	C, t[6] = madd2(x[2], y[6], t[6], C)
	C, t[7] = madd2(x[2], y[7], t[7], C)
	C, t[8] = madd2(x[2], y[8], t[8], C)
	C, t[9] = madd2(x[2], y[9], t[9], C)

	t[10], D = bits.Add64(t[10], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], y[0], t[0])
	C, t[1] = madd2(x[3], y[1], t[1], C)
	C, t[2] = madd2(x[3], y[2], t[2], C)
	C, t[3] = madd2(x[3], y[3], t[3], C)
	C, t[4] = madd2(x[3], y[4], t[4], C)
	C, t[5] = madd2(x[3], y[5], t[5], C)
	C, t[6] = madd2(x[3], y[6], t[6], C)
	C, t[7] = madd2(x[3], y[7], t[7], C)
	C, t[8] = madd2(x[3], y[8], t[8], C)
	C, t[9] = madd2(x[3], y[9], t[9], C)

	t[10], D = bits.Add64(t[10], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[4], y[0], t[0])
	C, t[1] = madd2(x[4], y[1], t[1], C)
	C, t[2] = madd2(x[4], y[2], t[2], C)
	C, t[3] = madd2(x[4], y[3], t[3], C)
	C, t[4] = madd2(x[4], y[4], t[4], C)
	C, t[5] = madd2(x[4], y[5], t[5], C)
	C, t[6] = madd2(x[4], y[6], t[6], C)
	C, t[7] = madd2(x[4], y[7], t[7], C)
	C, t[8] = madd2(x[4], y[8], t[8], C)
	C, t[9] = madd2(x[4], y[9], t[9], C)

	t[10], D = bits.Add64(t[10], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[5], y[0], t[0])
	C, t[1] = madd2(x[5], y[1], t[1], C)
	C, t[2] = madd2(x[5], y[2], t[2], C)
	C, t[3] = madd2(x[5], y[3], t[3], C)
	C, t[4] = madd2(x[5], y[4], t[4], C)
	C, t[5] = madd2(x[5], y[5], t[5], C)
	C, t[6] = madd2(x[5], y[6], t[6], C)
	C, t[7] = madd2(x[5], y[7], t[7], C)
	C, t[8] = madd2(x[5], y[8], t[8], C)
	C, t[9] = madd2(x[5], y[9], t[9], C)

	t[10], D = bits.Add64(t[10], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[6], y[0], t[0])
	C, t[1] = madd2(x[6], y[1], t[1], C)
	C, t[2] = madd2(x[6], y[2], t[2], C)
	C, t[3] = madd2(x[6], y[3], t[3], C)
	C, t[4] = madd2(x[6], y[4], t[4], C)
	C, t[5] = madd2(x[6], y[5], t[5], C)
	C, t[6] = madd2(x[6], y[6], t[6], C)
	C, t[7] = madd2(x[6], y[7], t[7], C)
	C, t[8] = madd2(x[6], y[8], t[8], C)
	C, t[9] = madd2(x[6], y[9], t[9], C)

	t[10], D = bits.Add64(t[10], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[7], y[0], t[0])
	C, t[1] = madd2(x[7], y[1], t[1], C)
	C, t[2] = madd2(x[7], y[2], t[2], C)
	C, t[3] = madd2(x[7], y[3], t[3], C)
	C, t[4] = madd2(x[7], y[4], t[4], C)
	C, t[5] = madd2(x[7], y[5], t[5], C)
	C, t[6] = madd2(x[7], y[6], t[6], C)
	C, t[7] = madd2(x[7], y[7], t[7], C)
	C, t[8] = madd2(x[7], y[8], t[8], C)
	C, t[9] = madd2(x[7], y[9], t[9], C)

	t[10], D = bits.Add64(t[10], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[8], y[0], t[0])
	C, t[1] = madd2(x[8], y[1], t[1], C)
	C, t[2] = madd2(x[8], y[2], t[2], C)
	C, t[3] = madd2(x[8], y[3], t[3], C)
	C, t[4] = madd2(x[8], y[4], t[4], C)
	C, t[5] = madd2(x[8], y[5], t[5], C)
	C, t[6] = madd2(x[8], y[6], t[6], C)
	C, t[7] = madd2(x[8], y[7], t[7], C)
	C, t[8] = madd2(x[8], y[8], t[8], C)
	C, t[9] = madd2(x[8], y[9], t[9], C)

	t[10], D = bits.Add64(t[10], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[9], y[0], t[0])
	C, t[1] = madd2(x[9], y[1], t[1], C)
	C, t[2] = madd2(x[9], y[2], t[2], C)
	C, t[3] = madd2(x[9], y[3], t[3], C)
	C, t[4] = madd2(x[9], y[4], t[4], C)
	C, t[5] = madd2(x[9], y[5], t[5], C)
	C, t[6] = madd2(x[9], y[6], t[6], C)
	C, t[7] = madd2(x[9], y[7], t[7], C)
	C, t[8] = madd2(x[9], y[8], t[8], C)
	C, t[9] = madd2(x[9], y[9], t[9], C)

	t[10], D = bits.Add64(t[10], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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

	if D != 0 && t[10] == 0 {
		// reduction was not necessary
		copy(z[:], t[:10])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled704(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[11]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[11]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[11]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[11]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [12]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

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
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)
	C, t[3] = madd2(x[1], y[3], t[3], C)
	C, t[4] = madd2(x[1], y[4], t[4], C)
	C, t[5] = madd2(x[1], y[5], t[5], C)
	C, t[6] = madd2(x[1], y[6], t[6], C)
	C, t[7] = madd2(x[1], y[7], t[7], C)
	C, t[8] = madd2(x[1], y[8], t[8], C)
	C, t[9] = madd2(x[1], y[9], t[9], C)
	C, t[10] = madd2(x[1], y[10], t[10], C)

	t[11], D = bits.Add64(t[11], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)
	C, t[3] = madd2(x[2], y[3], t[3], C)
	C, t[4] = madd2(x[2], y[4], t[4], C)
	C, t[5] = madd2(x[2], y[5], t[5], C)
	C, t[6] = madd2(x[2], y[6], t[6], C)
	C, t[7] = madd2(x[2], y[7], t[7], C)
	C, t[8] = madd2(x[2], y[8], t[8], C)
	C, t[9] = madd2(x[2], y[9], t[9], C)
	C, t[10] = madd2(x[2], y[10], t[10], C)

	t[11], D = bits.Add64(t[11], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], y[0], t[0])
	C, t[1] = madd2(x[3], y[1], t[1], C)
	C, t[2] = madd2(x[3], y[2], t[2], C)
	C, t[3] = madd2(x[3], y[3], t[3], C)
	C, t[4] = madd2(x[3], y[4], t[4], C)
	C, t[5] = madd2(x[3], y[5], t[5], C)
	C, t[6] = madd2(x[3], y[6], t[6], C)
	C, t[7] = madd2(x[3], y[7], t[7], C)
	C, t[8] = madd2(x[3], y[8], t[8], C)
	C, t[9] = madd2(x[3], y[9], t[9], C)
	C, t[10] = madd2(x[3], y[10], t[10], C)

	t[11], D = bits.Add64(t[11], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[4], y[0], t[0])
	C, t[1] = madd2(x[4], y[1], t[1], C)
	C, t[2] = madd2(x[4], y[2], t[2], C)
	C, t[3] = madd2(x[4], y[3], t[3], C)
	C, t[4] = madd2(x[4], y[4], t[4], C)
	C, t[5] = madd2(x[4], y[5], t[5], C)
	C, t[6] = madd2(x[4], y[6], t[6], C)
	C, t[7] = madd2(x[4], y[7], t[7], C)
	C, t[8] = madd2(x[4], y[8], t[8], C)
	C, t[9] = madd2(x[4], y[9], t[9], C)
	C, t[10] = madd2(x[4], y[10], t[10], C)

	t[11], D = bits.Add64(t[11], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[5], y[0], t[0])
	C, t[1] = madd2(x[5], y[1], t[1], C)
	C, t[2] = madd2(x[5], y[2], t[2], C)
	C, t[3] = madd2(x[5], y[3], t[3], C)
	C, t[4] = madd2(x[5], y[4], t[4], C)
	C, t[5] = madd2(x[5], y[5], t[5], C)
	C, t[6] = madd2(x[5], y[6], t[6], C)
	C, t[7] = madd2(x[5], y[7], t[7], C)
	C, t[8] = madd2(x[5], y[8], t[8], C)
	C, t[9] = madd2(x[5], y[9], t[9], C)
	C, t[10] = madd2(x[5], y[10], t[10], C)

	t[11], D = bits.Add64(t[11], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[6], y[0], t[0])
	C, t[1] = madd2(x[6], y[1], t[1], C)
	C, t[2] = madd2(x[6], y[2], t[2], C)
	C, t[3] = madd2(x[6], y[3], t[3], C)
	C, t[4] = madd2(x[6], y[4], t[4], C)
	C, t[5] = madd2(x[6], y[5], t[5], C)
	C, t[6] = madd2(x[6], y[6], t[6], C)
	C, t[7] = madd2(x[6], y[7], t[7], C)
	C, t[8] = madd2(x[6], y[8], t[8], C)
	C, t[9] = madd2(x[6], y[9], t[9], C)
	C, t[10] = madd2(x[6], y[10], t[10], C)

	t[11], D = bits.Add64(t[11], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[7], y[0], t[0])
	C, t[1] = madd2(x[7], y[1], t[1], C)
	C, t[2] = madd2(x[7], y[2], t[2], C)
	C, t[3] = madd2(x[7], y[3], t[3], C)
	C, t[4] = madd2(x[7], y[4], t[4], C)
	C, t[5] = madd2(x[7], y[5], t[5], C)
	C, t[6] = madd2(x[7], y[6], t[6], C)
	C, t[7] = madd2(x[7], y[7], t[7], C)
	C, t[8] = madd2(x[7], y[8], t[8], C)
	C, t[9] = madd2(x[7], y[9], t[9], C)
	C, t[10] = madd2(x[7], y[10], t[10], C)

	t[11], D = bits.Add64(t[11], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[8], y[0], t[0])
	C, t[1] = madd2(x[8], y[1], t[1], C)
	C, t[2] = madd2(x[8], y[2], t[2], C)
	C, t[3] = madd2(x[8], y[3], t[3], C)
	C, t[4] = madd2(x[8], y[4], t[4], C)
	C, t[5] = madd2(x[8], y[5], t[5], C)
	C, t[6] = madd2(x[8], y[6], t[6], C)
	C, t[7] = madd2(x[8], y[7], t[7], C)
	C, t[8] = madd2(x[8], y[8], t[8], C)
	C, t[9] = madd2(x[8], y[9], t[9], C)
	C, t[10] = madd2(x[8], y[10], t[10], C)

	t[11], D = bits.Add64(t[11], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[9], y[0], t[0])
	C, t[1] = madd2(x[9], y[1], t[1], C)
	C, t[2] = madd2(x[9], y[2], t[2], C)
	C, t[3] = madd2(x[9], y[3], t[3], C)
	C, t[4] = madd2(x[9], y[4], t[4], C)
	C, t[5] = madd2(x[9], y[5], t[5], C)
	C, t[6] = madd2(x[9], y[6], t[6], C)
	C, t[7] = madd2(x[9], y[7], t[7], C)
	C, t[8] = madd2(x[9], y[8], t[8], C)
	C, t[9] = madd2(x[9], y[9], t[9], C)
	C, t[10] = madd2(x[9], y[10], t[10], C)

	t[11], D = bits.Add64(t[11], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[10], y[0], t[0])
	C, t[1] = madd2(x[10], y[1], t[1], C)
	C, t[2] = madd2(x[10], y[2], t[2], C)
	C, t[3] = madd2(x[10], y[3], t[3], C)
	C, t[4] = madd2(x[10], y[4], t[4], C)
	C, t[5] = madd2(x[10], y[5], t[5], C)
	C, t[6] = madd2(x[10], y[6], t[6], C)
	C, t[7] = madd2(x[10], y[7], t[7], C)
	C, t[8] = madd2(x[10], y[8], t[8], C)
	C, t[9] = madd2(x[10], y[9], t[9], C)
	C, t[10] = madd2(x[10], y[10], t[10], C)

	t[11], D = bits.Add64(t[11], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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

	if D != 0 && t[11] == 0 {
		// reduction was not necessary
		copy(z[:], t[:11])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled768(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[12]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[12]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[12]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [13]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

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
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)
	C, t[3] = madd2(x[1], y[3], t[3], C)
	C, t[4] = madd2(x[1], y[4], t[4], C)
	C, t[5] = madd2(x[1], y[5], t[5], C)
	C, t[6] = madd2(x[1], y[6], t[6], C)
	C, t[7] = madd2(x[1], y[7], t[7], C)
	C, t[8] = madd2(x[1], y[8], t[8], C)
	C, t[9] = madd2(x[1], y[9], t[9], C)
	C, t[10] = madd2(x[1], y[10], t[10], C)
	C, t[11] = madd2(x[1], y[11], t[11], C)

	t[12], D = bits.Add64(t[12], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)
	C, t[3] = madd2(x[2], y[3], t[3], C)
	C, t[4] = madd2(x[2], y[4], t[4], C)
	C, t[5] = madd2(x[2], y[5], t[5], C)
	C, t[6] = madd2(x[2], y[6], t[6], C)
	C, t[7] = madd2(x[2], y[7], t[7], C)
	C, t[8] = madd2(x[2], y[8], t[8], C)
	C, t[9] = madd2(x[2], y[9], t[9], C)
	C, t[10] = madd2(x[2], y[10], t[10], C)
	C, t[11] = madd2(x[2], y[11], t[11], C)

	t[12], D = bits.Add64(t[12], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], y[0], t[0])
	C, t[1] = madd2(x[3], y[1], t[1], C)
	C, t[2] = madd2(x[3], y[2], t[2], C)
	C, t[3] = madd2(x[3], y[3], t[3], C)
	C, t[4] = madd2(x[3], y[4], t[4], C)
	C, t[5] = madd2(x[3], y[5], t[5], C)
	C, t[6] = madd2(x[3], y[6], t[6], C)
	C, t[7] = madd2(x[3], y[7], t[7], C)
	C, t[8] = madd2(x[3], y[8], t[8], C)
	C, t[9] = madd2(x[3], y[9], t[9], C)
	C, t[10] = madd2(x[3], y[10], t[10], C)
	C, t[11] = madd2(x[3], y[11], t[11], C)

	t[12], D = bits.Add64(t[12], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[4], y[0], t[0])
	C, t[1] = madd2(x[4], y[1], t[1], C)
	C, t[2] = madd2(x[4], y[2], t[2], C)
	C, t[3] = madd2(x[4], y[3], t[3], C)
	C, t[4] = madd2(x[4], y[4], t[4], C)
	C, t[5] = madd2(x[4], y[5], t[5], C)
	C, t[6] = madd2(x[4], y[6], t[6], C)
	C, t[7] = madd2(x[4], y[7], t[7], C)
	C, t[8] = madd2(x[4], y[8], t[8], C)
	C, t[9] = madd2(x[4], y[9], t[9], C)
	C, t[10] = madd2(x[4], y[10], t[10], C)
	C, t[11] = madd2(x[4], y[11], t[11], C)

	t[12], D = bits.Add64(t[12], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[5], y[0], t[0])
	C, t[1] = madd2(x[5], y[1], t[1], C)
	C, t[2] = madd2(x[5], y[2], t[2], C)
	C, t[3] = madd2(x[5], y[3], t[3], C)
	C, t[4] = madd2(x[5], y[4], t[4], C)
	C, t[5] = madd2(x[5], y[5], t[5], C)
	C, t[6] = madd2(x[5], y[6], t[6], C)
	C, t[7] = madd2(x[5], y[7], t[7], C)
	C, t[8] = madd2(x[5], y[8], t[8], C)
	C, t[9] = madd2(x[5], y[9], t[9], C)
	C, t[10] = madd2(x[5], y[10], t[10], C)
	C, t[11] = madd2(x[5], y[11], t[11], C)

	t[12], D = bits.Add64(t[12], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[6], y[0], t[0])
	C, t[1] = madd2(x[6], y[1], t[1], C)
	C, t[2] = madd2(x[6], y[2], t[2], C)
	C, t[3] = madd2(x[6], y[3], t[3], C)
	C, t[4] = madd2(x[6], y[4], t[4], C)
	C, t[5] = madd2(x[6], y[5], t[5], C)
	C, t[6] = madd2(x[6], y[6], t[6], C)
	C, t[7] = madd2(x[6], y[7], t[7], C)
	C, t[8] = madd2(x[6], y[8], t[8], C)
	C, t[9] = madd2(x[6], y[9], t[9], C)
	C, t[10] = madd2(x[6], y[10], t[10], C)
	C, t[11] = madd2(x[6], y[11], t[11], C)

	t[12], D = bits.Add64(t[12], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[7], y[0], t[0])
	C, t[1] = madd2(x[7], y[1], t[1], C)
	C, t[2] = madd2(x[7], y[2], t[2], C)
	C, t[3] = madd2(x[7], y[3], t[3], C)
	C, t[4] = madd2(x[7], y[4], t[4], C)
	C, t[5] = madd2(x[7], y[5], t[5], C)
	C, t[6] = madd2(x[7], y[6], t[6], C)
	C, t[7] = madd2(x[7], y[7], t[7], C)
	C, t[8] = madd2(x[7], y[8], t[8], C)
	C, t[9] = madd2(x[7], y[9], t[9], C)
	C, t[10] = madd2(x[7], y[10], t[10], C)
	C, t[11] = madd2(x[7], y[11], t[11], C)

	t[12], D = bits.Add64(t[12], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[8], y[0], t[0])
	C, t[1] = madd2(x[8], y[1], t[1], C)
	C, t[2] = madd2(x[8], y[2], t[2], C)
	C, t[3] = madd2(x[8], y[3], t[3], C)
	C, t[4] = madd2(x[8], y[4], t[4], C)
	C, t[5] = madd2(x[8], y[5], t[5], C)
	C, t[6] = madd2(x[8], y[6], t[6], C)
	C, t[7] = madd2(x[8], y[7], t[7], C)
	C, t[8] = madd2(x[8], y[8], t[8], C)
	C, t[9] = madd2(x[8], y[9], t[9], C)
	C, t[10] = madd2(x[8], y[10], t[10], C)
	C, t[11] = madd2(x[8], y[11], t[11], C)

	t[12], D = bits.Add64(t[12], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[9], y[0], t[0])
	C, t[1] = madd2(x[9], y[1], t[1], C)
	C, t[2] = madd2(x[9], y[2], t[2], C)
	C, t[3] = madd2(x[9], y[3], t[3], C)
	C, t[4] = madd2(x[9], y[4], t[4], C)
	C, t[5] = madd2(x[9], y[5], t[5], C)
	C, t[6] = madd2(x[9], y[6], t[6], C)
	C, t[7] = madd2(x[9], y[7], t[7], C)
	C, t[8] = madd2(x[9], y[8], t[8], C)
	C, t[9] = madd2(x[9], y[9], t[9], C)
	C, t[10] = madd2(x[9], y[10], t[10], C)
	C, t[11] = madd2(x[9], y[11], t[11], C)

	t[12], D = bits.Add64(t[12], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[10], y[0], t[0])
	C, t[1] = madd2(x[10], y[1], t[1], C)
	C, t[2] = madd2(x[10], y[2], t[2], C)
	C, t[3] = madd2(x[10], y[3], t[3], C)
	C, t[4] = madd2(x[10], y[4], t[4], C)
	C, t[5] = madd2(x[10], y[5], t[5], C)
	C, t[6] = madd2(x[10], y[6], t[6], C)
	C, t[7] = madd2(x[10], y[7], t[7], C)
	C, t[8] = madd2(x[10], y[8], t[8], C)
	C, t[9] = madd2(x[10], y[9], t[9], C)
	C, t[10] = madd2(x[10], y[10], t[10], C)
	C, t[11] = madd2(x[10], y[11], t[11], C)

	t[12], D = bits.Add64(t[12], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[11], y[0], t[0])
	C, t[1] = madd2(x[11], y[1], t[1], C)
	C, t[2] = madd2(x[11], y[2], t[2], C)
	C, t[3] = madd2(x[11], y[3], t[3], C)
	C, t[4] = madd2(x[11], y[4], t[4], C)
	C, t[5] = madd2(x[11], y[5], t[5], C)
	C, t[6] = madd2(x[11], y[6], t[6], C)
	C, t[7] = madd2(x[11], y[7], t[7], C)
	C, t[8] = madd2(x[11], y[8], t[8], C)
	C, t[9] = madd2(x[11], y[9], t[9], C)
	C, t[10] = madd2(x[11], y[10], t[10], C)
	C, t[11] = madd2(x[11], y[11], t[11], C)

	t[12], D = bits.Add64(t[12], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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

	if D != 0 && t[12] == 0 {
		// reduction was not necessary
		copy(z[:], t[:12])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled832(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[13]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[13]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[13]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[13]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [14]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

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
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)
	C, t[3] = madd2(x[1], y[3], t[3], C)
	C, t[4] = madd2(x[1], y[4], t[4], C)
	C, t[5] = madd2(x[1], y[5], t[5], C)
	C, t[6] = madd2(x[1], y[6], t[6], C)
	C, t[7] = madd2(x[1], y[7], t[7], C)
	C, t[8] = madd2(x[1], y[8], t[8], C)
	C, t[9] = madd2(x[1], y[9], t[9], C)
	C, t[10] = madd2(x[1], y[10], t[10], C)
	C, t[11] = madd2(x[1], y[11], t[11], C)
	C, t[12] = madd2(x[1], y[12], t[12], C)

	t[13], D = bits.Add64(t[13], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)
	C, t[3] = madd2(x[2], y[3], t[3], C)
	C, t[4] = madd2(x[2], y[4], t[4], C)
	C, t[5] = madd2(x[2], y[5], t[5], C)
	C, t[6] = madd2(x[2], y[6], t[6], C)
	C, t[7] = madd2(x[2], y[7], t[7], C)
	C, t[8] = madd2(x[2], y[8], t[8], C)
	C, t[9] = madd2(x[2], y[9], t[9], C)
	C, t[10] = madd2(x[2], y[10], t[10], C)
	C, t[11] = madd2(x[2], y[11], t[11], C)
	C, t[12] = madd2(x[2], y[12], t[12], C)

	t[13], D = bits.Add64(t[13], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], y[0], t[0])
	C, t[1] = madd2(x[3], y[1], t[1], C)
	C, t[2] = madd2(x[3], y[2], t[2], C)
	C, t[3] = madd2(x[3], y[3], t[3], C)
	C, t[4] = madd2(x[3], y[4], t[4], C)
	C, t[5] = madd2(x[3], y[5], t[5], C)
	C, t[6] = madd2(x[3], y[6], t[6], C)
	C, t[7] = madd2(x[3], y[7], t[7], C)
	C, t[8] = madd2(x[3], y[8], t[8], C)
	C, t[9] = madd2(x[3], y[9], t[9], C)
	C, t[10] = madd2(x[3], y[10], t[10], C)
	C, t[11] = madd2(x[3], y[11], t[11], C)
	C, t[12] = madd2(x[3], y[12], t[12], C)

	t[13], D = bits.Add64(t[13], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[4], y[0], t[0])
	C, t[1] = madd2(x[4], y[1], t[1], C)
	C, t[2] = madd2(x[4], y[2], t[2], C)
	C, t[3] = madd2(x[4], y[3], t[3], C)
	C, t[4] = madd2(x[4], y[4], t[4], C)
	C, t[5] = madd2(x[4], y[5], t[5], C)
	C, t[6] = madd2(x[4], y[6], t[6], C)
	C, t[7] = madd2(x[4], y[7], t[7], C)
	C, t[8] = madd2(x[4], y[8], t[8], C)
	C, t[9] = madd2(x[4], y[9], t[9], C)
	C, t[10] = madd2(x[4], y[10], t[10], C)
	C, t[11] = madd2(x[4], y[11], t[11], C)
	C, t[12] = madd2(x[4], y[12], t[12], C)

	t[13], D = bits.Add64(t[13], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[5], y[0], t[0])
	C, t[1] = madd2(x[5], y[1], t[1], C)
	C, t[2] = madd2(x[5], y[2], t[2], C)
	C, t[3] = madd2(x[5], y[3], t[3], C)
	C, t[4] = madd2(x[5], y[4], t[4], C)
	C, t[5] = madd2(x[5], y[5], t[5], C)
	C, t[6] = madd2(x[5], y[6], t[6], C)
	C, t[7] = madd2(x[5], y[7], t[7], C)
	C, t[8] = madd2(x[5], y[8], t[8], C)
	C, t[9] = madd2(x[5], y[9], t[9], C)
	C, t[10] = madd2(x[5], y[10], t[10], C)
	C, t[11] = madd2(x[5], y[11], t[11], C)
	C, t[12] = madd2(x[5], y[12], t[12], C)

	t[13], D = bits.Add64(t[13], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[6], y[0], t[0])
	C, t[1] = madd2(x[6], y[1], t[1], C)
	C, t[2] = madd2(x[6], y[2], t[2], C)
	C, t[3] = madd2(x[6], y[3], t[3], C)
	C, t[4] = madd2(x[6], y[4], t[4], C)
	C, t[5] = madd2(x[6], y[5], t[5], C)
	C, t[6] = madd2(x[6], y[6], t[6], C)
	C, t[7] = madd2(x[6], y[7], t[7], C)
	C, t[8] = madd2(x[6], y[8], t[8], C)
	C, t[9] = madd2(x[6], y[9], t[9], C)
	C, t[10] = madd2(x[6], y[10], t[10], C)
	C, t[11] = madd2(x[6], y[11], t[11], C)
	C, t[12] = madd2(x[6], y[12], t[12], C)

	t[13], D = bits.Add64(t[13], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[7], y[0], t[0])
	C, t[1] = madd2(x[7], y[1], t[1], C)
	C, t[2] = madd2(x[7], y[2], t[2], C)
	C, t[3] = madd2(x[7], y[3], t[3], C)
	C, t[4] = madd2(x[7], y[4], t[4], C)
	C, t[5] = madd2(x[7], y[5], t[5], C)
	C, t[6] = madd2(x[7], y[6], t[6], C)
	C, t[7] = madd2(x[7], y[7], t[7], C)
	C, t[8] = madd2(x[7], y[8], t[8], C)
	C, t[9] = madd2(x[7], y[9], t[9], C)
	C, t[10] = madd2(x[7], y[10], t[10], C)
	C, t[11] = madd2(x[7], y[11], t[11], C)
	C, t[12] = madd2(x[7], y[12], t[12], C)

	t[13], D = bits.Add64(t[13], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[8], y[0], t[0])
	C, t[1] = madd2(x[8], y[1], t[1], C)
	C, t[2] = madd2(x[8], y[2], t[2], C)
	C, t[3] = madd2(x[8], y[3], t[3], C)
	C, t[4] = madd2(x[8], y[4], t[4], C)
	C, t[5] = madd2(x[8], y[5], t[5], C)
	C, t[6] = madd2(x[8], y[6], t[6], C)
	C, t[7] = madd2(x[8], y[7], t[7], C)
	C, t[8] = madd2(x[8], y[8], t[8], C)
	C, t[9] = madd2(x[8], y[9], t[9], C)
	C, t[10] = madd2(x[8], y[10], t[10], C)
	C, t[11] = madd2(x[8], y[11], t[11], C)
	C, t[12] = madd2(x[8], y[12], t[12], C)

	t[13], D = bits.Add64(t[13], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[9], y[0], t[0])
	C, t[1] = madd2(x[9], y[1], t[1], C)
	C, t[2] = madd2(x[9], y[2], t[2], C)
	C, t[3] = madd2(x[9], y[3], t[3], C)
	C, t[4] = madd2(x[9], y[4], t[4], C)
	C, t[5] = madd2(x[9], y[5], t[5], C)
	C, t[6] = madd2(x[9], y[6], t[6], C)
	C, t[7] = madd2(x[9], y[7], t[7], C)
	C, t[8] = madd2(x[9], y[8], t[8], C)
	C, t[9] = madd2(x[9], y[9], t[9], C)
	C, t[10] = madd2(x[9], y[10], t[10], C)
	C, t[11] = madd2(x[9], y[11], t[11], C)
	C, t[12] = madd2(x[9], y[12], t[12], C)

	t[13], D = bits.Add64(t[13], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[10], y[0], t[0])
	C, t[1] = madd2(x[10], y[1], t[1], C)
	C, t[2] = madd2(x[10], y[2], t[2], C)
	C, t[3] = madd2(x[10], y[3], t[3], C)
	C, t[4] = madd2(x[10], y[4], t[4], C)
	C, t[5] = madd2(x[10], y[5], t[5], C)
	C, t[6] = madd2(x[10], y[6], t[6], C)
	C, t[7] = madd2(x[10], y[7], t[7], C)
	C, t[8] = madd2(x[10], y[8], t[8], C)
	C, t[9] = madd2(x[10], y[9], t[9], C)
	C, t[10] = madd2(x[10], y[10], t[10], C)
	C, t[11] = madd2(x[10], y[11], t[11], C)
	C, t[12] = madd2(x[10], y[12], t[12], C)

	t[13], D = bits.Add64(t[13], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[11], y[0], t[0])
	C, t[1] = madd2(x[11], y[1], t[1], C)
	C, t[2] = madd2(x[11], y[2], t[2], C)
	C, t[3] = madd2(x[11], y[3], t[3], C)
	C, t[4] = madd2(x[11], y[4], t[4], C)
	C, t[5] = madd2(x[11], y[5], t[5], C)
	C, t[6] = madd2(x[11], y[6], t[6], C)
	C, t[7] = madd2(x[11], y[7], t[7], C)
	C, t[8] = madd2(x[11], y[8], t[8], C)
	C, t[9] = madd2(x[11], y[9], t[9], C)
	C, t[10] = madd2(x[11], y[10], t[10], C)
	C, t[11] = madd2(x[11], y[11], t[11], C)
	C, t[12] = madd2(x[11], y[12], t[12], C)

	t[13], D = bits.Add64(t[13], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[12], y[0], t[0])
	C, t[1] = madd2(x[12], y[1], t[1], C)
	C, t[2] = madd2(x[12], y[2], t[2], C)
	C, t[3] = madd2(x[12], y[3], t[3], C)
	C, t[4] = madd2(x[12], y[4], t[4], C)
	C, t[5] = madd2(x[12], y[5], t[5], C)
	C, t[6] = madd2(x[12], y[6], t[6], C)
	C, t[7] = madd2(x[12], y[7], t[7], C)
	C, t[8] = madd2(x[12], y[8], t[8], C)
	C, t[9] = madd2(x[12], y[9], t[9], C)
	C, t[10] = madd2(x[12], y[10], t[10], C)
	C, t[11] = madd2(x[12], y[11], t[11], C)
	C, t[12] = madd2(x[12], y[12], t[12], C)

	t[13], D = bits.Add64(t[13], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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

	if D != 0 && t[13] == 0 {
		// reduction was not necessary
		copy(z[:], t[:13])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled896(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[14]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[14]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[14]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[14]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [15]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

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
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)
	C, t[3] = madd2(x[1], y[3], t[3], C)
	C, t[4] = madd2(x[1], y[4], t[4], C)
	C, t[5] = madd2(x[1], y[5], t[5], C)
	C, t[6] = madd2(x[1], y[6], t[6], C)
	C, t[7] = madd2(x[1], y[7], t[7], C)
	C, t[8] = madd2(x[1], y[8], t[8], C)
	C, t[9] = madd2(x[1], y[9], t[9], C)
	C, t[10] = madd2(x[1], y[10], t[10], C)
	C, t[11] = madd2(x[1], y[11], t[11], C)
	C, t[12] = madd2(x[1], y[12], t[12], C)
	C, t[13] = madd2(x[1], y[13], t[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)
	C, t[3] = madd2(x[2], y[3], t[3], C)
	C, t[4] = madd2(x[2], y[4], t[4], C)
	C, t[5] = madd2(x[2], y[5], t[5], C)
	C, t[6] = madd2(x[2], y[6], t[6], C)
	C, t[7] = madd2(x[2], y[7], t[7], C)
	C, t[8] = madd2(x[2], y[8], t[8], C)
	C, t[9] = madd2(x[2], y[9], t[9], C)
	C, t[10] = madd2(x[2], y[10], t[10], C)
	C, t[11] = madd2(x[2], y[11], t[11], C)
	C, t[12] = madd2(x[2], y[12], t[12], C)
	C, t[13] = madd2(x[2], y[13], t[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], y[0], t[0])
	C, t[1] = madd2(x[3], y[1], t[1], C)
	C, t[2] = madd2(x[3], y[2], t[2], C)
	C, t[3] = madd2(x[3], y[3], t[3], C)
	C, t[4] = madd2(x[3], y[4], t[4], C)
	C, t[5] = madd2(x[3], y[5], t[5], C)
	C, t[6] = madd2(x[3], y[6], t[6], C)
	C, t[7] = madd2(x[3], y[7], t[7], C)
	C, t[8] = madd2(x[3], y[8], t[8], C)
	C, t[9] = madd2(x[3], y[9], t[9], C)
	C, t[10] = madd2(x[3], y[10], t[10], C)
	C, t[11] = madd2(x[3], y[11], t[11], C)
	C, t[12] = madd2(x[3], y[12], t[12], C)
	C, t[13] = madd2(x[3], y[13], t[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[4], y[0], t[0])
	C, t[1] = madd2(x[4], y[1], t[1], C)
	C, t[2] = madd2(x[4], y[2], t[2], C)
	C, t[3] = madd2(x[4], y[3], t[3], C)
	C, t[4] = madd2(x[4], y[4], t[4], C)
	C, t[5] = madd2(x[4], y[5], t[5], C)
	C, t[6] = madd2(x[4], y[6], t[6], C)
	C, t[7] = madd2(x[4], y[7], t[7], C)
	C, t[8] = madd2(x[4], y[8], t[8], C)
	C, t[9] = madd2(x[4], y[9], t[9], C)
	C, t[10] = madd2(x[4], y[10], t[10], C)
	C, t[11] = madd2(x[4], y[11], t[11], C)
	C, t[12] = madd2(x[4], y[12], t[12], C)
	C, t[13] = madd2(x[4], y[13], t[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[5], y[0], t[0])
	C, t[1] = madd2(x[5], y[1], t[1], C)
	C, t[2] = madd2(x[5], y[2], t[2], C)
	C, t[3] = madd2(x[5], y[3], t[3], C)
	C, t[4] = madd2(x[5], y[4], t[4], C)
	C, t[5] = madd2(x[5], y[5], t[5], C)
	C, t[6] = madd2(x[5], y[6], t[6], C)
	C, t[7] = madd2(x[5], y[7], t[7], C)
	C, t[8] = madd2(x[5], y[8], t[8], C)
	C, t[9] = madd2(x[5], y[9], t[9], C)
	C, t[10] = madd2(x[5], y[10], t[10], C)
	C, t[11] = madd2(x[5], y[11], t[11], C)
	C, t[12] = madd2(x[5], y[12], t[12], C)
	C, t[13] = madd2(x[5], y[13], t[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[6], y[0], t[0])
	C, t[1] = madd2(x[6], y[1], t[1], C)
	C, t[2] = madd2(x[6], y[2], t[2], C)
	C, t[3] = madd2(x[6], y[3], t[3], C)
	C, t[4] = madd2(x[6], y[4], t[4], C)
	C, t[5] = madd2(x[6], y[5], t[5], C)
	C, t[6] = madd2(x[6], y[6], t[6], C)
	C, t[7] = madd2(x[6], y[7], t[7], C)
	C, t[8] = madd2(x[6], y[8], t[8], C)
	C, t[9] = madd2(x[6], y[9], t[9], C)
	C, t[10] = madd2(x[6], y[10], t[10], C)
	C, t[11] = madd2(x[6], y[11], t[11], C)
	C, t[12] = madd2(x[6], y[12], t[12], C)
	C, t[13] = madd2(x[6], y[13], t[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[7], y[0], t[0])
	C, t[1] = madd2(x[7], y[1], t[1], C)
	C, t[2] = madd2(x[7], y[2], t[2], C)
	C, t[3] = madd2(x[7], y[3], t[3], C)
	C, t[4] = madd2(x[7], y[4], t[4], C)
	C, t[5] = madd2(x[7], y[5], t[5], C)
	C, t[6] = madd2(x[7], y[6], t[6], C)
	C, t[7] = madd2(x[7], y[7], t[7], C)
	C, t[8] = madd2(x[7], y[8], t[8], C)
	C, t[9] = madd2(x[7], y[9], t[9], C)
	C, t[10] = madd2(x[7], y[10], t[10], C)
	C, t[11] = madd2(x[7], y[11], t[11], C)
	C, t[12] = madd2(x[7], y[12], t[12], C)
	C, t[13] = madd2(x[7], y[13], t[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[8], y[0], t[0])
	C, t[1] = madd2(x[8], y[1], t[1], C)
	C, t[2] = madd2(x[8], y[2], t[2], C)
	C, t[3] = madd2(x[8], y[3], t[3], C)
	C, t[4] = madd2(x[8], y[4], t[4], C)
	C, t[5] = madd2(x[8], y[5], t[5], C)
	C, t[6] = madd2(x[8], y[6], t[6], C)
	C, t[7] = madd2(x[8], y[7], t[7], C)
	C, t[8] = madd2(x[8], y[8], t[8], C)
	C, t[9] = madd2(x[8], y[9], t[9], C)
	C, t[10] = madd2(x[8], y[10], t[10], C)
	C, t[11] = madd2(x[8], y[11], t[11], C)
	C, t[12] = madd2(x[8], y[12], t[12], C)
	C, t[13] = madd2(x[8], y[13], t[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[9], y[0], t[0])
	C, t[1] = madd2(x[9], y[1], t[1], C)
	C, t[2] = madd2(x[9], y[2], t[2], C)
	C, t[3] = madd2(x[9], y[3], t[3], C)
	C, t[4] = madd2(x[9], y[4], t[4], C)
	C, t[5] = madd2(x[9], y[5], t[5], C)
	C, t[6] = madd2(x[9], y[6], t[6], C)
	C, t[7] = madd2(x[9], y[7], t[7], C)
	C, t[8] = madd2(x[9], y[8], t[8], C)
	C, t[9] = madd2(x[9], y[9], t[9], C)
	C, t[10] = madd2(x[9], y[10], t[10], C)
	C, t[11] = madd2(x[9], y[11], t[11], C)
	C, t[12] = madd2(x[9], y[12], t[12], C)
	C, t[13] = madd2(x[9], y[13], t[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[10], y[0], t[0])
	C, t[1] = madd2(x[10], y[1], t[1], C)
	C, t[2] = madd2(x[10], y[2], t[2], C)
	C, t[3] = madd2(x[10], y[3], t[3], C)
	C, t[4] = madd2(x[10], y[4], t[4], C)
	C, t[5] = madd2(x[10], y[5], t[5], C)
	C, t[6] = madd2(x[10], y[6], t[6], C)
	C, t[7] = madd2(x[10], y[7], t[7], C)
	C, t[8] = madd2(x[10], y[8], t[8], C)
	C, t[9] = madd2(x[10], y[9], t[9], C)
	C, t[10] = madd2(x[10], y[10], t[10], C)
	C, t[11] = madd2(x[10], y[11], t[11], C)
	C, t[12] = madd2(x[10], y[12], t[12], C)
	C, t[13] = madd2(x[10], y[13], t[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[11], y[0], t[0])
	C, t[1] = madd2(x[11], y[1], t[1], C)
	C, t[2] = madd2(x[11], y[2], t[2], C)
	C, t[3] = madd2(x[11], y[3], t[3], C)
	C, t[4] = madd2(x[11], y[4], t[4], C)
	C, t[5] = madd2(x[11], y[5], t[5], C)
	C, t[6] = madd2(x[11], y[6], t[6], C)
	C, t[7] = madd2(x[11], y[7], t[7], C)
	C, t[8] = madd2(x[11], y[8], t[8], C)
	C, t[9] = madd2(x[11], y[9], t[9], C)
	C, t[10] = madd2(x[11], y[10], t[10], C)
	C, t[11] = madd2(x[11], y[11], t[11], C)
	C, t[12] = madd2(x[11], y[12], t[12], C)
	C, t[13] = madd2(x[11], y[13], t[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[12], y[0], t[0])
	C, t[1] = madd2(x[12], y[1], t[1], C)
	C, t[2] = madd2(x[12], y[2], t[2], C)
	C, t[3] = madd2(x[12], y[3], t[3], C)
	C, t[4] = madd2(x[12], y[4], t[4], C)
	C, t[5] = madd2(x[12], y[5], t[5], C)
	C, t[6] = madd2(x[12], y[6], t[6], C)
	C, t[7] = madd2(x[12], y[7], t[7], C)
	C, t[8] = madd2(x[12], y[8], t[8], C)
	C, t[9] = madd2(x[12], y[9], t[9], C)
	C, t[10] = madd2(x[12], y[10], t[10], C)
	C, t[11] = madd2(x[12], y[11], t[11], C)
	C, t[12] = madd2(x[12], y[12], t[12], C)
	C, t[13] = madd2(x[12], y[13], t[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[13], y[0], t[0])
	C, t[1] = madd2(x[13], y[1], t[1], C)
	C, t[2] = madd2(x[13], y[2], t[2], C)
	C, t[3] = madd2(x[13], y[3], t[3], C)
	C, t[4] = madd2(x[13], y[4], t[4], C)
	C, t[5] = madd2(x[13], y[5], t[5], C)
	C, t[6] = madd2(x[13], y[6], t[6], C)
	C, t[7] = madd2(x[13], y[7], t[7], C)
	C, t[8] = madd2(x[13], y[8], t[8], C)
	C, t[9] = madd2(x[13], y[9], t[9], C)
	C, t[10] = madd2(x[13], y[10], t[10], C)
	C, t[11] = madd2(x[13], y[11], t[11], C)
	C, t[12] = madd2(x[13], y[12], t[12], C)
	C, t[13] = madd2(x[13], y[13], t[13], C)

	t[14], D = bits.Add64(t[14], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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

	if D != 0 && t[14] == 0 {
		// reduction was not necessary
		copy(z[:], t[:14])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled960(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[15]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[15]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[15]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[15]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [16]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

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
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)
	C, t[3] = madd2(x[1], y[3], t[3], C)
	C, t[4] = madd2(x[1], y[4], t[4], C)
	C, t[5] = madd2(x[1], y[5], t[5], C)
	C, t[6] = madd2(x[1], y[6], t[6], C)
	C, t[7] = madd2(x[1], y[7], t[7], C)
	C, t[8] = madd2(x[1], y[8], t[8], C)
	C, t[9] = madd2(x[1], y[9], t[9], C)
	C, t[10] = madd2(x[1], y[10], t[10], C)
	C, t[11] = madd2(x[1], y[11], t[11], C)
	C, t[12] = madd2(x[1], y[12], t[12], C)
	C, t[13] = madd2(x[1], y[13], t[13], C)
	C, t[14] = madd2(x[1], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)
	C, t[3] = madd2(x[2], y[3], t[3], C)
	C, t[4] = madd2(x[2], y[4], t[4], C)
	C, t[5] = madd2(x[2], y[5], t[5], C)
	C, t[6] = madd2(x[2], y[6], t[6], C)
	C, t[7] = madd2(x[2], y[7], t[7], C)
	C, t[8] = madd2(x[2], y[8], t[8], C)
	C, t[9] = madd2(x[2], y[9], t[9], C)
	C, t[10] = madd2(x[2], y[10], t[10], C)
	C, t[11] = madd2(x[2], y[11], t[11], C)
	C, t[12] = madd2(x[2], y[12], t[12], C)
	C, t[13] = madd2(x[2], y[13], t[13], C)
	C, t[14] = madd2(x[2], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], y[0], t[0])
	C, t[1] = madd2(x[3], y[1], t[1], C)
	C, t[2] = madd2(x[3], y[2], t[2], C)
	C, t[3] = madd2(x[3], y[3], t[3], C)
	C, t[4] = madd2(x[3], y[4], t[4], C)
	C, t[5] = madd2(x[3], y[5], t[5], C)
	C, t[6] = madd2(x[3], y[6], t[6], C)
	C, t[7] = madd2(x[3], y[7], t[7], C)
	C, t[8] = madd2(x[3], y[8], t[8], C)
	C, t[9] = madd2(x[3], y[9], t[9], C)
	C, t[10] = madd2(x[3], y[10], t[10], C)
	C, t[11] = madd2(x[3], y[11], t[11], C)
	C, t[12] = madd2(x[3], y[12], t[12], C)
	C, t[13] = madd2(x[3], y[13], t[13], C)
	C, t[14] = madd2(x[3], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[4], y[0], t[0])
	C, t[1] = madd2(x[4], y[1], t[1], C)
	C, t[2] = madd2(x[4], y[2], t[2], C)
	C, t[3] = madd2(x[4], y[3], t[3], C)
	C, t[4] = madd2(x[4], y[4], t[4], C)
	C, t[5] = madd2(x[4], y[5], t[5], C)
	C, t[6] = madd2(x[4], y[6], t[6], C)
	C, t[7] = madd2(x[4], y[7], t[7], C)
	C, t[8] = madd2(x[4], y[8], t[8], C)
	C, t[9] = madd2(x[4], y[9], t[9], C)
	C, t[10] = madd2(x[4], y[10], t[10], C)
	C, t[11] = madd2(x[4], y[11], t[11], C)
	C, t[12] = madd2(x[4], y[12], t[12], C)
	C, t[13] = madd2(x[4], y[13], t[13], C)
	C, t[14] = madd2(x[4], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[5], y[0], t[0])
	C, t[1] = madd2(x[5], y[1], t[1], C)
	C, t[2] = madd2(x[5], y[2], t[2], C)
	C, t[3] = madd2(x[5], y[3], t[3], C)
	C, t[4] = madd2(x[5], y[4], t[4], C)
	C, t[5] = madd2(x[5], y[5], t[5], C)
	C, t[6] = madd2(x[5], y[6], t[6], C)
	C, t[7] = madd2(x[5], y[7], t[7], C)
	C, t[8] = madd2(x[5], y[8], t[8], C)
	C, t[9] = madd2(x[5], y[9], t[9], C)
	C, t[10] = madd2(x[5], y[10], t[10], C)
	C, t[11] = madd2(x[5], y[11], t[11], C)
	C, t[12] = madd2(x[5], y[12], t[12], C)
	C, t[13] = madd2(x[5], y[13], t[13], C)
	C, t[14] = madd2(x[5], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[6], y[0], t[0])
	C, t[1] = madd2(x[6], y[1], t[1], C)
	C, t[2] = madd2(x[6], y[2], t[2], C)
	C, t[3] = madd2(x[6], y[3], t[3], C)
	C, t[4] = madd2(x[6], y[4], t[4], C)
	C, t[5] = madd2(x[6], y[5], t[5], C)
	C, t[6] = madd2(x[6], y[6], t[6], C)
	C, t[7] = madd2(x[6], y[7], t[7], C)
	C, t[8] = madd2(x[6], y[8], t[8], C)
	C, t[9] = madd2(x[6], y[9], t[9], C)
	C, t[10] = madd2(x[6], y[10], t[10], C)
	C, t[11] = madd2(x[6], y[11], t[11], C)
	C, t[12] = madd2(x[6], y[12], t[12], C)
	C, t[13] = madd2(x[6], y[13], t[13], C)
	C, t[14] = madd2(x[6], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[7], y[0], t[0])
	C, t[1] = madd2(x[7], y[1], t[1], C)
	C, t[2] = madd2(x[7], y[2], t[2], C)
	C, t[3] = madd2(x[7], y[3], t[3], C)
	C, t[4] = madd2(x[7], y[4], t[4], C)
	C, t[5] = madd2(x[7], y[5], t[5], C)
	C, t[6] = madd2(x[7], y[6], t[6], C)
	C, t[7] = madd2(x[7], y[7], t[7], C)
	C, t[8] = madd2(x[7], y[8], t[8], C)
	C, t[9] = madd2(x[7], y[9], t[9], C)
	C, t[10] = madd2(x[7], y[10], t[10], C)
	C, t[11] = madd2(x[7], y[11], t[11], C)
	C, t[12] = madd2(x[7], y[12], t[12], C)
	C, t[13] = madd2(x[7], y[13], t[13], C)
	C, t[14] = madd2(x[7], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[8], y[0], t[0])
	C, t[1] = madd2(x[8], y[1], t[1], C)
	C, t[2] = madd2(x[8], y[2], t[2], C)
	C, t[3] = madd2(x[8], y[3], t[3], C)
	C, t[4] = madd2(x[8], y[4], t[4], C)
	C, t[5] = madd2(x[8], y[5], t[5], C)
	C, t[6] = madd2(x[8], y[6], t[6], C)
	C, t[7] = madd2(x[8], y[7], t[7], C)
	C, t[8] = madd2(x[8], y[8], t[8], C)
	C, t[9] = madd2(x[8], y[9], t[9], C)
	C, t[10] = madd2(x[8], y[10], t[10], C)
	C, t[11] = madd2(x[8], y[11], t[11], C)
	C, t[12] = madd2(x[8], y[12], t[12], C)
	C, t[13] = madd2(x[8], y[13], t[13], C)
	C, t[14] = madd2(x[8], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[9], y[0], t[0])
	C, t[1] = madd2(x[9], y[1], t[1], C)
	C, t[2] = madd2(x[9], y[2], t[2], C)
	C, t[3] = madd2(x[9], y[3], t[3], C)
	C, t[4] = madd2(x[9], y[4], t[4], C)
	C, t[5] = madd2(x[9], y[5], t[5], C)
	C, t[6] = madd2(x[9], y[6], t[6], C)
	C, t[7] = madd2(x[9], y[7], t[7], C)
	C, t[8] = madd2(x[9], y[8], t[8], C)
	C, t[9] = madd2(x[9], y[9], t[9], C)
	C, t[10] = madd2(x[9], y[10], t[10], C)
	C, t[11] = madd2(x[9], y[11], t[11], C)
	C, t[12] = madd2(x[9], y[12], t[12], C)
	C, t[13] = madd2(x[9], y[13], t[13], C)
	C, t[14] = madd2(x[9], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[10], y[0], t[0])
	C, t[1] = madd2(x[10], y[1], t[1], C)
	C, t[2] = madd2(x[10], y[2], t[2], C)
	C, t[3] = madd2(x[10], y[3], t[3], C)
	C, t[4] = madd2(x[10], y[4], t[4], C)
	C, t[5] = madd2(x[10], y[5], t[5], C)
	C, t[6] = madd2(x[10], y[6], t[6], C)
	C, t[7] = madd2(x[10], y[7], t[7], C)
	C, t[8] = madd2(x[10], y[8], t[8], C)
	C, t[9] = madd2(x[10], y[9], t[9], C)
	C, t[10] = madd2(x[10], y[10], t[10], C)
	C, t[11] = madd2(x[10], y[11], t[11], C)
	C, t[12] = madd2(x[10], y[12], t[12], C)
	C, t[13] = madd2(x[10], y[13], t[13], C)
	C, t[14] = madd2(x[10], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[11], y[0], t[0])
	C, t[1] = madd2(x[11], y[1], t[1], C)
	C, t[2] = madd2(x[11], y[2], t[2], C)
	C, t[3] = madd2(x[11], y[3], t[3], C)
	C, t[4] = madd2(x[11], y[4], t[4], C)
	C, t[5] = madd2(x[11], y[5], t[5], C)
	C, t[6] = madd2(x[11], y[6], t[6], C)
	C, t[7] = madd2(x[11], y[7], t[7], C)
	C, t[8] = madd2(x[11], y[8], t[8], C)
	C, t[9] = madd2(x[11], y[9], t[9], C)
	C, t[10] = madd2(x[11], y[10], t[10], C)
	C, t[11] = madd2(x[11], y[11], t[11], C)
	C, t[12] = madd2(x[11], y[12], t[12], C)
	C, t[13] = madd2(x[11], y[13], t[13], C)
	C, t[14] = madd2(x[11], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[12], y[0], t[0])
	C, t[1] = madd2(x[12], y[1], t[1], C)
	C, t[2] = madd2(x[12], y[2], t[2], C)
	C, t[3] = madd2(x[12], y[3], t[3], C)
	C, t[4] = madd2(x[12], y[4], t[4], C)
	C, t[5] = madd2(x[12], y[5], t[5], C)
	C, t[6] = madd2(x[12], y[6], t[6], C)
	C, t[7] = madd2(x[12], y[7], t[7], C)
	C, t[8] = madd2(x[12], y[8], t[8], C)
	C, t[9] = madd2(x[12], y[9], t[9], C)
	C, t[10] = madd2(x[12], y[10], t[10], C)
	C, t[11] = madd2(x[12], y[11], t[11], C)
	C, t[12] = madd2(x[12], y[12], t[12], C)
	C, t[13] = madd2(x[12], y[13], t[13], C)
	C, t[14] = madd2(x[12], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[13], y[0], t[0])
	C, t[1] = madd2(x[13], y[1], t[1], C)
	C, t[2] = madd2(x[13], y[2], t[2], C)
	C, t[3] = madd2(x[13], y[3], t[3], C)
	C, t[4] = madd2(x[13], y[4], t[4], C)
	C, t[5] = madd2(x[13], y[5], t[5], C)
	C, t[6] = madd2(x[13], y[6], t[6], C)
	C, t[7] = madd2(x[13], y[7], t[7], C)
	C, t[8] = madd2(x[13], y[8], t[8], C)
	C, t[9] = madd2(x[13], y[9], t[9], C)
	C, t[10] = madd2(x[13], y[10], t[10], C)
	C, t[11] = madd2(x[13], y[11], t[11], C)
	C, t[12] = madd2(x[13], y[12], t[12], C)
	C, t[13] = madd2(x[13], y[13], t[13], C)
	C, t[14] = madd2(x[13], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[14], y[0], t[0])
	C, t[1] = madd2(x[14], y[1], t[1], C)
	C, t[2] = madd2(x[14], y[2], t[2], C)
	C, t[3] = madd2(x[14], y[3], t[3], C)
	C, t[4] = madd2(x[14], y[4], t[4], C)
	C, t[5] = madd2(x[14], y[5], t[5], C)
	C, t[6] = madd2(x[14], y[6], t[6], C)
	C, t[7] = madd2(x[14], y[7], t[7], C)
	C, t[8] = madd2(x[14], y[8], t[8], C)
	C, t[9] = madd2(x[14], y[9], t[9], C)
	C, t[10] = madd2(x[14], y[10], t[10], C)
	C, t[11] = madd2(x[14], y[11], t[11], C)
	C, t[12] = madd2(x[14], y[12], t[12], C)
	C, t[13] = madd2(x[14], y[13], t[13], C)
	C, t[14] = madd2(x[14], y[14], t[14], C)

	t[15], D = bits.Add64(t[15], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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

	if D != 0 && t[15] == 0 {
		// reduction was not necessary
		copy(z[:], t[:15])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}

func MulMontUnrolled1024(ctx *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[16]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[16]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[16]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[16]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [17]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

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
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], y[0], t[0])
	C, t[1] = madd2(x[1], y[1], t[1], C)
	C, t[2] = madd2(x[1], y[2], t[2], C)
	C, t[3] = madd2(x[1], y[3], t[3], C)
	C, t[4] = madd2(x[1], y[4], t[4], C)
	C, t[5] = madd2(x[1], y[5], t[5], C)
	C, t[6] = madd2(x[1], y[6], t[6], C)
	C, t[7] = madd2(x[1], y[7], t[7], C)
	C, t[8] = madd2(x[1], y[8], t[8], C)
	C, t[9] = madd2(x[1], y[9], t[9], C)
	C, t[10] = madd2(x[1], y[10], t[10], C)
	C, t[11] = madd2(x[1], y[11], t[11], C)
	C, t[12] = madd2(x[1], y[12], t[12], C)
	C, t[13] = madd2(x[1], y[13], t[13], C)
	C, t[14] = madd2(x[1], y[14], t[14], C)
	C, t[15] = madd2(x[1], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], y[0], t[0])
	C, t[1] = madd2(x[2], y[1], t[1], C)
	C, t[2] = madd2(x[2], y[2], t[2], C)
	C, t[3] = madd2(x[2], y[3], t[3], C)
	C, t[4] = madd2(x[2], y[4], t[4], C)
	C, t[5] = madd2(x[2], y[5], t[5], C)
	C, t[6] = madd2(x[2], y[6], t[6], C)
	C, t[7] = madd2(x[2], y[7], t[7], C)
	C, t[8] = madd2(x[2], y[8], t[8], C)
	C, t[9] = madd2(x[2], y[9], t[9], C)
	C, t[10] = madd2(x[2], y[10], t[10], C)
	C, t[11] = madd2(x[2], y[11], t[11], C)
	C, t[12] = madd2(x[2], y[12], t[12], C)
	C, t[13] = madd2(x[2], y[13], t[13], C)
	C, t[14] = madd2(x[2], y[14], t[14], C)
	C, t[15] = madd2(x[2], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], y[0], t[0])
	C, t[1] = madd2(x[3], y[1], t[1], C)
	C, t[2] = madd2(x[3], y[2], t[2], C)
	C, t[3] = madd2(x[3], y[3], t[3], C)
	C, t[4] = madd2(x[3], y[4], t[4], C)
	C, t[5] = madd2(x[3], y[5], t[5], C)
	C, t[6] = madd2(x[3], y[6], t[6], C)
	C, t[7] = madd2(x[3], y[7], t[7], C)
	C, t[8] = madd2(x[3], y[8], t[8], C)
	C, t[9] = madd2(x[3], y[9], t[9], C)
	C, t[10] = madd2(x[3], y[10], t[10], C)
	C, t[11] = madd2(x[3], y[11], t[11], C)
	C, t[12] = madd2(x[3], y[12], t[12], C)
	C, t[13] = madd2(x[3], y[13], t[13], C)
	C, t[14] = madd2(x[3], y[14], t[14], C)
	C, t[15] = madd2(x[3], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[4], y[0], t[0])
	C, t[1] = madd2(x[4], y[1], t[1], C)
	C, t[2] = madd2(x[4], y[2], t[2], C)
	C, t[3] = madd2(x[4], y[3], t[3], C)
	C, t[4] = madd2(x[4], y[4], t[4], C)
	C, t[5] = madd2(x[4], y[5], t[5], C)
	C, t[6] = madd2(x[4], y[6], t[6], C)
	C, t[7] = madd2(x[4], y[7], t[7], C)
	C, t[8] = madd2(x[4], y[8], t[8], C)
	C, t[9] = madd2(x[4], y[9], t[9], C)
	C, t[10] = madd2(x[4], y[10], t[10], C)
	C, t[11] = madd2(x[4], y[11], t[11], C)
	C, t[12] = madd2(x[4], y[12], t[12], C)
	C, t[13] = madd2(x[4], y[13], t[13], C)
	C, t[14] = madd2(x[4], y[14], t[14], C)
	C, t[15] = madd2(x[4], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[5], y[0], t[0])
	C, t[1] = madd2(x[5], y[1], t[1], C)
	C, t[2] = madd2(x[5], y[2], t[2], C)
	C, t[3] = madd2(x[5], y[3], t[3], C)
	C, t[4] = madd2(x[5], y[4], t[4], C)
	C, t[5] = madd2(x[5], y[5], t[5], C)
	C, t[6] = madd2(x[5], y[6], t[6], C)
	C, t[7] = madd2(x[5], y[7], t[7], C)
	C, t[8] = madd2(x[5], y[8], t[8], C)
	C, t[9] = madd2(x[5], y[9], t[9], C)
	C, t[10] = madd2(x[5], y[10], t[10], C)
	C, t[11] = madd2(x[5], y[11], t[11], C)
	C, t[12] = madd2(x[5], y[12], t[12], C)
	C, t[13] = madd2(x[5], y[13], t[13], C)
	C, t[14] = madd2(x[5], y[14], t[14], C)
	C, t[15] = madd2(x[5], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[6], y[0], t[0])
	C, t[1] = madd2(x[6], y[1], t[1], C)
	C, t[2] = madd2(x[6], y[2], t[2], C)
	C, t[3] = madd2(x[6], y[3], t[3], C)
	C, t[4] = madd2(x[6], y[4], t[4], C)
	C, t[5] = madd2(x[6], y[5], t[5], C)
	C, t[6] = madd2(x[6], y[6], t[6], C)
	C, t[7] = madd2(x[6], y[7], t[7], C)
	C, t[8] = madd2(x[6], y[8], t[8], C)
	C, t[9] = madd2(x[6], y[9], t[9], C)
	C, t[10] = madd2(x[6], y[10], t[10], C)
	C, t[11] = madd2(x[6], y[11], t[11], C)
	C, t[12] = madd2(x[6], y[12], t[12], C)
	C, t[13] = madd2(x[6], y[13], t[13], C)
	C, t[14] = madd2(x[6], y[14], t[14], C)
	C, t[15] = madd2(x[6], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[7], y[0], t[0])
	C, t[1] = madd2(x[7], y[1], t[1], C)
	C, t[2] = madd2(x[7], y[2], t[2], C)
	C, t[3] = madd2(x[7], y[3], t[3], C)
	C, t[4] = madd2(x[7], y[4], t[4], C)
	C, t[5] = madd2(x[7], y[5], t[5], C)
	C, t[6] = madd2(x[7], y[6], t[6], C)
	C, t[7] = madd2(x[7], y[7], t[7], C)
	C, t[8] = madd2(x[7], y[8], t[8], C)
	C, t[9] = madd2(x[7], y[9], t[9], C)
	C, t[10] = madd2(x[7], y[10], t[10], C)
	C, t[11] = madd2(x[7], y[11], t[11], C)
	C, t[12] = madd2(x[7], y[12], t[12], C)
	C, t[13] = madd2(x[7], y[13], t[13], C)
	C, t[14] = madd2(x[7], y[14], t[14], C)
	C, t[15] = madd2(x[7], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[8], y[0], t[0])
	C, t[1] = madd2(x[8], y[1], t[1], C)
	C, t[2] = madd2(x[8], y[2], t[2], C)
	C, t[3] = madd2(x[8], y[3], t[3], C)
	C, t[4] = madd2(x[8], y[4], t[4], C)
	C, t[5] = madd2(x[8], y[5], t[5], C)
	C, t[6] = madd2(x[8], y[6], t[6], C)
	C, t[7] = madd2(x[8], y[7], t[7], C)
	C, t[8] = madd2(x[8], y[8], t[8], C)
	C, t[9] = madd2(x[8], y[9], t[9], C)
	C, t[10] = madd2(x[8], y[10], t[10], C)
	C, t[11] = madd2(x[8], y[11], t[11], C)
	C, t[12] = madd2(x[8], y[12], t[12], C)
	C, t[13] = madd2(x[8], y[13], t[13], C)
	C, t[14] = madd2(x[8], y[14], t[14], C)
	C, t[15] = madd2(x[8], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[9], y[0], t[0])
	C, t[1] = madd2(x[9], y[1], t[1], C)
	C, t[2] = madd2(x[9], y[2], t[2], C)
	C, t[3] = madd2(x[9], y[3], t[3], C)
	C, t[4] = madd2(x[9], y[4], t[4], C)
	C, t[5] = madd2(x[9], y[5], t[5], C)
	C, t[6] = madd2(x[9], y[6], t[6], C)
	C, t[7] = madd2(x[9], y[7], t[7], C)
	C, t[8] = madd2(x[9], y[8], t[8], C)
	C, t[9] = madd2(x[9], y[9], t[9], C)
	C, t[10] = madd2(x[9], y[10], t[10], C)
	C, t[11] = madd2(x[9], y[11], t[11], C)
	C, t[12] = madd2(x[9], y[12], t[12], C)
	C, t[13] = madd2(x[9], y[13], t[13], C)
	C, t[14] = madd2(x[9], y[14], t[14], C)
	C, t[15] = madd2(x[9], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[10], y[0], t[0])
	C, t[1] = madd2(x[10], y[1], t[1], C)
	C, t[2] = madd2(x[10], y[2], t[2], C)
	C, t[3] = madd2(x[10], y[3], t[3], C)
	C, t[4] = madd2(x[10], y[4], t[4], C)
	C, t[5] = madd2(x[10], y[5], t[5], C)
	C, t[6] = madd2(x[10], y[6], t[6], C)
	C, t[7] = madd2(x[10], y[7], t[7], C)
	C, t[8] = madd2(x[10], y[8], t[8], C)
	C, t[9] = madd2(x[10], y[9], t[9], C)
	C, t[10] = madd2(x[10], y[10], t[10], C)
	C, t[11] = madd2(x[10], y[11], t[11], C)
	C, t[12] = madd2(x[10], y[12], t[12], C)
	C, t[13] = madd2(x[10], y[13], t[13], C)
	C, t[14] = madd2(x[10], y[14], t[14], C)
	C, t[15] = madd2(x[10], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[11], y[0], t[0])
	C, t[1] = madd2(x[11], y[1], t[1], C)
	C, t[2] = madd2(x[11], y[2], t[2], C)
	C, t[3] = madd2(x[11], y[3], t[3], C)
	C, t[4] = madd2(x[11], y[4], t[4], C)
	C, t[5] = madd2(x[11], y[5], t[5], C)
	C, t[6] = madd2(x[11], y[6], t[6], C)
	C, t[7] = madd2(x[11], y[7], t[7], C)
	C, t[8] = madd2(x[11], y[8], t[8], C)
	C, t[9] = madd2(x[11], y[9], t[9], C)
	C, t[10] = madd2(x[11], y[10], t[10], C)
	C, t[11] = madd2(x[11], y[11], t[11], C)
	C, t[12] = madd2(x[11], y[12], t[12], C)
	C, t[13] = madd2(x[11], y[13], t[13], C)
	C, t[14] = madd2(x[11], y[14], t[14], C)
	C, t[15] = madd2(x[11], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[12], y[0], t[0])
	C, t[1] = madd2(x[12], y[1], t[1], C)
	C, t[2] = madd2(x[12], y[2], t[2], C)
	C, t[3] = madd2(x[12], y[3], t[3], C)
	C, t[4] = madd2(x[12], y[4], t[4], C)
	C, t[5] = madd2(x[12], y[5], t[5], C)
	C, t[6] = madd2(x[12], y[6], t[6], C)
	C, t[7] = madd2(x[12], y[7], t[7], C)
	C, t[8] = madd2(x[12], y[8], t[8], C)
	C, t[9] = madd2(x[12], y[9], t[9], C)
	C, t[10] = madd2(x[12], y[10], t[10], C)
	C, t[11] = madd2(x[12], y[11], t[11], C)
	C, t[12] = madd2(x[12], y[12], t[12], C)
	C, t[13] = madd2(x[12], y[13], t[13], C)
	C, t[14] = madd2(x[12], y[14], t[14], C)
	C, t[15] = madd2(x[12], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[13], y[0], t[0])
	C, t[1] = madd2(x[13], y[1], t[1], C)
	C, t[2] = madd2(x[13], y[2], t[2], C)
	C, t[3] = madd2(x[13], y[3], t[3], C)
	C, t[4] = madd2(x[13], y[4], t[4], C)
	C, t[5] = madd2(x[13], y[5], t[5], C)
	C, t[6] = madd2(x[13], y[6], t[6], C)
	C, t[7] = madd2(x[13], y[7], t[7], C)
	C, t[8] = madd2(x[13], y[8], t[8], C)
	C, t[9] = madd2(x[13], y[9], t[9], C)
	C, t[10] = madd2(x[13], y[10], t[10], C)
	C, t[11] = madd2(x[13], y[11], t[11], C)
	C, t[12] = madd2(x[13], y[12], t[12], C)
	C, t[13] = madd2(x[13], y[13], t[13], C)
	C, t[14] = madd2(x[13], y[14], t[14], C)
	C, t[15] = madd2(x[13], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[14], y[0], t[0])
	C, t[1] = madd2(x[14], y[1], t[1], C)
	C, t[2] = madd2(x[14], y[2], t[2], C)
	C, t[3] = madd2(x[14], y[3], t[3], C)
	C, t[4] = madd2(x[14], y[4], t[4], C)
	C, t[5] = madd2(x[14], y[5], t[5], C)
	C, t[6] = madd2(x[14], y[6], t[6], C)
	C, t[7] = madd2(x[14], y[7], t[7], C)
	C, t[8] = madd2(x[14], y[8], t[8], C)
	C, t[9] = madd2(x[14], y[9], t[9], C)
	C, t[10] = madd2(x[14], y[10], t[10], C)
	C, t[11] = madd2(x[14], y[11], t[11], C)
	C, t[12] = madd2(x[14], y[12], t[12], C)
	C, t[13] = madd2(x[14], y[13], t[13], C)
	C, t[14] = madd2(x[14], y[14], t[14], C)
	C, t[15] = madd2(x[14], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[15], y[0], t[0])
	C, t[1] = madd2(x[15], y[1], t[1], C)
	C, t[2] = madd2(x[15], y[2], t[2], C)
	C, t[3] = madd2(x[15], y[3], t[3], C)
	C, t[4] = madd2(x[15], y[4], t[4], C)
	C, t[5] = madd2(x[15], y[5], t[5], C)
	C, t[6] = madd2(x[15], y[6], t[6], C)
	C, t[7] = madd2(x[15], y[7], t[7], C)
	C, t[8] = madd2(x[15], y[8], t[8], C)
	C, t[9] = madd2(x[15], y[9], t[9], C)
	C, t[10] = madd2(x[15], y[10], t[10], C)
	C, t[11] = madd2(x[15], y[11], t[11], C)
	C, t[12] = madd2(x[15], y[12], t[12], C)
	C, t[13] = madd2(x[15], y[13], t[13], C)
	C, t[14] = madd2(x[15], y[14], t[14], C)
	C, t[15] = madd2(x[15], y[15], t[15], C)

	t[16], D = bits.Add64(t[16], C, 0)
	// m = t[0]n'[0] mod W
	m = t[0] * ctx.MontParamInterleaved
	// -----------------------------------
	// Second loop
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

	if D != 0 && t[16] == 0 {
		// reduction was not necessary
		copy(z[:], t[:16])
	} /* else {
	    panic("not worst case performance")
	}*/

	return nil
}
