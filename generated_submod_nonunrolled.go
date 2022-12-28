package mont_arith

import (
	"math/bits"
	"unsafe"
    "errors"
    "fmt"
)






func SubModNonUnrolled64(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [1]uint64 { 0}

    for i := 0; i < 1; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 1; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled128(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [2]uint64 { 0, 0}

    for i := 0; i < 2; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 2; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled192(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [3]uint64 { 0, 0, 0}

    for i := 0; i < 3; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 3; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled256(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [4]uint64 { 0, 0, 0, 0}

    for i := 0; i < 4; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 4; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled320(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [5]uint64 { 0, 0, 0, 0, 0}

    for i := 0; i < 5; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 5; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled384(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [6]uint64 { 0, 0, 0, 0, 0, 0}

    for i := 0; i < 6; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 6; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled448(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [7]uint64 { 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 7; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 7; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled512(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [8]uint64 { 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 8; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 8; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled576(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [9]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 9; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 9; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled640(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [10]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 10; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 10; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled704(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [11]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 11; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 11; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled768(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [12]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 12; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 12; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled832(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [13]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 13; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 13; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled896(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [14]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 14; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 14; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled960(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [15]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 15; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 15; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}






func SubModNonUnrolled1024(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
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


	var c, c1 uint64
	tmp := [16]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 16; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 16; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

	if c == 0 {
		copy(z, tmp[:])
    }/*else {
        panic("not worst case performance")
    }*/

	return nil
}
