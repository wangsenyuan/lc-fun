package p4023

import "testing"

func runSample(t *testing.T, n int, start int, requests []int, expect int) {
	res := elevatorRequests(n, start, requests)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, start := 6, 4
	requests := []int{1, 5}
	expect := 6
	runSample(t, n, start, requests, expect)
}

func TestSample2(t *testing.T) {
	n, start := 8, 3
	requests := []int{3, 7, 1}
	expect := 10
	runSample(t, n, start, requests, expect)
}

func TestSample3(t *testing.T) {
	n, start := 10, 5
	requests := []int{0, 2, 9}
	expect := 22
	runSample(t, n, start, requests, expect)
}
