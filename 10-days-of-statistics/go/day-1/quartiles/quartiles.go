package quartiles

import (
	"fmt"
	"sort"
)

func sum(values []int) int {
	val := 0
	for _, v := range values {
		val += v
	}

	return val
}

func median(numbers []int) int {
	size := len(numbers)
	if size%2 == 0 {
		fmt.Println(numbers[size/2])
		return sum(numbers[size/2-1:size/2+1]) / 2
	}
	return numbers[size/2]
}

func Quartiles(n int, numbers []int) (q1, q2, q3 int) {
	q1, q2, q3 = 0, 0, 0
	sort.Ints(numbers)
	q1 = median(numbers[:n/2])
	q2 = median(numbers)
	if n%2 == 0 {
		q3 = median(numbers[n/2:])
	} else {
		q3 = median(numbers[n/2+1:])
	}

	return
}
