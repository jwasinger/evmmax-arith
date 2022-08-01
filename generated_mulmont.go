



package mont_arith

import (
	"math/bits"
    "math/big"
    "unsafe"
    "errors"
    "fmt"
)

// madd0 hi = a*b + c (discards lo bits)
func madd0(a, b, c uint64) (uint64) {
	var carry, lo uint64
	hi, lo := bits.Mul64(a, b)
	_, carry = bits.Add64(lo, c, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return hi
}

// madd1 hi, lo = a*b + c
func madd1(a, b, c uint64) (uint64, uint64) {
	var carry uint64
	hi, lo := bits.Mul64(a, b)
	lo, carry = bits.Add64(lo, c, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return hi, lo
}

// madd2 hi, lo = a*b + c + d
func madd2(a, b, c, d uint64) (uint64, uint64) {
	var carry uint64
	hi, lo := bits.Mul64(a, b)
	c, carry = bits.Add64(c, d, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	lo, carry = bits.Add64(lo, c, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return hi, lo
}

func madd3(a, b, c, d, e uint64) (uint64, uint64) {
	var carry uint64
    var c_uint uint64
	hi, lo := bits.Mul64(a, b)
	c_uint, carry = bits.Add64(c, d, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	lo, carry = bits.Add64(lo, c_uint, 0)
	hi, _ = bits.Add64(hi, e, carry)
	return hi, lo
}

/*
 * begin mulmont implementations
 */

func mulMont64(f *Field, outBytes, xBytes, yBytes []byte) error {
	var product [2]uint64
	var c uint64
    mod := f.Modulus
    modinv := f.MontParamInterleaved

    x := (*[1]uint64)(unsafe.Pointer(&xBytes[0]))[:]
    y := (*[1]uint64)(unsafe.Pointer(&yBytes[0]))[:]
    out := (*[1]uint64)(unsafe.Pointer(&outBytes[0]))[:]

    if x[0] >= mod[0] || y[0] >= mod[0] {
        return errors.New(fmt.Sprintf("x/y gte modulus"))
    }

	product[1], product[0] = bits.Mul64(x[0], y[0])
	m := product[0] * modinv
	c, _ = madd1(m, mod[0], product[0])
	out[0] = c + product[1]

	if out[0] > mod[0] {
		out[0] = c - mod[0]
	}
    return nil
}




func mulMont128(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [3]uint64
	var D uint64
	var m, C uint64

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
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

    // TODO this shows up here, but I can't find reference to it in any paper
    // that references CIOS. is this just a quick hack for the final subtraction?
	if t[2] != 0 {
		// we need to reduce, we have a result on 3 words
		var b uint64
		z[0],b = bits.Sub64(t[0], mod[0], 0)
				z[1], _ = bits.Sub64(t[1], mod[1], b)
		return nil
	}

	// copy t into z
		z[0] = t[0]
		z[1] = t[1]

	// final subtraction, overwriting z if z > mod
	if GTE(z, mod) {
		C = 0
		for i := 0; i < 2; i++ {
			z[i], C = bits.Sub64(z[i], mod[i], C)
		}
	}

	return nil
}




func mulMont192(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[3]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [4]uint64
	var D uint64
	var m, C uint64

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
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

    // TODO this shows up here, but I can't find reference to it in any paper
    // that references CIOS. is this just a quick hack for the final subtraction?
	if t[3] != 0 {
		// we need to reduce, we have a result on 4 words
		var b uint64
		z[0],b = bits.Sub64(t[0], mod[0], 0)
				z[1], b = bits.Sub64(t[1], mod[1], b)
				z[2], _ = bits.Sub64(t[2], mod[2], b)
		return nil
	}

	// copy t into z
		z[0] = t[0]
		z[1] = t[1]
		z[2] = t[2]

	// final subtraction, overwriting z if z > mod
	if GTE(z, mod) {
		C = 0
		for i := 0; i < 3; i++ {
			z[i], C = bits.Sub64(z[i], mod[i], C)
		}
	}

	return nil
}




func mulMont256(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[4]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [5]uint64
	var D uint64
	var m, C uint64

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
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

    // TODO this shows up here, but I can't find reference to it in any paper
    // that references CIOS. is this just a quick hack for the final subtraction?
	if t[4] != 0 {
		// we need to reduce, we have a result on 5 words
		var b uint64
		z[0],b = bits.Sub64(t[0], mod[0], 0)
				z[1], b = bits.Sub64(t[1], mod[1], b)
				z[2], b = bits.Sub64(t[2], mod[2], b)
				z[3], _ = bits.Sub64(t[3], mod[3], b)
		return nil
	}

	// copy t into z
		z[0] = t[0]
		z[1] = t[1]
		z[2] = t[2]
		z[3] = t[3]

	// final subtraction, overwriting z if z > mod
	if GTE(z, mod) {
		C = 0
		for i := 0; i < 4; i++ {
			z[i], C = bits.Sub64(z[i], mod[i], C)
		}
	}

	return nil
}




func mulMont320(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[5]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[5]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[5]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[5]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [6]uint64
	var D uint64
	var m, C uint64

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
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

    // TODO this shows up here, but I can't find reference to it in any paper
    // that references CIOS. is this just a quick hack for the final subtraction?
	if t[5] != 0 {
		// we need to reduce, we have a result on 6 words
		var b uint64
		z[0],b = bits.Sub64(t[0], mod[0], 0)
				z[1], b = bits.Sub64(t[1], mod[1], b)
				z[2], b = bits.Sub64(t[2], mod[2], b)
				z[3], b = bits.Sub64(t[3], mod[3], b)
				z[4], _ = bits.Sub64(t[4], mod[4], b)
		return nil
	}

	// copy t into z
		z[0] = t[0]
		z[1] = t[1]
		z[2] = t[2]
		z[3] = t[3]
		z[4] = t[4]

	// final subtraction, overwriting z if z > mod
	if GTE(z, mod) {
		C = 0
		for i := 0; i < 5; i++ {
			z[i], C = bits.Sub64(z[i], mod[i], C)
		}
	}

	return nil
}




func mulMont384(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [7]uint64
	var D uint64
	var m, C uint64

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
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

    // TODO this shows up here, but I can't find reference to it in any paper
    // that references CIOS. is this just a quick hack for the final subtraction?
	if t[6] != 0 {
		// we need to reduce, we have a result on 7 words
		var b uint64
		z[0],b = bits.Sub64(t[0], mod[0], 0)
				z[1], b = bits.Sub64(t[1], mod[1], b)
				z[2], b = bits.Sub64(t[2], mod[2], b)
				z[3], b = bits.Sub64(t[3], mod[3], b)
				z[4], b = bits.Sub64(t[4], mod[4], b)
				z[5], _ = bits.Sub64(t[5], mod[5], b)
		return nil
	}

	// copy t into z
		z[0] = t[0]
		z[1] = t[1]
		z[2] = t[2]
		z[3] = t[3]
		z[4] = t[4]
		z[5] = t[5]

	// final subtraction, overwriting z if z > mod
	if GTE(z, mod) {
		C = 0
		for i := 0; i < 6; i++ {
			z[i], C = bits.Sub64(z[i], mod[i], C)
		}
	}

	return nil
}




func mulMont448(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[7]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [8]uint64
	var D uint64
	var m, C uint64

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
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

    // TODO this shows up here, but I can't find reference to it in any paper
    // that references CIOS. is this just a quick hack for the final subtraction?
	if t[7] != 0 {
		// we need to reduce, we have a result on 8 words
		var b uint64
		z[0],b = bits.Sub64(t[0], mod[0], 0)
				z[1], b = bits.Sub64(t[1], mod[1], b)
				z[2], b = bits.Sub64(t[2], mod[2], b)
				z[3], b = bits.Sub64(t[3], mod[3], b)
				z[4], b = bits.Sub64(t[4], mod[4], b)
				z[5], b = bits.Sub64(t[5], mod[5], b)
				z[6], _ = bits.Sub64(t[6], mod[6], b)
		return nil
	}

	// copy t into z
		z[0] = t[0]
		z[1] = t[1]
		z[2] = t[2]
		z[3] = t[3]
		z[4] = t[4]
		z[5] = t[5]
		z[6] = t[6]

	// final subtraction, overwriting z if z > mod
	if GTE(z, mod) {
		C = 0
		for i := 0; i < 7; i++ {
			z[i], C = bits.Sub64(z[i], mod[i], C)
		}
	}

	return nil
}




func mulMont512(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[8]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[8]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[8]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[8]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [9]uint64
	var D uint64
	var m, C uint64

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
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

    // TODO this shows up here, but I can't find reference to it in any paper
    // that references CIOS. is this just a quick hack for the final subtraction?
	if t[8] != 0 {
		// we need to reduce, we have a result on 9 words
		var b uint64
		z[0],b = bits.Sub64(t[0], mod[0], 0)
				z[1], b = bits.Sub64(t[1], mod[1], b)
				z[2], b = bits.Sub64(t[2], mod[2], b)
				z[3], b = bits.Sub64(t[3], mod[3], b)
				z[4], b = bits.Sub64(t[4], mod[4], b)
				z[5], b = bits.Sub64(t[5], mod[5], b)
				z[6], b = bits.Sub64(t[6], mod[6], b)
				z[7], _ = bits.Sub64(t[7], mod[7], b)
		return nil
	}

	// copy t into z
		z[0] = t[0]
		z[1] = t[1]
		z[2] = t[2]
		z[3] = t[3]
		z[4] = t[4]
		z[5] = t[5]
		z[6] = t[6]
		z[7] = t[7]

	// final subtraction, overwriting z if z > mod
	if GTE(z, mod) {
		C = 0
		for i := 0; i < 8; i++ {
			z[i], C = bits.Sub64(z[i], mod[i], C)
		}
	}

	return nil
}




func mulMont576(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [10]uint64
	var D uint64
	var m, C uint64

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
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

    // TODO this shows up here, but I can't find reference to it in any paper
    // that references CIOS. is this just a quick hack for the final subtraction?
	if t[9] != 0 {
		// we need to reduce, we have a result on 10 words
		var b uint64
		z[0],b = bits.Sub64(t[0], mod[0], 0)
				z[1], b = bits.Sub64(t[1], mod[1], b)
				z[2], b = bits.Sub64(t[2], mod[2], b)
				z[3], b = bits.Sub64(t[3], mod[3], b)
				z[4], b = bits.Sub64(t[4], mod[4], b)
				z[5], b = bits.Sub64(t[5], mod[5], b)
				z[6], b = bits.Sub64(t[6], mod[6], b)
				z[7], b = bits.Sub64(t[7], mod[7], b)
				z[8], _ = bits.Sub64(t[8], mod[8], b)
		return nil
	}

	// copy t into z
		z[0] = t[0]
		z[1] = t[1]
		z[2] = t[2]
		z[3] = t[3]
		z[4] = t[4]
		z[5] = t[5]
		z[6] = t[6]
		z[7] = t[7]
		z[8] = t[8]

	// final subtraction, overwriting z if z > mod
	if GTE(z, mod) {
		C = 0
		for i := 0; i < 9; i++ {
			z[i], C = bits.Sub64(z[i], mod[i], C)
		}
	}

	return nil
}




func mulMont640(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[10]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [11]uint64
	var D uint64
	var m, C uint64

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
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

    // TODO this shows up here, but I can't find reference to it in any paper
    // that references CIOS. is this just a quick hack for the final subtraction?
	if t[10] != 0 {
		// we need to reduce, we have a result on 11 words
		var b uint64
		z[0],b = bits.Sub64(t[0], mod[0], 0)
				z[1], b = bits.Sub64(t[1], mod[1], b)
				z[2], b = bits.Sub64(t[2], mod[2], b)
				z[3], b = bits.Sub64(t[3], mod[3], b)
				z[4], b = bits.Sub64(t[4], mod[4], b)
				z[5], b = bits.Sub64(t[5], mod[5], b)
				z[6], b = bits.Sub64(t[6], mod[6], b)
				z[7], b = bits.Sub64(t[7], mod[7], b)
				z[8], b = bits.Sub64(t[8], mod[8], b)
				z[9], _ = bits.Sub64(t[9], mod[9], b)
		return nil
	}

	// copy t into z
		z[0] = t[0]
		z[1] = t[1]
		z[2] = t[2]
		z[3] = t[3]
		z[4] = t[4]
		z[5] = t[5]
		z[6] = t[6]
		z[7] = t[7]
		z[8] = t[8]
		z[9] = t[9]

	// final subtraction, overwriting z if z > mod
	if GTE(z, mod) {
		C = 0
		for i := 0; i < 10; i++ {
			z[i], C = bits.Sub64(z[i], mod[i], C)
		}
	}

	return nil
}




func mulMont704(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[11]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[11]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[11]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[11]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [12]uint64
	var D uint64
	var m, C uint64

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
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

    // TODO this shows up here, but I can't find reference to it in any paper
    // that references CIOS. is this just a quick hack for the final subtraction?
	if t[11] != 0 {
		// we need to reduce, we have a result on 12 words
		var b uint64
		z[0],b = bits.Sub64(t[0], mod[0], 0)
				z[1], b = bits.Sub64(t[1], mod[1], b)
				z[2], b = bits.Sub64(t[2], mod[2], b)
				z[3], b = bits.Sub64(t[3], mod[3], b)
				z[4], b = bits.Sub64(t[4], mod[4], b)
				z[5], b = bits.Sub64(t[5], mod[5], b)
				z[6], b = bits.Sub64(t[6], mod[6], b)
				z[7], b = bits.Sub64(t[7], mod[7], b)
				z[8], b = bits.Sub64(t[8], mod[8], b)
				z[9], b = bits.Sub64(t[9], mod[9], b)
				z[10], _ = bits.Sub64(t[10], mod[10], b)
		return nil
	}

	// copy t into z
		z[0] = t[0]
		z[1] = t[1]
		z[2] = t[2]
		z[3] = t[3]
		z[4] = t[4]
		z[5] = t[5]
		z[6] = t[6]
		z[7] = t[7]
		z[8] = t[8]
		z[9] = t[9]
		z[10] = t[10]

	// final subtraction, overwriting z if z > mod
	if GTE(z, mod) {
		C = 0
		for i := 0; i < 11; i++ {
			z[i], C = bits.Sub64(z[i], mod[i], C)
		}
	}

	return nil
}




func mulMont768(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[12]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[12]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[12]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [13]uint64
	var D uint64
	var m, C uint64

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
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

    // TODO this shows up here, but I can't find reference to it in any paper
    // that references CIOS. is this just a quick hack for the final subtraction?
	if t[12] != 0 {
		// we need to reduce, we have a result on 13 words
		var b uint64
		z[0],b = bits.Sub64(t[0], mod[0], 0)
				z[1], b = bits.Sub64(t[1], mod[1], b)
				z[2], b = bits.Sub64(t[2], mod[2], b)
				z[3], b = bits.Sub64(t[3], mod[3], b)
				z[4], b = bits.Sub64(t[4], mod[4], b)
				z[5], b = bits.Sub64(t[5], mod[5], b)
				z[6], b = bits.Sub64(t[6], mod[6], b)
				z[7], b = bits.Sub64(t[7], mod[7], b)
				z[8], b = bits.Sub64(t[8], mod[8], b)
				z[9], b = bits.Sub64(t[9], mod[9], b)
				z[10], b = bits.Sub64(t[10], mod[10], b)
				z[11], _ = bits.Sub64(t[11], mod[11], b)
		return nil
	}

	// copy t into z
		z[0] = t[0]
		z[1] = t[1]
		z[2] = t[2]
		z[3] = t[3]
		z[4] = t[4]
		z[5] = t[5]
		z[6] = t[6]
		z[7] = t[7]
		z[8] = t[8]
		z[9] = t[9]
		z[10] = t[10]
		z[11] = t[11]

	// final subtraction, overwriting z if z > mod
	if GTE(z, mod) {
		C = 0
		for i := 0; i < 12; i++ {
			z[i], C = bits.Sub64(z[i], mod[i], C)
		}
	}

	return nil
}
// NOTE: this assumes that x and y are in Montgomery form and can produce unexpected results when they are not
func MulMontNonInterleaved(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	// length x == y assumed

	product := new(big.Int)
	x := LEBytesToInt(x_bytes)
	y := LEBytesToInt(y_bytes)

	if x.Cmp(f.ModulusNonInterleaved) > 0 || y.Cmp(f.ModulusNonInterleaved) > 0 {
		return errors.New("x/y >= modulus")
	}

	// m <- ((x*y mod R)N`) mod R
	product.Mul(x, y)
	x.And(product, f.mask)
	x.Mul(x, f.MontParamNonInterleaved)
	x.And(x, f.mask)

	// t <- (T + mN) / R
	x.Mul(x, f.ModulusNonInterleaved)
	x.Add(x, product)
	x.Rsh(x, f.NumLimbs*64)

	if x.Cmp(f.ModulusNonInterleaved) >= 0 {
		x.Sub(x, f.ModulusNonInterleaved)
	}

	copy(out_bytes, LimbsToLEBytes(IntToLimbs(x, f.NumLimbs)))

	return nil
}
