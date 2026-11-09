package p4109

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, nums []int, target []int, expect []int) {
	res := countSuffix(nums, target)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample nums %v target %v, expect %v, but got %v", nums, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 2, 3}
	target := []int{2, 2, 3, 2, 1, 3, 2, 2}
	expect := []int{1, 1, 0, 0, 0, 1}
	runSample(t, nums, target, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 6, 5}
	target := []int{5, 6, 5, 5, 6, 5}
	expect := []int{2, 1, 0, 2}
	runSample(t, nums, target, expect)
}
