package utils

import "math"

func CheckMAE(x []float64, input []float64, targets []float64, len int) (float64, float64) {
	sum := 0.0
	max := 0.0
	elm := 0.0
	for i := range len {
		diff := math.Abs(x[i] - targets[i])
		sum += diff
		if diff > max {
			elm = input[i]
		}
	}
	avg := sum / float64(len)
	return elm, avg
}

func CheckMRE(x []float64, input []float64, targets []float64, len int) (float64, float64) {
	sum := 0.0
	max := 0.0
	elm := 0.0
	for i := range len {
		diff := math.Abs(1 - x[i]/targets[i])
		sum += diff
		if diff > max {
			max = diff
			elm = input[i]
		}
	}
	avg := sum / float64(len)
	return elm, avg
}

// Linspace generates a slice of n evenly spaced values between start and stop.
func Linspace(start, stop float64, n int) []float64 {
	if n < 2 {
		panic("n must be at least 2")
	}
	// if stop > 1000 {
	// 	stop = 1000
	// }
	step := (stop - start) / float64(n-1)
	linspace := make([]float64, n)

	for i := 0; i < n; i++ {
		linspace[i] = start + step*float64(i)
	}

	return linspace
}

// Mean
func Mean(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}

	sum := 0.0
	for _, value := range arr {
		sum += value
	}
	return sum / float64(len(arr))
}

// Standard Deviation
func stdDev(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}

	m := Mean(arr)
	var sumSquares float64
	for _, value := range arr {
		diff := value - m
		sumSquares += diff * diff
	}
	return math.Sqrt(sumSquares / float64(len(arr)))
}
