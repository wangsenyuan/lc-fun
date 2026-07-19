package p3995

import "testing"

func runSample(t *testing.T, source string, target string, rules [][]string, costs []int, expect int) {
	res := minCost(source, target, rules, costs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	source := "hello"
	target := "world"
	rules := [][]string{{"he", "wo"}, {"llo", "rld"}}
	costs := []int{3, 4}
	expect := 7
	runSample(t, source, target, rules, costs, expect)
}

func TestSample2(t *testing.T) {
	source := "cat"
	target := "dog"
	rules := [][]string{{"c*t", "dog"}}
	costs := []int{2}
	expect := 3
	runSample(t, source, target, rules, costs, expect)
}

func TestSample3(t *testing.T) {
	source := "test"
	target := "next"
	rules := [][]string{
		{"*e*t", "next"},
	}
	costs := []int{4}
	expect := 6
	runSample(t, source, target, rules, costs, expect)
}
