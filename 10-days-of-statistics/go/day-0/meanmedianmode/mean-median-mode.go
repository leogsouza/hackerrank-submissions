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

func Mode(n []int) int {
	size := len(n)
	sort.Ints(n)
	smallest := n[0]
	m := make(map[int]int, size)

	for _, v := range n {
		m[v]++
	}

	max := 0
	num := 0
	for _, v := range n {
		if m[v] > max {
			max = m[v]
			num = v
		}
	}

	if smallest == max {
		return smallest
	}

	return num
}
