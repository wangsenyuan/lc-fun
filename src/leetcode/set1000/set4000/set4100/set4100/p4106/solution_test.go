package p4106

import (
	"slices"
	"testing"
)

func equalSegments(got [][]string, expect [][]string) bool {
	if len(got) != len(expect) {
		return false
	}
	for i := range got {
		if !slices.Equal(got[i], expect[i]) {
			return false
		}
	}
	return true
}

func runSample(t *testing.T, s1 string, s2 string, expect [][]string) {
	res := findMismatchedSegments(s1, s2)

	if !equalSegments(res, expect) {
		t.Fatalf("Sample s1=%s s2=%s, expect %v, but got %v", s1, s2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	expect := [][]string{{"ar", "mp"}, {"h", "y"}}
	runSample(t, "earth", "empty", expect)
}

func TestSample2(t *testing.T) {
	expect := [][]string{{"ac", "na"}}
	runSample(t, "action", "nation", expect)
}

func TestSample3(t *testing.T) {
	runSample(t, "hot", "hot", nil)
}
