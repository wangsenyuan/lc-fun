package p4122

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, requiredSignals []int, expect []int64) {
	res := earliestActivationTimes(n, edges, requiredSignals)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1, 1}, {0, 2, 1}}
	requiredSignals := []int{1, 1, 2}
	expect := []int64{0, 1, -1}
	runSample(t, n, edges, requiredSignals, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	edges := [][]int{{0, 1, 1}, {0, 2, 1}, {1, 3, 2}, {2, 3, 2}}
	requiredSignals := []int{3, 1, 1, 2}
	expect := []int64{0, 1, 1, 3}
	runSample(t, n, edges, requiredSignals, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	edges := [][]int{{0, 1, 2}, {1, 2, 3}, {0, 2, 5}, {2, 3, 1}}
	requiredSignals := []int{3, 1, 2, 1}
	expect := []int64{0, 2, 5, 6}
	runSample(t, n, edges, requiredSignals, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	edges := [][]int{{0, 1, 2}, {1, 0, 1}, {1, 2, 1}, {2, 3, 1}, {3, 2, 1}}
	requiredSignals := []int{3, 1, 2, 1}
	expect := []int64{0, 2, -1, -1}
	runSample(t, n, edges, requiredSignals, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	edges := [][]int{{0, 1, 20}, {0, 2, 14}, {0, 3, 5}, {1, 4, 20}, {1, 2, 5}}
	requiredSignals := []int{1, 1, 2, 1, 1}
	expect := []int64{0, 20, 25, 5, 40}
	runSample(t, n, edges, requiredSignals, expect)
}

func TestSample6(t *testing.T) {
	n := 5
	edges := [][]int{{0, 3, 29}, {4, 3, 2}, {0, 4, 5}}
	requiredSignals := []int{1, 1, 1, 2, 1}
	expect := []int64{0, -1, -1, 29, 5}
	runSample(t, n, edges, requiredSignals, expect)
}

func TestSample7(t *testing.T) {
	n := 5
	edges := [][]int{{0, 1, 1}, {0, 2, 9}, {0, 3, 19}, {2, 4, 14}, {1, 2, 1}}
	requiredSignals := []int{1, 1, 1, 1, 1}
	expect := []int64{0, 1, 2, 19, 16}
	runSample(t, n, edges, requiredSignals, expect)
}

func TestSample8(t *testing.T) {
	n := 5
	edges := [][]int{{3, 4, 17}, {2, 4, 11}, {4, 0, 6}, {0, 3, 11}, {3, 1, 4}}
	requiredSignals := []int{1, 1, 1, 1, 1}
	expect := []int64{0, 15, -1, 11, 28}
	runSample(t, n, edges, requiredSignals, expect)
}
