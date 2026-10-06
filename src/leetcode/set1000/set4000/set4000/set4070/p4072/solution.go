package p4072

const inf = 1 << 60

func maxAlternatingSum1(nums []int) int64 {

	n := len(nums)
	mem := make([][][]int, n)
	for i := range n {
		mem[i] = make([][]int, 2)
		for j := range 2 {
			mem[i][j] = make([]int, 2)
			for k := range 2 {
				mem[i][j][k] = -inf
			}
		}
	}

	var f func(i int, j int, r int) int

	f = func(i int, j int, r int) (res int) {
		if i == n {
			return -inf
		}
		if mem[i][j][r] > -inf {
			return mem[i][j][r]
		}

		defer func() {
			mem[i][j][r] = res
		}()
		x := nums[i]
		if r == 1 {
			x = -x
		}
		if j == 0 {
			res = max(f(i+1, 0, r^1), 0) + x
		} else {
			// 删除i+1后的, 或者删除i
			res = max(f(i+1, 1, r^1)+x, f(i+1, 0, r))
		}
		return res
	}

	res := -inf
	for i := range n {
		res = max(res, f(i, 0, 0), f(i, 1, 0))
	}

	return int64(res)
}

func maxAlternatingSum(nums []int) int64 {
	// n := len(nums)
	f00, f01, f10, f11 := -inf, -inf, -inf, -inf
	ans := -inf
	for i := len(nums) - 1; i >= 0; i-- {
		x := nums[i]
		f10, f11 = max(f11+x, f00), max(f10-x, f01)
		f00, f01 = max(f01, 0)+x, max(f00, 0)-x
		ans = max(ans, f00, f10)
	}
	return int64(ans)
}
