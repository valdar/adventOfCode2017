package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/valdar/adventOfCode2017/utils"
)

type coordinates struct {
	i int
	j int
}

func main() {
	caseSelection := os.Args[1]
	input, err := strconv.Atoi(os.Args[2])
	utils.Check(err)

	switch {
	case caseSelection == "A":
		fmt.Printf("The value to shortcircuit the vortex is %d\n", FindShortCircuitValue(CalcSpinning(input, 2017), 2017))
	case caseSelection == "B":
		fmt.Println("Calculating, pleas wait...")
		fmt.Printf("The value to shortcircuit FOR GOOD the vortex is %d\n", SolveB(input, 50000000, 0))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func FindShortCircuitValue(circularBuffer []int, valueBeforeShortCircuit int) int {
	for index, value := range circularBuffer {
		if value == valueBeforeShortCircuit {
			return circularBuffer[(index+1)%len(circularBuffer)]
		}
	}
	return -1
}

func CalcSpinning(stepForward int, lastNumberAddedInTheSpinns int) []int {
	result := []int{0}
	currPosition := 0
	for i := 1; i <= lastNumberAddedInTheSpinns; i++ {
		fmt.Printf("Calculating the [%d] spin of [%d]\r", i, lastNumberAddedInTheSpinns)
		endingSpinnPosition := (stepForward + currPosition) % len(result)
		firstHalf := append([]int{}, result[:endingSpinnPosition+1]...)
		secondHalf := []int{}
		if endingSpinnPosition+1 < len(result) {
			secondHalf = append(secondHalf, result[endingSpinnPosition+1:]...)
		}

		result = append([]int{}, firstHalf...)
		result = append(result, i)
		result = append(result, secondHalf...)
		currPosition = endingSpinnPosition + 1
	}
	return result
}

func SolveB(stepForward int, lastNumberAddedInTheSpinns int, target int) int {
	lastPassTo0 := 0
	currPosition := 0
	for i := 1; i <= lastNumberAddedInTheSpinns; i++ {

		endingSpinnPosition := (stepForward + currPosition) % i

		if endingSpinnPosition == target {
			lastPassTo0 = i
		}

		currPosition = endingSpinnPosition + 1
	}
	return lastPassTo0
}
