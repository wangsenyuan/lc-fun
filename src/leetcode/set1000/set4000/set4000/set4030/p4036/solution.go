package p4036

func largestString(nums []int) []string {
	var res []string

	play := func(num int) string {
		var buf []byte
		w := 25
		for num > 0 {
			for 1<<w > num {
				w--
			}
			// 1 << w <= num
			buf = append(buf, byte(w+'a'))
			num -= 1 << w
		}
		return string(buf)
	}

	for _, cur := range nums {
		res = append(res, play(cur))
	}
	return res
}
