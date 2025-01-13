package examples_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/hm-choi/he-sqrt/engine"
	"github.com/hm-choi/he-sqrt/utils"
)

func experiment3() {
	// Set the two parameters LogN, Scale factor.
	LogN, SCALE := 17, 40
	// Generate test vector
	S := 65536
	values := make([]float64, S)
	result := make([]float64, S)
	start, mid, end := 0.001, 1.0, 1000.0

	values1 := utils.Linspace(start, mid, S/2)
	values2 := utils.Linspace(mid, end, S/2)
	for i := range S {
		if i < S/2 {
			values[i] = values1[i]
		} else {
			values[i] = values2[i-S/2]
		}
		result[i] = math.Sqrt(values[i])
	}
	fmt.Println("Input: ", values[:2], values[S-2:])
	fmt.Println("result:", result[:2], result[S-2:])

	degrees := []int{10, 11, 12, 13}
	for _, deg := range degrees {
		fmt.Println("+====================================+")
		fmt.Println("TEST1 (CryptoSqrt): ")
		fmt.Println("+====================================+")
		LEVEL := deg + 1
		params := engine.GetParam(LogN, LEVEL, SCALE)
		fmt.Println("LogPQ: ", params.LogQP(), ", Level: ", LEVEL)

		ecd, enc, dec, eval := engine.GetMudules(params)
		x0 := engine.EnC(params, ecd, enc, values)
		START_TIME := time.Now()
		y0 := engine.CryptoSqrt(eval, params, x0, deg, start, end)

		fmt.Println("Time (CryptoSqrt): ", time.Since(START_TIME))
		sqrt := engine.DeC(params, ecd, dec, y0)
		_, mean := utils.CheckMAE(sqrt, values, result, S)
		fmt.Println("MAE: ", mean)
		_, mean = utils.CheckMRE(sqrt, values, result, S)
		fmt.Println("MRE: ", mean)

	}

}

func TestExperiment3(t *testing.T) {
	experiment3()
}
