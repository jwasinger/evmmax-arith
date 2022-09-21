package mont_arith

import (
	"fmt"
	"math/big"
	"testing"
)

var MaxLimbsEVMMAX uint = 64

func benchmarkMulMont(b *testing.B, limbCount uint, preset ArithPreset) {
	mod := GenTestModulus(limbCount)
	montCtx := NewField(preset)

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

	x := LimbsToInt(mod)
	x = x.Sub(x, big.NewInt(100))

	y := new(big.Int).SetBytes(LimbsToInt(mod).Bytes())
	y = y.Sub(y, big.NewInt(100))

	outLimbs := make([]uint64, montCtx.NumLimbs)
	xLimbs := IntToLimbs(x, limbCount)
	yLimbs := IntToLimbs(y, limbCount)

	outBytes := LimbsToLEBytes(outLimbs)
	xBytes := LimbsToLEBytes(xLimbs)
	yBytes := LimbsToLEBytes(yLimbs)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		montCtx.MulMont(montCtx, outBytes, xBytes, yBytes)
	}
}

func benchmarkAddMod(b *testing.B, limbCount uint, preset ArithPreset) {
	modLimbs := MaxModulus(limbCount)
	mod := LimbsToInt(modLimbs)
	montCtx := NewField(preset)

	// worst-case performance: unecessary final subtraction
	err := montCtx.SetMod(modLimbs)
	if err != nil {
		panic("error")
	}
	x := new(big.Int).SetBytes(mod.Bytes())
	x = x.Sub(x, big.NewInt(2))
	y := big.NewInt(1)
	outBytes := make([]byte, limbCount*8)
	xBytes := LimbsToLEBytes(IntToLimbs(x, limbCount))
	yBytes := LimbsToLEBytes(IntToLimbs(y, limbCount))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		montCtx.AddMod(montCtx, outBytes, xBytes, yBytes)
	}
}

func benchmarkSubMod(b *testing.B, limbCount uint, preset ArithPreset) {
	modLimbs := MaxModulus(limbCount)
	montCtx := NewField(preset)

	// worst-case performance: unecessary final subtraction
	err := montCtx.SetMod(modLimbs)
	if err != nil {
		panic("error")
	}
	x := big.NewInt(1)
	y := big.NewInt(0)
	outBytes := make([]byte, limbCount*8)
	xBytes := LimbsToLEBytes(IntToLimbs(x, limbCount))
	yBytes := LimbsToLEBytes(IntToLimbs(y, limbCount))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		montCtx.SubMod(montCtx, outBytes, xBytes, yBytes)
	}
}

func benchmarkSetMod(b *testing.B, limbCount uint, preset ArithPreset) {
	modLimbs := SmolModulus(limbCount)
	montCtx := NewField(preset)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		montCtx.SetMod(modLimbs)
	}
}

func benchmarkOp(b *testing.B, opName string, opFn func(*testing.B, uint, ArithPreset)) {
	presets := AllPresets()

	for presetIdx := 0; presetIdx < len(presets); presetIdx++ {
		preset := presets[presetIdx]
		for i := uint(1); i <= MaxLimbsEVMMAX; i++ {
			b.Run(fmt.Sprintf("%s/%s/%d-bit", opName, preset.name, i*64), func(b *testing.B) {
				opFn(b, i, preset)
			})
		}
	}
}

func BenchmarkAddMod(b *testing.B) {
	benchmarkOp(b, "addmod", benchmarkAddMod)
}

func BenchmarkSubMod(b *testing.B) {
	benchmarkOp(b, "submod", benchmarkSubMod)
}

func BenchmarkMulMont(b *testing.B) {
	benchmarkOp(b, "mulmont", benchmarkMulMont)
}

func BenchmarkSetMod(b *testing.B) {
	benchmarkOp(b, "setmod", benchmarkSetMod)
}
