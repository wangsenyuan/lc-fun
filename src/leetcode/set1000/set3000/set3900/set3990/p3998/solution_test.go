package p3998

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, s string, words []string, expect []bool) {
	res := transformStr(s, words)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "101"
	words := []string{"1?1", "0?1", "0?0"}
	expect := []bool{true, true, false}
	runSample(t, s, words, expect)
}

func TestSample2(t *testing.T) {
	s := "1100"
	words := []string{"0011", "11?1", "1?1?"}
	expect := []bool{true, false, true}
	runSample(t, s, words, expect)
}

func TestSample3(t *testing.T) {
	s := "001"
	words := []string{"?1?", "01?"}
	expect := []bool{false, false}
	runSample(t, s, words, expect)
}
