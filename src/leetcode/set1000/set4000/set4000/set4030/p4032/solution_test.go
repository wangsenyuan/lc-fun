package p4032

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := longestSubarray(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{7, 6, 10, 12, 11}
	k := 3
	expect := 3
	runSample(t, nums, k, expect)
}


