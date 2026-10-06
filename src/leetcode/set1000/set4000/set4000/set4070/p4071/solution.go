package p4071

func minRotations(n int, s string) int {
	var sum int
	var x int
	for i := range n {
		y := int(s[i] - '0')
		d := abs(x - y)
		sum += min(d, 10-d)
		x = y
	}
	last := int(s[n-1] - '0')
	var best int
	for i := range n {
		y := int(s[i] - '0')
		var z int
		if i > 0 {
			z = int(s[i-1] - '0')
		}
		d1 := abs(y - z)
		w1 := min(d1, 10-d1)
		d2 := abs(z - last)
		w2 := min(d2, 10-d2)
		best = min(best, w2-w1)
	}
	return sum + best
}

func abs(x int) int {
	return max(x, -x)
}
