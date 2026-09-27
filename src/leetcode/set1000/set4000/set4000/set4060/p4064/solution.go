package p4064

import "sort"

func longestSubarray(nums []int, k int) int {
	first := make([]int, k)
	last := make([]int, k)
	n := len(nums)
	for i := range k {
		first[i] = n + 1
		last[i] = -1
	}

	mod := func(num int) int {
		num %= k
		if num < 0 {
			num += k
		}
		return num
	}

	pos := make([][]int, k)
	pos[0] = append(pos[0], 0)
	var sum int
	first[0] = 0
	last[0] = 0
	for i, v := range nums {
		sum = (sum + mod(v)) % k
		first[sum] = min(first[sum], i+1)
		last[sum] = i + 1
		pos[mod(v)] = append(pos[mod(v)], i+1)
	}

	var best int

	for u, l := range first {
		if l < n {
			for v, r := range last {
				if l < r {
					if u == v {
						best = max(best, r-l)
					} else {
						w := v
						if v < u {
							w += k
						}
						diff := w - u
						for range 2 {
							if diff%2 == 0 && diff/2 < k {
								i := sort.SearchInts(pos[diff/2], l+1)
								if i < len(pos[diff/2]) && pos[diff/2][i] <= r {
									best = max(best, r-l)
								}
							}
							diff += k
						}
					}
				}
			}
		}
	}

	return best
}
