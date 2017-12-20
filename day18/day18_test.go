package main

import (
	"testing"
)

func TestSolveA(t *testing.T) {
	cases := []struct {
		inProgram []instruction
		want      int
	}{
		{[]instruction{
			{"set", 0, "a", 1, "", false},
			{"add", 0, "a", 2, "", false},
			{"mul", 0, "a", 0, "a", false},
			{"mod", 0, "a", 5, "", false},
			{"snd", 0, "a", 0, "", true},
			{"set", 0, "a", 0, "", false},
			{"rcv", 0, "a", 0, "", true},
			{"jgz", 0, "a", -1, "", false},
			{"set", 0, "a", 1, "", false},
			{"jgz", 0, "a", -2, "", false},
		}, 4},
	}
	for _, c := range cases {
		got := SolveA(c.inProgram, map[string]int{}, "RECOVERY")
		if got != c.want {
			t.Errorf("SolveA( %v ) = %d expected %d", c.inProgram, got, c.want)
		}
	}
}

func TestSolveB(t *testing.T) {
	cases := []struct {
		inProgram []instruction
		want      int
	}{
		{[]instruction{
			{"snd", 0, "a", 0, "", true},
			{"snd", 0, "b", 0, "", true},
			{"snd", 0, "p", 0, "", true},
			{"rcv", 0, "a", 0, "", true},
			{"rcv", 0, "b", 0, "", true},
			{"rcv", 0, "c", 0, "", true},
			{"rcv", 0, "d", 0, "", true},
		}, 3},
	}
	for _, c := range cases {
		got := SolveB(c.inProgram)
		if got != c.want {
			t.Errorf("SolveB( %v ) = %d expected %d", c.inProgram, got, c.want)
		}
	}
}
