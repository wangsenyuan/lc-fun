package p4013

import "slices"

func countRatioSubarrays(nums []int, a int, b int) int64 {
	// x/y <= a / b => b * x - a * y <= 0
	n := len(nums)
	sum := make([]int, n+1)
	for i, v := range nums {
		sum[i+1] = sum[i]
		if v%2 == 1 {
			sum[i+1] += a
		} else {
			sum[i+1] -= b
		}
	}

	var f func(arr []int) int
	f = func(arr []int) int {
		if len(arr) <= 1 {
			return 0
		}
		mid := len(arr) / 2
		L := slices.Clone(arr[:mid])
		R := slices.Clone(arr[mid:])
		res := f(L) + f(R)
		var j, k int
		n := len(arr)
		for i := range n {
			if k == len(R) || j < len(L) && L[j] <= R[k] {
				arr[i] = L[j]
				j++
			} else {
				res += j
				arr[i] = R[k]
				k++
			}
		}

		return res
	}

	res := f(sum)

	return int64(res)
}
