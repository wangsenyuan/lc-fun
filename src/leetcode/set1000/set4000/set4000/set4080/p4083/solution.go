package p4083

import "slices"

func minDeletions(nums []int) int {
	n := len(nums)
	nums = slices.Compact(nums)
	return n - len(nums)
}
