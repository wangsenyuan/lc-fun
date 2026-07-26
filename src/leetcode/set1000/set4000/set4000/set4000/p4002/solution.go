package p4002

func countValidSequences(n int, k int) int {
	const mod = 1e9 + 7

	add := func(a int, b int) int {
		a += b
		if a >= mod {
			a -= mod
		}
		return a
	}

	sub := func(a int, b int) int {
		return add(a, mod-b)
	}

	mul := func(a int, b int) int {
		return a * b % mod
	}

	pow := func(a int, b int) int {
		res := 1
		for b > 0 {
			if b&1 == 1 {
				res = mul(res, a)
			}
			a = mul(a, a)
			b >>= 1
		}
		return res
	}

	F := make([]int, n+1)
	I := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(F[i-1], i)
	}
	I[n] = pow(F[n], mod-2)
	for i := n - 1; i >= 0; i-- {
		I[i] = mul(I[i+1], i+1)
	}

	nCr := func(n int, r int) int {
		if r < 0 || n < r {
			return 0
		}
		return mul(F[n], mul(I[r], I[n-r]))
	}

	res := nCr(n-1, k-1)
	if (n-k)&1 == 0 {
		res = sub(res, nCr((n-k)/2+k-1, k-1))
	}
	return res
}
