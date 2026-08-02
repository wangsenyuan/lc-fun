package p4013

import "testing"

func runSample(t *testing.T, nums []int, a int, b int, expect int) {
	res := countRatioSubarrays(nums, a, b)
	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 1, 2}
	a := 3
	b := 2
	expect := 7
	runSample(t, nums, a, b, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 2, 1}
	a := 2
	b := 1
	expect := 3
	runSample(t, nums, a, b, expect)
}
