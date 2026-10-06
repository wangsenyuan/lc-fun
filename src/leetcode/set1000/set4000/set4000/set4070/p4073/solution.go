package p4073

func countGoodStrings(n int64) int {
	m := matrix{{1, 1}, {1, 0}}
	f0 := matrix{{2}, {0}}
	res := powMat(m, int(n)-1, f0)

	return res[0][0]
}

type matrix [][]int

func newMatrix(r int, c int) matrix {
	res := make(matrix, r)
	for i := range r {
		res[i] = make([]int, c)
	}
	return res
}

const mod = 1e9 + 7

func add(a int, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a int, b int) int {
	return (a * b) % mod
}

func mulMat(a matrix, b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))

	for i := range len(a) {
		for j := range len(b[0]) {
			for k := range len(b) {
				c[i][j] = add(c[i][j], mul(a[i][k], b[k][j]))
			}
		}
	}
	return c
}

func powMat(a matrix, n int, res matrix) matrix {
	for n > 0 {
		if n&1 == 1 {
			res = mulMat(a, res)
		}
		a = mulMat(a, a)
		n >>= 1
	}
	return res
}
