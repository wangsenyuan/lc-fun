package p4084

func lastVisitedCell(grid [][]byte) []int {
	n := len(grid)
	m := len(grid[0])
	marked := make([][]bool, n)
	for i := range n {
		marked[i] = make([]bool, m)
	}

	var r, c int
	for {
		marked[r][c] = true
		nr, nc := r, c
		switch grid[r][c] {
		case 'U':
			nr--
		case 'D':
			nr++
		case 'R':
			nc++
		default:
			nc--
		}
		if nr < 0 || nr == n || nc < 0 || nc == m || marked[nr][nc] {
			return []int{r, c}
		}
		r, c = nr, nc
	}
}
