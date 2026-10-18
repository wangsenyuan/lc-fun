package p4084

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, grid [][]byte, expect []int) {
	res := lastVisitedCell(grid)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample %q, expect %v, but got %v", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]byte{
		{'R', 'D'},
		{'U', 'L'},
	}
	runSample(t, grid, []int{1, 0})
}

func TestSample2(t *testing.T) {
	grid := [][]byte{
		{'R', 'R', 'R'},
	}
	runSample(t, grid, []int{0, 2})
}
