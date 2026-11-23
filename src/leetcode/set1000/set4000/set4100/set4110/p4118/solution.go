package p4118

import "slices"

const inf = 1 << 60

func minOperations(nums []int) int {
	w := slices.Max(nums)
	n := len(nums)
	w = max(w, n) * 2
	play := func(arr []int) [][]int {
		dp := make([][]int, n)
		for i := range n {
			dp[i] = make([]int, w+1)
			for j := range w + 1 {
				dp[i][j] = inf
			}
		}

		for j := 1; j <= nums[0]; j++ {
			dp[0][j] = nums[0] - j
		}

		clear(dp[0][nums[0]:])

		for i := 1; i < n; i++ {
			v := arr[i]
			for j := i + 1; j <= w; j++ {
				dp[i][j] = abs(v-j) + dp[i-1][j-1]
				dp[i][j] = min(dp[i][j], dp[i][j-1])
			}
		}

		return dp
	}

	dp := play(nums)
	slices.Reverse(nums)
	fp := play(nums)
	slices.Reverse(fp)
	slices.Reverse(nums)

	ans := inf

	for i := 1; i < n-1; i++ {
		for v := max(i+1, n-i); v <= w; v++ {
			ans = min(ans, abs(v-nums[i])+dp[i-1][v-1]+fp[i+1][v-1])
		}
	}

	if ans == inf {
		return -1
	}

	return ans
}

func abs(num int) int {
	return max(num, -num)
}
