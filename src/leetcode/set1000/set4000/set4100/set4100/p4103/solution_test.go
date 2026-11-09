package p4103

import "testing"

func runSample(t *testing.T, types []int, colors []int, groupSize []int, expect bool) {
	t.Helper()
	res := canRemoveAll(types, colors, groupSize)
	if res != expect {
		t.Fatalf("Sample types %v colors %v groupSize %v expect %v, but got %v", types, colors, groupSize, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{0, 1, 1, 0}, []int{5, 3, 3, 5}, []int{2, 2}, true)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{0, 1, 0}, []int{9, 9, 9}, []int{3, 2}, false)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{2, 0, 1, 2}, []int{7, 5, 4, 7}, []int{3, 1, 2}, false)
}
