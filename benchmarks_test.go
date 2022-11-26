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

	x := new(big.Int).SetBytes(mod)
	x = x.Sub(x, big.NewInt(100))

	y := new(big.Int).SetBytes(mod)
	y = y.Sub(y, big.NewInt(100))

    outBytes := make([]byte, limbCount * 8)
	xBytes := PadBytes(x.Bytes(), montCtx.ElementSize)
	yBytes := PadBytes(y.Bytes(), montCtx.ElementSize)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		montCtx.MulMont(montCtx, outBytes, xBytes, yBytes)
	}
}

func benchmarkAddMod(b *testing.B, limbCount uint, preset ArithPreset) {
	// worst-case performance: unecessary final subtraction
    // TODO verify this again
	mod := MaxModulus(limbCount)
	montCtx := NewField(preset)

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}
	x := new(big.Int).SetBytes(mod)
	x = x.Sub(x, big.NewInt(2))
	y := big.NewInt(1)
	outBytes := make([]byte, limbCount*8)
	xBytes := PadBytes(x.Bytes(), montCtx.ElementSize)
	yBytes := PadBytes(y.Bytes(), montCtx.ElementSize)

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

type opFn func(*testing.B, uint, ArithPreset)

func BenchmarkOps(b *testing.B) {
	ops := []string{"addmod", "submod", "mulmont", "setmod"}
	presets := AllPresets()

	for opsIdx := 0; opsIdx < len(ops); opsIdx++ {
		op := ops[opsIdx]
		for presetIdx := 0; presetIdx < len(presets); presetIdx++ {
			preset := presets[presetIdx]
			if preset.benchRanges[op].min == 0 {
				continue
			}

			for limbCount := uint(1); limbCount <= 100000; {
				cluster := true
				var dist uint
				// bench every 3 if it's under 100
				if limbCount < 64 {
					dist = 1
				} else if limbCount < 100 {
					dist = 5
					cluster = false
				} else if limbCount < 1000 {
					dist = 50
				} else if limbCount < 10000 {
					dist = 500
				} else if limbCount < 100000 {
					dist = 5000
				}
				// with cluster samples:
				// bench every 30 if it's under 1000
				// bench every 300 if it's under 10000
				// bench every 3000 if it's under 100000

				var fn opFn
				switch op {
				case "mulmont":
					fn = benchmarkMulMont
				case "addmod":
					fn = benchmarkAddMod
				case "submod":
					fn = benchmarkSubMod
				case "setmod":
					fn = benchmarkSetMod
				}
				_ = cluster // TODO?
				const samplesPerBench = 10
				for i := 0; i < samplesPerBench; i++ {
					b.Run(fmt.Sprintf("%s_%s_%d", preset.name, op, limbCount*64), func(b *testing.B) {
						fn(b, limbCount, preset)
					})
				}
				limbCount += dist
			}
		}
	}
}
