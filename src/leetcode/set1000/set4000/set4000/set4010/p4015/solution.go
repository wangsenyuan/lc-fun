package p4015

func weightedSum(parent []int, nums []int) int64 {
	n := len(parent)
	adj := make([][]int, n)

	for i := 1; i < n; i++ {
		adj[parent[i]] = append(adj[parent[i]], i)
	}

	var sum int
	var res int
	var dfs func(u int, d int) int

	dfs = func(u int, d int) int {
		h := d
		sum += nums[u]
		res += d * nums[u]
		for _, v := range adj[u] {
			h = max(h, dfs(v, d+1))
		}
		return h
	}

	h := dfs(0, 0)

	res = sum*(h+1) - res

	return int64(res)
}
