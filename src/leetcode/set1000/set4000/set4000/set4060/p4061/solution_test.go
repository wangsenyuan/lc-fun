package p4061

import "testing"

func runSample(t *testing.T, source []int, target []int, expect int) {
	res := minQueenMoves(source, target)

	if res != expect {
		t.Fatalf("Sample source %v target %v, expect %d, but got %d", source, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	source := []int{8, 1}
	target := []int{1, 8}
	runSample(t, source, target, 1)
}

func TestSample2(t *testing.T) {
	source := []int{4, 2}
	target := []int{1, 3}
	runSample(t, source, target, 2)
}

func TestSample3(t *testing.T) {
	source := []int{1, 1}
	target := []int{1, 1}
	runSample(t, source, target, 0)
}
