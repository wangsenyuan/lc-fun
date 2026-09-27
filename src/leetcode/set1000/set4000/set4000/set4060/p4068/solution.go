package p4068

import (
	"slices"
	"sort"
)

const inf = 1 << 60

func maxEarnings(meetings [][]int) int64 {
	var nums []int
	for _, cur := range meetings {
		nums = append(nums, cur[0], cur[1])
	}
	slices.Sort(nums)
	nums = slices.Compact(nums)
	m := len(nums)
	openAt := make([][]int, m)
	for i, cur := range meetings {
		j := sort.SearchInts(nums, cur[0])
		openAt[j] = append(openAt[j], i)
	}
	dp := make(segtree, 2*m)

	for i := range m {
		dp.update(i, -inf)
	}

	var best int
	for i := range m {
		s := nums[i]
		w := max(dp.get(0, i+1)+s, 0)

		for _, j := range openAt[i] {
			e := meetings[j][1]
			r := meetings[j][2]
			best = max(best, w+r)
			k := sort.SearchInts(nums, e)
			v := dp.get(k, k+1)
			if v < w+r-e {
				dp.update(k, w+r-e)
			}
		}
	}

	return int64(best)
}

type segtree []int

func (tr segtree) update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p] = v
	for p > 1 {
		tr[p>>1] = max(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr segtree) get(l int, r int) int {
	n := len(tr) / 2
	l += n
	r += n
	res := -inf
	for l < r {
		if l&1 == 1 {
			res = max(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
