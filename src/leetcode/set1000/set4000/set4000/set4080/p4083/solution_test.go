package p4083

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minDeletions(nums)

	if res != expect {
		t.Fatalf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 2, 2, 2, 3}
	runSample(t, nums, 3)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 4, 9, 4}
	runSample(t, nums, 1)
}
