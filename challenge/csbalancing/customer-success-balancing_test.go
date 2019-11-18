package csbalancing

import "testing"

func TestCsRush(t *testing.T) {
	data := []struct {
		n          int32
		m          int32
		css        [][]int32
		customers  [][]int32
		vacant_css []int32
		expected   int32
	}{
		{4, 6, [][]int32{{1, 60}, {2, 20}, {3, 95}, {4, 75}},
			[][]int32{{1, 90}, {2, 20}, {3, 70}, {4, 40}, {5, 60}, {6, 10}}, []int32{2, 4}, 1},
		{4, 6, [][]int32{{1, 60}, {2, 20}, {3, 95}, {4, 100}},
			[][]int32{{1, 90}, {2, 20}, {3, 70}, {4, 40}, {5, 60}, {6, 10}}, []int32{1, 3}, 4},
	}

	for _, tt := range data {
		if got := csRush(tt.n, tt.m, tt.css, tt.customers, tt.vacant_css); got != tt.expected {
			t.Errorf("CsRush expected %d but got %d", tt.expected, got)
		}
	}
}
