package p4026

import "testing"

func runSample(t *testing.T, skills string, station string, expect int) {
	res := maximumGap(skills, station)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	skills := "aa"
	station := "aaaa"
	expect := 3
	runSample(t, skills, station, expect)
}

func TestSample2(t *testing.T) {
	skills := "xyz"
	station := "xyzz"
	expect := 2
	runSample(t, skills, station, expect)
}

func TestSample3(t *testing.T) {
	skills := "cbc"
	station := "cbcdbc"
	expect := 4
	runSample(t, skills, station, expect)
}
