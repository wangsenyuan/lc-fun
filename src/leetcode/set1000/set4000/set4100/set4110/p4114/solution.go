package p4114

func countConsistent(nums []int) int64 {
	n := len(nums)
	tr := build(n)
	R := make([]int, n)
	var stack []int
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			stack = stack[:len(stack)-1]
		}
		tr.set(i, nums[i])
		tmp := tr.find(nums[i])
		R[i] = tmp[0]
		if tmp[1]&nums[i] == nums[i] {
			R[i]++
		}
		if len(stack) > 0 {
			R[i] = min(R[i], stack[len(stack)-1])
		}
		stack = append(stack, i)
	}
	var ans int
	tr = build(n)

	for i := 0; i < n; i++ {
		tr.set(n-i-1, nums[i])
		tmp := tr.find(nums[i])
		l := tmp[0]
		l = n - l - 1
		if tmp[1]&nums[i] == nums[i] {
			l--
		}
		ans += (R[i] - i) * (i - l)
	}
	return int64(ans)
}

type tree []int

func build(n int) tree {
	tr := make(tree, 4*n)
	for i := range tr {
		tr[i] = -1
	}
	return tr
}

func (tr tree) set(pos int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			tr[i] = v
			return
		}
		mid := (l + r) >> 1
		if pos < mid {
			f(i*2, l, mid)
		} else {
			f(i*2+1, mid, r)
		}
		tr[i] = tr[i*2] & tr[i*2+1]
	}
	f(1, 0, len(tr)/4)
}

func (tr tree) find(val int) []int {
	var f func(i int, l int, r int) []int
	f = func(i int, l int, r int) []int {
		if l+1 == r {
			return []int{l, tr[i]}
		}
		mid := (l + r) >> 1
		if tr[i*2]&val != val {
			return f(i*2, l, mid)
		}
		return f(i*2+1, mid, r)
	}
	return f(1, 0, len(tr)/4)
}
