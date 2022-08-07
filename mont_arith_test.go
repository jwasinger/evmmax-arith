package mont_arith

import (
	"fmt"
	"math/big"
    "math/rand"
	"testing"
)

func testMulMont(t *testing.T, limbCount uint) {
	mod := GenTestModulus(limbCount)

	montCtx := NewField(DefaultPreset())

	err := montCtx.SetMod(mod)
	if err != nil {
		panic(err)
	}
	x := LimbsToInt(mod)
	x = x.Sub(x, big.NewInt(1))
	y := LimbsToInt(mod)
	y = y.Sub(y, big.NewInt(1))

	expected := new(big.Int)
	expected.Mul(x, y)
	expected.Mul(expected, montCtx.RInv())
	expected.Mod(expected, LimbsToInt(mod))

    outBytes := make([]byte, montCtx.NumLimbs * 8)
	xLimbs := IntToLimbs(x, montCtx.NumLimbs)
	yLimbs := IntToLimbs(y, montCtx.NumLimbs)

	if err := montCtx.MulMont(montCtx, outBytes, LimbsToLEBytes(xLimbs), LimbsToLEBytes(yLimbs)); err != nil {
        t.Fatal(err)
    }

	result := LEBytesToInt(outBytes)
	if result.Cmp(expected) != 0 {
		t.Fatalf("result (%x) != expected (%x)\n", result, expected)
	}
}

func randBigInt(r *rand.Rand, modulus *big.Int, limbCount uint) *big.Int {
    resBytes := make([]byte, limbCount * 8)
    for i := 0; i < int(limbCount) * 8; i++ {
        resBytes[i] = byte(r.Int())
    }

    res := new(big.Int).SetBytes(resBytes)
    return res.Mod(res, modulus)
}

func TestMulMontBLS12831(t *testing.T) {
	montCtx := NewField(DefaultPreset())
    modInt, _ := new(big.Int).SetString("1a0111ea397fe69a4b1ba7b6434bacd764774b84f38512bf6730d2a0f6b0f6241eabfffeb153ffffb9feffffffffaaab", 16)

    var limbCount uint = 6
    mod := IntToLimbs(modInt, limbCount)
    montCtx.SetMod(mod)

    s := rand.NewSource(42)
    r := rand.New(s)

    x := IntToLimbs(randBigInt(r, LimbsToInt(montCtx.Modulus), limbCount), limbCount)
    montX := montCtx.ToMont(x)
    if !Eq(montCtx.ToNorm(montX), x) {
        panic("mont form should have correct normal form")
    }

    y := IntToLimbs(randBigInt(r, LimbsToInt(montCtx.Modulus), limbCount), limbCount)
    montY := montCtx.ToMont(y)
    if !Eq(montCtx.ToNorm(montY), y) {
        panic("mont form should have correct normal form")
    }

    out := make([]uint64, limbCount)
    montCtx.MulMont(montCtx, LimbsToLEBytes(out), LimbsToLEBytes(x), LimbsToLEBytes(y))
    // TODO assert that the result is correct
}

func TestMulMont(t *testing.T) {
	test := func(t *testing.T, name string, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			// test x/y >= modulus
			t.Run(fmt.Sprintf("%s/%d-bit", name, i*64), func(t *testing.T) {
				testMulMont(t, uint(i))
			})
		}
	}

    // TODO 64bit mulmont broken rn
	test(t, "gnark-mulnocarry-unrolled", 2, 12)
}

func testSubMod(t *testing.T, limbCount uint) {
	mod := GenTestModulus(limbCount)
	montCtx := NewField(DefaultPreset())
	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}
    one := big.NewInt(1)
    x := LimbsToInt(mod)
    x.Sub(x, one)
    xLimbs := IntToLimbs(x, montCtx.NumLimbs)
    oneLimbs := IntToLimbs(one, montCtx.NumLimbs)

    resultBytes := make([]byte, limbCount * 8)
    expected := new(big.Int)
    expected.Sub(one, x).Mod(expected, montCtx.ModulusNonInterleaved)

    // test where final addition happens
    if err := montCtx.SubMod(montCtx, resultBytes, LimbsToLEBytes(oneLimbs), LimbsToLEBytes(xLimbs)); err != nil {
        t.Fatal(err)
    }

    result := LEBytesToInt(resultBytes)

    if result.Cmp(expected) != 0 {
		t.Fatalf("result (%x) != expected (%x)\n", result, expected)
    }
    // test where final addition doesn't happen
    expected = new(big.Int)
    expected.Sub(x, one).Mod(expected, montCtx.ModulusNonInterleaved)
    if err = montCtx.SubMod(montCtx, resultBytes, LimbsToLEBytes(xLimbs), LimbsToLEBytes(oneLimbs)); err != nil {
        t.Fatal(err)
    }
    result = LEBytesToInt(resultBytes)
    if result.Cmp(expected) != 0 {
		t.Fatalf("result (%x) != expected (%x)\n", result, expected)
    }
}

func TestSubMod(t *testing.T) {
    test := func(t *testing.T, name string, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
            t.Run(fmt.Sprintf("%s/%d-bit", name, i*64), func(t *testing.T) {
                testSubMod(t, uint(i))
            })
        }
    }
    test(t, "submod", 1, 12)
}

/*
func genWorstCaseAddModInputs(limbCount uint) ([]uint64, []uint64) {

}
*/

func testAddMod(t *testing.T, limbCount uint) {
	mod := GenTestModulus(limbCount)
	montCtx := NewField(DefaultPreset())
	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}
    one := big.NewInt(1)
    two := big.NewInt(2)
    x := LimbsToInt(mod)
    x.Sub(x, two)
    xLimbs := IntToLimbs(x, montCtx.NumLimbs)
    oneLimbs := IntToLimbs(one, montCtx.NumLimbs)
    twoLimbs := IntToLimbs(two, montCtx.NumLimbs)

    resultBytes := make([]byte, limbCount * 8)
    expected := new(big.Int)
    expected.Add(one, x).Mod(expected, montCtx.ModulusNonInterleaved)

    // TODO test where final subtraction doesn't 
    if err := montCtx.AddMod(montCtx, resultBytes, LimbsToLEBytes(oneLimbs), LimbsToLEBytes(xLimbs)); err != nil {
        t.Fatal(err)
    }

    result := LEBytesToInt(resultBytes)

    if result.Cmp(expected) != 0 {
		t.Fatalf("result (%x) != expected (%x)\n", result, expected)
    }
    // TODO test where final subtraction does happen
    expected = big.NewInt(0)
    if err = montCtx.AddMod(montCtx, resultBytes, LimbsToLEBytes(xLimbs), LimbsToLEBytes(twoLimbs)); err != nil {
        t.Fatal(err)
    }
    result = LEBytesToInt(resultBytes)
    if result.Cmp(expected) != 0 {
		t.Fatalf("result (%x) != expected (%x)\n", result, expected)
    }
}

func TestAddMod(t *testing.T) {
    test := func(t *testing.T, name string, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
            t.Run(fmt.Sprintf("%s/%d-bit", name, i*64), func(t *testing.T) {
                testAddMod(t, uint(i))
            })
        }
    }
    test(t, "addmod", 1, 12)
}
