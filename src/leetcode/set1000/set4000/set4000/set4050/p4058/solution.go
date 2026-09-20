package p4058

func maxValue(nums []int) int64 {
	n := len(nums)
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if i&1 == 0 {
			suf[i] = nums[i] + suf[i+1]
		} else {
			suf[i] = -nums[i] + suf[i+1]
		}
	}

	// 考虑交换长度是偶数长度, 比如 [a, b, c, d], left shift even
	// 那么其实是没有区别的. 所以只能是奇数长度 [a, b, c, d] => [b, c, d, a]
	// 那么原来需要 +的, 现在变成-
	// 长度3没有影响; 就考虑1即可
	// pref[l] + suf[r+1] + 这个区间内的
	// pref[r] - pref[l] > 0 不需要替换
	// pref[r] < pref[l] => 2 * pref[l] - pref[r] + suf[r+1]
	// 所以对于r来说, 就是找到最大的 pref[l]
	ans := suf[0]

	best := []int{-(1 << 30), 0}

	var sum int
	for i := range n {
		if i&1 == 0 {
			sum += nums[i]
		} else {
			sum -= nums[i]
		}
		ans = max(ans, 2*best[i&1]-sum+suf[i+1])
		best[i&1] = max(best[i&1], sum)
	}

	return int64(ans)
}
