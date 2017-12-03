package main

import "testing"

func TestDoTheSumPass1(t *testing.T) {
	cases := []struct {
		in   string
		pass int
		want int
	}{
		{"3345673", 1, 6},
		{"434555673", 1, 10},
	}
	for _, c := range cases {
		got := DoTheSum(c.in, c.pass)
		if got != c.want {
			t.Errorf("DoTheSum(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDoTheSumDifferentPasses(t *testing.T) {
	cases := []struct {
		in   string
		pass int
		want int
	}{
		{"1212", 2, 6},
		{"1221", 2, 0},
		{"123425", 3, 4},
		{"123123", 3, 12},
		{"12131415", 4, 4},
	}
	for _, c := range cases {
		got := DoTheSum(c.in, c.pass)
		if got != c.want {
			t.Errorf("DoTheSum(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
