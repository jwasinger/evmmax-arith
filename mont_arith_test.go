package mont_arith

import (
	"fmt"
	"math/big"
    "math/rand"
	"testing"
)

func testMulModMont(t *testing.T, limbCount uint) {
	mod := GenTestModulus(limbCount)

	montCtx := NewField()

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}
	x := LimbsToInt(mod)
	x = x.Sub(x, big.NewInt(10))
	y := LimbsToInt(mod)
	y = y.Sub(y, big.NewInt(15))

	// convert to montgomery form
	x.Mul(x, montCtx.RVal())
	x.Mod(x, LimbsToInt(mod))

	y.Mul(y, montCtx.RVal())
	y.Mod(y, LimbsToInt(mod))

	expected := new(big.Int)
	expected.Mul(x, y)
	expected.Mul(expected, montCtx.RInv())
	expected.Mod(expected, LimbsToInt(mod))

	outLimbs := make(nat, montCtx.NumLimbs)
	xLimbs := IntToLimbs(x, montCtx.NumLimbs)
	yLimbs := IntToLimbs(y, montCtx.NumLimbs)

	montCtx.MulModMont(outLimbs, xLimbs, yLimbs)

	result := LimbsToInt(outLimbs)
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

func TestMulModMontBLS12831(t *testing.T) {
	montCtx := NewField()
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

    out := make(nat, limbCount)
    montCtx.MulModMont(out, x, y)
    // TODO assert that the result is correct
}

func benchmarkMulModMont(b *testing.B, limbCount uint) {
	mod := MaxModulus(limbCount)
	montCtx := NewField()

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

	x := big.NewInt(2)
	x.Lsh(x, (limbCount*64)-10)

	y := big.NewInt(2)
	y.Lsh(y, (limbCount*64)-10)

	// convert x/y to montgomery

	outLimbs := make(nat, montCtx.NumLimbs)
	xLimbs := IntToLimbs(x, limbCount)
	yLimbs := IntToLimbs(y, limbCount)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		montCtx.MulModMont(outLimbs, xLimbs, yLimbs)
	}
}

func TestMulModMont(t *testing.T) {
	test := func(t *testing.T, name string, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			// test x/y >= modulus
			t.Run(fmt.Sprintf("%s/%d-bit", name, i*64), func(t *testing.T) {
				testMulModMont(t, uint(i))
			})
		}
	}

	test(t, "gnark-mulnocarry-unrolled", 1, 12)
}

func BenchmarkMulModMont(b *testing.B) {
	bench := func(b *testing.B, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			b.Run(fmt.Sprintf("%d-bit", i*64), func(b *testing.B) {
				benchmarkMulModMont(b, uint(i))
			})
		}
	}

	bench(b, 1, 12)
}

func testAddMod(t *testing.T, limbCount uint) {
    // TODO test where final addition is too much
    // TODO test where final addition is correct
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

func testSubMod(t *testing.T, limbCount uint) {
	mod := GenTestModulus(limbCount)
	montCtx := NewField()
	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}
    one := big.NewInt(1)
    x := LimbsToInt(mod)
    x.Sub(x, one)
    xLimbs := IntToLimbs(x, montCtx.NumLimbs)
    oneLimbs := IntToLimbs(one, montCtx.NumLimbs)

    resultLimbs := make(nat, limbCount)
    expected := new(big.Int)
    expected.Sub(one, x).Mod(expected, montCtx.ModulusNonInterleaved)

    // test where final addition happens
    montCtx.SubMod(resultLimbs, oneLimbs, xLimbs)
    result := LimbsToInt(resultLimbs)

    if result.Cmp(expected) != 0 {
		t.Fatalf("result (%x) != expected (%x)\n", result, expected)
    }
    // test where final addition doesn't happen
    expected = new(big.Int)
    expected.Sub(x, one).Mod(expected, montCtx.ModulusNonInterleaved)
    montCtx.SubMod(resultLimbs, xLimbs, oneLimbs)
    result = LimbsToInt(resultLimbs)
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

func BenchmarkAddMod(b *testing.B) {

}

func BenchmarkSubMod(b *testing.B) {

}
