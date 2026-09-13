package p4055

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := shadowPairs(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 4, 2, 5}
	expect := 5
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 7, 8, 9}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 4, 4}
	expect := 2
	runSample(t, nums, expect)
}
