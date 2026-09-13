package p4048

import (
	"cmp"
	"slices"
)

func countSpecialIntegers(nums []int) int {
	type pair struct {
		first  int
		second int
	}
	n := len(nums)
	arr := make([]pair, n)
	for i, v := range nums {
		arr[i] = pair{v, i}
	}
	slices.SortFunc(arr, func(a pair, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})

	var res int

	for i := 0; i < n; {
		j := i
		var diff int
		ok := true
		for i < n && arr[i].first == arr[j].first {
			if i > j {
				if diff > 0 && arr[i].second-arr[i-1].second != diff {
					ok = false
				}
				diff = arr[i].second - arr[i-1].second
			}
			i++
		}
		if i-j == 3 && ok {
			res++
		}
	}
	return res
}
