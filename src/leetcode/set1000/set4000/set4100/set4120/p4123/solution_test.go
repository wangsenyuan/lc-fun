package p4123

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minCost(nums)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, -7, 3, -5}
	expect := -7
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 0, -5, 2}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{7, -8, 1}
	expect := 8
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{46, 16, 51, -47, -53}
	expect := -38
	runSample(t, nums, expect)
}
