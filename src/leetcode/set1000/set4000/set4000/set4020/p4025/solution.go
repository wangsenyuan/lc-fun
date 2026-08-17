package p4025

import "slices"

func minPenalty(period int, lights []int, arrivalTime []int) int {
	w := slices.Max(lights)
	var res int
	for _, v := range arrivalTime {
		r := v % period
		if r >= w {
			res = max(res, period-r)
		}
	}
	return res
}
