package p4093

import "slices"

func maximizeSum(nums []int, k int) int64 {
	var sum int
	var arr []int
	for _, num := range nums {
		sum += num
		w := transform(num)
		arr = append(arr, w-num)
	}
	slices.Sort(arr)
	slices.Reverse(arr)
	for i := range k {
		sum += arr[i]
	}
	return int64(sum)
}

func transform(num int) int {
	var arr []int
	for n := num; n > 0; n /= 10 {
		arr = append(arr, n%10)
	}
	slices.Sort(arr)
	var res int
	for _, a := range slices.Backward(arr) {
		res = res*10 + a
	}
	return res
}
