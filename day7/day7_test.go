package main

import (
	"testing"
)

func TestFindStartProgram(t *testing.T) {
	cases := []struct {
		in   map[string]Program
		want string
	}{
		{map[string]Program{
			"pbga": {"pbga", 66, []string{}},
			"xhth": {"xhth", 57, []string{}},
			"ebii": {"ebii", 61, []string{}},
			"havc": {"havc", 66, []string{}},
			"ktlj": {"ktlj", 57, []string{}},
			"fwft": {"fwft", 72, []string{"ktlj", "cntj", "xhth"}},
			"qoyq": {"qoyq", 66, []string{}},
			"padx": {"padx", 45, []string{"pbga", "havc", "qoyq"}},
			"tknk": {"tknk", 41, []string{"ugml", "padx", "fwft"}},
			"jptl": {"jptl", 61, []string{}},
			"ugml": {"ugml", 68, []string{"gyxo", "ebii", "jptl"}},
			"gyxo": {"gyxo", 61, []string{}},
			"cntj": {"cntj", 57, []string{}},
		}, "tknk"},
	}
	for _, c := range cases {
		got := FindStartProgram(c.in)
		if got != c.want {
			t.Errorf("FindStartProgram(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TesFindTheImbalanceDifference(t *testing.T) {
	cases := []struct {
		in   map[string]Program
		want int
	}{
		{map[string]Program{
			"pbga": {"pbga", 66, []string{}},
			"xhth": {"xhth", 57, []string{}},
			"ebii": {"ebii", 61, []string{}},
			"havc": {"havc", 66, []string{}},
			"ktlj": {"ktlj", 57, []string{}},
			"fwft": {"fwft", 72, []string{"ktlj", "cntj", "xhth"}},
			"qoyq": {"qoyq", 66, []string{}},
			"padx": {"padx", 45, []string{"pbga", "havc", "qoyq"}},
			"tknk": {"tknk", 41, []string{"ugml", "padx", "fwft"}},
			"jptl": {"jptl", 61, []string{}},
			"ugml": {"ugml", 68, []string{"gyxo", "ebii", "jptl"}},
			"gyxo": {"gyxo", 61, []string{}},
			"cntj": {"cntj", 57, []string{}},
		}, 8},
	}
	for _, c := range cases {
		got := FindTheImbalanceDifference(c.in)
		if got != c.want {
			t.Errorf("FindTheImbalanceDifference(%v) == %d want %d", c.in, got, c.want)
		}
	}
}
