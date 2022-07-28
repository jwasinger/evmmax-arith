package mont_arith

import (
	"math/bits"
	"unsafe"
    "errors"
    "fmt"
)




// TODO check unrolled speed
func AddModNonUnrolled64(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[1]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[1]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
    mod := f.Modulus

    if GTE(x, y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }

    var c uint64 = 0
    tmp := make([]uint64, len(mod))

    for i := 0; i < 1; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    c = 0
    for i := 0; i < 1; i++ {
        z[i], c = bits.Sub64(tmp[i], mod[i], c)
    }

    // final sub was unnecessary
    if c != 0 {
        copy(z, tmp[:])
    }
    return nil
}




// TODO check unrolled speed
func AddModNonUnrolled128(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
    mod := f.Modulus

    if GTE(x, y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }

    var c uint64 = 0
    tmp := make([]uint64, len(mod))

    for i := 0; i < 2; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    c = 0
    for i := 0; i < 2; i++ {
        z[i], c = bits.Sub64(tmp[i], mod[i], c)
    }

    // final sub was unnecessary
    if c != 0 {
        copy(z, tmp[:])
    }
    return nil
}




// TODO check unrolled speed
func AddModNonUnrolled192(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
    mod := f.Modulus

    if GTE(x, y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }

    var c uint64 = 0
    tmp := make([]uint64, len(mod))

    for i := 0; i < 3; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    c = 0
    for i := 0; i < 3; i++ {
        z[i], c = bits.Sub64(tmp[i], mod[i], c)
    }

    // final sub was unnecessary
    if c != 0 {
        copy(z, tmp[:])
    }
    return nil
}




// TODO check unrolled speed
func AddModNonUnrolled256(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
    mod := f.Modulus

    if GTE(x, y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }

    var c uint64 = 0
    tmp := make([]uint64, len(mod))

    for i := 0; i < 4; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    c = 0
    for i := 0; i < 4; i++ {
        z[i], c = bits.Sub64(tmp[i], mod[i], c)
    }

    // final sub was unnecessary
    if c != 0 {
        copy(z, tmp[:])
    }
    return nil
}




// TODO check unrolled speed
func AddModNonUnrolled320(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[5]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[5]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[5]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
    mod := f.Modulus

    if GTE(x, y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }

    var c uint64 = 0
    tmp := make([]uint64, len(mod))

    for i := 0; i < 5; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    c = 0
    for i := 0; i < 5; i++ {
        z[i], c = bits.Sub64(tmp[i], mod[i], c)
    }

    // final sub was unnecessary
    if c != 0 {
        copy(z, tmp[:])
    }
    return nil
}




// TODO check unrolled speed
func AddModNonUnrolled384(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
    mod := f.Modulus

    if GTE(x, y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }

    var c uint64 = 0
    tmp := make([]uint64, len(mod))

    for i := 0; i < 6; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    c = 0
    for i := 0; i < 6; i++ {
        z[i], c = bits.Sub64(tmp[i], mod[i], c)
    }

    // final sub was unnecessary
    if c != 0 {
        copy(z, tmp[:])
    }
    return nil
}




// TODO check unrolled speed
func AddModNonUnrolled448(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
    mod := f.Modulus

    if GTE(x, y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }

    var c uint64 = 0
    tmp := make([]uint64, len(mod))

    for i := 0; i < 7; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    c = 0
    for i := 0; i < 7; i++ {
        z[i], c = bits.Sub64(tmp[i], mod[i], c)
    }

    // final sub was unnecessary
    if c != 0 {
        copy(z, tmp[:])
    }
    return nil
}




// TODO check unrolled speed
func AddModNonUnrolled512(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[8]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[8]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[8]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
    mod := f.Modulus

    if GTE(x, y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }

    var c uint64 = 0
    tmp := make([]uint64, len(mod))

    for i := 0; i < 8; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    c = 0
    for i := 0; i < 8; i++ {
        z[i], c = bits.Sub64(tmp[i], mod[i], c)
    }

    // final sub was unnecessary
    if c != 0 {
        copy(z, tmp[:])
    }
    return nil
}




// TODO check unrolled speed
func AddModNonUnrolled576(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
    mod := f.Modulus

    if GTE(x, y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }

    var c uint64 = 0
    tmp := make([]uint64, len(mod))

    for i := 0; i < 9; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    c = 0
    for i := 0; i < 9; i++ {
        z[i], c = bits.Sub64(tmp[i], mod[i], c)
    }

    // final sub was unnecessary
    if c != 0 {
        copy(z, tmp[:])
    }
    return nil
}




// TODO check unrolled speed
func AddModNonUnrolled640(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
    mod := f.Modulus

    if GTE(x, y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }

    var c uint64 = 0
    tmp := make([]uint64, len(mod))

    for i := 0; i < 10; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    c = 0
    for i := 0; i < 10; i++ {
        z[i], c = bits.Sub64(tmp[i], mod[i], c)
    }

    // final sub was unnecessary
    if c != 0 {
        copy(z, tmp[:])
    }
    return nil
}




// TODO check unrolled speed
func AddModNonUnrolled704(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[11]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[11]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[11]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
    mod := f.Modulus

    if GTE(x, y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }

    var c uint64 = 0
    tmp := make([]uint64, len(mod))

    for i := 0; i < 11; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    c = 0
    for i := 0; i < 11; i++ {
        z[i], c = bits.Sub64(tmp[i], mod[i], c)
    }

    // final sub was unnecessary
    if c != 0 {
        copy(z, tmp[:])
    }
    return nil
}




// TODO check unrolled speed
func AddModNonUnrolled768(f *Field, out_bytes, x_bytes, y_bytes []byte) (error) {
	x := (*[12]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[12]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
    mod := f.Modulus

    if GTE(x, y, mod) {
        return errors.New(fmt.Sprintf("input greater than or equal to modulus"))
    }

    var c uint64 = 0
    tmp := make([]uint64, len(mod))

    for i := 0; i < 12; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    c = 0
    for i := 0; i < 12; i++ {
        z[i], c = bits.Sub64(tmp[i], mod[i], c)
    }

    // final sub was unnecessary
    if c != 0 {
        copy(z, tmp[:])
    }
    return nil
}
