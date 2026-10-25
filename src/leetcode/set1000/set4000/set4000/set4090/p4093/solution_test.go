package p4093

import "testing"

func runSample(t *testing.T, nums []int, k, expect int) {
	res := maximizeSum(nums, k)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	nums := []int{102, 90, 14}
	k := 2
	expect := 341
	runSample(t, nums, k, expect)
}
