package p4007

import "testing"

func runSample(t *testing.T, planks []int, expect int) {
	res := maximumWidth(planks)

	if res != expect {
		t.Fatalf("Sample %v, expect %d, but got %d", planks, expect, res)
	}
}

func TestSample1(t *testing.T) {
	planks := []int{59, 37, 55, 90, 94}
	expect := 2
	runSample(t, planks, expect)
}

func TestSinglesAndPairsCanShareTheTargetHeight(t *testing.T) {
	planks := []int{2, 2, 2, 2, 4, 4, 4}
	expect := 5
	runSample(t, planks, expect)
}

func TestEqualHeightPairs(t *testing.T) {
	planks := []int{3, 3, 3, 3, 6, 6, 6, 6, 6}
	expect := 7
	runSample(t, planks, expect)
}

func TestDifferentHeightPairsUseTheSmallerFrequency(t *testing.T) {
	planks := []int{1, 1, 1, 4, 4, 5, 5}
	expect := 4
	runSample(t, planks, expect)
}
