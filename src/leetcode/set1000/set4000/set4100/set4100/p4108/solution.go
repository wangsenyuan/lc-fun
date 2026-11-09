package p4108

func minCost(n int, colors string, edges [][]int) int64 {
	sz := make([]int, n)
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	var base int
	var mx int

	dp := make([]int, n)

	var dfs func(int, int)
	dfs = func(p int, u int) {
		sz[u] = 1
		for _, v := range g[u] {
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}

		if colors[u] == 'R' {
			base += sz[u]
			dp[u] = sz[u]
		} else {
			dp[u] = -sz[u]
		}

		for _, v := range g[u] {
			if p != v {
				dp[u] += max(0, dp[v])
			}
		}

		mx = max(mx, dp[u])
	}

	dfs(-1, 0)

	return int64(base - mx)
}
