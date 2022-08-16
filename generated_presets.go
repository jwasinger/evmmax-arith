package mont_arith

type ArithPreset struct {
	AddModImpls  []arithFunc
	SubModImpls  []arithFunc
	MulMontImpls []arithFunc
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
		AddModNonUnrolled1088,
		AddModNonUnrolled1152,
		AddModNonUnrolled1216,
		AddModNonUnrolled1280,
		AddModNonUnrolled1344,
		AddModNonUnrolled1408,
		AddModNonUnrolled1472,
		AddModNonUnrolled1536,
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
		SubModNonUnrolled1088,
		SubModNonUnrolled1152,
		SubModNonUnrolled1216,
		SubModNonUnrolled1280,
		SubModNonUnrolled1344,
		SubModNonUnrolled1408,
		SubModNonUnrolled1472,
		SubModNonUnrolled1536,
	}
	mulMontImpls := []arithFunc{
		mulMont64,
		mulMont128,
		mulMont192,
		mulMont256,
		mulMont320,
		MulMont384_asm,
		mulMont448,
		mulMont512,
		mulMont576,
		mulMont640,
		mulMont704,
		mulMont768,
		mulMont832,
		mulMont896,
		mulMont960,
		mulMont1024,
		mulMont1088,
		mulMont1152,
		mulMont1216,
		mulMont1280,
		mulMont1344,
		mulMont1408,
		mulMont1472,
		mulMont1536,
	}

	return ArithPreset{addModImpls, subModImpls, mulMontImpls}
}

func DefaultPreset() ArithPreset {
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
		AddModNonUnrolled1088,
		AddModNonUnrolled1152,
		AddModNonUnrolled1216,
		AddModNonUnrolled1280,
		AddModNonUnrolled1344,
		AddModNonUnrolled1408,
		AddModNonUnrolled1472,
		AddModNonUnrolled1536,
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
		SubModNonUnrolled1088,
		SubModNonUnrolled1152,
		SubModNonUnrolled1216,
		SubModNonUnrolled1280,
		SubModNonUnrolled1344,
		SubModNonUnrolled1408,
		SubModNonUnrolled1472,
		SubModNonUnrolled1536,
	}
	mulMontImpls := []arithFunc{
		mulMont64,
		mulMont128,
		mulMont192,
		mulMont256,
		mulMont320,
		mulMont384,
		mulMont448,
		mulMont512,
		mulMont576,
		mulMont640,
		mulMont704,
		mulMont768,
		mulMont832,
		mulMont896,
		mulMont960,
		mulMont1024,
		mulMont1088,
		mulMont1152,
		mulMont1216,
		mulMont1280,
		mulMont1344,
		mulMont1408,
		mulMont1472,
		mulMont1536,
	}

	return ArithPreset{addModImpls, subModImpls, mulMontImpls}
}
