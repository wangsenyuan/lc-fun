package p4078

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := resilientSubarray(nums, k)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 4, 6, 3}
	k := 2
	expect := 3
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 4, 7, 1}
	k := 3
	expect := 4
	runSample(t, nums, k, expect)
}
