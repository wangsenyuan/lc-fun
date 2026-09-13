package p4051

import (
	"slices"
	"sort"
)

func distantSubarrays(nums []int, goal int, k int) int64 {
	if k == 0 {
		n := len(nums)
		return int64(n * (n + 1) / 2)
	}
	var sum int
	var arr []int
	arr = append(arr, 0)
	for _, v := range nums {
		sum += v
		arr = append(arr, sum)
	}

	slices.Sort(arr)
	arr = slices.Compact(arr)
	m := len(arr)
	bit := make(BIT, m+3)
	bit.update(sort.SearchInts(arr, 0), 1)
	sum = 0
	var res int
	for _, v := range nums {
		sum += v
		// let diff = sum - s1
		// diff - goal >= k or goal - diff <= k
		// sum - s1 - goal >= k => s1 <= sum - goal - k and sum - s1 >= goal
		// => s1 <= min(sum - goal, sum - goal - k) => s1 = sum - goal - k  (k >= 0)
		l := sort.SearchInts(arr, sum-goal-k)
		if l < m && arr[l] > sum-goal-k {
			l--
		}
		res += bit.get(l)
		// goal - (sum - s1) >= k, goal - k >= sum - s1 => s1 >= sum + k - goal
		r := sort.SearchInts(arr, sum+k-goal)
		res += bit.getRange(r, m)

		bit.update(sort.SearchInts(arr, sum), 1)
	}

	return int64(res)
}

type BIT []int

func (bit BIT) update(p int, v int) {
	p++
	for p < len(bit) {
		bit[p] += v
		p += p & -p
	}
}

func (bit BIT) get(p int) int {
	var res int
	p++
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}

func (bit BIT) getRange(l int, r int) int {
	return bit.get(r) - bit.get(l-1)
}
