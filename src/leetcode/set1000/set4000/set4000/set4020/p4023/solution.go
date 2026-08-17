package p4023

import (
	"slices"
	"sort"
)

func elevatorRequests(n int, start int, requests []int) int64 {
	slices.Sort(requests)
	pos := sort.SearchInts(requests, start)
	if pos == len(requests) || requests[pos] != start {
		requests = append(requests, start)
		slices.Sort(requests)
	}

	const inf = 1 << 60
	m := len(requests)
	dp := make([][][2]int, m)
	for i := range m {
		dp[i] = make([][2]int, m)
		for j := range m {
			for d := range 2 {
				dp[i][j][d] = inf
			}
		}
	}

	k := sort.SearchInts(requests, start)
	dp[k][k][0] = 0
	dp[k][k][1] = 0

	for w := 2; w <= m; w++ {
		for i := 0; i+w <= m; i++ {
			j := i + w - 1
			t := m + 1 - w
			dp[i][j][0] = min(
				dp[i+1][j][0]+(requests[i+1]-requests[i])*t,
				dp[i+1][j][1]+(requests[j]-requests[i])*t,
			)
			dp[i][j][1] = min(
				dp[i][j-1][0]+(requests[j]-requests[i])*t,
				dp[i][j-1][1]+(requests[j]-requests[j-1])*t,
			)
		}
	}

	ans := min(dp[0][m-1][0], dp[0][m-1][1])
	return int64(ans)
}
