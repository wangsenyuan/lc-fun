package p4088

func peakValleyDifference(nums []int) int {
	var s1, s2 int
	for i, v := range nums {
		if i > 0 && i < len(nums)-1 {
			if v > nums[i-1] && v > nums[i+1] {
				s1 += v
			}
			if v < nums[i-1] && v < nums[i+1] {
				s2 += v
			}
		}
	}
	return s1 - s2
}
