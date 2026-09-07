package p4045

import "testing"

func runSample(t *testing.T, position []int, speed []int, distance int, expect int) {
	res := countGroups(position, speed, distance)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	position := []int{1, 5, 6, 20}
	speed := []int{4, 3, 2, 3}
	distance := 1
	expect := 2
	runSample(t, position, speed, distance, expect)
}

func TestSample2(t *testing.T) {
	position := []int{1, 5, 9}
	speed := []int{3, 2, 2}
	distance := 2
	expect := 2
	runSample(t, position, speed, distance, expect)
}
