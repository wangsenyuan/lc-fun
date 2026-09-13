package p4051

import "testing"

func runSample(t *testing.T, nums []int, goal int, k int, expect int) {
	res := distantSubarrays(nums, goal, k)
	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 1}
	goal := 4
	k := 1
	expect := 5
	runSample(t, nums, goal, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, -1, 3}
	goal := 2
	k := 2
	expect := 2
	runSample(t, nums, goal, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{-3, 1, 2}
	goal := 0
	k := 3
	expect := 2
	runSample(t, nums, goal, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{16, 26, 41, 20, -25, 18}
	goal := -7
	k := 0
	expect := 21
	runSample(t, nums, goal, k, expect)
}
