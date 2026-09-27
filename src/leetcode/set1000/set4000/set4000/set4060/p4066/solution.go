package p4066

func maxEqualAdjacentPairs(nums []int) int {
	n := len(nums)

	var ans int
	type pair struct {
		first  int
		second int
	}
	freq := make(map[pair]int)

	for i := range n - 1 {
		x, y := nums[i], nums[i+1]
		if x > y {
			x, y = y, x
		}
		if x == y {
			ans++
		} else {
			freq[pair{x, y}]++
		}
	}
	var add int
	for _, v := range freq {
		add = max(add, v)
	}
	return ans + add
}
