package p3985

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	ans := getSum(nums)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{10, 10}
	expect := int64(20)
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 2, 1, 5, 6}
	expect := int64(9)
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{7, 1, 2, 1, 7, 3, 4, 3, 4}
	expect := int64(18)
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expect := int64(5)
	runSample(t, nums, expect)
}
