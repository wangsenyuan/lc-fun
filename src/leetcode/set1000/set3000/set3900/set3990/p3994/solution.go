package p3994

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func minAdjacentSwaps(nums []int, a int, b int) int {
	cnt := make([]int, 3)
	var res int
	for _, v := range nums {
		if v > b {
			cnt[2]++
		} else if v >= a {
			res = add(res, cnt[2])
			cnt[1]++
		} else {
			res = add(res, cnt[1])
			res = add(res, cnt[2])
			cnt[0]++
		}
	}

	return res
}
