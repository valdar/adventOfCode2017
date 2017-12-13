package main

import (
	"testing"
)

func TestSolveA(t *testing.T) {
	cases := []struct {
		in   [][]string
		want int
	}{
		{[][]string{
			[]string{"b", "inc", "5", "if", "a", ">", "1"},
			[]string{"a", "inc", "1", "if", "b", "<", "5"},
			[]string{"c", "dec", "-10", "if", "a", ">=", "1"},
			[]string{"c", "inc", "-20", "if", "c", "==", "10"},
		}, 1},
	}
	for _, c := range cases {
		registryState, _ := ComputeInstructions(c.in)
		got := CalcMax(registryState)
		if got != c.want {
			t.Errorf("CalcMax(ComputeInstructions(%v)) == %d want %d", c.in, got, c.want)
		}
	}
}

func TestSolveB(t *testing.T) {
	cases := []struct {
		in   [][]string
		want int
	}{
		{[][]string{
			[]string{"b", "inc", "5", "if", "a", ">", "1"},
			[]string{"a", "inc", "1", "if", "b", "<", "5"},
			[]string{"c", "dec", "-10", "if", "a", ">=", "1"},
			[]string{"c", "inc", "-20", "if", "c", "==", "10"},
		}, 10},
	}
	for _, c := range cases {
		_, got := ComputeInstructions(c.in)
		if got != c.want {
			t.Errorf("ComputeInstructions(%v) == %d want %d", c.in, got, c.want)
		}
	}
}
