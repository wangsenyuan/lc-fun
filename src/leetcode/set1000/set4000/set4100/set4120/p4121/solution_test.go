package p4121

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minCost(nums)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 3, 3, 6, 3, 3}
	expect := 9
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{3, 4, 4, 5, 4, 4, 3}
	expect := 20
	runSample(t, nums, expect)
}
