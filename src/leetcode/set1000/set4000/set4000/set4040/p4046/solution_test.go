package p4046

import "testing"

func runSample(t *testing.T, grid [][]int, k int, expect int) {
	res := minCost(grid, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSampele1(t *testing.T) {
	grid := [][]int{
		{2, 7, 3}, {1, 4, 5},
	}
	k := 1
	expect := 12
	runSample(t, grid, k, expect)
}

func TestSampele2(t *testing.T) {
	grid := [][]int{
		{4, 1, 9}, {3, 2, 5}, {4, 8, 6},
	}
	k := 2
	expect := 20
	runSample(t, grid, k, expect)
}
