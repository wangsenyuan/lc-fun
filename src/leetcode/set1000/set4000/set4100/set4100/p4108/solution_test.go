package p4108

import "testing"

func runSample(t *testing.T, n int, colors string, edges [][]int, expect int64) {
	res := minCost(n, colors, edges)

	if res != expect {
		t.Fatalf("Sample n=%d colors=%s edges %v, expect %d, but got %d", n, colors, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}}
	runSample(t, 3, "RBR", edges, 1)
}

func TestSample2(t *testing.T) {
	edges := [][]int{{0, 1}, {0, 2}}
	runSample(t, 3, "RBR", edges, 0)
}

func TestSample3(t *testing.T) {
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}}
	runSample(t, 4, "BRRR", edges, 2)
}

func TestSample4(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}}
	runSample(t, 3, "BBB", edges, 0)
}
