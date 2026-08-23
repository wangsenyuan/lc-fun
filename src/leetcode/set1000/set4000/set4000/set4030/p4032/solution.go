package p4032

import "slices"

func longestSubarray(nums []int, k int) int {
	x := slices.Max(nums)
	lpf := make([]int, x+1)
	var primes []int

	for i := 2; i <= x; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j > x {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	freq := make([]int, x+1)
	var cnt int

	add := func(num int) {
		for num > 1 {
			w := lpf[num]
			freq[w]++
			if freq[w] == 1 {
				cnt++
			}
			for num%w == 0 {
				num /= w
			}
		}
	}

	rem := func(num int) {
		for num > 1 {
			w := lpf[num]
			freq[w]--
			if freq[w] == 0 {
				cnt--
			}
			for num%w == 0 {
				num /= w
			}
		}
	}

	var best int
	for l, r := 0, 0; r < len(nums); r++ {
		add(nums[r])
		for cnt > k {
			rem(nums[l])
			l++
		}
		best = max(best, r-l+1)
	}
	return best
}
