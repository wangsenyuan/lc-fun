package p4033

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, nums []int, k int, queries [][]int, expect []bool) {
	res := validSubarrays(nums, k, queries)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 2, 1}
	k := 2
	queries := [][]int{
		{0, 1}, {0, 3}, {1, 2},
	}
	expect := []bool{false, true, false}
	runSample(t, nums, k, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 3, 3}
	k := 1
	queries := [][]int{
		{1, 2}, {0, 2},
	}
	expect := []bool{true, false}
	runSample(t, nums, k, queries, expect)
}
