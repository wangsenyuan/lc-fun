package p4106

func findMismatchedSegments(s1 string, s2 string) [][]string {
	var ans [][]string

	n := len(s1)

	for i := 0; i < n; {
		if s1[i] == s2[i] {
			i++
			continue
		}
		j := i
		for i < n && s1[i] != s2[i] {
			i++
		}
		ans = append(ans, []string{s1[j:i], s2[j:i]})
	}
	return ans
}
