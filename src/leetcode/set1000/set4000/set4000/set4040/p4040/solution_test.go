package p4040

import "testing"

func runSample(t *testing.T, nums []int, sum int, expect int) {
	res := minOperations(nums, sum)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, 6, 10}
	sum := 4
	expect := 3
	runSample(t, nums, sum, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{10, 2}
	sum := 13
	expect := 3
	runSample(t, nums, sum, expect)
}
