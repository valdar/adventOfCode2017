package main

import (
	"reflect"
	"testing"
)

func TestCalcOccurencies(t *testing.T) {
	cases := []struct {
		in   []string
		want map[string]int
	}{
		{[]string{"aa", "bb", "cc", "dd"}, map[string]int{"aa": 1, "bb": 1, "cc": 1, "dd": 1}},
		{[]string{"aa", "bb", "cc", "aa"}, map[string]int{"aa": 2, "bb": 1, "cc": 1}},
		{[]string{"aa", "bb", "cc", "aaa"}, map[string]int{"aa": 1, "bb": 1, "cc": 1, "aaa": 1}},
	}
	for _, c := range cases {
		got := make(map[string]int)
		got = CalcOccurencies(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CalcOccurencies(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestCheckAnagram(t *testing.T) {
	cases := []struct {
		in   []string
		want bool
	}{
		{[]string{"abcde", "fghij"}, false},
		{[]string{"abcde", "xyz", "ecdab"}, true},
		{[]string{"a", "ab", "abc", "abd", "abf", "abj"}, false},
		{[]string{"iiii", "oiii", "ooii", "oooi", "oooo"}, false},
		{[]string{"oiii", "ioii", "iioi", "iiio"}, true},
	}
	for _, c := range cases {
		got := CheckAnagram(c.in)
		if got != c.want {
			t.Errorf("CalcOccurencies(%v) == %t want %t", c.in, got, c.want)
		}
	}
}
