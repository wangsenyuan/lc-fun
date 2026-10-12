package p4078

func resilientSubarray(nums []int, k int) int {
	var best int

	n := len(nums)
	for i := 0; i < n; {
		j := i
		w := nums[j] % k
		for i < n && nums[i]%k == w {
			i++
		}
		// x * w % k = 0
		if w == 0 {
			best = max(best, i-j)
		} else {
			x := k / gcd(w, k)
			y := (i - j - 1) / x * x
			best = max(best, y+1)
		}
	}
	return best
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
