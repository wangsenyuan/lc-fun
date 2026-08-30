package p4037

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxValidSplits(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{10, 30, 15, 10}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 10, 14}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 4}
	expect := 0
	runSample(t, nums, expect)
}
