package p4003

import "testing"

func runSample(t *testing.T, m int, n int, penalty [][]int, expect int64) {
	res := minCost(m, n, penalty)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, [][]int{
		{5, 3}, {1, 4},
	}, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 2, [][]int{
		{0, 7}, {3, 2},
	}, 7)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 3, [][]int{
		{8, 0, 9}, {7, 4, 1},
	}, 12)
}
