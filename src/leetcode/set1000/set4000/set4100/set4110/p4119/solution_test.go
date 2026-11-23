package p4119

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minSetBits(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{6, 7, 3}
	k := 2
	expect := 1
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 5, 12}
	k := 1
	expect := 2
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 1}
	k := 2
	expect := -1
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{9, 55, 48, 29, 10}
	k := 1
	expect := 2
	runSample(t, nums, k, expect)
}
