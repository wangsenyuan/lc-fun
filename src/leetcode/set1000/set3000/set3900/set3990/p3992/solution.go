package p3992

func rearrangeString(s string, x byte, y byte) string {
	cnt := make([]int, 2)
	buf := []byte(s)
	var pos int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case y:
			cnt[0]++
		case x:
			cnt[1]++
		default:
			buf[pos] = s[i]
			pos++
		}
	}
	for range cnt[0] {
		buf[pos] = y
		pos++
	}
	for range cnt[1] {
		buf[pos] = x
		pos++
	}
	return string(buf)
}
