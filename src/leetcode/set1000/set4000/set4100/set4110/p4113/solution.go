package p4113

func maxDifference(nums []int) int64 {
	n := len(nums)
	dp := make([]int, n+1)
	// dp[i] = 如果当前人操作, 它能得到的最优解

	play := func(cur int, sum int, i int) int {
		if cur%3 == 0 {
			// make turn
			return sum - dp[i]
		}
		// still take control
		return dp[i]
	}

	var sum int
	for i := n - 1; i >= 0; i-- {
		sum += nums[i]
		if i >= n-2 {
			// take all
			dp[i] = sum
		} else {
			dp[i] = nums[i] + play(nums[i], sum-nums[i], i+1)
			dp[i] = max(dp[i], nums[i]+nums[i+1]+play(nums[i]+nums[i+1], sum-(nums[i]+nums[i+1]), i+2))
		}
	}
	return int64(dp[0] - (sum - dp[0]))
}
