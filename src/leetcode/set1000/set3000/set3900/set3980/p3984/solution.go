package p3984

import "slices"

func divisibleGame(nums []int) int {
	mx := slices.Max(nums)
	lpf := make([]int, mx+1)
	var primes []int
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}

		for _, p := range primes {
			if i*p > mx {
				break
			}
			lpf[i*p] = p
			if i%p == 0 {
				break
			}
		}
	}

	mark := make([]int, mx+1)

	best := -(1 << 60)
	k := 1
	for i, v := range nums {
		if v == 1 {
			if best < -v {
				best = -v
				k = 2
			}
			continue
		}
		// 如果区间以i开始
		var fs []int
		for u := v; u > 1; u /= lpf[u] {
			x := lpf[u]
			if mark[x] == 0 {
				fs = append(fs, x)
			}
			mark[x]++
		}

		sum := make([]int, len(fs))

		for j := i; j < len(nums); j++ {
			for w, x := range fs {
				if nums[j]%x == 0 {
					sum[w] += nums[j]
				} else {
					sum[w] -= nums[j]
				}
				if sum[w] > best || sum[w] == best && x < k {
					best = sum[w]
					k = x
				}
			}
		}

		for _, x := range fs {
			mark[x] = 0
		}
	}

	best = (best + 1_000_000_007) % 1_000_000_007

	return best * k % 1_000_000_007
}
