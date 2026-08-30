package p4037

import "slices"

func maxValidSplits(nums []int) int {
	play := func(arr []int) int {
		suf := make([]int, len(arr))
		for i := len(arr) - 1; i >= 0; i-- {
			suf[i] = arr[i]
			if i+1 < len(arr) {
				suf[i] = gcd(arr[i], suf[i+1])
			}
		}
		var pref int
		var cnt int
		for i := 0; i+1 < len(arr); i++ {
			pref = gcd(pref, arr[i])
			if pref == suf[i+1] {
				cnt++
			}
		}
		return cnt
	}

	best := play(nums)
	var pref int
	for i := 0; i < len(nums); i++ {
		cur := gcd(pref, nums[i])
		if pref != cur {
			tmp := slices.Clone(nums[:i])
			tmp = append(tmp, nums[i+1:]...)
			best = max(best, play(tmp))
		}
		pref = cur
	}
	var next int
	for i := len(nums) - 1; i >= 0; i-- {
		cur := gcd(next, nums[i])
		if cur != next {
			tmp := slices.Clone(nums[:i])
			tmp = append(tmp, nums[i+1:]...)
			best = max(best, play(tmp))
		}
		next = cur
	}

	return best
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
