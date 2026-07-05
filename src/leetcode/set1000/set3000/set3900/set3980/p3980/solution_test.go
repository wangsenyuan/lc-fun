package p3980

import "testing"

func runSample(t *testing.T, s1 string, s2 string, expect int) {
	res := minOperations(s1, s2)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s1 := "11"
	s2 := "00"
	expect := 1
	runSample(t, s1, s2, expect)
}

func TestSample2(t *testing.T) {
	s1 := "01"
	s2 := "10"
	expect := 3
	runSample(t, s1, s2, expect)
}

func TestSample3(t *testing.T) {
	s1 := "1"
	s2 := "0"
	expect := -1
	runSample(t, s1, s2, expect)
}

func TestExhaustiveSmallAgainstBruteForce(t *testing.T) {
	for n := 1; n <= 5; n++ {
		limit := 1 << n
		for a := 0; a < limit; a++ {
			s1 := formatBits(a, n)
			for b := 0; b < limit; b++ {
				s2 := formatBits(b, n)
				expect := bruteForce(s1, s2)
				runSample(t, s1, s2, expect)
			}
		}
	}
}

func formatBits(mask int, n int) string {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if mask>>(n-1-i)&1 == 1 {
			buf[i] = '1'
		} else {
			buf[i] = '0'
		}
	}
	return string(buf)
}

func bruteForce(s1 string, s2 string) int {
	n := len(s1)
	start := parseBits(s1)
	target := parseBits(s2)
	dist := make([]int, 1<<n)
	for i := range dist {
		dist[i] = -1
	}
	queue := []int{start}
	dist[start] = 0

	for head := 0; head < len(queue); head++ {
		state := queue[head]
		if state == target {
			return dist[state]
		}

		for i := 0; i < n; i++ {
			if state>>(n-1-i)&1 == 0 {
				next := state | (1 << (n - 1 - i))
				if dist[next] < 0 {
					dist[next] = dist[state] + 1
					queue = append(queue, next)
				}
			}
		}

		for i := 0; i+1 < n; i++ {
			left := 1 << (n - 1 - i)
			right := 1 << (n - 2 - i)
			if state&left != 0 && state&right != 0 {
				next := state &^ left &^ right
				if dist[next] < 0 {
					dist[next] = dist[state] + 1
					queue = append(queue, next)
				}
			}
		}
	}

	return -1
}

func parseBits(s string) int {
	var res int
	for i := range s {
		res <<= 1
		if s[i] == '1' {
			res++
		}
	}
	return res
}
