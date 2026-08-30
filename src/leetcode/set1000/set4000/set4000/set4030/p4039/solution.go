package p4039

import "slices"

const mod = 1e9 + 7

func add(a, b int) int {
	return (a + b) % mod
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func sumDecoded(nums []int64) int {
	var res int
	for _, x := range nums {
		num := int(x)
		w := num % 10
		d := num / 10
		var arr []int
		for i := d; i > 0; i /= 10 {
			arr = append(arr, i%10)
		}
		slices.Reverse(arr)
		var x int
		for range w {
			x = x*10 + arr[0]
			arr = arr[1:]
		}
		var y int
		for _, v := range arr {
			y = y*10 + v
		}
		// pow(x, y, mod)
		res = add(res, pow(x, y))
	}
	return res
}
