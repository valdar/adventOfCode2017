package main

import "testing"

func TestCalcCheckSum(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{[]string{"5", "1", "9", "5"}, 8},
		{[]string{"7", "5", "3"}, 4},
		{[]string{"2", "4", "6", "8"}, 6},
	}
	for _, c := range cases {
		got := CalcCheckSum(c.in)
		if got != c.want {
			t.Errorf("CalcCheckSum(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestCalcCheckSumByEvenDivision(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{[]string{"5", "9", "2", "8"}, 4},
		{[]string{"9", "4", "7", "3"}, 3},
		{[]string{"3", "8", "6", "5"}, 2},
	}
	for _, c := range cases {
		got := CalcCheckSumByEvenDivision(c.in)
		if got != c.want {
			t.Errorf("CalcCheckSum(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
