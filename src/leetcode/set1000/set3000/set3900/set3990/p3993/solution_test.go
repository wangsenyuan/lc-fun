package p3993

import "testing"

func runSample(t *testing.T, n int, s int, m int, expect int) {
	res := maximumValue(n, s, m)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 3, 5, 12)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 4, 3, 7)
}
