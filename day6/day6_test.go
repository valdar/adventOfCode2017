package main

import "testing"

func TestCalcStepsSameConfig(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{0, 2, 7, 0}, 5},
	}
	for _, c := range cases {
		got, _ := CalcStepsSameConfig(c.in)
		if got != c.want {
			t.Errorf("CalcStepsSameConfig(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestCalcStepsSameConfigB(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{0, 2, 7, 0}, 4},
	}
	for _, c := range cases {
		_, alreadySeen := CalcStepsSameConfig(c.in)
		got, _ := CalcStepsSameConfig(alreadySeen)
		if got != c.want {
			t.Errorf("CalcStepsSameConfig(CalcStepsSameConfig(%v)) == %d, want %d", c.in, got, c.want)
		}
	}
}
