package p4050

const inf = 1 << 60

func minDays(n int) int {
	dp := make([]int, n+1)

	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = inf
		for j := 1; j*(j+1)/2 <= i; j++ {
			w := dp[i-j*(j+1)/2] + j
			if i > j*(j+1)/2 {
				w++
			}
			dp[i] = min(dp[i], w)
		}
	}

	return dp[n]
}
