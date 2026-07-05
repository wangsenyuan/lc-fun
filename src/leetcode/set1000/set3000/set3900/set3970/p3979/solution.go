package p3979

func maxValidPairSum(nums []int, k int) int {
	var best int
	n := len(nums)
	var ans int
	for i := range n {
		if i >= k {
			best = max(best, nums[i-k])
			ans = max(ans, best+nums[i])
		}
	}

	return ans
}
