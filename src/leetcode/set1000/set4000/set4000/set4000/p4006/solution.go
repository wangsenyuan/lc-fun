package p4006

func countValidPrefixes(s string) int {
	var res int
	cnt := make([]int, 2)
	for i := range s {
		cnt[int(s[i]-'0')]++
		if abs(cnt[0]-cnt[1]) <= 1 {
			res++
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
