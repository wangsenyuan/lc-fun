package p4055

import (
	"math/bits"
	"slices"
	"sort"
)

func shadowPairs(nums []int) int {
	arr := slices.Clone(nums)
	slices.Sort(arr)
	arr = slices.Compact(arr)

	for i, v := range nums {
		nums[i] = sort.SearchInts(arr, v)
	}
	m := len(arr)
	mp := bits.Len(uint(m)) - 1
	var res int

	for p := mp; p >= 0; p-- {
		tot := 1 << (mp - p)
		small := make([][]int, tot)
		big := make([][]int, tot)
		for i, v := range nums {
			g := v >> (p + 1)
			flag := (v >> p) & 1
			if flag > 0 {
				for len(big[g]) > 0 && nums[last(big[g])] >= v {
					big[g] = big[g][:len(big[g])-1]
				}
				lim := -1
				if len(big[g]) > 0 {
					lim = last(big[g])
				}
				big[g] = append(big[g], i)
				j := sort.SearchInts(small[g], lim)
				res += len(small[g]) - j
			} else {
				for len(small[g]) > 0 && nums[last(small[g])] < v {
					small[g] = small[g][:len(small[g])-1]
				}
				small[g] = append(small[g], i)
			}
		}
	}

	return res
}

func last(arr []int) int {
	return arr[len(arr)-1]
}
