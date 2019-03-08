package quartiles

import (
	"testing"
)

func TestQuartiles(t *testing.T) {
	data := []struct {
		n          int
		numbers    []int
		q1, q2, q3 int
	}{
		{9, []int{3, 7, 8, 5, 12, 14, 21, 13, 18}, 6, 12, 16},
		{10, []int{3, 7, 8, 5, 12, 14, 21, 15, 18, 14}, 7, 13, 15},
	}

	for _, tt := range data {
		q1, q2, q3 := Quartiles(tt.n, tt.numbers)

		if q1 != tt.q1 || q2 != tt.q2 || q3 != tt.q3 {
			t.Errorf("Quartiles expected %d, %d, %d but got %d, %d, %d",
				tt.q1, tt.q2, tt.q3, q1, q2, q3)
		}
	}
}
