package p4064

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := longestSubarray(nums, k)

	if res != expect {
		t.Fatalf("Sample %v k=%d, expect %d, but got %d", nums, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 1, 2}
	k := 3
	expect := 3
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 3, 4}
	k := 7
	expect := 2
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 2, 5}
	k := 6
	expect := 2
	runSample(t, nums, k, expect)
}
