package p4089

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var ans int

	n := len(nums)
	for i := 0; i < n; {
		if nums[i] < minK || nums[i] > maxK {
			i++
			continue
		}
		j := i
		for i < n && nums[i] >= minK && nums[i] <= maxK {
			i++
		}
		c := i - j
		ans += c * (c + 1) / 2
	}
	return int64(ans)
}
