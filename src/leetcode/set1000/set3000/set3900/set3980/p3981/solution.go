package p3981

func interleaveCharacters(word1 string, word2 string, target string) int {
	n := len(word1)
	m := len(word2)

	const mod = 1e9 + 7

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	dp[0][0] = 1

	for _, c := range target {
		rowPrefix := make([][]int, n+1)
		for i := range rowPrefix {
			rowPrefix[i] = make([]int, m+1)
			for j := 0; j <= m; j++ {
				rowPrefix[i][j] = dp[i][j]
				if i > 0 {
					rowPrefix[i][j] = (rowPrefix[i][j] + rowPrefix[i-1][j]) % mod
				}
			}
		}

		colPrefix := make([][]int, n+1)
		for i := range colPrefix {
			colPrefix[i] = make([]int, m+1)
			for j := 0; j <= m; j++ {
				colPrefix[i][j] = dp[i][j]
				if j > 0 {
					colPrefix[i][j] = (colPrefix[i][j] + colPrefix[i][j-1]) % mod
				}
			}
		}

		ndp := make([][]int, n+1)
		for i := range ndp {
			ndp[i] = make([]int, m+1)
		}

		for i := 0; i < n; i++ {
			if rune(word1[i]) == c {
				for j := 0; j <= m; j++ {
					ndp[i+1][j] = (ndp[i+1][j] + rowPrefix[i][j]) % mod
				}
			}
		}

		for j := 0; j < m; j++ {
			if rune(word2[j]) == c {
				for i := 0; i <= n; i++ {
					ndp[i][j+1] = (ndp[i][j+1] + colPrefix[i][j]) % mod
				}
			}
		}

		dp = ndp
	}

	var ans int

	for i := range n + 1 {
		for j := range m + 1 {
			if i > 0 && j > 0 {
				ans = (ans + dp[i][j]) % mod
			}
		}
	}

	return ans
}
