package p4046

import "container/heap"

const inf = 1 << 60

func minCost(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][][][4]int, n)
	for i := range n {
		dp[i] = make([][][4]int, m)
		for j := range m {
			dp[i][j] = make([][4]int, k+1)
			for x := range k + 1 {
				for w := range 4 {
					dp[i][j][x][w] = inf
				}
			}
		}
	}
	// 0 for top, 1 for right, 2 for bot, 3 for left
	dp[0][0][0][0] = grid[0][0]
	dp[0][0][0][3] = grid[0][0]

	var dd = [][]int{
		{1, 0}, {0, -1}, {-1, 0}, {0, 1},
	}

	var pq PQ

	heap.Push(&pq, state{r: 0, c: 0, dir: 0, turn: 0, value: grid[0][0]})
	heap.Push(&pq, state{r: 0, c: 0, dir: 3, turn: 0, value: grid[0][0]})

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(state)
		if dp[cur.r][cur.c][cur.turn][cur.dir] != cur.value {
			continue
		}
		if cur.r == n-1 && cur.c == m-1 {
			return cur.value
		}
		for dir, d := range dd {
			r1, c1 := cur.r+d[0], cur.c+d[1]
			if r1 >= 0 && r1 < n && c1 >= 0 && c1 < m {
				turn1 := cur.turn
				if dir != cur.dir {
					turn1++
				}
				if turn1 <= k && dp[r1][c1][turn1][dir] > cur.value+grid[r1][c1] {
					dp[r1][c1][turn1][dir] = cur.value + grid[r1][c1]
					heap.Push(&pq, state{r: r1, c: c1, dir: dir, turn: turn1, value: dp[r1][c1][turn1][dir]})
				}
			}
		}
	}

	return -1
}

type state struct {
	r     int
	c     int
	dir   int
	turn  int
	value int
}

type PQ []state

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i int, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq PQ) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	s := x.(state)
	*pq = append(*pq, s)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}
