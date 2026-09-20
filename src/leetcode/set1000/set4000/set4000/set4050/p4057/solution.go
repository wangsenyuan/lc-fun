package p4057

import (
	"cmp"
	"slices"
)

func countIntersectingIntervals(intervals [][]int) int64 {
	slices.SortFunc(intervals, func(a []int, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	// a[i][0] <= a[j][0]
	var nums []int
	for _, cur := range intervals {
		nums = append(nums, cur[0], cur[1])
	}
	slices.Sort(nums)
	nums = slices.Compact(nums)

	k := len(nums)
	bit := make(BIT, k+2)

	var res int

	for i, cur := range intervals {
		j, _ := slices.BinarySearch(nums, cur[0])
		res += i - bit.pre(j-1)
		j, _ = slices.BinarySearch(nums, cur[1])
		bit.update(j, 1)
	}

	return int64(res)
}

type BIT []int

func (bit BIT) update(i int, v int) {
	for i++; i < len(bit); i += i & -i {
		bit[i] += v
	}
}

func (bit BIT) pre(i int) int {
	var res int
	for i++; i > 0; i -= i & -i {
		res += bit[i]
	}
	return res
}
