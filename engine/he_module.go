package engine

import (
	"fmt"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
)

func GetParam(LogN int, LEVEL int, SCALE int) ckks.Parameters {
	LogQ := make([]int, LEVEL+1)
	LogQ[0] = 60
	for i := range LEVEL {
		LogQ[i+1] = SCALE
	}

	params, _ := ckks.NewParametersFromLiteral(
		ckks.ParametersLiteral{
			LogN:            LogN,                  // log2(ring degree)
			LogQ:            LogQ,                  // log2(primes Q) (ciphertext modulus)
			LogP:            []int{61, 61, 61, 61}, // log2(primes P) (auxiliary modulus)
			LogDefaultScale: SCALE,                 // log2(scale)
			// RingType:        ring.ConjugateInvariant,
		})
	return params
}

func GetMudules(params ckks.Parameters) (*ckks.Encoder, *rlwe.Encryptor, *rlwe.Decryptor, *ckks.Evaluator) {
	// Key Generator
	kgen := rlwe.NewKeyGenerator(params)
	sk, pk := kgen.GenKeyPairNew()
	ecd := ckks.NewEncoder(params)
	enc := rlwe.NewEncryptor(params, pk)
	dec := rlwe.NewDecryptor(params, sk)
	rlk := kgen.GenRelinearizationKeyNew(sk)
	evk := rlwe.NewMemEvaluationKeySet(rlk)
	eval := ckks.NewEvaluator(params, evk)

	galEls := []uint64{
		params.GaloisElementForComplexConjugation(),
	}
	eval = eval.WithKey(rlwe.NewMemEvaluationKeySet(rlk, kgen.GenGaloisKeysNew(galEls, sk)...))

	fmt.Println("sk.BinarySize()", sk.BinarySize())
	fmt.Println("pk.BinarySize()", pk.BinarySize())
	fmt.Println("rlk.BinarySize()", rlk.BinarySize())
	fmt.Println("gk", kgen.GenGaloisKeysNew(galEls, sk)[0].BinarySize())
	return ecd, enc, dec, eval
}

func EnC(params ckks.Parameters, ecd *ckks.Encoder, enc *rlwe.Encryptor, x []float64) *rlwe.Ciphertext {
	pt := ckks.NewPlaintext(params, params.MaxLevel())
	ecd.Encode(x, pt)
	ct, _ := enc.EncryptNew(pt)
	return ct
}

func DeC(params ckks.Parameters, ecd *ckks.Encoder, dec *rlwe.Decryptor, ct *rlwe.Ciphertext) []float64 {
	result := make([]float64, params.MaxSlots())
	ecd.Decode(dec.DecryptNew(ct), result)
	return result
}
