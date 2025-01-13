package engine

import (
	"fmt"
	"math"
	"math/big"

	"github.com/tuneinsight/lattigo/v6/circuits/ckks/minimax"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
)

func GetParams(d int) (float64, float64, float64, float64, float64) {
	var k1, k2, x1, x2, P float64
	if d == 7 {
		k1, k2, x1, x2, P = 0.128, 1.6645, 0.3111, 343.6645, 0.9053
	} else if d == 8 {
		k1, k2, x1, x2, P = 0.0855, 1.6876, 0.1322, 340.035, 0.3887
	} else if d == 9 {
		k1, k2, x1, x2, P = 0.0702, 1.6958, 0.0876, 338.781, 0.2587
	} else {
		fmt.Printf("The depth d=%d must be one of 7, 8, or 9.\n", d)
		panic("")
	}

	return k1, k2, x1, x2, P
}

// Assume that x0 is normalized
// // x0 = x/(b-a), target = P/(b-a), Calculate Step(target - x0)
// func Step(eval *ckks.Evaluator, params ckks.Parameters, x0 *rlwe.Ciphertext, target *rlwe.Ciphertext, d int) *rlwe.Ciphertext {
// 	result, _ := eval.SubNew(target, x0)

// 	for _ = range d {
// 		result = f3(eval, params, result)
// 	}
// 	eval.Add(result, 1.0, result)
// 	eval.Mul(result, 0.5, result)
// 	eval.Rescale(result, result)
// 	return result
// }

func TwoLineApprox(eval *ckks.Evaluator, params ckks.Parameters, ecd *ckks.Encoder, enc *rlwe.Encryptor, x0 *rlwe.Ciphertext, d int, a float64, b float64) *rlwe.Ciphertext {
	_, k2, x1, x2, P := GetParams(d)
	a, b = 0.001, 1000.0
	L1, err := eval.MulNew(x0, -0.5*k2*math.Pow(x1, -1.5))
	ReturnErr(err)
	L2, err := eval.MulNew(x0, -0.5*k2*math.Pow(x2, -1.5))
	ReturnErr(err)
	err = eval.Rescale(L1, L1)
	ReturnErr(err)
	err = eval.Rescale(L2, L2)
	ReturnErr(err)
	err = eval.Add(L1, 1.5*k2/math.Sqrt(x1), L1)
	ReturnErr(err)
	err = eval.Add(L2, 1.5*k2/math.Sqrt(x2), L2)
	ReturnErr(err)

	pt1 := ckks.NewPlaintext(params, params.MaxLevel())
	values1 := make([]float64, params.MaxSlots())
	for idx := range len(values1) {
		values1[idx] = 1.0
	}
	err = ecd.Encode(values1, pt1)
	ReturnErr(err)
	// target, err := enc.EncryptNew(pt1)
	// ReturnErr(err)
	ct_1, err := enc.EncryptNew(pt1)
	ReturnErr(err)

	polys := minimax.NewPolynomial(DefaultPolynomialForSign)
	minimaxEvl := minimax.NewEvaluator(params, eval, nil)
	CmpEval := NewEvaluator(params, minimaxEvl, polys)

	x, err := eval.MulNew(x0, 1.0/(b-a))
	ReturnErr(err)
	err = eval.Rescale(x, x)
	ReturnErr(err)

	x, err = eval.SubNew(x, P/(b-a))
	ReturnErr(err)
	beta, err := CmpEval.Step(x)
	ReturnErr(err)

	beta_1, err := eval.SubNew(ct_1, beta)
	ReturnErr(err)

	result1, err := eval.MulRelinNew(L1, beta_1)
	ReturnErr(err)
	err = eval.Rescale(result1, result1)
	ReturnErr(err)

	result2, err := eval.MulRelinNew(L2, beta)
	ReturnErr(err)
	err = eval.Rescale(result2, result2)
	ReturnErr(err)
	result, err := eval.AddNew(result1, result2)
	ReturnErr(err)

	return result
}

// It consumes 4 levels
func f3(eval *ckks.Evaluator, params ckks.Parameters, x0 *rlwe.Ciphertext) *rlwe.Ciphertext {
	k1, k2, k3, k4 := 35.0/16, -35.0/16, 21.0/16, -5.0/16
	x_1, x_3, x_5, x_7 := getXs(eval, x0)

	result, _ := eval.MulNew(x_1, k1)
	tmp2, _ := eval.MulNew(x_3, k2)
	tmp3, _ := eval.MulNew(x_5, k3)
	tmp4, _ := eval.MulNew(x_7, k4)
	eval.Rescale(result, result)
	eval.Rescale(tmp2, tmp2)
	eval.Rescale(tmp3, tmp3)
	eval.Rescale(tmp4, tmp4)

	eval.Add(result, tmp2, result)
	eval.Add(result, tmp3, result)
	eval.Add(result, tmp4, result)

	if params.RingType() == ring.Standard {
		result.Scale = result.Scale.Mul(rlwe.NewScale(2))
		result_conj, _ := eval.ConjugateNew(result)
		eval.Add(result, result_conj, result)
		result.Scale = x0.Scale
	}

	return result
}

// It consumes 4 levels
func g3(eval *ckks.Evaluator, params ckks.Parameters, x0 *rlwe.Ciphertext) *rlwe.Ciphertext {
	k1, k2, k3, k4 := 4589.0/1024, -16577.0/1024, 25614.0/1024, -12860.0/1024
	x_1, x_3, x_5, x_7 := getXs(eval, x0)

	result, _ := eval.MulNew(x_1, k1)
	tmp2, _ := eval.MulNew(x_3, k2)
	tmp3, _ := eval.MulNew(x_5, k3)
	tmp4, _ := eval.MulNew(x_7, k4)
	eval.Rescale(result, result)
	eval.Rescale(tmp2, tmp2)
	eval.Rescale(tmp3, tmp3)
	eval.Rescale(tmp4, tmp4)

	eval.Add(result, tmp2, result)
	eval.Add(result, tmp3, result)
	eval.Add(result, tmp4, result)

	if params.RingType() == ring.Standard {
		result.Scale = result.Scale.Mul(rlwe.NewScale(2))
		result_conj, _ := eval.ConjugateNew(result)
		eval.Add(result, result_conj, result)
		result.Scale = x0.Scale
	}
	return result
}

func getXs(eval *ckks.Evaluator, x0 *rlwe.Ciphertext) (*rlwe.Ciphertext, *rlwe.Ciphertext, *rlwe.Ciphertext, *rlwe.Ciphertext) {
	x_1 := x0.CopyNew()

	// Consume Level 1
	x_2, _ := eval.MulRelinNew(x_1, x_1)
	eval.Rescale(x_2, x_2)

	// Consume Level 2
	x_3, _ := eval.MulRelinNew(x_2, x_1)
	eval.Rescale(x_3, x_3)

	// Consume Level 2
	x_4, _ := eval.MulRelinNew(x_2, x_2)
	eval.Rescale(x_4, x_4)

	// Consume Level 3
	x_5, _ := eval.MulRelinNew(x_2, x_3)
	eval.Rescale(x_5, x_5)

	// Consume Level 3
	x_7, _ := eval.MulRelinNew(x_4, x_3)
	eval.Rescale(x_7, x_7)

	return x_1, x_3, x_5, x_7
}

// Step evaluates f(x) = 1 if x > 0, 0 if x < 0, else 0.5 (i.e. (sign+1)/2).
// This will ensure that step.Scale = params.DefaultScale().
func (eval Evaluator) Step(op0 *rlwe.Ciphertext) (step *rlwe.Ciphertext, err error) {

	n := len(eval.MinimaxCompositeSignPolynomial)

	stepPoly := make([]bignum.Polynomial, n)

	for i := 0; i < n; i++ {
		stepPoly[i] = eval.MinimaxCompositeSignPolynomial[i]
	}

	half := new(big.Float).SetFloat64(0.5)

	// (x+1)/2
	lastPoly := eval.MinimaxCompositeSignPolynomial[n-1].Clone()
	for i := range lastPoly.Coeffs {
		lastPoly.Coeffs[i][0].Mul(lastPoly.Coeffs[i][0], half)
	}
	lastPoly.Coeffs[0][0].Add(lastPoly.Coeffs[0][0], half)

	stepPoly[n-1] = lastPoly

	return MinimaxEvaluate(eval, op0, stepPoly)
}

type Evaluator struct {
	Parameters ckks.Parameters
	*minimax.Evaluator
	MinimaxCompositeSignPolynomial minimax.Polynomial
}

// Evaluate evaluates the provided MinimaxCompositePolynomial on the input ciphertext.
func MinimaxEvaluate(eval Evaluator, ct *rlwe.Ciphertext, mcp minimax.Polynomial) (res *rlwe.Ciphertext, err error) {

	params := eval.Parameters
	res = ct.CopyNew()

	for _, poly := range mcp {
		// Define the scale that res must have after the polynomial evaluation.
		// If we use the regular CKKS (with complex values), we chose a scale to be
		// half of the desired scale, so that (x + conj(x)/2) has the correct scale.
		var targetScale rlwe.Scale
		if params.RingType() == ring.Standard {
			targetScale = params.DefaultScale().Div(rlwe.NewScale(2))
		} else {
			targetScale = params.DefaultScale()
		}

		// Evaluate the polynomial
		if res, err = eval.PolyEval.Evaluate(res, poly, targetScale); err != nil {
			return nil, fmt.Errorf("evaluate polynomial: %w", err)
		}

		// Clean the imaginary part (else it tends to explode)
		if params.RingType() == ring.Standard {

			// Reassigns the scale back to the original one
			res.Scale = res.Scale.Mul(rlwe.NewScale(2))

			var resConj *rlwe.Ciphertext
			if resConj, err = eval.ConjugateNew(res); err != nil {
				return
			}

			if err = eval.Add(res, resConj, res); err != nil {
				return
			}
		}
	}

	res.Scale = ct.Scale

	return
}

func NewEvaluator(params ckks.Parameters, eval *minimax.Evaluator, signPoly ...minimax.Polynomial) *Evaluator {
	if len(signPoly) == 1 {
		return &Evaluator{
			Parameters:                     params,
			Evaluator:                      eval,
			MinimaxCompositeSignPolynomial: signPoly[0],
		}
	} else {
		return &Evaluator{
			Parameters:                     params,
			Evaluator:                      eval,
			MinimaxCompositeSignPolynomial: minimax.NewPolynomial(DefaultPolynomialForSign),
		}
	}
}

// 29 = 3 + 4 + 4 + 4 + 3 + 5 + 3 + 3
var DefaultPolynomialForSign = [][]string{{"0", "0.6390324059720205", "0", "-0.2198072442832513", "0", "0.1414406903374961", "0", "-0.5606601964716950"},
	{"0", "0.6371916849818620", "0", "-0.2138182920938642", "0", "0.1300528543273972", "0", "-0.0948905401988566", "0", "0.0760465373114826", "0", "-0.0647752466148176", "0", "0.0577934985612991", "0", "-0.5275291175103518"},
	{"0", "0.6377206727878986", "0", "-0.2139936398068216", "0", "0.1301568728941797", "0", "-0.0949635425470285", "0", "0.0761019418378052", "0", "-0.0648191338716430", "0", "0.0578291285745938", "0", "-0.5271294130801401"},
	{"0", "0.6443950923061463", "0", "-0.2162055771803981", "0", "0.1314684297456403", "0", "-0.0958833600353299", "0", "0.0767993066598502", "0", "-0.0653707358190023", "0", "0.0582760544840787", "0", "-0.5220846192664677"},
	{"0", "0.6811845934958654", "0", "-0.2334032891661299", "0", "0.1490228759436919", "0", "-0.5303235460737469"},
	{"0", "1.2724046571429532", "0", "-0.4219130535252118", "0", "0.2504961101202538", "0", "-0.1761103958158040", "0", "0.1340921891162222", "0", "-0.1068120717134459", "0", "0.0874944606348786", "0", "-0.0729814769275110", "0", "0.0616046472097950", "0", "-0.0524007824752258", "0", "0.0447759297248556", "0", "-0.0383446741769455", "0", "0.0328466191095667", "0", "-0.0281000581726133", "0", "0.0239749328696791", "0", "-0.0724805452224001"},
	minimax.CoeffsSignX4Cheby,
	minimax.CoeffsSignX4Cheby,
}

func ReturnErr(err error) {
	if err != nil {
		panic(err)
	}
}
