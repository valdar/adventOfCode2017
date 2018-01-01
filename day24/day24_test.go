package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolveA(t *testing.T) {
	s := "0/2\n2/2\n2/3\n3/4\n3/5\n0/1\n10/1\n9/10\n"
	b := strings.NewReader(s)
	parts := parse(bufio.NewReader(b))

	cases := []struct {
		in   []port
		want int
	}{
		{parts, 31},
	}
	for _, c := range cases {
		got := SolveA(c.in)
		if got != c.want {
			t.Errorf("SolveA( %v ) = %d expected %d", c.in, got, c.want)
		}
	}
}

func TestSolveB(t *testing.T) {
	s := "0/2\n2/2\n2/3\n3/4\n3/5\n0/1\n10/1\n9/10\n"
	b := strings.NewReader(s)
	parts := parse(bufio.NewReader(b))

	cases := []struct {
		in   []port
		want int
	}{
		{parts, 19},
	}
	for _, c := range cases {
		got := SolveB(c.in)
		if got != c.want {
			t.Errorf("SolveB( %v ) = %d expected %d", c.in, got, c.want)
		}
	}
}
