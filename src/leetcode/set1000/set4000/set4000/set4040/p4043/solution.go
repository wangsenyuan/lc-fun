package p4043

func countRotations(s string, k int) int {
	n := len(s)
	var x int
	for i := range n - 1 {
		if s[i] == s[i+1] {
			x++
		}
	}

	var res int

	for i := range n {
		if k == x {
			res++
		}
		// 如果 s[i] 是头部
		j := (i + 1) % n
		if s[i] == s[j] {
			x--
		}
		j = (i - 1 + n) % n
		if s[i] == s[j] {
			x++
		}
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}
