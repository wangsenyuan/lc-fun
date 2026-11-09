package p4102

func countValidSubstrings(s string) int64 {
	n := len(s)
	var ans int
	var l int
	for r := range n {
		if r >= 2 && s[r-2:r+1] == "abc" {
			l = r - 1
		}
		ans += r - l + 1
	}
	return int64(ans)
}
