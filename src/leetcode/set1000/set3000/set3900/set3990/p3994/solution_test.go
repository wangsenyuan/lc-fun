package p3994

import "testing"

func runSample(t *testing.T, nums []int, a int, b int, expect int) {
	res := minAdjacentSwaps(nums, a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 2, 4, 5, 6}
	a := 3
	b := 4
	expect := 1
	runSample(t, nums, a, b, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{9, 7, 5, 3}
	a := 4
	b := 8
	expect := 5
	runSample(t, nums, a, b, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{3, 7, 5, 9}
	a := 4
	b := 8
	expect := 0
	runSample(t, nums, a, b, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{945, 461}
	a := 380
	b := 868
	expect := 1
	runSample(t, nums, a, b, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{242, 106, 150}
	a := 203
	b := 263
	expect := 2
	runSample(t, nums, a, b, expect)
}
