package p4065

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := rearrangeArray(nums)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample %v, expect %v, but got %v", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 3, 2, 1, 3}
	expect := []int{1, 2, 3, 1, 3, 3}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{7, 7, 4, 4, 4}
	expect := []int{4, 7, 4, 7, 4}
	runSample(t, nums, expect)
}
