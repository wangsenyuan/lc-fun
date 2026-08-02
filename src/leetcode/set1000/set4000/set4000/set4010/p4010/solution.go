package p4010

func maxPairStrength(nums []int) int64 {
	var best int

	for i, v := range nums {
		for j := range i {
			g := gcd(nums[j], v)
			g *= g
			tmp := nums[j] * v / g
			best = max(best, tmp)
		}
	}
	return int64(best)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
