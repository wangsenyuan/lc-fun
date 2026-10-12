package p4075

import (
	"strconv"
)

type BusBooking struct {
	rows      []int
	take      [][]int
	total     int
	totalTime int
}

func Constructor(n int) BusBooking {
	rows := make([]int, n)
	take := make([][]int, n)
	for i := range n {
		take[i] = make([]int, 4)
	}
	return BusBooking{rows, take, 0, 0}
}

func (this *BusBooking) Toggle(seat string) {
	col := int(seat[0] - 'A')
	// 0, 3 靠窗
	row, _ := strconv.Atoi(seat[1:])
	row--
	flag := this.rows[row]
	cur := this.totalTime
	if (flag>>col)&1 == 0 {
		if col > 0 && col < 3 {
			this.totalTime++
		} else if col == 0 {
			this.totalTime++

			if flag&2 == 2 {
				this.totalTime += 2
			}
		} else {
			// col == 3
			this.totalTime++
			if flag&4 == 4 {
				this.totalTime += 2
			}
		}
		this.take[row][col] = this.totalTime - cur
		this.total++
	} else {
		this.total--
		this.totalTime -= this.take[row][col]
	}

	this.rows[row] ^= 1 << col
}

func (this *BusBooking) GetTotalTime() int {
	return this.totalTime
}

func (this *BusBooking) GetMinTime() int {
	return this.total
}
