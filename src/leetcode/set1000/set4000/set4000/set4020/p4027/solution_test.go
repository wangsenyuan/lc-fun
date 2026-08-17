package p4027

import "testing"

func runSample(t *testing.T, n int, start int, requests [][]int, expect int) {
	ans := elevatorRequests(n, start, requests)
	if ans != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	n := 9
	start := 0
	requests := [][]int{
		{0, 8},
		{6, 5},
	}
	expect := 9
	runSample(t, n, start, requests, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	start := 5
	requests := [][]int{
		{1, 7},
		{7, 3},
	}
	expect := 7
	runSample(t, n, start, requests, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	start := 3
	requests := [][]int{
		{0, 5}, {0, 1}, {6, 3},
	}
	expect := 8
	runSample(t, n, start, requests, expect)
}
