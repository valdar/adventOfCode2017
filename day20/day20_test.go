package main

import (
	"testing"
)

func TestSolveA(t *testing.T) {
	cases := []struct {
		in   []particel
		want int
	}{
		{[]particel{
			{1, vector{3, 0, 0}, vector{2, 0, 0}, vector{-1, 0, 0}},
			{0, vector{4, 0, 0}, vector{0, 0, 0}, vector{-2, 0, 0}},
		}, 1},
	}
	for _, c := range cases {
		got := SolveA(c.in)
		if got.id != c.want {
			t.Errorf("SolveA( %v ) = %d expected %d", c.in, got.id, c.want)
		}
	}
}

// p=<-6,0,0>, v=< 3,0,0>, a=< 0,0,0>
// p=<-4,0,0>, v=< 2,0,0>, a=< 0,0,0>
// p=<-2,0,0>, v=< 1,0,0>, a=< 0,0,0>
// p=< 3,0,0>, v=<-1,0,0>, a=< 0,0,0>
func TestSolveB(t *testing.T) {
	cases := []struct {
		in   []particel
		want int
	}{
		{[]particel{
			{0, vector{-6, 0, 0}, vector{3, 0, 0}, vector{0, 0, 0}},
			{1, vector{-4, 0, 0}, vector{2, 0, 0}, vector{0, 0, 0}},
			{2, vector{-2, 0, 0}, vector{1, 0, 0}, vector{0, 0, 0}},
			{3, vector{3, 0, 0}, vector{-1, 0, 0}, vector{0, 0, 0}},
		}, 1},
	}
	for _, c := range cases {
		got := SolveB(c.in)
		if got != c.want {
			t.Errorf("SolveB( %v ) = %d expected %d", c.in, got, c.want)
		}
	}
}
