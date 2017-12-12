package main

import (
	"testing"
)

func TestCaclScore(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}
	for _, c := range cases {
		got, _ := CaclScoreAndGarbage(c.in)
		if got != c.want {
			t.Errorf("CaclScoreAndGarbage(%v) == %d want %d", c.in, got, c.want)
		}
	}
}

func TestCaclGarbage(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"{<>}", 0},
		{"{<random characters>}", 17},
		{"{<<<<>}", 3},
		{"{<{!>}>}", 2},
		{"{<!!>}", 0},
		{"{<!!!>>}", 0},
		{"{<{o\"i!a,<{i<a>}", 10},
	}
	for _, c := range cases {
		_, got := CaclScoreAndGarbage(c.in)
		if got != c.want {
			t.Errorf("CaclScoreAndGarbage(%v) == %d want %d", c.in, got, c.want)
		}
	}
}
