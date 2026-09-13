package p4054

func shadowPairs(nums []int) int64 {
	// (i, j), nums[i] < nums[j], 且i, j中间不存在k, nums[k] < nums[i]
	type pair struct {
		first  int
		second int
	}
	var stack []pair
	stack = append(stack, pair{0, 0})
	var ans int
	var sz int
	for _, v := range nums {
		for stack[len(stack)-1].first > v {
			sz -= stack[len(stack)-1].second
			stack = stack[:len(stack)-1]
		}
		ans += sz
		if stack[len(stack)-1].first == v {
			ans -= stack[len(stack)-1].second
			stack[len(stack)-1].second++
		} else {
			stack = append(stack, pair{v, 1})
		}
		sz++
	}
	return int64(ans)
}
