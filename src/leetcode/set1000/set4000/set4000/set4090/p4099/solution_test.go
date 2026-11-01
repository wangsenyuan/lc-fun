package p4099

import "testing"

func runSample(t *testing.T, nums []int, locked []int, k int, expect int) {
	res := minOperations(nums, locked, k)

	if res != expect {
		t.Fatalf("Sample nums %v locked %v k %d, expect %d, but got %d", nums, locked, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 3, 4}
	locked := []int{0, 1, 0, 1}
	runSample(t, nums, locked, 1, 2)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 1, 2}
	locked := []int{1, 0, 1}
	runSample(t, nums, locked, 2, 4)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 3, 1}
	locked := []int{1, 0, 0}
	runSample(t, nums, locked, 0, -1)
}
