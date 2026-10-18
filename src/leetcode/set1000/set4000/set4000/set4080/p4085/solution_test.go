package p4085

import "testing"

func runSample(t *testing.T, edges [][]int, colors string, expect int) {
	res := findColoredRoot(edges, colors)

	if res != expect {
		t.Fatalf("Sample edges %v colors %s, expect %d, but got %d", edges, colors, expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{{3, 1}, {3, 2}, {1, 4}, {2, 0}}
	colors := "BRRWB"
	runSample(t, edges, colors, 3)
}

func TestSample2(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}}
	colors := "RBRB"
	runSample(t, edges, colors, -1)
}
