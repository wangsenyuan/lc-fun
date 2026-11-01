package p4099

func minOperations(nums []int, locked []int, k int) int {
	n := len(nums)
	flag := make([]int, n)
	maxValue := -1
	for i := 0; i < n; i++ {
		if maxValue > nums[i] {
			flag[i] = 1
		}
		maxValue = max(maxValue, nums[i])
	}
	minValue := 1 << 30
	for i := n - 1; i >= 0; i-- {
		if minValue < nums[i] {
			flag[i] = 1
		}
		minValue = min(minValue, nums[i])
	}
	var res int
	for i := range n {
		if flag[i] == 1 && locked[i] == 1 {
			k--
			res++
			if k < 0 {
				return -1
			}
		}
	}
	res += countInversions(nums)
	return res
}

func countInversions(nums []int) int {
	n := len(nums)
	f := make(fenwick, n+3)
	var res int
	for _, v := range nums {
		res += f.query(v+1, n)
		f.update(v, 1)
	}

	return res
}

type fenwick []int

func (f fenwick) update(i int, v int) {
	i++
	for i < len(f) {
		f[i] += v
		i += i & -i
	}
}

func (f fenwick) pre(i int) int {
	i++
	var res int
	for i > 0 {
		res += f[i]
		i -= i & -i
	}
	return res
}

func (f fenwick) query(l int, r int) int {
	return f.pre(r) - f.pre(l-1)
}
