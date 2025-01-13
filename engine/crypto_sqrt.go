package engine

import (
	"math"

	"github.com/tuneinsight/lattigo/v6/circuits/ckks/polynomial"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
)

func F3(x float64) (y float64) {
	if x > -1.0 {
		return math.Sqrt(x+1) * math.Sqrt(500)
	} else {
		return 0
	}
}

func CryptoSqrt(eval *ckks.Evaluator, params ckks.Parameters, ct *rlwe.Ciphertext, d int, A float64, B float64) *rlwe.Ciphertext {
	x := ct.CopyNew()
	eval.Mul(x, 2.0/B, x)
	eval.Rescale(x, x)

	scaled_ct, _ := eval.SubNew(x, 1)
	gcbsp := GetChebyshevPoly(1.0, int(math.Pow(2, float64(d))-2), F3)
	poly := polynomial.NewPolynomial(gcbsp)

	polyEval := polynomial.NewEvaluator(params, eval)
	p2, _ := polyEval.Evaluate(scaled_ct, poly, params.DefaultScale().Div(rlwe.NewScale(2)))

	p2.Scale = p2.Scale.Mul(rlwe.NewScale(2))
	conj, _ := eval.ConjugateNew(p2)
	eval.Add(p2, conj, p2)

	p2.Scale = scaled_ct.Scale
	return p2
}
