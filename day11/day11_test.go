package main

import (
	"testing"
)

func TestCalcNextHex(t *testing.T) {
	cases := []struct {
		inX   int
		inY   int
		inDir string
		wantX int
		wantY int
	}{
		{0, 0, "n", 0, 1},
		{0, 0, "ne", 1, 0},
		{0, 0, "se", 1, -1},
		{0, 0, "s", 0, -1},
		{0, 0, "sw", -1, -0},
		{0, 0, "nw", -1, 1},
	}
	for _, c := range cases {
		gotX, gotY := CalcNextHex(c.inX, c.inY, c.inDir)
		if gotX != c.wantX || gotY != c.wantY {
			t.Errorf("CalcNextHex(%d,%d,%v) == %d,%d want %d,%d", c.inX, c.inY, c.inDir, gotX, gotY, c.wantX, c.wantY)
		}
	}
}

func TestA(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 0},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3},
	}
	for _, c := range cases {
		endX, endy, _ := CalcEndingHexAndMax(c.in)
		got := CalcCellDistance(0, 0, endX, endy)
		if got != c.want {
			t.Errorf("Resolve A on input %v, got ditance %d, expected %d", c.in, got, c.want)
		}
	}
}

func TestB(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 2},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3},
	}
	for _, c := range cases {
		_, _, got := CalcEndingHexAndMax(c.in)
		if got != c.want {
			t.Errorf("Resolve A on input %v, got ditance %d, expected %d", c.in, got, c.want)
		}
	}
}
