package p4123

func minCost(nums []int) int64 {
	n := len(nums)
	if n == 2 {
		return int64(nums[0]) + int64(nums[1])
	}
	return int64(nums[0]) + int64(nums[n-1]) + harvest(nums[1:n-1])
}

func harvest(nums []int) int64 {
	var f int64
	var best int64
	has := false
	for i, v := range nums {
		next := f
		if i >= 1 && has {
			next = min(next, best+int64(v))
		}
		cand := f + int64(v)
		if !has || cand < best {
			best = cand
			has = true
		}
		f = next
	}
	return f
}
