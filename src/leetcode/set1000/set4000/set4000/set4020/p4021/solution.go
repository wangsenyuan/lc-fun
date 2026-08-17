package p4021

const inf = 1 << 60

func minOperations(s string) int {
	// s[i] = s[n-1-i]
	best := inf

	n := len(s)
	for i := range n {
		// s[i]是 最左边的
		var cnt int
		for l, r := i, i+n-1; l < r; l, r = l+1, r-1 {
			x := int(s[l%n] - 'a')
			y := int(s[r%n] - 'a')
			if x > y {
				x, y = y, x
			}
			cnt += min(y-x, 26-(y-x))
		}
		best = min(best, cnt+i)
	}

	return best
}
