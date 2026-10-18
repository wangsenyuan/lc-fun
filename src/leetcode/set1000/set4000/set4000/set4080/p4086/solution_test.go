package p4086

import "testing"

func runSample(t *testing.T, availableIntervals [][]int, targetIntervals [][]int, expect int) {
	res := minSelection(availableIntervals, targetIntervals)

	if res != expect {
		t.Fatalf("Sample available %v target %v, expect %d, but got %d", availableIntervals, targetIntervals, expect, res)
	}
}

func TestSample1(t *testing.T) {
	available := [][]int{{1, 4}, {1, 7}, {3, 9}, {5, 10}, {10, 15}}
	target := [][]int{{1, 3}, {2, 5}, {7, 8}, {10, 13}, {11, 15}, {12, 13}}
	runSample(t, available, target, 3)
}

func TestSample2(t *testing.T) {
	available := [][]int{{4, 6}, {1, 4}, {2, 5}}
	target := [][]int{{3, 4}, {2, 5}, {6, 7}}
	runSample(t, available, target, -1)
}

func TestSample3(t *testing.T) {
	available := [][]int{{1, 2}, {3, 4}}
	target := [][]int{{1, 4}}
	runSample(t, available, target, -1)
}
