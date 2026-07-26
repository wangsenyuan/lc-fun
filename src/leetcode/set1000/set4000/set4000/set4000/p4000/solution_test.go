package p4000

import "testing"

func runSample(t *testing.T, n int, s int, expect int) {
	res := largestInteger(n, s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 9, 90)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 19, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 0, 0)
}
