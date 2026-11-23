package p4117

import "slices"

func minRange(a int, b int, c int) int {
	arr := []int{a, b, c}
	slices.Sort(arr)
	res := arr[2] - arr[0]

	res = min(res, arr[1])

	return res
}
