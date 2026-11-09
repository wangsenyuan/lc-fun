package p4105

func repeatedStringMatch(a string, b string, x int, y int) bool {
	n := len(a)
	m := len(b)
	if n*x < m*y {
		return false
	}
	w := min(lcm(n, m), m*y)

	for i := 0; i < n && i+m*y <= n*x; i++ {
		ok := true

		for j := range w {
			if a[(i+j)%n] != b[j%m] {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}

	return false
}

func gcd(a int, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a int, b int) int {
	g := gcd(a, b)
	return b / g * a
}
