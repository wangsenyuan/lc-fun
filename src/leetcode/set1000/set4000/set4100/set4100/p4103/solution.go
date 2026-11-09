package p4103

func canRemoveAll(types []int, colors []int, groupSize []int) bool {
	n := len(types)

	type item struct {
		typ int
		cnt int
		col int
	}

	var stack []item

	for i := range n {
		if len(stack) > 0 && stack[len(stack)-1].typ == types[i] && stack[len(stack)-1].col == colors[i] {
			stack[len(stack)-1].cnt++
		} else {
			stack = append(stack, item{typ: types[i], cnt: 1, col: colors[i]})
		}
		if len(stack) > 0 {
			last := stack[len(stack)-1]
			if last.cnt == groupSize[last.typ] {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
