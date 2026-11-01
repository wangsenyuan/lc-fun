package p4098

import "testing"

func runSample(t *testing.T, n int, index int, value int, expect int64) {
	res := minimumMaxValue(n, index, value)

	if res != expect {
		t.Fatalf("Sample n=%d index=%d value=%d, expect %d, but got %d", n, index, value, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 2, 8, 10)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 1, 3, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 0, 3, 6)
}
