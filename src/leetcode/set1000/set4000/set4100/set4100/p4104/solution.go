package p4104

func countValidSubstrings(s string) int64 {
	pos := make([][]int, 2)

	check := func() bool {
		if len(pos[0]) == 0 || len(pos[1]) == 0 {
			return false
		}
		i := pos[0][0]
		j := pos[1][len(pos[1])-1]
		return i < j
	}

	rem := func(i int) {
		x := int(s[i] - 'a')
		if x == 0 {
			pos[0] = pos[0][1:]
		} else if x == 1 {
			pos[1] = pos[1][1:]
		}
	}

	add := func(i int) {
		x := int(s[i] - 'a')
		if x == 0 {
			pos[0] = append(pos[0], i)
		} else if x == 1 {
			pos[1] = append(pos[1], i)
		}
	}

	var ans int
	var l int
	for i := range s {
		add(i)
		for s[i] == 'c' && check() {
			rem(l)
			l++
		}
		ans += i - l + 1
	}

	return int64(ans)
}
