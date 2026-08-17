package p4026

func maximumGap(skill string, station string) int {
	// dp[i] 表示前i个人, 最小的下标
	n := len(skill)
	if n == 1 {
		return 0
	}
	dp := make([]int, n)
	lines := make([][]int, 26)
	m := len(station)
	for i := range m {
		x := int(station[i] - 'a')
		lines[x] = append(lines[x], i)
	}
	for i := range n {
		x := int(skill[i] - 'a')
		prev := -1
		if i > 0 {
			prev = dp[i-1] + 1
		}
		for lines[x][0] < prev {
			lines[x] = lines[x][1:]
		}
		// 答案存在的情况下, 肯定会存在一个lines[x]
		dp[i] = lines[x][0]
		lines[x] = lines[x][1:]
	}

	for x := range 26 {
		lines[x] = lines[x][:0]
	}

	for i := m - 1; i >= 0; i-- {
		x := int(station[i] - 'a')
		lines[x] = append(lines[x], i)
	}
	fp := m

	var best int

	for i := n - 1; i > 0; i-- {
		x := int(skill[i] - 'a')
		fp--
		for lines[x][0] > fp {
			lines[x] = lines[x][1:]
		}
		fp = lines[x][0]
		best = max(best, fp-dp[i-1])
		lines[x] = lines[x][1:]
	}

	return best
}
