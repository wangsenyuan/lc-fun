package p4027

import (
	"cmp"
	"slices"
	"sort"
)

func elevatorRequests(n int, start int, requests [][]int) int64 {
	// m := len(requests)
	// m <= 16
	requests = append(requests, []int{-1, start})

	slices.SortFunc(requests, func(a []int, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	var floors []int
	for _, cur := range requests {
		floors = append(floors, cur[1])
	}

	const inf = 1 << 60

	slices.Sort(floors)
	floors = slices.Compact(floors)

	m := len(requests)
	k := len(floors)
	// dp[mask][i] 表示到时刻requests[i][0]时的最优解
	dp := make([][]int, 1<<m)
	for i := range dp {
		dp[i] = make([]int, k)
		for j := range k {
			dp[i][j] = inf
		}
	}

	dp[1][sort.SearchInts(floors, start)] = 0

	for mask := 1; mask < 1<<m; mask++ {
		for i := range k {
			if dp[mask][i] < inf {
				for j := range m {
					if (mask>>j)&1 == 0 {
						dist := abs(requests[j][1] - floors[i])
						i1 := sort.SearchInts(floors, requests[j][1])
						dp[mask|(1<<j)][i1] = min(dp[mask|(1<<j)][i1], max(requests[j][0], dp[mask][i]+dist))
					}
				}
			}
		}
	}

	ans := inf

	for i := range k {
		ans = min(ans, dp[(1<<m)-1][i])
	}

	return int64(ans)
}

func abs(num int) int {
	return max(num, -num)
}
