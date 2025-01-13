package engine

import (
	"fmt"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
)

/*
This method supports Newton–Raphson Iteration Algorithm for Inverse Square Root.
It supports a floating point approximation. (Without encryption)
- Input
  - x0: Initial number
  - d: Iteration number
  - y0: Initial point. If y is nil then y is set to 1.0

- Output
  - y: Result of approximation
*/
func NewtonInvSqrt(x0 float64, d int, y0 float64) float64 {
	y := y0
	for _ = range d {
		y = 0.5 * y * (3 - x0*y*y)
	}
	return y
}

/*
Homomorphic Encryption version of Newton–Raphson Iteration Algorithm for Inverse Square Root.
It does not supports bootstrapping, the level of x0, y0 must larger than 2d and 1 respectively.
*/
func HENewtonInvSqrt(eval *ckks.Evaluator, x0 *rlwe.Ciphertext, d int, y *rlwe.Ciphertext) *rlwe.Ciphertext {
	var requiredLevel = 2*d + 1
	if x0.Level() < requiredLevel {
		fmt.Printf("Level of x0 %d must be larger or equal than %d.\n", x0.Level(), requiredLevel)
	}

	// If d > 1, then it reduces the total level consumption as d-1.
	// x_half = 0.5 * x0
	// It requires only 1 level consumption.
	x_half, _ := eval.MulNew(x0, 0.5)
	eval.Rescale(x_half, x_half)

	// In this iteration, only 2d levels are consumed.
	for _ = range d {
		// 4. Calculate 1.5 * y
		y_, _ := eval.MulNew(y, 1.5)
		eval.Rescale(y_, y_)

		// 1. Calculate 0.5 * x0 * y = x_half * y
		xy, _ := eval.MulRelinNew(x_half, y)
		eval.Rescale(xy, xy)

		// 2. Calculate y * y
		yy, _ := eval.MulRelinNew(y, y)
		eval.Rescale(yy, yy)

		// 3. Calculate 0.5 * x0 * y * y * y
		xyyy, _ := eval.MulRelinNew(xy, yy)
		eval.Rescale(xyyy, xyyy)

		// 5. Calculate 0.5 * y * (3 - x0 * y * y) = 1.5 * y - 0.5 * x0 * y * y * y
		eval.Sub(y_, xyyy, y)
	}
	return y
}

func OptimizedHENewtonInvSqrt(eval *ckks.Evaluator, x0 *rlwe.Ciphertext, d int, y0 *rlwe.Ciphertext) *rlwe.Ciphertext {
	var requiredLevel = 2*d + 2
	if x0.Level() < requiredLevel {
		fmt.Printf("Level of x0 %d must be larger or equal than %d.\n", x0.Level(), requiredLevel)
	}

	if y0 == nil {

	}
	y := y0.CopyNew()

	x_half := x0.CopyNew()
	// x_half, _ := eval.MulNew(x0, 0)
	// eval.Rescale(x_half, x_half)
	eval.Mul(y, 0.25, y)
	eval.Rescale(y, y)

	// In this iteration, only 2d levels are consumed.
	for _ = range d {
		xy, _ := eval.MulRelinNew(x_half, y)
		eval.Rescale(xy, xy)
		conj, _ := eval.ConjugateNew(xy)
		eval.Add(xy, conj, xy)

		// 2. Calculate y * y
		yy, _ := eval.MulRelinNew(y, y)
		eval.Rescale(yy, yy)
		conj, _ = eval.ConjugateNew(yy)
		eval.Add(yy, conj, yy)

		// 3. Calculate 0.5 * x0 * y * y * y
		xyyy, _ := eval.MulRelinNew(xy, yy)
		eval.Rescale(xyyy, xyyy)
		conj, _ = eval.ConjugateNew(xyyy)
		eval.Add(xyyy, conj, xyyy)

		// 4. Calculate 1.5 * y
		y_, _ := eval.MulNew(y, 3.0/2)
		eval.Rescale(y_, y_)

		// 5. Calculate 0.5 * y * (3 - x0 * y * y) = 1.5 * y - 0.5 * x0 * y * y * y
		eval.Sub(y_, xyyy, y)
	}
	eval.Add(y, y, y)
	eval.Add(y, y, y)
	return y
}
