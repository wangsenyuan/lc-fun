package p3983

import "testing"

func runSample(t *testing.T, s string, w string, expect bool) {
	res := canMakeSubsequence(s, w)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `cat`
	w := "chat"
	expect := true
	runSample(t, s, w, expect)
}

func TestSample2(t *testing.T) {
	s := `plane`
	w := "apple"
	expect := false
	runSample(t, s, w, expect)
}

func TestSample3(t *testing.T) {
	s := `jj`
	w := "khj"
	expect := true
	runSample(t, s, w, expect)
}
