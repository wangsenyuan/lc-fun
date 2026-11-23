package p4116

func assignUmbrella(rain []int) []string {
	cnt := make([]int, 2)
	ans := make([]string, len(rain))

	for i, v := range rain {
		if v == 0 {
			ans[i] = "None"
			continue
		}
		if i+1 < len(rain) && rain[i] == rain[i+1] {
			// alice & bob
			if cnt[0] <= cnt[1] {
				ans[i] = "Alice"
				cnt[0]++
			} else {
				ans[i] = "Bob"
				cnt[1]++
			}
		} else {
			ans[i] = "Bob"
			cnt[1]++
		}
	}

	return ans
}
