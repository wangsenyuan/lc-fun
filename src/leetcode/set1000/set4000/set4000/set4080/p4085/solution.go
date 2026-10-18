package p4085

func findColoredRoot(edges [][]int, colors string) int {
	// R -> B -> W -> R
	n := len(colors)
	g := make([][]int, n)
	inDegree := make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		if colors[u] == colors[v] {
			return -1
		}
		if colors[u] == 'R' && colors[v] == 'W' ||
			colors[u] == 'B' && colors[v] == 'R' ||
			colors[u] == 'W' && colors[v] == 'B' {
			u, v = v, u
		}
		g[u] = append(g[u], v)
		inDegree[v]++
	}
	root := -1
	for u := range n {
		if inDegree[u] == 0 {
			if root != -1 {
				return -1
			}
			root = u
		}
	}

	var que []int
	que = append(que, root)

	for i := 0; i < len(que); i++ {
		u := que[i]
		if len(g[u]) > 2 {
			return -1
		}
		for _, v := range g[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				que = append(que, v)
			}
		}
	}

	if len(que) != n {
		return -1
	}

	return root
}
