package p4073

import "testing"

func runSample(t *testing.T, n int64, expect int) {
	res := countGoodStrings(n)

	if res != expect {
		t.Fatalf("Sample n=%d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 2)
}
