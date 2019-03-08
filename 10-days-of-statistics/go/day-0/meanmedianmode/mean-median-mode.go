package meanmedianmode

func Median(n []int) float64 {
	size := float64(len(n))
	t := 0.0
	for _, v := range n {
		t += float64(v)
	}

	return t / size
}
