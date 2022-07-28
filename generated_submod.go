package mont_arith

import (
	"math/bits"
	"unsafe"
    "errors"
    "fmt"
)




func SubModUnrolled64(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[1]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[1]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[1]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [1]uint64 { 0}
			tmp[0], c = bits.Sub64(x[0], y[0], 0)
			z[0], c1 = bits.Add64(tmp[0], mod[0], 0)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModNonUnrolled64(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[1]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[1]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[1]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [1]uint64 { 0}

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func SubModUnrolled128(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [2]uint64 { 0, 0}
			tmp[0], c = bits.Sub64(x[0], y[0], 0)
			tmp[1], c = bits.Sub64(x[1], y[1], c)
			z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
			z[1], c1  = bits.Add64(tmp[1], mod[1], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModNonUnrolled128(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [2]uint64 { 0, 0}

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func SubModUnrolled192(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[3]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [3]uint64 { 0, 0, 0}
			tmp[0], c = bits.Sub64(x[0], y[0], 0)
			tmp[1], c = bits.Sub64(x[1], y[1], c)
			tmp[2], c = bits.Sub64(x[2], y[2], c)
			z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
			z[1], c1  = bits.Add64(tmp[1], mod[1], c1)
			z[2], c1  = bits.Add64(tmp[2], mod[2], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModNonUnrolled192(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[3]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [3]uint64 { 0, 0, 0}

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func SubModUnrolled256(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[4]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [4]uint64 { 0, 0, 0, 0}
			tmp[0], c = bits.Sub64(x[0], y[0], 0)
			tmp[1], c = bits.Sub64(x[1], y[1], c)
			tmp[2], c = bits.Sub64(x[2], y[2], c)
			tmp[3], c = bits.Sub64(x[3], y[3], c)
			z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
			z[1], c1  = bits.Add64(tmp[1], mod[1], c1)
			z[2], c1  = bits.Add64(tmp[2], mod[2], c1)
			z[3], c1  = bits.Add64(tmp[3], mod[3], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModNonUnrolled256(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[4]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [4]uint64 { 0, 0, 0, 0}

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func SubModUnrolled320(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[5]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[5]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[5]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[5]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [5]uint64 { 0, 0, 0, 0, 0}
			tmp[0], c = bits.Sub64(x[0], y[0], 0)
			tmp[1], c = bits.Sub64(x[1], y[1], c)
			tmp[2], c = bits.Sub64(x[2], y[2], c)
			tmp[3], c = bits.Sub64(x[3], y[3], c)
			tmp[4], c = bits.Sub64(x[4], y[4], c)
			z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
			z[1], c1  = bits.Add64(tmp[1], mod[1], c1)
			z[2], c1  = bits.Add64(tmp[2], mod[2], c1)
			z[3], c1  = bits.Add64(tmp[3], mod[3], c1)
			z[4], c1  = bits.Add64(tmp[4], mod[4], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModNonUnrolled320(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[5]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[5]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[5]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[5]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [5]uint64 { 0, 0, 0, 0, 0}

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func SubModUnrolled384(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [6]uint64 { 0, 0, 0, 0, 0, 0}
			tmp[0], c = bits.Sub64(x[0], y[0], 0)
			tmp[1], c = bits.Sub64(x[1], y[1], c)
			tmp[2], c = bits.Sub64(x[2], y[2], c)
			tmp[3], c = bits.Sub64(x[3], y[3], c)
			tmp[4], c = bits.Sub64(x[4], y[4], c)
			tmp[5], c = bits.Sub64(x[5], y[5], c)
			z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
			z[1], c1  = bits.Add64(tmp[1], mod[1], c1)
			z[2], c1  = bits.Add64(tmp[2], mod[2], c1)
			z[3], c1  = bits.Add64(tmp[3], mod[3], c1)
			z[4], c1  = bits.Add64(tmp[4], mod[4], c1)
			z[5], c1  = bits.Add64(tmp[5], mod[5], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModNonUnrolled384(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [6]uint64 { 0, 0, 0, 0, 0, 0}

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func SubModUnrolled448(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[7]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[6] >= mod[6] || y[6] >= mod[6] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [7]uint64 { 0, 0, 0, 0, 0, 0, 0}
			tmp[0], c = bits.Sub64(x[0], y[0], 0)
			tmp[1], c = bits.Sub64(x[1], y[1], c)
			tmp[2], c = bits.Sub64(x[2], y[2], c)
			tmp[3], c = bits.Sub64(x[3], y[3], c)
			tmp[4], c = bits.Sub64(x[4], y[4], c)
			tmp[5], c = bits.Sub64(x[5], y[5], c)
			tmp[6], c = bits.Sub64(x[6], y[6], c)
			z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
			z[1], c1  = bits.Add64(tmp[1], mod[1], c1)
			z[2], c1  = bits.Add64(tmp[2], mod[2], c1)
			z[3], c1  = bits.Add64(tmp[3], mod[3], c1)
			z[4], c1  = bits.Add64(tmp[4], mod[4], c1)
			z[5], c1  = bits.Add64(tmp[5], mod[5], c1)
			z[6], c1  = bits.Add64(tmp[6], mod[6], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModNonUnrolled448(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[7]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[6] >= mod[6] || y[6] >= mod[6] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [7]uint64 { 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func SubModUnrolled512(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[8]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[8]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[8]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[8]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[6] >= mod[6] || y[6] >= mod[6] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[7] >= mod[7] || y[7] >= mod[7] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [8]uint64 { 0, 0, 0, 0, 0, 0, 0, 0}
			tmp[0], c = bits.Sub64(x[0], y[0], 0)
			tmp[1], c = bits.Sub64(x[1], y[1], c)
			tmp[2], c = bits.Sub64(x[2], y[2], c)
			tmp[3], c = bits.Sub64(x[3], y[3], c)
			tmp[4], c = bits.Sub64(x[4], y[4], c)
			tmp[5], c = bits.Sub64(x[5], y[5], c)
			tmp[6], c = bits.Sub64(x[6], y[6], c)
			tmp[7], c = bits.Sub64(x[7], y[7], c)
			z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
			z[1], c1  = bits.Add64(tmp[1], mod[1], c1)
			z[2], c1  = bits.Add64(tmp[2], mod[2], c1)
			z[3], c1  = bits.Add64(tmp[3], mod[3], c1)
			z[4], c1  = bits.Add64(tmp[4], mod[4], c1)
			z[5], c1  = bits.Add64(tmp[5], mod[5], c1)
			z[6], c1  = bits.Add64(tmp[6], mod[6], c1)
			z[7], c1  = bits.Add64(tmp[7], mod[7], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModNonUnrolled512(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[8]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[8]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[8]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[8]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[6] >= mod[6] || y[6] >= mod[6] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[7] >= mod[7] || y[7] >= mod[7] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [8]uint64 { 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func SubModUnrolled576(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[6] >= mod[6] || y[6] >= mod[6] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[7] >= mod[7] || y[7] >= mod[7] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[8] >= mod[8] || y[8] >= mod[8] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [9]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0}
			tmp[0], c = bits.Sub64(x[0], y[0], 0)
			tmp[1], c = bits.Sub64(x[1], y[1], c)
			tmp[2], c = bits.Sub64(x[2], y[2], c)
			tmp[3], c = bits.Sub64(x[3], y[3], c)
			tmp[4], c = bits.Sub64(x[4], y[4], c)
			tmp[5], c = bits.Sub64(x[5], y[5], c)
			tmp[6], c = bits.Sub64(x[6], y[6], c)
			tmp[7], c = bits.Sub64(x[7], y[7], c)
			tmp[8], c = bits.Sub64(x[8], y[8], c)
			z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
			z[1], c1  = bits.Add64(tmp[1], mod[1], c1)
			z[2], c1  = bits.Add64(tmp[2], mod[2], c1)
			z[3], c1  = bits.Add64(tmp[3], mod[3], c1)
			z[4], c1  = bits.Add64(tmp[4], mod[4], c1)
			z[5], c1  = bits.Add64(tmp[5], mod[5], c1)
			z[6], c1  = bits.Add64(tmp[6], mod[6], c1)
			z[7], c1  = bits.Add64(tmp[7], mod[7], c1)
			z[8], c1  = bits.Add64(tmp[8], mod[8], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModNonUnrolled576(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[6] >= mod[6] || y[6] >= mod[6] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[7] >= mod[7] || y[7] >= mod[7] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[8] >= mod[8] || y[8] >= mod[8] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [9]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func SubModUnrolled640(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[10]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[6] >= mod[6] || y[6] >= mod[6] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[7] >= mod[7] || y[7] >= mod[7] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[8] >= mod[8] || y[8] >= mod[8] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[9] >= mod[9] || y[9] >= mod[9] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [10]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			tmp[0], c = bits.Sub64(x[0], y[0], 0)
			tmp[1], c = bits.Sub64(x[1], y[1], c)
			tmp[2], c = bits.Sub64(x[2], y[2], c)
			tmp[3], c = bits.Sub64(x[3], y[3], c)
			tmp[4], c = bits.Sub64(x[4], y[4], c)
			tmp[5], c = bits.Sub64(x[5], y[5], c)
			tmp[6], c = bits.Sub64(x[6], y[6], c)
			tmp[7], c = bits.Sub64(x[7], y[7], c)
			tmp[8], c = bits.Sub64(x[8], y[8], c)
			tmp[9], c = bits.Sub64(x[9], y[9], c)
			z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
			z[1], c1  = bits.Add64(tmp[1], mod[1], c1)
			z[2], c1  = bits.Add64(tmp[2], mod[2], c1)
			z[3], c1  = bits.Add64(tmp[3], mod[3], c1)
			z[4], c1  = bits.Add64(tmp[4], mod[4], c1)
			z[5], c1  = bits.Add64(tmp[5], mod[5], c1)
			z[6], c1  = bits.Add64(tmp[6], mod[6], c1)
			z[7], c1  = bits.Add64(tmp[7], mod[7], c1)
			z[8], c1  = bits.Add64(tmp[8], mod[8], c1)
			z[9], c1  = bits.Add64(tmp[9], mod[9], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModNonUnrolled640(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[10]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[6] >= mod[6] || y[6] >= mod[6] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[7] >= mod[7] || y[7] >= mod[7] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[8] >= mod[8] || y[8] >= mod[8] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[9] >= mod[9] || y[9] >= mod[9] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [10]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func SubModUnrolled704(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[11]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[11]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[11]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[11]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[6] >= mod[6] || y[6] >= mod[6] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[7] >= mod[7] || y[7] >= mod[7] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[8] >= mod[8] || y[8] >= mod[8] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[9] >= mod[9] || y[9] >= mod[9] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[10] >= mod[10] || y[10] >= mod[10] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [11]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			tmp[0], c = bits.Sub64(x[0], y[0], 0)
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
			z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
			z[1], c1  = bits.Add64(tmp[1], mod[1], c1)
			z[2], c1  = bits.Add64(tmp[2], mod[2], c1)
			z[3], c1  = bits.Add64(tmp[3], mod[3], c1)
			z[4], c1  = bits.Add64(tmp[4], mod[4], c1)
			z[5], c1  = bits.Add64(tmp[5], mod[5], c1)
			z[6], c1  = bits.Add64(tmp[6], mod[6], c1)
			z[7], c1  = bits.Add64(tmp[7], mod[7], c1)
			z[8], c1  = bits.Add64(tmp[8], mod[8], c1)
			z[9], c1  = bits.Add64(tmp[9], mod[9], c1)
			z[10], c1  = bits.Add64(tmp[10], mod[10], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModNonUnrolled704(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[11]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[11]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[11]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[11]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[6] >= mod[6] || y[6] >= mod[6] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[7] >= mod[7] || y[7] >= mod[7] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[8] >= mod[8] || y[8] >= mod[8] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[9] >= mod[9] || y[9] >= mod[9] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[10] >= mod[10] || y[10] >= mod[10] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [11]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func SubModUnrolled768(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[12]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[12]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[12]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[6] >= mod[6] || y[6] >= mod[6] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[7] >= mod[7] || y[7] >= mod[7] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[8] >= mod[8] || y[8] >= mod[8] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[9] >= mod[9] || y[9] >= mod[9] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[10] >= mod[10] || y[10] >= mod[10] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[11] >= mod[11] || y[11] >= mod[11] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [12]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			tmp[0], c = bits.Sub64(x[0], y[0], 0)
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
			z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
			z[1], c1  = bits.Add64(tmp[1], mod[1], c1)
			z[2], c1  = bits.Add64(tmp[2], mod[2], c1)
			z[3], c1  = bits.Add64(tmp[3], mod[3], c1)
			z[4], c1  = bits.Add64(tmp[4], mod[4], c1)
			z[5], c1  = bits.Add64(tmp[5], mod[5], c1)
			z[6], c1  = bits.Add64(tmp[6], mod[6], c1)
			z[7], c1  = bits.Add64(tmp[7], mod[7], c1)
			z[8], c1  = bits.Add64(tmp[8], mod[8], c1)
			z[9], c1  = bits.Add64(tmp[9], mod[9], c1)
			z[10], c1  = bits.Add64(tmp[10], mod[10], c1)
			z[11], c1  = bits.Add64(tmp[11], mod[11], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModNonUnrolled768(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[12]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[12]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[12]uint64)(unsafe.Pointer(&f.Modulus[0]))[:]

    // TODO move bounds check into its own template?
        if x[0] >= mod[0] || y[0] >= mod[0] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[1] >= mod[1] || y[1] >= mod[1] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[2] >= mod[2] || y[2] >= mod[2] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[3] >= mod[3] || y[3] >= mod[3] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[4] >= mod[4] || y[4] >= mod[4] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[5] >= mod[5] || y[5] >= mod[5] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[6] >= mod[6] || y[6] >= mod[6] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[7] >= mod[7] || y[7] >= mod[7] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[8] >= mod[8] || y[8] >= mod[8] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[9] >= mod[9] || y[9] >= mod[9] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[10] >= mod[10] || y[10] >= mod[10] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }
        if x[11] >= mod[11] || y[11] >= mod[11] {
            return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
        }

	var c, c1 uint64
	tmp := [12]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

