package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolveA(t *testing.T) {
	s := "..#\n#..\n...\n"
	b := strings.NewReader(s)
	grid, size := parseA(bufio.NewReader(b))

	cases := []struct {
		in              map[string]bool
		inNumberOfBurst int
		want            int
	}{
		{grid, 70, 41},
	}
	for _, c := range cases {
		got := SolveA(c.in, size/2, size/2, c.inNumberOfBurst)
		if got != c.want {
			t.Errorf("SolveA( %v, %d, %d, %d ) = %d expected %d", c.in, size/2, size/2, c.inNumberOfBurst, got, c.want)
		}
	}
}

func TestSolveB(t *testing.T) {
	s := "..#\n#..\n...\n"
	b := strings.NewReader(s)
	grid, size := parseB(bufio.NewReader(b))

	cases := []struct {
		in              map[string]string
		inNumberOfBurst int
		want            int
	}{
		{grid, 100, 26},
	}
	for _, c := range cases {
		got := SolveB(c.in, size/2, size/2, c.inNumberOfBurst)
		if got != c.want {
			t.Errorf("SolveB( %v, %d, %d, %d ) = %d expected %d", c.in, size/2, size/2, c.inNumberOfBurst, got, c.want)
		}
	}
}
