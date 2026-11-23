package p4120

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countDominantElements(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 10, 2}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 4, 2}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{7}
	expect := 1
	runSample(t, nums, expect)
}
