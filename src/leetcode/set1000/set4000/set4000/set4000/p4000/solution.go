package p4000

func largestInteger(n int, s int) int {
	// digits sum = s, len(digits) = n
	if n*9 < s {
		return -1
	}

	if s == 0 {
		return 0
	}

	var res int

	for i := range n {
		end := 0
		if i == 0 {
			end = 1
		}
		w := -1
		for d := 9; d >= end; d-- {
			if s-d >= 0 && s-d <= (n-1)*9 {
				w = d
				break
			}
		}
		if w < 0 {
			return -1
		}
		s -= w
		res = res*10 + w
	}

	return res
}
