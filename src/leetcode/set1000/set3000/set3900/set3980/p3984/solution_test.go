package p3984

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := divisibleGame(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 4, 6, 8}
	expect := 36
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 1, 2}
	expect := 6
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1}
	expect := 1000000005
	runSample(t, nums, expect)
}
