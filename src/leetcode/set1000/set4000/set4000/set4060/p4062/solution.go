package p4062

func canTransform(source []int, target []int) bool {
	return sum(source) == sum(target)
}

func sum(arr []int) int {
	var res int
	for _, v := range arr {
		res += v
	}
	return res
}
