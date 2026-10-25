package p4092

import "math/bits"

func calculateScore(a int, b int) int {
	w := bits.OnesCount(uint(a))
	v := bits.OnesCount(uint(b))

	if w >= v {
		return w * v
	}
	return v / w
}
