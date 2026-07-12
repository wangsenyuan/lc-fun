package p3989

import "slices"

func maxConsistentColumns(grid [][]int, limit int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([]int, m)

	for c := range m {
		dp[c] = 1
		for c1 := range c {
			var diff int
			for r := range n {
				diff = max(diff, abs(grid[r][c]-grid[r][c1]))
			}
			if diff <= limit {
				dp[c] = max(dp[c], dp[c1]+1)
			}
		}
	}
	return slices.Max(dp)
}

func abs(num int) int {
	return max(num, -num)
}
