



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




var Zero2Limbs []uint = make([]uint, 2, 2)

/* NOTE: addmod/submod/mulmodmont assume:
    len(z) == len(x) == len(y) == len(mod)
    and
    x < mod, y < mod
*/

func mulMont128(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
    	var t [2]uint64
	var c [3]uint64
	modinv := ctx.MontParamInterleaved

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
		// round 0
			v := x[0]
			c[1], c[0] = bits.Mul64(v, y[0])
			m := c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd1(v, y[1], c[1])
					t[1], t[0]  = madd3(m, mod[1], c[0], c[2], c[1])
		// round 1
			v = x[1]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1],  c[1], t[1])
					z[1], z[0] = madd3(m, mod[1], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 2; i++ {
		t[i], c[0] = bits.Sub64(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}




var Zero3Limbs []uint = make([]uint, 3, 3)

/* NOTE: addmod/submod/mulmodmont assume:
    len(z) == len(x) == len(y) == len(mod)
    and
    x < mod, y < mod
*/

func mulMont192(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[3]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
    	var t [3]uint64
	var c [3]uint64
	modinv := ctx.MontParamInterleaved

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
		// round 0
			v := x[0]
			c[1], c[0] = bits.Mul64(v, y[0])
			m := c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd1(v, y[1], c[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd1(v, y[2], c[1])
					t[2], t[1]  = madd3(m, mod[2], c[0], c[2], c[1])
		// round 1
			v = x[1]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					t[2], t[1] = madd3(m, mod[2], c[0], c[2], c[1])
		// round 2
			v = x[2]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1],  c[1], t[1])
					c[2], z[0] = madd2(m, mod[1],  c[2], c[0])
				c[1], c[0] = madd2(v, y[2],  c[1], t[2])
					z[2], z[1] = madd3(m, mod[2], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 3; i++ {
		t[i], c[0] = bits.Sub64(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}




var Zero4Limbs []uint = make([]uint, 4, 4)

/* NOTE: addmod/submod/mulmodmont assume:
    len(z) == len(x) == len(y) == len(mod)
    and
    x < mod, y < mod
*/

func mulMont256(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[4]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
    	var t [4]uint64
	var c [4]uint64
	modinv := ctx.MontParamInterleaved

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
		// round 0
			v := x[0]
			c[1], c[0] = bits.Mul64(v, y[0])
			m := c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd1(v, y[1], c[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd1(v, y[2], c[1])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd1(v, y[3], c[1])
					t[3], t[2]  = madd3(m, mod[3], c[0], c[2], c[1])
		// round 1
			v = x[1]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					t[3], t[2] = madd3(m, mod[3], c[0], c[2], c[1])
		// round 2
			v = x[2]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					t[3], t[2] = madd3(m, mod[3], c[0], c[2], c[1])
		// round 3
			v = x[3]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1],  c[1], t[1])
					c[2], z[0] = madd2(m, mod[1],  c[2], c[0])
				c[1], c[0] = madd2(v, y[2],  c[1], t[2])
					c[2], z[1] = madd2(m, mod[2],  c[2], c[0])
				c[1], c[0] = madd2(v, y[3],  c[1], t[3])
					z[3], z[2] = madd3(m, mod[3], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 4; i++ {
		t[i], c[0] = bits.Sub64(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}




var Zero5Limbs []uint = make([]uint, 5, 5)

/* NOTE: addmod/submod/mulmodmont assume:
    len(z) == len(x) == len(y) == len(mod)
    and
    x < mod, y < mod
*/

func mulMont320(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[5]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[5]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[5]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[5]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
    	var t [5]uint64
	var c [5]uint64
	modinv := ctx.MontParamInterleaved

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
		// round 0
			v := x[0]
			c[1], c[0] = bits.Mul64(v, y[0])
			m := c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd1(v, y[1], c[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd1(v, y[2], c[1])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd1(v, y[3], c[1])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd1(v, y[4], c[1])
					t[4], t[3]  = madd3(m, mod[4], c[0], c[2], c[1])
		// round 1
			v = x[1]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					t[4], t[3] = madd3(m, mod[4], c[0], c[2], c[1])
		// round 2
			v = x[2]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					t[4], t[3] = madd3(m, mod[4], c[0], c[2], c[1])
		// round 3
			v = x[3]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					t[4], t[3] = madd3(m, mod[4], c[0], c[2], c[1])
		// round 4
			v = x[4]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1],  c[1], t[1])
					c[2], z[0] = madd2(m, mod[1],  c[2], c[0])
				c[1], c[0] = madd2(v, y[2],  c[1], t[2])
					c[2], z[1] = madd2(m, mod[2],  c[2], c[0])
				c[1], c[0] = madd2(v, y[3],  c[1], t[3])
					c[2], z[2] = madd2(m, mod[3],  c[2], c[0])
				c[1], c[0] = madd2(v, y[4],  c[1], t[4])
					z[4], z[3] = madd3(m, mod[4], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 5; i++ {
		t[i], c[0] = bits.Sub64(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}




var Zero6Limbs []uint = make([]uint, 6, 6)

/* NOTE: addmod/submod/mulmodmont assume:
    len(z) == len(x) == len(y) == len(mod)
    and
    x < mod, y < mod
*/

func mulMont384(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
    	var t [6]uint64
	var c [6]uint64
	modinv := ctx.MontParamInterleaved

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
		// round 0
			v := x[0]
			c[1], c[0] = bits.Mul64(v, y[0])
			m := c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd1(v, y[1], c[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd1(v, y[2], c[1])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd1(v, y[3], c[1])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd1(v, y[4], c[1])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd1(v, y[5], c[1])
					t[5], t[4]  = madd3(m, mod[5], c[0], c[2], c[1])
		// round 1
			v = x[1]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					t[5], t[4] = madd3(m, mod[5], c[0], c[2], c[1])
		// round 2
			v = x[2]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					t[5], t[4] = madd3(m, mod[5], c[0], c[2], c[1])
		// round 3
			v = x[3]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					t[5], t[4] = madd3(m, mod[5], c[0], c[2], c[1])
		// round 4
			v = x[4]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					t[5], t[4] = madd3(m, mod[5], c[0], c[2], c[1])
		// round 5
			v = x[5]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1],  c[1], t[1])
					c[2], z[0] = madd2(m, mod[1],  c[2], c[0])
				c[1], c[0] = madd2(v, y[2],  c[1], t[2])
					c[2], z[1] = madd2(m, mod[2],  c[2], c[0])
				c[1], c[0] = madd2(v, y[3],  c[1], t[3])
					c[2], z[2] = madd2(m, mod[3],  c[2], c[0])
				c[1], c[0] = madd2(v, y[4],  c[1], t[4])
					c[2], z[3] = madd2(m, mod[4],  c[2], c[0])
				c[1], c[0] = madd2(v, y[5],  c[1], t[5])
					z[5], z[4] = madd3(m, mod[5], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 6; i++ {
		t[i], c[0] = bits.Sub64(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}




var Zero7Limbs []uint = make([]uint, 7, 7)

/* NOTE: addmod/submod/mulmodmont assume:
    len(z) == len(x) == len(y) == len(mod)
    and
    x < mod, y < mod
*/

func mulMont448(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[7]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
    	var t [7]uint64
	var c [7]uint64
	modinv := ctx.MontParamInterleaved

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
		// round 0
			v := x[0]
			c[1], c[0] = bits.Mul64(v, y[0])
			m := c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd1(v, y[1], c[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd1(v, y[2], c[1])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd1(v, y[3], c[1])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd1(v, y[4], c[1])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd1(v, y[5], c[1])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd1(v, y[6], c[1])
					t[6], t[5]  = madd3(m, mod[6], c[0], c[2], c[1])
		// round 1
			v = x[1]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					t[6], t[5] = madd3(m, mod[6], c[0], c[2], c[1])
		// round 2
			v = x[2]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					t[6], t[5] = madd3(m, mod[6], c[0], c[2], c[1])
		// round 3
			v = x[3]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					t[6], t[5] = madd3(m, mod[6], c[0], c[2], c[1])
		// round 4
			v = x[4]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					t[6], t[5] = madd3(m, mod[6], c[0], c[2], c[1])
		// round 5
			v = x[5]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					t[6], t[5] = madd3(m, mod[6], c[0], c[2], c[1])
		// round 6
			v = x[6]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1],  c[1], t[1])
					c[2], z[0] = madd2(m, mod[1],  c[2], c[0])
				c[1], c[0] = madd2(v, y[2],  c[1], t[2])
					c[2], z[1] = madd2(m, mod[2],  c[2], c[0])
				c[1], c[0] = madd2(v, y[3],  c[1], t[3])
					c[2], z[2] = madd2(m, mod[3],  c[2], c[0])
				c[1], c[0] = madd2(v, y[4],  c[1], t[4])
					c[2], z[3] = madd2(m, mod[4],  c[2], c[0])
				c[1], c[0] = madd2(v, y[5],  c[1], t[5])
					c[2], z[4] = madd2(m, mod[5],  c[2], c[0])
				c[1], c[0] = madd2(v, y[6],  c[1], t[6])
					z[6], z[5] = madd3(m, mod[6], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 7; i++ {
		t[i], c[0] = bits.Sub64(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}




var Zero8Limbs []uint = make([]uint, 8, 8)

/* NOTE: addmod/submod/mulmodmont assume:
    len(z) == len(x) == len(y) == len(mod)
    and
    x < mod, y < mod
*/

func mulMont512(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[8]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[8]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[8]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[8]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
    	var t [8]uint64
	var c [8]uint64
	modinv := ctx.MontParamInterleaved

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
		// round 0
			v := x[0]
			c[1], c[0] = bits.Mul64(v, y[0])
			m := c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd1(v, y[1], c[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd1(v, y[2], c[1])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd1(v, y[3], c[1])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd1(v, y[4], c[1])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd1(v, y[5], c[1])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd1(v, y[6], c[1])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd1(v, y[7], c[1])
					t[7], t[6]  = madd3(m, mod[7], c[0], c[2], c[1])
		// round 1
			v = x[1]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					t[7], t[6] = madd3(m, mod[7], c[0], c[2], c[1])
		// round 2
			v = x[2]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					t[7], t[6] = madd3(m, mod[7], c[0], c[2], c[1])
		// round 3
			v = x[3]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					t[7], t[6] = madd3(m, mod[7], c[0], c[2], c[1])
		// round 4
			v = x[4]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					t[7], t[6] = madd3(m, mod[7], c[0], c[2], c[1])
		// round 5
			v = x[5]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					t[7], t[6] = madd3(m, mod[7], c[0], c[2], c[1])
		// round 6
			v = x[6]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					t[7], t[6] = madd3(m, mod[7], c[0], c[2], c[1])
		// round 7
			v = x[7]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1],  c[1], t[1])
					c[2], z[0] = madd2(m, mod[1],  c[2], c[0])
				c[1], c[0] = madd2(v, y[2],  c[1], t[2])
					c[2], z[1] = madd2(m, mod[2],  c[2], c[0])
				c[1], c[0] = madd2(v, y[3],  c[1], t[3])
					c[2], z[2] = madd2(m, mod[3],  c[2], c[0])
				c[1], c[0] = madd2(v, y[4],  c[1], t[4])
					c[2], z[3] = madd2(m, mod[4],  c[2], c[0])
				c[1], c[0] = madd2(v, y[5],  c[1], t[5])
					c[2], z[4] = madd2(m, mod[5],  c[2], c[0])
				c[1], c[0] = madd2(v, y[6],  c[1], t[6])
					c[2], z[5] = madd2(m, mod[6],  c[2], c[0])
				c[1], c[0] = madd2(v, y[7],  c[1], t[7])
					z[7], z[6] = madd3(m, mod[7], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 8; i++ {
		t[i], c[0] = bits.Sub64(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}




var Zero9Limbs []uint = make([]uint, 9, 9)

/* NOTE: addmod/submod/mulmodmont assume:
    len(z) == len(x) == len(y) == len(mod)
    and
    x < mod, y < mod
*/

func mulMont576(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
    	var t [9]uint64
	var c [9]uint64
	modinv := ctx.MontParamInterleaved

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
		// round 0
			v := x[0]
			c[1], c[0] = bits.Mul64(v, y[0])
			m := c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd1(v, y[1], c[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd1(v, y[2], c[1])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd1(v, y[3], c[1])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd1(v, y[4], c[1])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd1(v, y[5], c[1])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd1(v, y[6], c[1])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd1(v, y[7], c[1])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd1(v, y[8], c[1])
					t[8], t[7]  = madd3(m, mod[8], c[0], c[2], c[1])
		// round 1
			v = x[1]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
		// round 2
			v = x[2]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
		// round 3
			v = x[3]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
		// round 4
			v = x[4]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
		// round 5
			v = x[5]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
		// round 6
			v = x[6]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
		// round 7
			v = x[7]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
		// round 8
			v = x[8]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1],  c[1], t[1])
					c[2], z[0] = madd2(m, mod[1],  c[2], c[0])
				c[1], c[0] = madd2(v, y[2],  c[1], t[2])
					c[2], z[1] = madd2(m, mod[2],  c[2], c[0])
				c[1], c[0] = madd2(v, y[3],  c[1], t[3])
					c[2], z[2] = madd2(m, mod[3],  c[2], c[0])
				c[1], c[0] = madd2(v, y[4],  c[1], t[4])
					c[2], z[3] = madd2(m, mod[4],  c[2], c[0])
				c[1], c[0] = madd2(v, y[5],  c[1], t[5])
					c[2], z[4] = madd2(m, mod[5],  c[2], c[0])
				c[1], c[0] = madd2(v, y[6],  c[1], t[6])
					c[2], z[5] = madd2(m, mod[6],  c[2], c[0])
				c[1], c[0] = madd2(v, y[7],  c[1], t[7])
					c[2], z[6] = madd2(m, mod[7],  c[2], c[0])
				c[1], c[0] = madd2(v, y[8],  c[1], t[8])
					z[8], z[7] = madd3(m, mod[8], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 9; i++ {
		t[i], c[0] = bits.Sub64(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}




var Zero10Limbs []uint = make([]uint, 10, 10)

/* NOTE: addmod/submod/mulmodmont assume:
    len(z) == len(x) == len(y) == len(mod)
    and
    x < mod, y < mod
*/

func mulMont640(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[10]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
    	var t [10]uint64
	var c [10]uint64
	modinv := ctx.MontParamInterleaved

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
		// round 0
			v := x[0]
			c[1], c[0] = bits.Mul64(v, y[0])
			m := c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd1(v, y[1], c[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd1(v, y[2], c[1])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd1(v, y[3], c[1])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd1(v, y[4], c[1])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd1(v, y[5], c[1])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd1(v, y[6], c[1])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd1(v, y[7], c[1])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd1(v, y[8], c[1])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd1(v, y[9], c[1])
					t[9], t[8]  = madd3(m, mod[9], c[0], c[2], c[1])
		// round 1
			v = x[1]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
		// round 2
			v = x[2]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
		// round 3
			v = x[3]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
		// round 4
			v = x[4]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
		// round 5
			v = x[5]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
		// round 6
			v = x[6]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
		// round 7
			v = x[7]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
		// round 8
			v = x[8]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
		// round 9
			v = x[9]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1],  c[1], t[1])
					c[2], z[0] = madd2(m, mod[1],  c[2], c[0])
				c[1], c[0] = madd2(v, y[2],  c[1], t[2])
					c[2], z[1] = madd2(m, mod[2],  c[2], c[0])
				c[1], c[0] = madd2(v, y[3],  c[1], t[3])
					c[2], z[2] = madd2(m, mod[3],  c[2], c[0])
				c[1], c[0] = madd2(v, y[4],  c[1], t[4])
					c[2], z[3] = madd2(m, mod[4],  c[2], c[0])
				c[1], c[0] = madd2(v, y[5],  c[1], t[5])
					c[2], z[4] = madd2(m, mod[5],  c[2], c[0])
				c[1], c[0] = madd2(v, y[6],  c[1], t[6])
					c[2], z[5] = madd2(m, mod[6],  c[2], c[0])
				c[1], c[0] = madd2(v, y[7],  c[1], t[7])
					c[2], z[6] = madd2(m, mod[7],  c[2], c[0])
				c[1], c[0] = madd2(v, y[8],  c[1], t[8])
					c[2], z[7] = madd2(m, mod[8],  c[2], c[0])
				c[1], c[0] = madd2(v, y[9],  c[1], t[9])
					z[9], z[8] = madd3(m, mod[9], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 10; i++ {
		t[i], c[0] = bits.Sub64(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}




var Zero11Limbs []uint = make([]uint, 11, 11)

/* NOTE: addmod/submod/mulmodmont assume:
    len(z) == len(x) == len(y) == len(mod)
    and
    x < mod, y < mod
*/

func mulMont704(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[11]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[11]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[11]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[11]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
    	var t [11]uint64
	var c [11]uint64
	modinv := ctx.MontParamInterleaved

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
		// round 0
			v := x[0]
			c[1], c[0] = bits.Mul64(v, y[0])
			m := c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd1(v, y[1], c[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd1(v, y[2], c[1])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd1(v, y[3], c[1])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd1(v, y[4], c[1])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd1(v, y[5], c[1])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd1(v, y[6], c[1])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd1(v, y[7], c[1])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd1(v, y[8], c[1])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd1(v, y[9], c[1])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd1(v, y[10], c[1])
					t[10], t[9]  = madd3(m, mod[10], c[0], c[2], c[1])
		// round 1
			v = x[1]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
		// round 2
			v = x[2]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
		// round 3
			v = x[3]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
		// round 4
			v = x[4]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
		// round 5
			v = x[5]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
		// round 6
			v = x[6]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
		// round 7
			v = x[7]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
		// round 8
			v = x[8]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
		// round 9
			v = x[9]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
		// round 10
			v = x[10]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1],  c[1], t[1])
					c[2], z[0] = madd2(m, mod[1],  c[2], c[0])
				c[1], c[0] = madd2(v, y[2],  c[1], t[2])
					c[2], z[1] = madd2(m, mod[2],  c[2], c[0])
				c[1], c[0] = madd2(v, y[3],  c[1], t[3])
					c[2], z[2] = madd2(m, mod[3],  c[2], c[0])
				c[1], c[0] = madd2(v, y[4],  c[1], t[4])
					c[2], z[3] = madd2(m, mod[4],  c[2], c[0])
				c[1], c[0] = madd2(v, y[5],  c[1], t[5])
					c[2], z[4] = madd2(m, mod[5],  c[2], c[0])
				c[1], c[0] = madd2(v, y[6],  c[1], t[6])
					c[2], z[5] = madd2(m, mod[6],  c[2], c[0])
				c[1], c[0] = madd2(v, y[7],  c[1], t[7])
					c[2], z[6] = madd2(m, mod[7],  c[2], c[0])
				c[1], c[0] = madd2(v, y[8],  c[1], t[8])
					c[2], z[7] = madd2(m, mod[8],  c[2], c[0])
				c[1], c[0] = madd2(v, y[9],  c[1], t[9])
					c[2], z[8] = madd2(m, mod[9],  c[2], c[0])
				c[1], c[0] = madd2(v, y[10],  c[1], t[10])
					z[10], z[9] = madd3(m, mod[10], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 11; i++ {
		t[i], c[0] = bits.Sub64(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}




var Zero12Limbs []uint = make([]uint, 12, 12)

/* NOTE: addmod/submod/mulmodmont assume:
    len(z) == len(x) == len(y) == len(mod)
    and
    x < mod, y < mod
*/

func mulMont768(ctx *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[12]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[12]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[12]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
    	var t [12]uint64
	var c [12]uint64
	modinv := ctx.MontParamInterleaved

    if GTE(x, mod) || GTE(y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }
		// round 0
			v := x[0]
			c[1], c[0] = bits.Mul64(v, y[0])
			m := c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd1(v, y[1], c[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd1(v, y[2], c[1])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd1(v, y[3], c[1])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd1(v, y[4], c[1])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd1(v, y[5], c[1])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd1(v, y[6], c[1])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd1(v, y[7], c[1])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd1(v, y[8], c[1])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd1(v, y[9], c[1])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd1(v, y[10], c[1])
					c[2], t[9] = madd2(m, mod[10], c[2], c[0])
				c[1], c[0] = madd1(v, y[11], c[1])
					t[11], t[10]  = madd3(m, mod[11], c[0], c[2], c[1])
		// round 1
			v = x[1]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					c[2], t[9] = madd2(m, mod[10], c[2], c[0])
				c[1], c[0] = madd2(v, y[11], c[1], t[11])
					t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
		// round 2
			v = x[2]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					c[2], t[9] = madd2(m, mod[10], c[2], c[0])
				c[1], c[0] = madd2(v, y[11], c[1], t[11])
					t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
		// round 3
			v = x[3]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					c[2], t[9] = madd2(m, mod[10], c[2], c[0])
				c[1], c[0] = madd2(v, y[11], c[1], t[11])
					t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
		// round 4
			v = x[4]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					c[2], t[9] = madd2(m, mod[10], c[2], c[0])
				c[1], c[0] = madd2(v, y[11], c[1], t[11])
					t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
		// round 5
			v = x[5]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					c[2], t[9] = madd2(m, mod[10], c[2], c[0])
				c[1], c[0] = madd2(v, y[11], c[1], t[11])
					t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
		// round 6
			v = x[6]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					c[2], t[9] = madd2(m, mod[10], c[2], c[0])
				c[1], c[0] = madd2(v, y[11], c[1], t[11])
					t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
		// round 7
			v = x[7]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					c[2], t[9] = madd2(m, mod[10], c[2], c[0])
				c[1], c[0] = madd2(v, y[11], c[1], t[11])
					t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
		// round 8
			v = x[8]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					c[2], t[9] = madd2(m, mod[10], c[2], c[0])
				c[1], c[0] = madd2(v, y[11], c[1], t[11])
					t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
		// round 9
			v = x[9]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					c[2], t[9] = madd2(m, mod[10], c[2], c[0])
				c[1], c[0] = madd2(v, y[11], c[1], t[11])
					t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
		// round 10
			v = x[10]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1], c[1], t[1])
					c[2], t[0] = madd2(m, mod[1], c[2], c[0])
				c[1], c[0] = madd2(v, y[2], c[1], t[2])
					c[2], t[1] = madd2(m, mod[2], c[2], c[0])
				c[1], c[0] = madd2(v, y[3], c[1], t[3])
					c[2], t[2] = madd2(m, mod[3], c[2], c[0])
				c[1], c[0] = madd2(v, y[4], c[1], t[4])
					c[2], t[3] = madd2(m, mod[4], c[2], c[0])
				c[1], c[0] = madd2(v, y[5], c[1], t[5])
					c[2], t[4] = madd2(m, mod[5], c[2], c[0])
				c[1], c[0] = madd2(v, y[6], c[1], t[6])
					c[2], t[5] = madd2(m, mod[6], c[2], c[0])
				c[1], c[0] = madd2(v, y[7], c[1], t[7])
					c[2], t[6] = madd2(m, mod[7], c[2], c[0])
				c[1], c[0] = madd2(v, y[8], c[1], t[8])
					c[2], t[7] = madd2(m, mod[8], c[2], c[0])
				c[1], c[0] = madd2(v, y[9], c[1], t[9])
					c[2], t[8] = madd2(m, mod[9], c[2], c[0])
				c[1], c[0] = madd2(v, y[10], c[1], t[10])
					c[2], t[9] = madd2(m, mod[10], c[2], c[0])
				c[1], c[0] = madd2(v, y[11], c[1], t[11])
					t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
		// round 11
			v = x[11]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1],  c[1], t[1])
					c[2], z[0] = madd2(m, mod[1],  c[2], c[0])
				c[1], c[0] = madd2(v, y[2],  c[1], t[2])
					c[2], z[1] = madd2(m, mod[2],  c[2], c[0])
				c[1], c[0] = madd2(v, y[3],  c[1], t[3])
					c[2], z[2] = madd2(m, mod[3],  c[2], c[0])
				c[1], c[0] = madd2(v, y[4],  c[1], t[4])
					c[2], z[3] = madd2(m, mod[4],  c[2], c[0])
				c[1], c[0] = madd2(v, y[5],  c[1], t[5])
					c[2], z[4] = madd2(m, mod[5],  c[2], c[0])
				c[1], c[0] = madd2(v, y[6],  c[1], t[6])
					c[2], z[5] = madd2(m, mod[6],  c[2], c[0])
				c[1], c[0] = madd2(v, y[7],  c[1], t[7])
					c[2], z[6] = madd2(m, mod[7],  c[2], c[0])
				c[1], c[0] = madd2(v, y[8],  c[1], t[8])
					c[2], z[7] = madd2(m, mod[8],  c[2], c[0])
				c[1], c[0] = madd2(v, y[9],  c[1], t[9])
					c[2], z[8] = madd2(m, mod[9],  c[2], c[0])
				c[1], c[0] = madd2(v, y[10],  c[1], t[10])
					c[2], z[9] = madd2(m, mod[10],  c[2], c[0])
				c[1], c[0] = madd2(v, y[11],  c[1], t[11])
					z[11], z[10] = madd3(m, mod[11], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 12; i++ {
		t[i], c[0] = bits.Sub64(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
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
