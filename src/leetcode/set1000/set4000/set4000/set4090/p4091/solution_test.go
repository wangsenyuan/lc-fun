package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := countMagicalPairs(s)

	if int64(expect) != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aba", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, "kkk", 4)
}

func TestSample3(t *testing.T) {
	runSample(t, "abcac", 7)
}
