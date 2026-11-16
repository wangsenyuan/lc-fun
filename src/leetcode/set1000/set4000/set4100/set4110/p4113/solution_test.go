package p4113

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := maxDifference(nums)

	if res != expect {
		t.Fatalf("Sample nums %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 6, 4}
	runSample(t, nums, 5)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 1, 2, 1, 2}
	runSample(t, nums, 12)
}

func TestSample3(t *testing.T) {
	nums := []int{3, 3, 4, 5}
	runSample(t, nums, -3)
}
