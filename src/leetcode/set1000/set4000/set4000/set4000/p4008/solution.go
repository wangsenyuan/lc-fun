package p4008

import "sort"

func minInitialStrength(monsters []int, boosts [][]int) int64 {
	n := len(monsters)
	bonus := make([]int, n+1)
	for _, cur := range boosts {
		l, r, v := cur[0], cur[1], cur[2]
		bonus[l] += v
		bonus[r+1] -= v
	}
	for i := 1; i < n; i++ {
		bonus[i] += bonus[i-1]
	}

	check := func(w int) bool {
		for i := range n {
			if w+bonus[i] < monsters[i] {
				return false
			}
			w = max(0, w-monsters[i])
		}
		return true
	}

	var r int
	for _, v := range monsters {
		r += v
	}

	ans := sort.Search(r+1, check)
	return int64(ans)
}
