package p4038

func countSpecialIntegers(nums []int) int {
	first := make(map[int]int)
	last := make(map[int]int)
	for i, v := range nums {
		if _, ok := first[v]; !ok {
			first[v] = i
		}
		last[v] = i
	}
	var ans int
	for i := 0; i < len(nums); {
		if first[nums[i]] != i {
			i++
			continue
		}
		j := i
		for i < len(nums) && nums[i] == nums[j] {
			i++
		}
		if last[nums[j]] == i-1 {
			ans++
		}
	}
	return ans
}
