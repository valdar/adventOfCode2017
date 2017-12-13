package main

import (
	"testing"
)

func TestCalcStepsOutA(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{0, 3, 0, 1, -3}, 5},
	}
	for _, c := range cases {
		got := CalcStepsOutA(c.in)
		if got != c.want {
			t.Errorf("CalcStepsOutA(%v) == %d want %d", c.in, got, c.want)
		}
	}
}

func TestCalcStepsOutB(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{0, 3, 0, 1, -3}, 10},
	}
	for _, c := range cases {
		got := CalcStepsOutB(c.in)
		if got != c.want {
			t.Errorf("CalcStepsOutB(%v) == %d want %d", c.in, got, c.want)
		}
	}
}
