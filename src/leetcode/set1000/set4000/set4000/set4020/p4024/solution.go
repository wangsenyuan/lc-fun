package p4024

func nearestDrone(drones [][]int, target []int) int {
	res := -1
	best := -1
	for i, cur := range drones {
		d := dist(cur, target)
		if cur[2] >= d {
			if best < 0 || d < best {
				res = i
				best = d
			}
		}
	}

	return res
}

func dist(a []int, b []int) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return abs(dx) + abs(dy)
}

func abs(num int) int {
	return max(num, -num)
}
