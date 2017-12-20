package main

import (
	"testing"
)

//     |
//     |  +--+
//     A  |  C
// F---|----E|--+
// |  |  |  D
// +B-+  +--+
func TestWalkTheMaze(t *testing.T) {
	cases := []struct {
		in        [][]string
		want      string
		wantSteps int
	}{
		{[][]string{
			[]string{" ", " ", " ", " ", "|", " ", " ", " ", " ", " ", " ", " ", " ", " "},
			[]string{" ", " ", " ", " ", "|", " ", " ", "+", "-", "-", "+", " ", " ", " "},
			[]string{" ", " ", " ", " ", "A", " ", " ", "|", " ", " ", "C", " ", " ", " "},
			[]string{"F", "-", "-", "-", "|", "-", "-", "-", "-", "E", "|", "-", "-", "+"},
			[]string{" ", " ", " ", " ", "|", " ", " ", "|", " ", " ", "|", " ", " ", "D"},
			[]string{" ", " ", " ", " ", "+", "B", "-", "+", " ", " ", "+", "-", "-", "+"},
		}, "ABCDEF", 38},
	}
	for _, c := range cases {
		gotLetters, gotSteps := WalkTheMaze(c.in)
		if gotLetters != c.want || gotSteps != c.wantSteps {
			t.Errorf("WalkTheMaze( %v ) = %v, %d expected %v, %d", c.in, gotLetters, gotSteps, c.want, c.wantSteps)
		}
	}
}
