package main

import "testing"

func TestCalcFullBlocks(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"flqrgnkx", 8108},
	}
	for _, c := range cases {
		got := CalcFullBlocks(c.in)
		if got != c.want {
			t.Errorf("CalcFullBlocks(%v) == %d want %d", c.in, got, c.want)
		}
	}
}

func TestCalcRegions(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"flqrgnkx", 1242},
	}
	for _, c := range cases {
		got := CalcRegions(c.in)
		if got != c.want {
			t.Errorf("CalcRegions(%v) == %d want %d", c.in, got, c.want)
		}
	}
}
