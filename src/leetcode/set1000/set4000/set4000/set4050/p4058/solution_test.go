package p4058

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func runSample(t *testing.T, nums []int, expect int64) {
	res := maxValue(nums)
	assert.Equal(t, expect, res)
}

func TestSample1(t *testing.T) {
	nums := []int{1, 5, 2}
	expect := int64(6)
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 4, 3}
	expect := int64(7)
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{9, 7}
	expect := int64(2)
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{-3, -2}
	expect := int64(1)
	runSample(t, nums, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{-10, -1, -9}
	expect := int64(0)
	runSample(t, nums, expect)
}
