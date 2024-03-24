package evmmax_arith

import (
	"bytes"
	"math"
	"math/big"
	"math/rand"
	"testing"
)

func randBigInt(r *rand.Rand, modulus *big.Int) *big.Int {
	modulusLen := len(modulus.Bytes())
	resBytes := make([]byte, modulusLen)
	for i := 0; i < modulusLen; i++ {
		resBytes[i] = byte(r.Int())
	}

	res := new(big.Int).SetBytes(resBytes)
	return res.Mod(res, modulus)
}

func TestMulMontBLS12831(t *testing.T) {
	modInt, _ := new(big.Int).SetString("1a0111ea397fe69a4b1ba7b6434bacd764774b84f38512bf6730d2a0f6b0f6241eabfffeb153ffffb9feffffffffaaab", 16)
	mod := modInt.Bytes()

	var limbCount uint = 6
	montCtx, _ := NewModulusState(mod, 256)
	elemSize := int(math.Ceil(float64(len(mod)) / 8.0))

	s := rand.NewSource(42)
	r := rand.New(s)

	/*
			x := PadBytes(randBigInt(r, modInt).Bytes(), uint64(elemSize))
			y := PadBytes(randBigInt(r, modInt).Bytes(), uint64(elemSize))
		xInt := big.NewInt(2)
		yInt := big.NewInt(3)
	*/
	xInt := randBigInt(r, modInt)
	yInt := randBigInt(r, modInt)

	x := PadBytes(xInt.Bytes(), uint64(elemSize)*8)
	y := PadBytes(yInt.Bytes(), uint64(elemSize)*8)
	expected := new(big.Int).Mul(xInt, yInt)
	expected.Mod(expected, montCtx.modulusInt)
	montCtx.Store(1, 1, x)
	montCtx.Store(2, 1, y)

	montCtx.MulMod(montCtx.modInv,
		montCtx.Modulus,
		montCtx.scratchSpace[0:elemSize],
		montCtx.scratchSpace[elemSize:2*elemSize],
		montCtx.scratchSpace[elemSize*2:elemSize*3])
	outBytes := make([]byte, limbCount*8)
	montCtx.Load(outBytes, 0, 1)
	if bytes.Compare(PadBytes(expected.Bytes(), uint64(elemSize)*8), outBytes) != 0 {
		t.Fatalf("result not matching")
	}
	// TODO assert that the result is correct
}
