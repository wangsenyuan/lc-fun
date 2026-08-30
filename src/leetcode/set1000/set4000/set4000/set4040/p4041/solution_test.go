package p4041

import "testing"

func runSample(t *testing.T, nums []int, sum int, expect int) {
	res := minOperations(nums, sum)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	nums := []int{10, 2}
	sum := 13
	expect := 3
	runSample(t, nums, sum, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 3}
	sum := 8
	expect := 2
	runSample(t, nums, sum, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 2}
	sum := 7
	expect := -1
	runSample(t, nums, sum, expect)
}
