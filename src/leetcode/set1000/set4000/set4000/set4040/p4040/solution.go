package p4040

import "slices"

const inf = 1 << 60

func minOperations(nums []int, sum int) int {
	dp := make([]int, sum+1)
	for i := range sum {
		dp[i+1] = inf
	}
	for _, num := range nums {
		ndp := slices.Clone(dp)
		for i := range sum {
			for x, c := num, 0; x+i <= sum; x *= 2 {
				ndp[x+i] = min(ndp[x+i], dp[i]+c)
				c++
			}
			for x, c := num, 0; x > 0; x /= 2 {
				if x+i <= sum {
					ndp[x+i] = min(ndp[x+i], dp[i]+c)
				}
				c++
			}
		}
		dp = ndp
	}

	if dp[sum] == inf {
		return -1
	}
	return dp[sum]
}
