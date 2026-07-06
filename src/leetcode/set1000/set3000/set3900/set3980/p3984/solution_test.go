package p3981

import "testing"

func runSample(t *testing.T, word1 string, word2 string, target string, expect int) {
	res := interleaveCharacters(word1, word2, target)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word1 := "abc"
	word2 := "bac"
	target := "abc"
	expect := 5

	runSample(t, word1, word2, target, expect)
}

func TestSample2(t *testing.T) {
	word1 := "cd"
	word2 := "cd"
	target := "ccd"
	expect := 4

	runSample(t, word1, word2, target, expect)
}

func TestSample3(t *testing.T) {
	word1 := "xy"
	word2 := "xy"
	target := "xyxy"
	expect := 2

	runSample(t, word1, word2, target, expect)
}

func TestSample4(t *testing.T) {
	word1 := "ab"
	word2 := "cde"
	target := "ade"
	expect := 1

	runSample(t, word1, word2, target, expect)
}

func TestSkipBeforeFirstMatch(t *testing.T) {
	word1 := "xa"
	word2 := "b"
	target := "ab"
	expect := 1

	runSample(t, word1, word2, target, expect)
}
