package p4070

func minRotations(s string) int {
	n := len(s)
	var sum int
	var x int
	for i := range n {
		y := int(s[i] - '0')
		d := abs(x - y)
		sum += min(d, 10-d)
		x = y
	}
	return sum
}

func abs(x int) int {
	return max(x, -x)
}
