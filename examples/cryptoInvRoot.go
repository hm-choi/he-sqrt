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
	LogN, LEVEL, d1, d2 := 17, 9+1+8, 9, 4

	for _, SCALE := range SCALE_LIST {
		fmt.Println("SCALE ", SCALE)
		params := engine.GetParam(LogN, LEVEL, SCALE)
		S := params.MaxSlots()
		fmt.Println("LogQP: ", params.LogQP())
		ecd, enc, dec, eval := engine.GetMudules(params)

		values := make([]float64, S)
		result := make([]float64, S)
		start, mid, end := 1e-6, 1e-2, 1e+2

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

		fmt.Println("Input: ", values[:2], values[params.MaxSlots()-2:])
		fmt.Println("result:", result[:2], result[params.MaxSlots()-2:])
		x0 := engine.EnC(params, ecd, enc, values)
		fmt.Println("x0 size", x0.BinarySize())
		TIME_LIST := make([]float64, 30)
		MAE_LIST := make([]float64, 30)
		MRE_LIST := make([]float64, 30)
		for idx := range 1 {
			START_TIME := time.Now()
			y0 := engine.CryptoInvRoot(eval, params, x0, d1, start, end, 1)
			invSqrts := engine.HENewtonInvSqrt(eval, x0, d2, y0)
			TIME_LIST[idx] = float64(time.Since(START_TIME)) / 1000000000.0
			fmt.Println("Time ", time.Since(START_TIME))
			invSqrt := engine.DeC(params, ecd, dec, invSqrts)
			aa, mean, max := utils.CheckMAE(invSqrt, values, result, S)
			MAE_LIST[idx] = mean
			fmt.Println("MAE(idx, mean, max): ", aa, mean, max)
			aa, mean, max = utils.CheckMRE(invSqrt, values, result, S)
			MRE_LIST[idx] = mean
			fmt.Println("MRE(idx, mean, max): ", aa, mean, max)
		}
		// fmt.Println("Scale, TIME Mean, Time Std: ", SCALE, mean(TIME_LIST), stdDev(TIME_LIST))
		// fmt.Println("Scale, MAE Mean, MAE Std:   ", SCALE, mean(MAE_LIST), stdDev(MAE_LIST))
		// fmt.Println("Scale, MRE Mean, MRE Std:   ", SCALE, mean(MRE_LIST), stdDev(MRE_LIST))
	}
}

// 평균 구하기
func mean(arr []float64) float64 {
	if len(arr) == 0 {
		return 0 // 배열이 비어있으면 0 반환
	}

	sum := 0.0
	for _, value := range arr {
		sum += value
	}
	return sum / float64(len(arr))
}

// 표준편차 구하기
func stdDev(arr []float64) float64 {
	if len(arr) == 0 {
		return 0 // 배열이 비어있으면 0 반환
	}

	m := mean(arr) // 평균 계산
	var sumSquares float64
	for _, value := range arr {
		diff := value - m
		sumSquares += diff * diff
	}
	return math.Sqrt(sumSquares / float64(len(arr)))
}
