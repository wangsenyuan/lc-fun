package p4121

func minCost(nums []int) int64 {
	n := len(nums)
	ans := 1 << 60

	for i := 1; i < n-1; {
		j := i
		for i < n-1 && nums[i] == nums[j] {
			i++
		}
		v := nums[j]
		pref := j
		if j == 1 && nums[0] == v {
			pref--
		}
		suf := n - i
		if i == n-1 && v == nums[n-1] {
			suf--
		}
		ans = min(ans, v*(pref+suf))
	}

	return int64(ans)
}
