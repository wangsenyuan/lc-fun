package p4104

import "testing"

func runSample(t *testing.T, s string, expect int64) {
	res := countValidSubstrings(s)

	if res != expect {
		t.Fatalf("Sample s=%s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abac", 9)
}

func TestSample2(t *testing.T) {
	runSample(t, "abc", 5)
}

func TestSample3(t *testing.T) {
	runSample(t, "ac", 3)
}
