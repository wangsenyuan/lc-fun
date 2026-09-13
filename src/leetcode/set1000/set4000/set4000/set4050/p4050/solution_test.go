package p4050

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := minDays(n)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 12, 7)
}
