package p3996

import "testing"

func runSample(t *testing.T, start []int, target []int, expect bool) {
	res := canReach(start, target)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	start := []int{1, 1}
	target := []int{2, 2}
	expect := true
	runSample(t, start, target, expect)
}

func TestSample2(t *testing.T) {
	start := []int{1, 1}
	target := []int{2, 3}
	expect := false
	runSample(t, start, target, expect)
}

func TestSameCell(t *testing.T) {
	start := []int{8, 8}
	target := []int{8, 8}
	expect := true
	runSample(t, start, target, expect)
}
