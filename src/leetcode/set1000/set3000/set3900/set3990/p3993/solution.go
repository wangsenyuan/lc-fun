package p3993

func maximumValue(n int, s int, m int) int64 {
	// s[0] = s
	// s[1] <= s[0] + m => s[1] = s[0] + m
	// 然后下一个就必须 s[2] >= s[1] - m = s[1] - 1
	if n == 1 {
		return int64(s)
	}

	play := func(x int, n int) int64 {
		// n肯定是偶数
		h := n / 2
		return int64(x + h*m - (h - 1))
	}

	if n&1 == 0 {
		return play(s, n)
	}
	// n是奇数
	return play(s, n-1)
}
