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
	}

	for _, tt := range data {
		if got := WeightedMean(tt.x, tt.w); got != tt.expected {
			t.Errorf("WeightedMeand expected %f but got %f", tt.expected, got)
		}
	}

}
