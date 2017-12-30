package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {

	s := "../.# => ##./#../...\n.#./..#/### => #..#/..../..../#..#\n"
	b := strings.NewReader(s)

	cases := []struct {
		inStart   [][]bool
		inMapping map[string][][]bool
		want      int
	}{
		{[][]bool{
			[]bool{false, true, false},
			[]bool{false, false, true},
			[]bool{true, true, true},
		}, parse(bufio.NewReader(b)), 12},
	}

	for _, c := range cases {
		got := Solve(c.inMapping, c.inStart, 2)
		if got != c.want {
			t.Errorf("Solve( %v, %v, %d ) = %d expected %d", c.inMapping, c.inStart, 3, got, c.want)
		}
	}
}
