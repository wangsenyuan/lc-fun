package p3987

const mod = 1000000007

func minimumCost(nums []int, k int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}

	cnt := (sum + k - 1) / k
	cnt--
	cnt %= mod

	return cnt * (1 + cnt) / 2 % mod
}
