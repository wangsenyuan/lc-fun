package p4096

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countStaircase(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, 8, 8, 2, 2, 2}
	expect := 9
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 6, 9, 9}
	expect := 5
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 3, 4, 4, 3, 4}
	expect := 8
	runSample(t, nums, expect)
}
