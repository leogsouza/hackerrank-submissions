package weightedmean

func WeightedMean(x []int, w []int) float64 {
	sumx := 0
	sumw := 0
	size := len(x)

	for i := 0; i < size; i++ {
		sumx += (x[i] * w[i])
		sumw += w[i]
	}

	return float64(sumx / sumw)
}
