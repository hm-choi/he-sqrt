package main

import (
	"fmt"
	"math"
	"time"

	"github.com/hm-choi/he-sqrt/engine"
	"github.com/hm-choi/he-sqrt/utils"
)

// 1 + 29 + 1+ 16 = 47
func main() {
	LogN, LEVEL, SCALE, d := 17, 47, 50, 8
	params := engine.GetParam(LogN, LEVEL, SCALE)
	S := params.MaxSlots()

	fmt.Println("LogQP: ", params.LogQP())
	ecd, enc, dec, eval := engine.GetMudules(params)

	result := make([]float64, S)
	values := make([]float64, S)
	start, mid, end := 0.001, 1.0, 1000.0

	values1 := utils.Linspace(start, mid, S/2)
	values2 := utils.Linspace(mid, end, S/2)
	for i := range S {
		if i < S/2 {
			values[i] = values1[i]
		} else {
			values[i] = values2[i-S/2]
		}
		// values[i] = start * float64(i+1)
		result[i] = 1.0 / math.Sqrt(values[i])
	}

	fmt.Println("origin[:6] ", values[:3], values[S-3:])
	fmt.Println("invSqrt[:6]", result[:3], result[S-3:])
	x0 := engine.EnC(params, ecd, enc, values)

	START_TIME := time.Now()
	y0 := engine.TwoLineApprox(eval, params, ecd, enc, x0, d, start, end)
	x_result := engine.DeC(params, ecd, dec, x0)
	y_result := engine.DeC(params, ecd, dec, y0)
	fmt.Println("x_result: ", x_result[:3], x_result[S-3:])
	fmt.Println("y_result: ", y_result[:3], y_result[S-3:])

	invSqrts := engine.HENewtonInvSqrt(eval, x0, d, y0)
	invSqrt := engine.DeC(params, ecd, dec, invSqrts)
	fmt.Println("invSqrt: ", invSqrt[:6])
	idx, mean, max := utils.CheckMAE(invSqrt, values, result)
	fmt.Println("MAE(idx, mean, max): ", idx, mean, max)
	idx, mean, max = utils.CheckMRE(invSqrt, values, result)
	fmt.Println("MRE(idx, mean, max): ", idx, mean, max)
	fmt.Println(invSqrts.Level())
	fmt.Println("Operation Time: ", time.Since(START_TIME))
}
