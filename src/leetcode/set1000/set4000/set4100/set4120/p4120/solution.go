package p4120

func countDominantElements(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	var ans int
	for i, v := range nums {
		var sum int
		var cnt int
		if i > 0 {
			sum += nums[i-1]
			cnt++
		}
		if i+1 < len(nums) {
			sum += nums[i+1]
			cnt++
		}
		if cnt*v >= sum {
			ans++
		}
	}
	return ans
}
