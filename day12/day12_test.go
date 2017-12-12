package main

import (
	"testing"
)

func TestFindReachableFromStarting(t *testing.T) {
	cases := []struct {
		inStartingPos int
		inPipeChart   map[int][]int
		want          int
	}{
		{0, map[int][]int{
			0: []int{2},
			1: []int{1},
			2: []int{0, 3, 4},
			3: []int{2, 4},
			4: []int{2, 3, 6},
			5: []int{6},
			6: []int{4, 5},
		}, 6},
	}
	for _, c := range cases {
		got := FindReachableFromStarting(c.inStartingPos, c.inPipeChart)
		if len(got) != c.want {
			t.Errorf("FindReachableFromStarting( %d, %v ) = %d [len(%v)] expected %d", c.inStartingPos, c.inPipeChart, len(got), got, c.want)
		}
	}
}

func TestFindReachableGroups(t *testing.T) {
	cases := []struct {
		inPipeChart map[int][]int
		want        int
	}{
		{map[int][]int{
			0: []int{2},
			1: []int{1},
			2: []int{0, 3, 4},
			3: []int{2, 4},
			4: []int{2, 3, 6},
			5: []int{6},
			6: []int{4, 5},
		}, 2},
	}
	for _, c := range cases {
		got := FindReachableGroups(c.inPipeChart)
		if got != c.want {
			t.Errorf("FindReachableGroups( %v ) = %d expected %d", c.inPipeChart, got, c.want)
		}
	}
}
