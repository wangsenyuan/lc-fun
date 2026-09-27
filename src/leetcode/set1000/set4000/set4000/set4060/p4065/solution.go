package p4065

import "slices"

func rearrangeArray(nums []int) []int {
	slices.Sort(nums)
	var arr [][]int
	for i := 0; i < len(nums); {
		j := i
		for i < len(nums) && nums[i] == nums[j] {
			i++
		}
		arr = append(arr, nums[j:i])
	}
	var ans []int
	for len(arr) > 0 {
		var next [][]int
		for _, cur := range arr {
			ans = append(ans, cur[0])
			cur = cur[1:]
			if len(cur) > 0 {
				next = append(next, cur)
			}
		}
		arr = next
	}

	return ans
}
