package evmmax_arith

type benchRange struct {
	min uint
	max uint
}

type ArithPreset struct {
	AddModImpls       []arithFunc
	SubModImpls       []arithFunc
	MulMontImpls      []arithFunc
	name              string
	mulMontCIOSCutoff uint

	benchRanges map[string]benchRange
}

func makeBenchRanges(addModMin, addModMax, subModMin, subModMax, mulMontMin, mulMontMax, setModMin, setModMax uint) map[string]benchRange {
	return map[string]benchRange{
		"addmod":  {addModMin, addModMax},
		"submod":  {subModMin, subModMax},
		"mulmont": {mulMontMin, mulMontMax},
		"setmod":  {setModMin, setModMax},
	}
}

func (a *ArithPreset) MaxLimbCount() uint {
	return uint(len(a.MulMontImpls))
}

// Preset same as default except it uses blst's go-asm impl of the arithmetic at 384bit widths
func Asm384Preset() ArithPreset {
	addModImpls := []arithFunc{
		AddModNonUnrolled64,
		AddModNonUnrolled128,
		AddModNonUnrolled192,
		AddModNonUnrolled256,
		AddModNonUnrolled320,
		AddMod384_asm,
		AddModNonUnrolled448,
		AddModNonUnrolled512,
		AddModNonUnrolled576,
		AddModNonUnrolled640,
		AddModNonUnrolled704,
		AddModNonUnrolled768,
		AddModNonUnrolled832,
		AddModNonUnrolled896,
		AddModNonUnrolled960,
		AddModNonUnrolled1024,
	}

	subModImpls := []arithFunc{
		SubModNonUnrolled64,
		SubModNonUnrolled128,
		SubModNonUnrolled192,
		SubModNonUnrolled256,
		SubModNonUnrolled320,
		SubMod384_asm,
		SubModNonUnrolled448,
		SubModNonUnrolled512,
		SubModNonUnrolled576,
		SubModNonUnrolled640,
		SubModNonUnrolled704,
		SubModNonUnrolled768,
		SubModNonUnrolled832,
		SubModNonUnrolled896,
		SubModNonUnrolled960,
		SubModNonUnrolled1024,
	}
	mulMontImpls := []arithFunc{
		MulMontNonUnrolled64,
		MulMontNonUnrolled128,
		MulMontNonUnrolled192,
		MulMontNonUnrolled256,
		MulMontNonUnrolled320,
		MulMont384_asm,
		MulMontNonUnrolled448,
		MulMontNonUnrolled512,
		MulMontNonUnrolled576,
		MulMontNonUnrolled640,
		MulMontNonUnrolled704,
		MulMontNonUnrolled768,
		MulMontNonUnrolled832,
		MulMontNonUnrolled896,
		MulMontNonUnrolled960,
		MulMontNonUnrolled1024,
	}

	return ArithPreset{addModImpls, subModImpls, mulMontImpls, "asm384", 16,
		makeBenchRanges(6, 6,
			6, 6,
			6, 6,
			0, 0),
	}
}

func NonUnrolledPreset() ArithPreset {
	addModImpls := []arithFunc{
		AddModNonUnrolled64,
		AddModNonUnrolled128,
		AddModNonUnrolled192,
		AddModNonUnrolled256,
		AddModNonUnrolled320,
		AddModNonUnrolled384,
		AddModNonUnrolled448,
		AddModNonUnrolled512,
		AddModNonUnrolled576,
		AddModNonUnrolled640,
		AddModNonUnrolled704,
		AddModNonUnrolled768,
		AddModNonUnrolled832,
		AddModNonUnrolled896,
		AddModNonUnrolled960,
		AddModNonUnrolled1024,
	}

	subModImpls := []arithFunc{
		SubModNonUnrolled64,
		SubModNonUnrolled128,
		SubModNonUnrolled192,
		SubModNonUnrolled256,
		SubModNonUnrolled320,
		SubModNonUnrolled384,
		SubModNonUnrolled448,
		SubModNonUnrolled512,
		SubModNonUnrolled576,
		SubModNonUnrolled640,
		SubModNonUnrolled704,
		SubModNonUnrolled768,
		SubModNonUnrolled832,
		SubModNonUnrolled896,
		SubModNonUnrolled960,
		SubModNonUnrolled1024,
	}
	mulMontImpls := []arithFunc{
		MulMontNonUnrolled64,
		MulMontNonUnrolled128,
		MulMontNonUnrolled192,
		MulMontNonUnrolled256,
		MulMontNonUnrolled320,
		MulMontNonUnrolled384,
		MulMontNonUnrolled448,
		MulMontNonUnrolled512,
		MulMontNonUnrolled576,
		MulMontNonUnrolled640,
		MulMontNonUnrolled704,
		MulMontNonUnrolled768,
		MulMontNonUnrolled832,
		MulMontNonUnrolled896,
		MulMontNonUnrolled960,
		MulMontNonUnrolled1024,
	}

	return ArithPreset{addModImpls, subModImpls, mulMontImpls, "non-unrolled", 16,
		makeBenchRanges(
			1, 16,
			1, 16,
			1, 16,
			1, 16),
	}
}

func GenericMulMontPreset() ArithPreset {
	addModImpls := []arithFunc{
		AddModNonUnrolled64,
		AddModNonUnrolled128,
		AddModNonUnrolled192,
		AddModNonUnrolled256,
		AddModNonUnrolled320,
		AddModNonUnrolled384,
		AddModNonUnrolled448,
		AddModNonUnrolled512,
		AddModNonUnrolled576,
		AddModNonUnrolled640,
		AddModNonUnrolled704,
		AddModNonUnrolled768,
		AddModNonUnrolled832,
		AddModNonUnrolled896,
		AddModNonUnrolled960,
		AddModNonUnrolled1024,
	}

	subModImpls := []arithFunc{
		SubModNonUnrolled64,
		SubModNonUnrolled128,
		SubModNonUnrolled192,
		SubModNonUnrolled256,
		SubModNonUnrolled320,
		SubModNonUnrolled384,
		SubModNonUnrolled448,
		SubModNonUnrolled512,
		SubModNonUnrolled576,
		SubModNonUnrolled640,
		SubModNonUnrolled704,
		SubModNonUnrolled768,
		SubModNonUnrolled832,
		SubModNonUnrolled896,
		SubModNonUnrolled960,
		SubModNonUnrolled1024,
	}
	mulMontImpls := []arithFunc{
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
		MulMontNonInterleaved,
	}

	return ArithPreset{addModImpls, subModImpls, mulMontImpls, "generic", 0,
		makeBenchRanges(
			1, 1000,
			1, 1000,
			64, 1000,
			64, 1000),
	}
}

func DefaultPreset() ArithPreset {
	return NonUnrolledPreset()
}

func AllPresets() []ArithPreset {
	return []ArithPreset{
		NonUnrolledPreset(),
		Asm384Preset(),
	}
}
