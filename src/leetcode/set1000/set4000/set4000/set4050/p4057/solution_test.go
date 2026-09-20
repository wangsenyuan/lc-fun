package p4057

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func runSample(t *testing.T, intervals [][]int, expect int64) {
	res := countIntersectingIntervals(intervals)
	assert.Equal(t, expect, res)
}

func TestSample1(t *testing.T) {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}}
	expect := int64(2)
	runSample(t, intervals, expect)
}

func TestSample2(t *testing.T) {
	intervals := [][]int{{1, 5}, {2, 4}, {3, 6}}
	expect := int64(3)
	runSample(t, intervals, expect)
}
