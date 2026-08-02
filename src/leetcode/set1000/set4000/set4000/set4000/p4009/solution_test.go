package p4009

import "testing"

func runSample(t *testing.T, demand []int, fuel []int, expect int) {
	res := minMaxWaitingTime(demand, fuel)

	if res != expect {
		t.Fatalf("Sample %v %v, expect %d, but got %d", demand, fuel, expect, res)
	}
}

func TestSample1(t *testing.T) {
	demand := []int{6, 8, 4, 6, 5}
	fuel := []int{16, 13}
	expect := 6
	runSample(t, demand, fuel, expect)
}

func TestSample2(t *testing.T) {
	demand := []int{10, 15}
	fuel := []int{12, 17}
	expect := 0
	runSample(t, demand, fuel, expect)
}

func TestSample3(t *testing.T) {
	demand := []int{10, 5}
	fuel := []int{8, 8}
	expect := -1
	runSample(t, demand, fuel, expect)
}

func TestSingleCar(t *testing.T) {
	demand := []int{7}
	fuel := []int{7, 1}
	expect := 0
	runSample(t, demand, fuel, expect)
}

func TestEarlierPumpCanStillBeBusyAfterSwitchingAway(t *testing.T) {
	demand := []int{100, 1, 1}
	fuel := []int{101, 1}
	expect := 100
	runSample(t, demand, fuel, expect)
}

func TestBusyPumpStateCannotBeCollapsedToMinimum(t *testing.T) {
	demand := []int{1, 1, 1, 2, 2, 1}
	fuel := []int{4, 4}
	expect := 1
	runSample(t, demand, fuel, expect)
}
