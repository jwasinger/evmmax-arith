package mont_arith

import (
	"github.com/jwasinger/evmmax-arith/arith384_asm"
	"unsafe"
)

func MulMont384_asm(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))
	mod := (*[6]uint64)(unsafe.Pointer(&f.Modulus[0]))

    // TODO bounds checks
	arith384_asm.MulMod384(z, x, y, mod, f.MontParamInterleaved)

	return nil
}

func AddMod384_asm(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))
	mod := (*[6]uint64)(unsafe.Pointer(&f.Modulus[0]))

    // TODO bounds checks
	arith384_asm.AddMod384(z, x, y, mod)

	return nil
}

func SubMod384_asm(f *Field, out_bytes, x_bytes, y_bytes []byte) error {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))
	mod := (*[6]uint64)(unsafe.Pointer(&f.Modulus[0]))

    // TODO bounds checks
	arith384_asm.SubMod384(z, x, y, mod)

	return nil
}
