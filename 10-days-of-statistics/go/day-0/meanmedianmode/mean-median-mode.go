package meanmedianmode

import (
	"sort"
)

func Mean(n []int) float64 {
	size := float64(len(n))
	t := 0.0
	for _, v := range n {
		t += float64(v)
	}

	return t / size
}

func Median(n []int) float64 {
	size := len(n)
	sort.Ints(n)
	median := Mean(n[size/2-1 : size/2+1])

	return median
}
