package p4043

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := countRotations(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aab"
	k := 1
	expect := 2
	runSample(t, s, k, expect)
}
