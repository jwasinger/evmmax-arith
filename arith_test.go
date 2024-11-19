package evmmax_arith

import (
	cryptorand "crypto/rand"
	"fmt"
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
	res.Mod(res, modulus)
	return res
}

const opRepeat = 10

func testOp(t *testing.T, op string, mod *big.Int) {
	fieldCtx, err := NewFieldContext(mod.Bytes(), 256)
	if err != nil {
		fmt.Printf("failed to instantiate modulus context: %+v\n", err)
	}
	elemSize := int(math.Ceil(float64(len(mod.Bytes())) / 8.0))

	s := rand.NewSource(42)
	r := rand.New(s)

	for i := 0; i < opRepeat; i++ {
		xInt := randBigInt(r, mod)
		yInt := randBigInt(r, mod)
		x := PadBytes(xInt.Bytes(), uint64(elemSize)*8)
		y := PadBytes(yInt.Bytes(), uint64(elemSize)*8)
		var expected *big.Int

		if err := fieldCtx.Store(1, 1, x); err != nil {
			t.Fatalf("error storing value: %v", err)
		}
		if err := fieldCtx.Store(2, 1, y); err != nil {
			t.Fatalf("error storing value: %v", err)
		}

		switch op {
		case "mul":
			fieldCtx.MulMod(0, 1, 1, 1, 2, 1, 1)
			expected = new(big.Int).Mul(xInt, yInt)
			expected.Mod(expected, fieldCtx.modulusInt)
		case "add":
			fieldCtx.AddMod(0, 1, 1, 1, 2, 1, 1)
			expected = new(big.Int).Add(xInt, yInt)
			expected.Mod(expected, mod)
		case "sub":
			fieldCtx.SubMod(0, 1, 1, 1, 2, 1, 1)
			expected = new(big.Int).Sub(xInt, yInt)
			expected.Mod(expected, mod)
		default:
			panic("unknown op")
		}

		resBytes := make([]byte, elemSize*8)
		fieldCtx.Load(resBytes, 0, 1)
		res := new(big.Int).SetBytes(resBytes)
		if res.Cmp(expected) != 0 {
			t.Fatalf("mismatch. received %s != expected %s\n", res.String(), expected.String())
		}
	}
}

func randOddModulus(size int) []byte {
	res := make([]byte, size)

	for {
		_, err := cryptorand.Read(res[:])
		if err != nil {
			panic(err)
		}
		if res[len(res)-1]%2 != 0 {
			return res
		}
	}
}

func randBinaryModulus(size int) []byte {
	modulus := big.NewInt(1)
	modulus.Lsh(modulus, uint(size*8))
	return modulus.Bytes()
}

func TestOps(t *testing.T) {
	for i := 1; i < 96; i++ {
		mod := new(big.Int).SetBytes(randOddModulus(i))
		t.Run(fmt.Sprintf("mulmod-odd-%dbyte", i), func(t *testing.T) {
			testOp(t, "mul", mod)
		})
		t.Run(fmt.Sprintf("addmod-odd-%dbyte", i), func(t *testing.T) {
			testOp(t, "add", mod)
		})
		t.Run(fmt.Sprintf("submod-odd-%dbyte", i), func(t *testing.T) {
			testOp(t, "sub", mod)
		})

		mod = new(big.Int).SetBytes(randBinaryModulus(i))
		t.Run(fmt.Sprintf("mulmod-binary-%dbyte", i), func(t *testing.T) {
			testOp(t, "mul", mod)
		})
		t.Run(fmt.Sprintf("addmod-binary-%dbyte", i), func(t *testing.T) {
			testOp(t, "add", mod)
		})
		t.Run(fmt.Sprintf("submod-binary-%dbyte", i), func(t *testing.T) {
			testOp(t, "sub", mod)
		})
	}
}
