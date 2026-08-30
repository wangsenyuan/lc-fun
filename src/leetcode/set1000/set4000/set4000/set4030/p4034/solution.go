package p4034

func minBishopMoves(source []int, target []int) int {
	if (source[0]+source[1])&1 != (target[0]+target[1])&1 {
		return -1
	}

	if source[0] == target[0] && source[1] == target[1] {
		return 0
	}

	dx := target[0] - source[0]
	dy := target[1] - source[1]

	if abs(dx) == abs(dy) {
		return 1
	}

	return 2
}

func abs(num int) int {
	return max(num, -num)
}
