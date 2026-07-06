package p3982

import "slices"

func maxDigitRange(nums []int) int {
	var maxDiff int
	for _, num := range nums {
		maxDiff = max(maxDiff, getDigitDiff(num))
	}

	var res int
	for _, num := range nums {
		if getDigitDiff(num) == maxDiff {
			res += num
		}
	}
	return res
}

func getDigitDiff(num int) int {
	var ds []int
	for num > 0 {
		ds = append(ds, num%10)
		num /= 10
	}
	x := slices.Min(ds)
	y := slices.Max(ds)
	return y - x
}
