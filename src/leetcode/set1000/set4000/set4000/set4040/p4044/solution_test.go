package p4044

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countGoodRotations(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 1, 2}
	expect := 0
	runSample(t, nums, expect)
}
