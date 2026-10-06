package p4072

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := maxAlternatingSum(nums)

	if res != expect {
		t.Fatalf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, -5, 1}
	runSample(t, nums, 11)
}

func TestSample2(t *testing.T) {
	nums := []int{10, -5, -100}
	runSample(t, nums, 110)
}

func TestSample3(t *testing.T) {
	nums := []int{4, 7}
	runSample(t, nums, 7)
}
