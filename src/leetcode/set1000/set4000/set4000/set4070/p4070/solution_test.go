package p4070

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minRotations(s)

	if res != expect {
		t.Fatalf("Sample s=%s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "0192837465", 25)
}

func TestSample2(t *testing.T) {
	runSample(t, "1200210200", 12)
}
