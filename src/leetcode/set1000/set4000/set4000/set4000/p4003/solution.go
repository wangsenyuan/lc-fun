package p4003

import "container/heap"

func minCost(m int, n int, penalty [][]int) int64 {
	grid := make([][][]*node, m)
	var pq PQ
	for i := range m {
		grid[i] = make([][]*node, n)
		for j := range n {
			grid[i][j] = make([]*node, 2)
			for d := range 2 {
				grid[i][j][d] = &node{
					r:        i,
					c:        j,
					d:        d,
					priority: inf,
				}
				heap.Push(&pq, grid[i][j][d])
			}
		}
	}
	pq.update(grid[0][0][0], 1)

	next := [][][]int{
		{
			{0, 1, 0},
			{1, 0, 0},
			{0, -1, 1},
			{-1, 0, 1},
		},
		{
			{0, 1, 1},
			{1, 0, 1},
			{0, -1, 0},
			{-1, 0, 0},
		},
	}

	for pq.Len() > 0 {
		nd := heap.Pop(&pq).(*node)
		r, c, d := nd.r, nd.c, nd.d

		for _, mv := range next[d] {
			nr, nc := r+mv[0], c+mv[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n {
				cost := nd.priority + mv[2]*penalty[r][c] + (nr+1)*(nc+1)
				if cost < grid[nr][nc][d^1].priority {
					pq.update(grid[nr][nc][d^1], cost)
				}
			}
		}
		if grid[r][c][d^1].priority > penalty[r][c]+nd.priority {
			pq.update(grid[r][c][d^1], penalty[r][c]+nd.priority)
		}
	}

	ans := min(grid[m-1][n-1][0].priority, grid[m-1][n-1][1].priority)

	return int64(ans)
}

const inf = 1 << 60

type node struct {
	r        int
	c        int
	d        int
	priority int
	index    int
}

type PQ []*node

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x any) {
	n := len(*pq)
	node := x.(*node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PQ) Pop() any {
	n := len(*pq)
	node := (*pq)[n-1]
	node.index = -1
	*pq = (*pq)[:n-1]
	return node
}

func (pq *PQ) update(node *node, priority int) {
	node.priority = priority
	heap.Fix(pq, node.index)
}
