package p4100

import "math/bits"

type edge struct {
	to   int
	cost int
}

func minActivationCost(edges [][]int, values []int) int {
	n := len(values)
	g := make([][]edge, n)

	for _, cur := range edges {
		u, v, w := cur[0], cur[1], cur[2]
		g[u] = append(g[u], edge{v, w})
		g[v] = append(g[v], edge{u, w})
	}

	const inf = 1 << 60

	best := inf

	var dfs func(p int, u int) [2]int
	dfs = func(p int, u int) (res [2]int) {
		res[0] = inf
		res[1] = inf
		ds := bits.OnesCount(uint(values[u]))
		res[ds&1] = 0

		for _, e := range g[u] {
			if e.to != p {
				sub := dfs(u, e.to)
				best = min(best, res[0]+sub[0]+e.cost, res[1]+sub[1]+e.cost)
				if ds&1 == 0 {
					res[0] = min(res[0], sub[0]+e.cost)
					res[1] = min(res[1], sub[1]+e.cost)
				} else {
					res[0] = min(res[0], sub[1]+e.cost)
					res[1] = min(res[1], sub[0]+e.cost)
				}
			}
		}

		return res
	}

	dfs(-1, 0)

	if best == inf {
		return -1
	}

	return best
}
