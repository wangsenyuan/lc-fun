package p4067

import (
	"slices"
	"sort"
)

func maxSubarray(nums []int) int {
	// 如果 a + b = c
	// 那么找到最近的c, 那么可以找到[a, c]的对

	x := slices.Max(nums)
	pos := make([][]int, x+1)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}
	n := len(nums)
	dp := make([]int, n)
	for i := range n {
		dp[i] = n
		for j := i + 1; j < dp[i]; j++ {
			c := nums[i] + nums[j]
			if c <= x {
				if k := sort.SearchInts(pos[c], j+1); k < len(pos[c]) {
					dp[i] = min(dp[i], pos[c][k])
				}
			}
			// 如果 nums[i] 是最大值
			c = abs(nums[i] - nums[j])
			if c > 0 {
				if k := sort.SearchInts(pos[c], j+1); k < len(pos[c]) {
					dp[i] = min(dp[i], pos[c][k])
				}
			}
		}
	}
	// 如果 l...r 中间存在一对, 就是无效的
	var res int
	for i := range n {
		r := dp[i]
		for j := i + 1; j < r; j++ {
			r = min(r, dp[j])
		}
		res = max(res, r-i)
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}
