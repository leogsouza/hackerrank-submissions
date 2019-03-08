package meanmedianmode

import (
	"testing"
)

func TestMedian(t *testing.T) {
	data := []struct {
		numbers  []int
		expected float64
	}{
		{[]int{10, 20, 30, 40}, 25},
	}

	for _, tt := range data {
		if got := Median(tt.numbers); got != tt.expected {
			t.Errorf("Median expected %f but got %f", tt.expected, got)
		}
	}
}
