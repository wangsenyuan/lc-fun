package p4062

import "testing"

func runSample(t *testing.T, source []int, target []int, expect bool) {
	res := canTransform(source, target)

	if res != expect {
		t.Fatalf("Sample source %v target %v, expect %t, but got %t", source, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	source := []int{1, 2, 3}
	target := []int{0, 2, 4}
	runSample(t, source, target, true)
}

func TestSample2(t *testing.T) {
	source := []int{-5, -5}
	target := []int{-15, 5}
	runSample(t, source, target, true)
}

func TestSample3(t *testing.T) {
	source := []int{1, 2, 1}
	target := []int{0, 2, 5}
	runSample(t, source, target, false)
}
