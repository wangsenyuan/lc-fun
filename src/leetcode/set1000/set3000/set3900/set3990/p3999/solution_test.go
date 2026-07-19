package p3999

import "testing"

func runSample(t *testing.T, words []string, expect int) {
	res := minimumGroups(words)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{
		"ntgwz", "zwntg",
	}
	expect := 1
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{
		"abc", "cab", "bac", "acb", "bca", "cba",
	}
	expect := 3
	runSample(t, words, expect)
}

func TestSample3(t *testing.T) {
	words := []string{
		"leet", "abb", "bab", "deed", "edde", "code", "bba",
	}
	expect := 5
	runSample(t, words, expect)
}
