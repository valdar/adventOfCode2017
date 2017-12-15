package main

import "testing"

func TestCalcSpiralLayerAndPosition(t *testing.T) {
	cases := []struct {
		in           int
		wantPosition int
		wantLayer    int
	}{
		{9, 8, 2},
		{2, 1, 2},
		{13, 4, 3},
		{25, 16, 3},
	}
	for _, c := range cases {
		gotPosition, gotLayer := CalcSpiralLayerAndPosition(c.in)
		if gotPosition != c.wantPosition || gotLayer != c.wantLayer {
			t.Errorf("CalcSpiralLayerAndPosition(%d) == %d,%d want %d,%d", c.in, gotPosition, gotLayer, c.wantPosition, c.wantLayer)
		}
	}
}

func TestCalcCoordinates(t *testing.T) {
	cases := []struct {
		inPosition int
		inLayern   int
		wantX      int
		wantY      int
	}{
		{4, 3, 5, 5},
		{7, 3, 2, 5},
		{1, 2, 3, 2},
		{8, 2, 3, 1},
		{10, 3, 1, 3},
	}
	for _, c := range cases {
		gotX, gotY := CalcCoordinates(c.inPosition, c.inLayern)
		if gotX != c.wantX || gotY != c.wantY {
			t.Errorf("CalcCoordinates(%d,%d) == %d,%d want %d,%d", c.inPosition, c.inLayern, gotX, gotY, c.wantX, c.wantY)
		}
	}
}

func TestSolveB(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{747, 806},
	}
	for _, c := range cases {
		got := SolveB(c.in)
		if got != c.want {
			t.Errorf("SolveB(%d) = %d want %d", c.in, got, c.want)
		}
	}
}
