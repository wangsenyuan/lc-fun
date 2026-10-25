package main

import "slices"

func countMagicalPairs(s string) int64 {
	n := len(s)
	pos := make([]int, 26)
	for i := range pos {
		pos[i] = -1
	}
	dp := make([][26]int, n)
	for i := range n {
		x := int(s[i] - 'a')
		pos[x] = i
		var todo []int
		for _, v := range pos {
			if v >= 0 {
				todo = append(todo, v)
			}
		}
		slices.Sort(todo)
		slices.Reverse(todo)
		for j, v := range todo {
			if j+1 < len(todo) {
				dp[i][j] = v - todo[j+1]
			} else {
				dp[i][j] = v + 1
			}
		}
	}

	var ans int

	for i := range pos {
		pos[i] = n
	}

	for i := n - 1; i > 0; i-- {
		x := int(s[i] - 'a')
		pos[x] = i
		var todo []int
		for _, v := range pos {
			if v < n {
				todo = append(todo, v)
			}
		}
		slices.Sort(todo)
		for j, v := range todo {
			w := n - v
			if j+1 < len(todo) {
				w = todo[j+1] - v
			}
			ans += w * dp[i-1][j]
		}
	}

	return int64(ans)
}
