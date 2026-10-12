package p4079

func threeFibonacciSum(n int) bool {
	if n == 1 {
		return false
	}
	a, b := 0, 1

	for {
		c := a + b
		if a+b+c > n {
			break
		}
		if a+b+c == n {
			return true
		}
		a, b = b, c
	}

	return false
}
