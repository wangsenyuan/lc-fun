package p4002

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := countValidSequences(n, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 3, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 2, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 5, 0)
}
