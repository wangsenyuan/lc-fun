package p4109

import (
	"math/rand/v2"
	"slices"
)

func countSuffix(nums []int, target []int) []int {

	mp := make([]int, 1e5+1)
	for i := range mp {
		mp[i] = rand.IntN(1 << 62)
	}

	var sum int
	for _, v := range nums {
		sum += mp[v]
	}

	arr := slices.Clone(nums)
	arr = append(arr, 0)
	arr = append(arr, target...)

	z := zFunction(arr)

	m := len(target)

	last := make([]int, m)

	for i := range m {
		last[i] = i
		if i > 0 && target[i-1] >= target[i] {
			last[i] = last[i-1]
		}
	}

	n := len(nums)

	z = z[n+1:]

	var ans []int

	var sum2 int
	for k := range n - 1 {
		sum2 += mp[target[k]]
	}

	play := func(k int) int {
		hi := min(z[k], n-1)
		lo := max(0, last[k+n-1]-k)
		return max(0, hi-lo+1)
	}

	for k := 0; k+n <= len(target); k++ {
		sum2 += mp[target[k+n-1]]

		if sum2 == sum {
			ans = append(ans, play(k))
		} else {
			ans = append(ans, 0)
		}

		sum2 -= mp[target[k]]
	}

	return ans
}

func zFunction(arr []int) []int {
	n := len(arr)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && arr[z[i]] == arr[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}
