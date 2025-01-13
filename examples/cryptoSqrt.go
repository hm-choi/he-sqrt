package main

import (
	"fmt"
	"math"
	"time"

	"github.com/hm-choi/he-sqrt/engine"
	"github.com/hm-choi/he-sqrt/utils"
)

func main() {
	SCALE_LIST := []int{40}
	degree := 12
	LogN, LEVEL, d1, d2 := 17, degree+8+2, degree, 4

	for _, SCALE := range SCALE_LIST {
		fmt.Println("SCALE ", SCALE)
		params := engine.GetParam(LogN, LEVEL, SCALE)
		S := params.MaxSlots()
		fmt.Println("LogQP: ", params.LogQP())
		ecd, enc, dec, eval := engine.GetMudules(params)

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
		x0 := engine.EnC(params, ecd, enc, values)

		TIME_LIST := make([]float64, 30)
		MAE_LIST := make([]float64, 30)
		MRE_LIST := make([]float64, 30)
		for idx := range 1 {
			START_TIME := time.Now()
			y0 := engine.CryptoInvRoot(eval, params, x0, d1, start, end, 1)
			invSqrts := engine.HENewtonInvSqrt(eval, x0, d2, y0)

			eval.MulRelin(invSqrts, x0, invSqrts)
			eval.Rescale(invSqrts, invSqrts)

			TIME_LIST[idx] = float64(time.Since(START_TIME)) / 1000000000.0
			fmt.Println("Time ", time.Since(START_TIME))
			invSqrt := engine.DeC(params, ecd, dec, invSqrts)
			fmt.Println(invSqrt[:3])
			fmt.Println(result[:3])
			aa, mean, max := utils.CheckMAE(invSqrt, values, result)
			MAE_LIST[idx] = mean
			fmt.Println("MAE(idx, mean, max): ", aa, mean, max)
			aa, mean, max = utils.CheckMRE(invSqrt, values, result)
			MRE_LIST[idx] = mean
			fmt.Println("MRE(idx, mean, max): ", aa, mean, max)
		}
	}
}
