package p4044

func countGoodRotations(nums []int) int {
	n := len(nums)
	var sum int
	for _, v := range nums {
		sum += v
	}

	var res int
	var sum1 int
	for i := range n / 2 {
		sum1 += nums[i]
	}
	for i := range n {
		if sum1*2 > sum {
			res++
		}
		sum1 -= nums[i]
		j := (i + n/2) % n
		sum1 += nums[j]
	}

	return res
}
