package weightedmean

import "math"

func WeightedMean(x []int, w []int) float64 {
	sumx := 0.0
	sumw := 0.0
	size := len(x)

	for i := 0; i < size; i++ {
		sumx += float64(x[i] * w[i])
		sumw += float64(w[i])
	}

	return roundTo(sumx/sumw, 10)
}

func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*float64(decimals)) / float64(decimals)
}
