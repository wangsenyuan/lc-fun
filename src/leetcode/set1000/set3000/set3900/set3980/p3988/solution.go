package p3988

func createGrid(m int, n int, k int) []string {
	// k <= 4
	// C(m + n - 1, n - 1)
	ans := make([]string, m)
	buf := make([]byte, n)

	play := func(w int) string {
		for i := range n {
			buf[i] = '#'
			if i >= w {
				buf[i] = '.'
			}
		}
		return string(buf)
	}
	if k == 1 {
		ans := make([]string, m)
		ans[0] = play(0)
		for i := 1; i < m; i++ {
			ans[i] = play(n - 1)
		}
		return ans
	}
	if m == 1 || n == 1 {
		return nil
	}
	if k == 2 {
		// .....
		// ###..
		// ####.
		ans[0] = play(0)
		ans[1] = play(n - 2)
		for i := 2; i < m; i++ {
			ans[i] = play(n - 1)
		}
		return ans
	}
	// k == 3
	if m >= k {
		// .....
		// ###..
		// ###..
		// ####.
		ans[0] = play(0)
		for i := 1; i < k; i++ {
			ans[i] = play(n - 2)
		}
		for i := k; i < m; i++ {
			ans[i] = play(n - 1)
		}
		return ans
	} else if n >= k {
		// ......
		// ###...
		// #####.
		ans[0] = play(0)
		ans[1] = play(n - k)
		for i := 2; i < m; i++ {
			ans[i] = play(n - 1)
		}
		return ans
	}
	if k == 4 && m >= 3 && n >= 3 {
		for i := range n {
			buf[i] = '#'
			if i < 2 {
				buf[i] = '.'
			}
		}
		ans[0] = string(buf)
		for i := range n {
			buf[i] = '#'
			if i < 3 {
				buf[i] = '.'
			}
		}
		ans[1] = string(buf)
		for i := range n {
			buf[i] = '#'
			if i >= 1 {
				buf[i] = '.'
			}
		}
		ans[2] = string(buf)
		for i := 3; i < m; i++ {
			ans[i] = play(n - 1)
		}
		return ans
	}

	return nil
}
