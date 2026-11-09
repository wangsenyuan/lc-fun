package p4102

import "testing"

func runSample(t *testing.T, s string, expect int64) {
	res := countValidSubstrings(s)

	if res != expect {
		t.Fatalf("Sample s=%s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abac", 10)
}

func TestSample2(t *testing.T) {
	runSample(t, "abc", 5)
}

func TestSample3(t *testing.T) {
	runSample(t, "ac", 3)
}

func TestSample4(t *testing.T) {
	runSample(t, "aabc", 8)
}
