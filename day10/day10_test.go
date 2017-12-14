package main

import (
	"reflect"
	"testing"

	"github.com/valdar/adventOfCode2017/day10/hashKnot"
)

func TestCalcOccurencies(t *testing.T) {
	cases := []struct {
		in          []int
		inStartPos  int
		inSkipSize  int
		inLength    int
		want        []int
		wantNextPos int
	}{
		{[]int{0, 1, 2, 3, 4}, 0, 0, 3, []int{2, 1, 0, 3, 4}, 3},
		{[]int{2, 1, 0, 3, 4}, 3, 1, 4, []int{4, 3, 0, 1, 2}, 3},
		{[]int{4, 3, 0, 1, 2}, 3, 2, 1, []int{4, 3, 0, 1, 2}, 1},
		{[]int{4, 3, 0, 1, 2}, 1, 3, 5, []int{3, 4, 2, 1, 0}, 4},
	}
	for _, c := range cases {
		originalIn := append([]int{}, c.in...)
		got := hashKnot.PerformStep(c.in, c.inStartPos, c.inSkipSize, c.inLength, len(c.in))
		if !reflect.DeepEqual(c.in, c.want) || got != c.wantNextPos {
			t.Errorf("PerformStep(%v, %d, %d, %d, %d) == %v, %d want %v, %d", originalIn, c.inStartPos, c.inSkipSize, c.inLength, len(c.in), c.in, got, c.want, c.wantNextPos)
		}
	}
}

func TestCalcHash(t *testing.T) {
	cases := []struct {
		in         []int
		inLenght   []string
		want       int
		wantStatus []int
	}{
		{[]int{0, 1, 2, 3, 4}, []string{"3", "4", "1", "5"}, 12, []int{3, 4, 2, 1, 0}},
	}
	for _, c := range cases {
		originalIn := append([]int{}, c.in...)
		got := hashKnot.CalcHash(c.in, c.inLenght)
		if !reflect.DeepEqual(c.in, c.wantStatus) || got != c.want {
			t.Errorf("CalcHash(%v, %v) == %v, %d want %v, %d", originalIn, c.inLenght, c.in, got, c.wantStatus, c.want)
		}
	}
}

func TestCalcHashKnot(t *testing.T) {
	inputSlice := make([]int, 256, 256)
	for i := 0; i < 256; i++ {
		inputSlice[i] = i
	}
	cases := []struct {
		in   string
		want string
	}{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "8fa34530c0cbce096407b4d7b298ce71"},
		{"1,2,3", "58d80f7f8865c66e2d78a61ce6af2536"},
		{"1,2,4", "bc0d6ccdc92c757405822f023cbd7ea1"},
	}
	for _, c := range cases {
		got := hashKnot.CalcHashKnot(inputSlice, c.in)
		if got != c.want {
			t.Errorf("CalcHashKnot(inputSlice, %v) == %v want %v", c.in, got, c.want)
		}
	}
}
