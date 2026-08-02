package p4009

func minMaxWaitingTime(demand []int, fuel []int) int {
	n := len(demand)
	if n == 0 {
		return -1
	}

	pref := make([]int, n+1)
	for i, x := range demand {
		pref[i+1] = pref[i] + x
	}

	k := maxServed(demand, fuel, pref)
	if k == 0 {
		return -1
	}
	if k == 1 {
		return 0
	}

	addState := func(dp map[state]bool, x int, limit int, used0 int, used1 int, pump int, pumpBusy int, otherBusy int) {
		if pump == 0 {
			if used0+x > fuel[0] {
				return
			}
			used0 += x
		} else {
			if used1+x > fuel[1] {
				return
			}
		}
		if pumpBusy > limit {
			return
		}
		dp[state{used0: used0, last: pump, otherBusy: max(0, otherBusy-pumpBusy)}] = true
	}

	check := func(limit int) bool {
		dp := make(map[state]bool)
		if demand[0] <= fuel[0] {
			dp[state{used0: demand[0]}] = true
		}
		if demand[0] <= fuel[1] {
			dp[state{last: 1}] = true
		}

		for i := 1; i < k; i++ {
			ndp := make(map[state]bool)
			for cur := range dp {
				used1 := pref[i] - cur.used0
				lastBusy := demand[i-1]
				if cur.last == 0 {
					addState(ndp, demand[i], limit, cur.used0, used1, 0, lastBusy, cur.otherBusy)
					addState(ndp, demand[i], limit, cur.used0, used1, 1, cur.otherBusy, lastBusy)
				} else {
					addState(ndp, demand[i], limit, cur.used0, used1, 0, cur.otherBusy, lastBusy)
					addState(ndp, demand[i], limit, cur.used0, used1, 1, lastBusy, cur.otherBusy)
				}
			}
			if len(ndp) == 0 {
				return false
			}
			dp = ndp
		}
		return true
	}

	l, r := 0, pref[k]
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func maxServed(demand []int, fuel []int, pref []int) int {
	dp := make([]bool, fuel[0]+1)
	dp[0] = true
	for i, v := range demand {
		ndp := make([]bool, fuel[0]+1)
		for u0 := range fuel[0] + 1 {
			u1 := pref[i] - u0
			if u0+v <= fuel[0] && dp[u0] {
				ndp[u0+v] = true
			}
			if u1+v <= fuel[1] && dp[u0] {
				ndp[u0] = true
			}
		}
		dp = ndp
		ok := false
		for i := range dp {
			if dp[i] {
				ok = true
				break
			}
		}
		if !ok {
			return i
		}
	}
	return len(demand)
}

type state struct {
	used0     int
	last      int
	otherBusy int
}
