package p4052

import "slices"

func cyclicShift(n int, grid [][]int, rowShift []int, colShift []int) [][]int {
	for i, v := range rowShift {
		shift(grid[i], v)
	}
	col := make([]int, n)

	for i, v := range colShift {
		for j := range n {
			col[j] = grid[j][i]
		}
		shift(col, v)
		for j := range n {
			grid[j][i] = col[j]
		}
	}
	return grid
}

func shift(a []int, k int) {
	slices.Reverse(a[:k])
	slices.Reverse(a[k:])
	slices.Reverse(a)
}
