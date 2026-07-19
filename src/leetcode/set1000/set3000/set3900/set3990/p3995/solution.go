package p3995

const inf = 1 << 60

func minCost(source string, target string, rules [][]string, costs []int) int {

	n := len(source)
	dp := make([]int, n+1)
	for i := range n + 1 {
		dp[i] = inf
	}

	dp[0] = 0

	m := len(rules)
	for i, cur := range rules {
		for j := range len(cur[0]) {
			if cur[0][j] == '*' {
				costs[i]++
			}
		}
	}

	for i := range n {
		if dp[i] == inf {
			continue
		}
		if source[i] == target[i] {
			dp[i+1] = min(dp[i+1], dp[i])
		}

		var todo []int
		for j := range m {
			todo = append(todo, j)
		}

		for j := 0; i+j < n && len(todo) > 0; j++ {
			var next []int
			for _, k := range todo {
				if (rules[k][0][j] == '*' || rules[k][0][j] == source[i+j]) &&
					target[i+j] == rules[k][1][j] {
					if j+1 == len(rules[k][0]) {
						dp[i+j+1] = min(dp[i+j+1], dp[i]+costs[k])
					} else {
						next = append(next, k)
					}
				}
			}
			todo = next
		}
	}

	if dp[n] == inf {
		return -1
	}

	return dp[n]
}
