package weightedmean

import (
	"testing"
)

func TestWeightedMean(t *testing.T) {
	data := []struct {
		x        []int
		w        []int
		expected float64
	}{
		{[]int{10, 40, 30, 50, 20}, []int{1, 2, 3, 4, 5}, 32.0},
		{[]int{10, 40, 30, 50, 20, 10, 40, 30, 50, 20}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 31.1},
	}

	for _, tt := range data {
		if got := WeightedMean(tt.x, tt.w); got != tt.expected {
			t.Errorf("WeightedMean expected %.1f but got %.1f", tt.expected, got)
		}
	}

}
