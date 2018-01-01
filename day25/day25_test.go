package main

import (
	"testing"
)

func TestSolveA(t *testing.T) {

	cases := []struct {
		inProgram map[string]nextState
		inSteps   int
		want      int
	}{
		{
			map[string]nextState{
				"A,0": nextState{"B", 1, RIGHT},
				"A,1": nextState{"B", 0, LEFT},
				"B,0": nextState{"A", 1, LEFT},
				"B,1": nextState{"A", 1, RIGHT},
			},
			6,
			3,
		},
	}
	for _, c := range cases {
		got := SolveA(c.inSteps, c.inProgram)
		if got != c.want {
			t.Errorf("SolveA( %d, %v ) = %d expected %d", c.inSteps, c.inProgram, got, c.want)
		}
	}
}
