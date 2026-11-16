package p4112

import "testing"

func runSample(t *testing.T, parent []int, op []int, score []int, expect int) {
	res := maxScore(parent, op, score)

	if res != expect {
		t.Fatalf("Sample parent %v, op %v, score %v, expect %d, but got %d", parent, op, score, expect, res)
	}
}

func TestSample1(t *testing.T) {
	parent := []int{-1, 0, 0, 0, 1, 2, 3, 3}
	op := []int{1, 1, 0, 0, 1, 1, 0, 1}
	score := []int{5, 6, 1, 2, 3, 2, 3, -3}
	runSample(t, parent, op, score, 16)
}

func TestSample2(t *testing.T) {
	parent := []int{-1, 0, 0}
	op := []int{0, 0, 0}
	score := []int{1, 2, 3}
	runSample(t, parent, op, score, 6)
}

func TestSample3(t *testing.T) {
	parent := []int{-1}
	op := []int{1}
	score := []int{-10}
	runSample(t, parent, op, score, -10)
}
