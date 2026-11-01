package p4097

import (
	"fmt"
	"strconv"
)

func minMergedInteger(nums []int) int {
	res := 1 << 60
	for i := 0; i < len(nums); i += 2 {
		tmp := concate(nums[i], nums[i+1])
		res = min(res, tmp)
	}
	return res
}

func concate(a int, b int) int {
	s := fmt.Sprintf("%d%d", a, b)
	res, _ := strconv.Atoi(s)
	return res
}
