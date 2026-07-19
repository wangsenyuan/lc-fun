package p3998

func transformStr(s string, strs []string) []bool {
	n := len(s)
	suf := make([]int, 2)
	for i := range n {
		suf[int(s[i]-'0')]++
	}
	buf := make([]byte, n)
	play := func(t string) bool {
		cnt := make([]int, 2)
		for i := range n {
			switch t[i] {
			case '0':
				cnt[0]++
			case '1':
				cnt[1]++
			}
		}

		marks := n - cnt[0] - cnt[1]

		if cnt[0] > suf[0] || suf[0] > cnt[0]+marks {
			return false
		}
		// len(t) = n
		// 1可以往后面移动, 0可以往前移动

		copy(buf, t)

		for i := range n {
			if t[i] == '?' {
				if cnt[0] < suf[0] {
					buf[i] = '0'
					cnt[0]++
				} else {
					buf[i] = '1'
					cnt[1]++
				}
			}
		}
		i, j := 0, 0
		for range suf[0] {
			for i < n && s[i] != '0' {
				i++
			}
			for j < n && buf[j] != '0' {
				j++
			}
			if i < j {
				return false
			}
			i++
			j++
		}

		return true
	}

	res := make([]bool, len(strs))

	for i, x := range strs {
		res[i] = play(x)
	}

	return res
}
