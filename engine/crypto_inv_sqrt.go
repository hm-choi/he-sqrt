package engine

import (

	// "github.com/hm-choi/he-sqrt/engine"
	"math"
	"math/big"

	"github.com/tuneinsight/lattigo/v6/circuits/ckks/polynomial"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
)

func CryptoInvSqrt(eval *ckks.Evaluator, params ckks.Parameters, ct *rlwe.Ciphertext, d int, A float64, B float64, types int) *rlwe.Ciphertext {
	x := ct.CopyNew()
	F := func(x float64) (y float64) {
		if x > -1.0 {
			return 1 / (math.Sqrt(x + 1.0))
		} else {
			return 0
		}
	}
	if types == 1 {
		eval.Mul(x, 2.0/B, x)
		eval.Rescale(x, x)
		F = func(x float64) (y float64) {
			if x > -1.0 {
				return 1 / math.Sqrt(B/2) / (math.Sqrt(x + 1.0))
			} else {
				return 0
			}
		}

	}

	scaled_ct, _ := eval.SubNew(x, 1)

	gcbsp := GetChebyshevPoly(1.0, int(math.Pow(2, float64(d))-2), F)

	poly := polynomial.NewPolynomial(gcbsp)

	polyEval := polynomial.NewEvaluator(params, eval)
	p2, _ := polyEval.Evaluate(scaled_ct, poly, params.DefaultScale().Div(rlwe.NewScale(2)))

	p2.Scale = p2.Scale.Mul(rlwe.NewScale(2))
	conj, _ := eval.ConjugateNew(p2)
	eval.Add(p2, conj, p2)

	p2.Scale = scaled_ct.Scale
	return p2
}

// GetChebyshevPoly returns the Chebyshev polynomial approximation of f the
// in the interval [-K, K] for the given degree.
func GetChebyshevPoly(K float64, degree int, f64 func(x float64) (y float64)) bignum.Polynomial {

	FBig := func(x *big.Float) (y *big.Float) {
		xF64, _ := x.Float64()
		return new(big.Float).SetPrec(x.Prec()).SetFloat64(f64(xF64))
	}

	var prec uint = 128

	interval := bignum.Interval{
		A:     *bignum.NewFloat(-K, prec),
		B:     *bignum.NewFloat(K, prec),
		Nodes: degree,
	}
	// Returns the polynomial.
	return bignum.ChebyshevApproximation(FBig, interval)
}
