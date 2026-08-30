package p4041

import "slices"

const inf = 1 << 60

func minOperations(nums []int, sum int) int {
	dp := make([]int, sum+1)
	for i := range sum + 1 {
		dp[i] = inf
	}
	dp[0] = 0
	for _, num := range nums {
		var todo [][2]int
		var cnt int
		var prev int
		for num > 0 {
			todo = append(todo, [2]int{num, cnt})

			if num*2 != prev {
				cnt2 := 1
				for x := 2 * num; x <= sum; x *= 2 {
					todo = append(todo, [2]int{x, cnt + cnt2})
					cnt2++
				}
			}

			prev = num
			num /= 2
			cnt++
		}
		ndp := slices.Clone(dp)
		for x := 0; x <= sum; x++ {
			for _, cur := range todo {
				if x+cur[0] <= sum {
					ndp[x+cur[0]] = min(ndp[x+cur[0]], dp[x]+cur[1])
				}
			}
		}
		dp = ndp
	}
	if dp[sum] == inf {
		return -1
	}
	return dp[sum]
}
