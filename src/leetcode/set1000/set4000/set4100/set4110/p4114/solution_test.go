package p4114

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := countConsistent(nums)

	if res != expect {
		t.Fatalf("Sample nums %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, 3, 1, 2}
	runSample(t, nums, 6)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1, 1}
	runSample(t, nums, 6)
}

func TestSample3(t *testing.T) {
	nums := []int{0}
	runSample(t, nums, 1)
}

func TestSample4(t *testing.T) {
	nums := []int{0, 4}
	runSample(t, nums, 3)
}
