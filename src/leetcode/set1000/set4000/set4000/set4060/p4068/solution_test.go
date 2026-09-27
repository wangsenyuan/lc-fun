package p4068

import "testing"

func runSample(t *testing.T, meetings [][]int, expect int64) {
	res := maxEarnings(meetings)

	if res != expect {
		t.Fatalf("Sample %v, expect %d, but got %d", meetings, expect, res)
	}
}

func TestSample1(t *testing.T) {
	meetings := [][]int{{2, 5, 4}, {6, 8, 3}}
	runSample(t, meetings, 8)
}

func TestSample2(t *testing.T) {
	meetings := [][]int{{3, 5, 4}, {4, 7, 8}, {8, 10, 3}}
	runSample(t, meetings, 12)
}

func TestSample3(t *testing.T) {
	meetings := [][]int{{1, 2, 2}, {4, 5, 2}, {7, 9, 3}}
	runSample(t, meetings, 11)
}

func TestSample4(t *testing.T) {
	meetings := [][]int{{7, 13, 10}, {4, 7, 8}}
	runSample(t, meetings, 18)
}
