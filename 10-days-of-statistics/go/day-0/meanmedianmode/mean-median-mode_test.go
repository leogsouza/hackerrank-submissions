package meanmedianmode

import (
	"testing"
)

func TestMean(t *testing.T) {
	data := []struct {
		numbers  []int
		expected float64
	}{
		{[]int{64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060}, 43900.6},
	}

	for _, tt := range data {
		if got := Mean(tt.numbers); got != tt.expected {
			t.Errorf("Mean expected %f but got %f", tt.expected, got)
		}
	}
}

func TestMedian(t *testing.T) {
	data := []struct {
		numbers  []int
		expected float64
	}{
		{[]int{64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060}, 44627.5},
	}

	for _, tt := range data {
		if got := Median(tt.numbers); got != tt.expected {
			t.Errorf("Median expected %f but got %f", tt.expected, got)
		}
	}
}
