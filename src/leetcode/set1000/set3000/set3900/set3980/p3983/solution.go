package p3983

import "slices"

func canMakeSubsequence(s string, t string) bool {
	// dp[i] = s[:i]能匹配的最短的t的前缀

	play := func(s string, t string) []int {
		n := len(s)
		var j int
		dp := make([]int, n)
		for i, x := range s {
			for j < len(t) && rune(t[j]) != x {
				j++
			}
			j = min(j+1, len(t))
			dp[i] = j
		}
		return dp
	}
	dp := play(s, t)
	fp := play(reverse(s), reverse(t))
	slices.Reverse(fp)
	n := len(s)
	for i := range n {
		var l, r int
		if i > 0 {
			l = dp[i-1]
		}
		if i+1 < n {
			r = fp[i+1]
		}
		if l+r+1 <= len(t) {
			return true
		}
	}

	return false
}

func reverse(s string) string {
	buf := []byte(s)
	slices.Reverse(buf)
	return string(buf)
}
