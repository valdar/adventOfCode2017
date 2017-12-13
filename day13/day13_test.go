package main

import (
	"testing"
)

func TestCalcScannerPosition(t *testing.T) {
	cases := []struct {
		inTime  int
		inRange int
		want    int
	}{
		{0, 3 - 1, 0},
		{1, 3 - 1, 1},
		{2, 3 - 1, 2},
		{3, 3 - 1, 1},
		{4, 3 - 1, 0},
		{5, 3 - 1, 1},
		{6, 3 - 1, 2},
	}
	for _, c := range cases {
		got := CalcScannerPosition(c.inTime, c.inRange)
		if got != c.want {
			t.Errorf("CalcScannerPosition( %d, %d ) = %d expected %d", c.inTime, c.inRange, got, c.want)
		}
	}
}

func TestCalcSeverity(t *testing.T) {
	cases := []struct {
		in   map[int]int
		want int
	}{
		{map[int]int{
			0: 3,
			1: 2,
			4: 4,
			6: 4,
		}, 24},
	}
	for _, c := range cases {
		got := CalcSeverity(c.in)
		if got != c.want {
			t.Errorf("CalcSeverity( %v ) = %d expected %d", c.in, got, c.want)
		}
	}
}

func TestCalcDelayToSeverity0(t *testing.T) {
	cases := []struct {
		in   map[int]int
		want int
	}{
		{map[int]int{
			0: 3,
			1: 2,
			4: 4,
			6: 4,
		}, 10},
	}
	for _, c := range cases {
		got := CalcDelayToSeverity0(c.in)
		if got != c.want {
			t.Errorf("CalcDelayToSeverity0( %v ) = %d expected %d", c.in, got, c.want)
		}
	}
}
