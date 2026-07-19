package p3999

func minimumGroups(words []string) int {
	cnt := make(map[string]int)

	for _, w := range words {
		w0, w1 := splitByPos(w)
		w0 = play(w0)
		w1 = play(w1)
		cnt[w0+w1]++
	}

	return len(cnt)
}

func splitByPos(s string) (string, string) {
	var buf0 []byte
	var buf1 []byte
	for i := range len(s) {
		if i&1 == 0 {
			buf0 = append(buf0, s[i])
		} else {
			buf1 = append(buf1, s[i])
		}
	}
	return string(buf0), string(buf1)
}

func play(s string) string {
	n := len(s)
	s += s
	// 注：如果要返回一个和原串不同的字符串，初始化 i=1, j=2
	i := 0
	for j := 1; j < n; {
		// 暴力比较：是 i 开头的字典序小，还是 j 开头的字典序小？
		// 相同就继续往后比，至多循环 n 次（如果循环 n 次，说明所有字母都相同，不用再比了）
		k := 0
		for k < n && s[i+k] == s[j+k] {
			k++
		}
		if k >= n {
			break
		}

		if s[i+k] < s[j+k] { // 注：如果求字典序最大，改成 >
			// 比如从 i 开始是 "aaab"，从 j 开始是 "aaac"
			// 从 i 开始比从 j 开始更小（排除 j）
			// 此外：
			// 从 i+1 开始比从 j+1 开始更小，所以从 j+1 开始不可能是答案，排除
			// 从 i+2 开始比从 j+2 开始更小，所以从 j+2 开始不可能是答案，排除
			// ……
			// 从 i+k 开始比从 j+k 开始更小，所以从 j+k 开始不可能是答案，排除
			// 所以下一个「可能是答案」的开始位置是 j+k+1
			j += k + 1
		} else {
			// 从 j 开始比从 i 开始更小，更新 i=j（也意味着我们排除了 i）
			// 此外：
			// 从 j+1 开始比从 i+1 开始更小，所以从 i+1 开始不可能是答案，排除
			// 从 j+2 开始比从 i+2 开始更小，所以从 i+2 开始不可能是答案，排除
			// ……
			// 从 j+k 开始比从 i+k 开始更小，所以从 i+k 开始不可能是答案，排除
			// 所以把 j 跳到 i+k+1，不过这可能比 j+1 小，所以与 j+1 取 max
			// 综上所述，下一个「可能是答案」的开始位置是 max(j+1, i+k+1)
			i, j = j, max(j, i+k)+1
		}
		// 每次要么排除 k+1 个与 i 相关的位置（这样的位置至多 n 个），要么排除 k+1 个与 j 相关的位置（这样的位置至多 n 个）
		// 所以上面关于 k 的循环，∑k <= 2n，所以二重循环的总循环次数是 O(n) 的
	}
	return s[i : i+n]

}
