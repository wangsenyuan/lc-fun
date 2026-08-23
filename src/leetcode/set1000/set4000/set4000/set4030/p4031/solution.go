package p4031

import "slices"

func findDisappearedNumbers(nums []int, lower int, upper int) [][]int {
	slices.Sort(nums)

	var res [][]int

	for _, v := range nums {
		if v < lower {
			continue
		}
		if v > upper {
			break
		}
		if lower+1 <= v {
			res = append(res, []int{lower, v - 1})
		}
		lower = v + 1
	}
	if lower <= upper {
		res = append(res, []int{lower, upper})
	}
	return res
}
