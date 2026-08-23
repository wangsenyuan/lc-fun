package p4033

import (
	"math/rand/v2"
	"slices"
)

func validSubarrays(nums []int, k int, queries [][]int) []bool {
	arr := slices.Clone(nums)
	slices.Sort(arr)
	arr = slices.Compact(arr)

	mp := make(map[int]uint64)

	for _, x := range arr {
		mp[x] = rand.Uint64()
	}
	n := len(nums)

	sum := make([]uint64, n+1)
	for i, v := range nums {
		w := mp[v]
		sum[i+1] = sum[i] ^ w
	}

	todo := make([][]int, n)
	for i, cur := range queries {
		r := cur[1]
		todo[r] = append(todo[r], i)
	}

	f1 := make(map[uint64]int)
	f2 := make(map[uint64]int)
	cnt := make(map[uint64]int)

	ans := make([]bool, len(queries))

	for l1, l2, r := 0, 0, 0; r < len(nums); r++ {
		f1[mp[nums[r]]]++
		f2[mp[nums[r]]]++

		for len(f1) >= k {
			w := mp[nums[l1]]
			f1[w]--
			if f1[w] == 0 {
				delete(f1, w)
			}
			if len(f1) < k {
				f1[w]++
				break
			}
			cnt[sum[l1]]++
			l1++
		}
		for len(f2) > k {
			w := mp[nums[l2]]
			f2[w]--
			if f2[w] == 0 {
				delete(f2, w)
			}
			if len(f2) == k {
				f2[w]++
				break
			}
			cnt[sum[l2]]--
			l2++
		}

		for _, i := range todo[r] {
			l := queries[i][0]
			if sum[r+1] == sum[l] && l2 <= l && l < l1 {
				ans[i] = true
			}
		}
	}
	return ans
}
