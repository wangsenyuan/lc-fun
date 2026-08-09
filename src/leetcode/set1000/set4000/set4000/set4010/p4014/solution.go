package p4014

import "slices"

func minPrice(prices []int, discounts []int) float64 {
	slices.Sort(prices)
	slices.Reverse(prices)
	slices.Sort(discounts)
	slices.Reverse(discounts)
	var res float64
	for i := 0; i < len(prices); i++ {
		if i < len(discounts) {
			res += float64(prices[i]) * float64(100-discounts[i]) / 100
		} else {
			res += float64(prices[i])
		}
	}
	return res
}
