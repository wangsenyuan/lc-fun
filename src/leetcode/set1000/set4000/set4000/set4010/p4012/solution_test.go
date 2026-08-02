package p4012

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, tasks []int, shifts []int, expect []int) {
	res := countTasks(tasks, shifts)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tasks := []int{1, 4, 4}
	shifts := []int{9, 1, 4}
	expect := []int{0, 2, 1}
	runSample(t, tasks, shifts, expect)
}

func TestSample2(t *testing.T) {
	tasks := []int{2, 3, 4}
	shifts := []int{20, 4, 5}
	expect := []int{0, 2, 0}
	runSample(t, tasks, shifts, expect)
}

func TestSample3(t *testing.T) {
	tasks := []int{4, 2}
	shifts := []int{3, 6, 1}
	expect := []int{2, 0, 2}
	runSample(t, tasks, shifts, expect)
}
