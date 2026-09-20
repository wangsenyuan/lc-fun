package p4059

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := largestPower(nums)
	assert.Equal(t, expect, res)
}

func TestSample1(t *testing.T) {
	nums := []int{7, 5}
	expect := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 2}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 1, 7}
	expect := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3}
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 11}
	// 1011
	// 0010
	expect := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 1}
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{5, 13}
	// 0101
	// 1101
	expect := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 0, 2}
	runSample(t, nums, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{6, 3}
	// 110
	// 011
	expect := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 0}
	runSample(t, nums, expect)
}
