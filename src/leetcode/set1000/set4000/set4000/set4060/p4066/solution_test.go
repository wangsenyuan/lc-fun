package p4066

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxEqualAdjacentPairs(nums)

	if res != expect {
		t.Fatalf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 2}
	runSample(t, nums, 2)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 1, 2, 1}
	runSample(t, nums, 4)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1, 1}
	runSample(t, nums, 2)
}

func TestSample4(t *testing.T) {
	nums := []int{2, 8, 2, 5, 5, 6}
	runSample(t, nums, 3)
}

func TestSample5(t *testing.T) {
	nums := []int{3, 9, 2, 2, 8, 3, 9, 8}
	runSample(t, nums, 3)
}
