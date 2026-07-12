package p3987

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minimumCost(nums, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000}
	k := 1
	expect := 630
	runSample(t, nums, k, expect)
}
