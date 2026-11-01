package p4100

import "testing"

func runSample(t *testing.T, edges [][]int, values []int, expect int) {
	res := minActivationCost(edges, values)

	if res != expect {
		t.Fatalf("Sample edges %v values %v, expect %d, but got %d", edges, values, expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{{0, 1, 4}, {0, 2, 6}, {0, 3, 3}, {3, 4, 5}}
	values := []int{3, 1, 7, 1, 3}
	runSample(t, edges, values, 7)
}

func TestSample2(t *testing.T) {
	edges := [][]int{{0, 1, 5}, {1, 2, 10}}
	values := []int{1, 7, 3}
	runSample(t, edges, values, 5)
}

func TestSample3(t *testing.T) {
	edges := [][]int{{0, 1, 5}, {0, 2, 7}, {0, 3, 9}}
	values := []int{1, 3, 5, 9}
	runSample(t, edges, values, -1)
}
