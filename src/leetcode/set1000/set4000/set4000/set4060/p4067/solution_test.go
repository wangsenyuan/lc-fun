package p4067

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxSubarray(nums)

	if res != expect {
		t.Fatalf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 5, 3, 2, 1}
	runSample(t, nums, 3)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 4, 5, 6}
	runSample(t, nums, 4)
}
