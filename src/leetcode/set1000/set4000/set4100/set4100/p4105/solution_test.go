package p4105

import "testing"

func runSample(t *testing.T, a string, b string, x int, y int, expect bool) {
	t.Helper()
	res := repeatedStringMatch(a, b, x, y)
	if res != expect {
		t.Fatalf("Sample a=%q b=%q x=%d y=%d expect %v, but got %v", a, b, x, y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "ab", "abab", 2, 1, true)
}

func TestSample2(t *testing.T) {
	runSample(t, "abc", "bc", 2, 3, false)
}

func TestSample3(t *testing.T) {
	runSample(t, "abcd", "cdab", 2, 1, true)
}
