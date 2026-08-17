package p4020

func elevatorRequests(n int, requests []int) int {
	var res int
	var prev int
	for _, v := range requests {
		res += abs(v - prev)
		prev = v
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}
