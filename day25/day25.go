package main

import (
	"fmt"
	"os"
	"strconv"
)

type nextState struct {
	stateName     string
	writeInput    int
	nextDirection direction
}

type direction string

const LEFT direction = "LEFT"
const RIGHT direction = "RIGHT"

func main() {
	caseSelection := os.Args[1]

	program := map[string]nextState{
		"A,0": nextState{"B", 1, RIGHT},
		"A,1": nextState{"C", 0, LEFT},
		"B,0": nextState{"A", 1, LEFT},
		"B,1": nextState{"D", 1, RIGHT},
		"C,0": nextState{"B", 0, LEFT},
		"C,1": nextState{"E", 0, LEFT},
		"D,0": nextState{"A", 1, RIGHT},
		"D,1": nextState{"B", 0, RIGHT},
		"E,0": nextState{"F", 1, LEFT},
		"E,1": nextState{"C", 1, LEFT},
		"F,0": nextState{"D", 1, RIGHT},
		"F,1": nextState{"A", 1, RIGHT},
	}

	switch {
	case caseSelection == "A":
		fmt.Printf("The diagnostic number after %d steps is %d\n", 12481997, SolveA(12481997, program))
	default:
		fmt.Printf("Invalid Selection, possible values: A\n")
	}
}

func SolveA(steps int, program map[string]nextState) int {
	tape := []int{0}
	currState := "A"
	cursorPosition := 0

	for step := 0; step < steps; step++ {
		currValue := tape[cursorPosition]
		next := program[getStateKey(currState, currValue)]
		tape[cursorPosition] = next.writeInput
		currState = next.stateName
		cursorPosition, tape = moveCursor(next.nextDirection, cursorPosition, tape)
	}

	result := 0
	for _, tapeCell := range tape {
		if tapeCell == 1 {
			result++
		}
	}
	return result
}

func moveCursor(dir direction, currentCursorPosition int, tape []int) (int, []int) {
	switch dir {
	case LEFT:
		if currentCursorPosition == 0 {
			return 0, append([]int{0}, tape...)
		}
		return currentCursorPosition - 1, tape
	case RIGHT:
		if currentCursorPosition == len(tape)-1 {
			return currentCursorPosition + 1, append(tape, 0)
		}
		return currentCursorPosition + 1, tape
	default:
		panic("Unsupported direction!")
	}
}

func getStateKey(stateName string, value int) string {
	return stateName + "," + strconv.Itoa(value)
}
