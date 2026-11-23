package p4118

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minOperations(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 2, 3}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{9, 4, 9}
	expect := 6
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{5, 7, 7, 4}
	expect := 1
	runSample(t, nums, expect)
}
