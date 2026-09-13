package p4053

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minOperations(nums)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{10, 12, 14, 16}
	expect := 9
	runSample(t, nums, expect)
}
