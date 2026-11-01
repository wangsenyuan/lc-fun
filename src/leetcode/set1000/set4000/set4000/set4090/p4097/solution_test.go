package p4097

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minMergedInteger(nums)

	if res != expect {
		t.Fatalf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{23, 55, 42, 21}
	runSample(t, nums, 2355)
}

func TestSample2(t *testing.T) {
	nums := []int{34, 7, 11, 11}
	runSample(t, nums, 347)
}

func TestSample3(t *testing.T) {
	nums := []int{9, 20}
	runSample(t, nums, 920)
}
