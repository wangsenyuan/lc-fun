package p3978

func isMiddleElementUnique(nums []int) bool {
	n := len(nums)
	mid := nums[n/2]

	for i := range n {
		if i != n/2 && nums[i] == mid {
			return false
		}
	}
	return true
}
