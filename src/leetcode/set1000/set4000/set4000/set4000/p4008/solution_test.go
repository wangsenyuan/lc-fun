package p4008

import "testing"

func runSample(t *testing.T, monsters []int, boosts [][]int, expect int64) {
	res := minInitialStrength(monsters, boosts)

	if res != expect {
		t.Fatalf("Sample %v %v, expect %d, but got %d", monsters, boosts, expect, res)
	}
}

func TestSample1(t *testing.T) {
	monsters := []int{5, 10, 15}
	boosts := [][]int{{1, 1, 10}}
	expect := int64(30)
	runSample(t, monsters, boosts, expect)
}

func TestSample2(t *testing.T) {
	monsters := []int{5, 10, 15}
	boosts := [][]int{{1, 2, 10}, {1, 2, 5}}
	expect := int64(5)
	runSample(t, monsters, boosts, expect)
}
