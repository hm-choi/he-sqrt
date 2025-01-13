package examples_test

import (
	"fmt"
	"testing"

	"github.com/hm-choi/he-sqrt/engine"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
)

func keySize(LEVEL int) {
	LogN, SCALE := 17, 40

	params := engine.GetParam(LogN, LEVEL, SCALE)
	kgen := rlwe.NewKeyGenerator(params)
	sk, pk := kgen.GenKeyPairNew()
	enc := rlwe.NewEncryptor(params, pk)
	rlk := kgen.GenRelinearizationKeyNew(sk)
	galEls := []uint64{
		params.GaloisElementForComplexConjugation(),
	}
	glk := kgen.GenGaloisKeysNew(galEls, sk)

	ct, _ := enc.EncryptNew(ckks.NewPlaintext(params, params.MaxLevel()))
	fmt.Println("SK   size (Bytes): ", sk.BinarySize())
	fmt.Println("Pk   size (Bytes): ", pk.BinarySize())
	fmt.Println("RLK  size (Bytes): ", rlk.BinarySize())
	fmt.Println("GLK  size (Bytes): ", glk[0].BinarySize())
	fmt.Println("CTXT size (Bytes): ", ct.BinarySize())
}
func experiment5() {
	fmt.Println("+====================================+")
	fmt.Println("TEST1 (CryptoInvSqrt): ")
	fmt.Println("+====================================+")
	keySize(19)

	fmt.Println("+====================================+")
	fmt.Println("TEST2 (Pivot-Tangent): ")
	fmt.Println("+====================================+")
	keySize(47)
}

func TestExperiment5(t *testing.T) {
	experiment5()
}
