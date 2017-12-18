package main

import "testing"

func TestsolveA(t *testing.T) {
	cases := []struct {
		inStepforward     int
		inLastAddedNumber int
		want              int
	}{
		{3, 2017, 638},
	}
	for _, c := range cases {
		got := FindShortCircuitValue(CalcSpinning(c.inStepforward, c.inLastAddedNumber), c.inLastAddedNumber)
		if got != c.want {
			t.Errorf("FindShortCircuitValue(CalcSpinning(%d,%d), %d) == %d want %d", c.inStepforward, c.inLastAddedNumber, c.inLastAddedNumber, got, c.want)
		}
	}
}

//FindShortCircuitValue(CalcSpinning(input, 2017), 2017))
