package p4071

import "testing"

func runSample(t *testing.T, n int, s string, expect int) {
	res := minRotations(n, s)

	if res != expect {
		t.Fatalf("Sample n=%d s=%s, expect %d, but got %d", n, s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, "1502", 9)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, "2916", 12)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, "4219", 6)
}
