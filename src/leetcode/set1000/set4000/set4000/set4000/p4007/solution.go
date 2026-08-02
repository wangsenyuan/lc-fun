package p4007

func maximumWidth(planks []int) int {
	freq := make(map[int]int)
	for _, x := range planks {
		freq[x]++
	}

	width := make(map[int]int)
	values := make([]int, 0, len(freq))
	for x, count := range freq {
		values = append(values, x)
		width[x] += count
		width[2*x] += count / 2
	}

	for i, x := range values {
		for _, y := range values[i+1:] {
			width[x+y] += min(freq[x], freq[y])
		}
	}

	best := 0
	for _, count := range width {
		best = max(best, count)
	}
	return best
}
