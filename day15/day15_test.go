package main

import "testing"

func TestCalcJudgeMatches(t *testing.T) {
	cases := []struct {
		inStartA int
		inStartB int
		want     int
	}{
		{65, 8921, 588},
	}
	for _, c := range cases {
		got := CalcJudgeMatches(c.inStartA, c.inStartB)
		if got != c.want {
			t.Errorf("CalcJudgeMatches(%d, %d) == %d want %d", c.inStartA, c.inStartB, got, c.want)
		}
	}
}

func TestCalcJudgeMatchesSecondPart(t *testing.T) {
	cases := []struct {
		inStartA int
		inStartB int
		want     int
	}{
		{65, 8921, 309},
	}
	for _, c := range cases {
		got := CalcJudgeMatchesSecondPart(c.inStartA, c.inStartB)
		if got != c.want {
			t.Errorf("CalcJudgeMatchesSecondPart(%d, %d) == %d want %d", c.inStartA, c.inStartB, got, c.want)
		}
	}
}
