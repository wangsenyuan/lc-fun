package p3980

const inf = 1 << 60

func minOperations(s1 string, s2 string) int {
	n := len(s1)
	dp := [2]int{inf, inf}
	dp[int(s1[0]-'0')] = 0

	for i := 0; i < n; i++ {
		target := int(s2[i] - '0')

		if i == n-1 {
			ans := inf
			for cur, cost := range dp {
				if cost >= inf {
					continue
				}
				if target == 0 {
					if cur == 0 {
						ans = min(ans, cost)
					}
				} else {
					ans = min(ans, cost+1-cur)
				}
			}
			if ans >= inf {
				return -1
			}
			return ans
		}

		nextOrig := int(s1[i+1] - '0')
		ndp := [2]int{inf, inf}

		for cur, cost := range dp {
			if cost >= inf {
				continue
			}

			if target == 0 {
				if cur == 0 {
					ndp[nextOrig] = min(ndp[nextOrig], cost)
				}
			} else {
				ndp[nextOrig] = min(ndp[nextOrig], cost+1-cur)
			}

			pairCost := 1 + 1 - cur + 1 - nextOrig
			if target == 1 {
				pairCost++
			}
			ndp[0] = min(ndp[0], cost+pairCost)
		}

		dp = ndp
	}

	return -1
}
