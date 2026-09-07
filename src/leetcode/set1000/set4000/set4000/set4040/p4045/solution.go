package p4045

func countGroups(position []int, speed []int, distance int) int {

	n := len(position)

	var res int
	for i := n - 1; i >= 0; {
		j := i
		res++
		i--
		for i >= 0 {
			if position[i+1]-position[i] > distance && speed[i] <= speed[j] {
				break
			}
			i--
		}
	}

	return res
}
