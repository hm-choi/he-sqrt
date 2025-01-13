package examples_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/hm-choi/he-sqrt/engine"
	"github.com/hm-choi/he-sqrt/utils"
)

func experiment4() {
	// Set the two parameters LogN, Scale factor.
	LogN, SCALE := 17, 40
	// Generate test vector
	S := 65536
	values := make([]float64, S)
	result := make([]float64, S)

	// If testCase == 0 => Test

	for testCase := range 3 {
		start, mid, end := 0.0001, 1.0, 10000.0
		if testCase == 1 {
			start, mid, end = 0.000001, 1.0, 10.0
		} else if testCase == 2 {
			start, mid, end = 0.01, 1.0, 1000000.0
		}

		values1 := utils.Linspace(start, mid, S/2)
		values2 := utils.Linspace(mid, end, S/2)
		for i := range S {
			if i < S/2 {
				values[i] = values1[i]
			} else {
				values[i] = values2[i-S/2]
			}
			result[i] = 1.0 / math.Sqrt(values[i])
		}
		fmt.Println("Input: ", values[:2], values[S-2:])
		fmt.Println("result:", result[:2], result[S-2:])

		////////////////////////////
		// TEST for CryptoInvRoot //
		////////////////////////////
		fmt.Println("+====================================+")
		if testCase == 0 {
			fmt.Println("Test on [1e-4, 1e+4]")
		} else if testCase == 1 {
			fmt.Println("Test on [1e-6, 1e+2]")
		} else if testCase == 2 {
			fmt.Println("Test on [1e-2, 1e+6]")
		}
		fmt.Println("+====================================+")
		LEVEL, d1, d2 := 21, 12, 4
		params := engine.GetParam(LogN, LEVEL, SCALE)
		fmt.Println("LogPQ: ", params.LogQP())

		ecd, enc, dec, eval := engine.GetMudules(params)
		x0 := engine.EnC(params, ecd, enc, values)
		START_TIME := time.Now()
		y0 := engine.CryptoInvRoot(eval, params, x0, d1, start, end, 1)
		invSqrts := engine.HENewtonInvSqrt(eval, x0, d2, y0)

		fmt.Println("Time (CryptoInvSqrt): ", time.Since(START_TIME))
		invSqrt := engine.DeC(params, ecd, dec, invSqrts)
		_, mean := utils.CheckMAE(invSqrt, values, result, S)
		fmt.Println("MAE: ", mean)
		_, mean = utils.CheckMRE(invSqrt, values, result, S)
		fmt.Println("MRE: ", mean)
	}
}

func TestExperiment4(t *testing.T) {
	experiment4()
}
