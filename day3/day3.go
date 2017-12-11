package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/valdar/adventOfCode2017/utils"
)

func main() {
	caseSelection := os.Args[1]
	input, err := strconv.Atoi(os.Args[2])
	utils.Check(err)

	switch {
	case caseSelection == "A":
		position, layer := CalcSpiralLayerAndPosition(input)
		fmt.Printf("The desired data sector is in position %d of the spiral layer %d\n", position, layer)
		x, y := CalcCoordinates(position, layer)
		fmt.Printf("Which means, plottin in a cartesian graph, position %d,%d\n", x, y)
		fmt.Printf("Data from square %d is carried %d\n", input, calcSteps(x, y, layer))
	case caseSelection == "B":
		//fmt.Printf("The checksum is %d\n", calcCS(CalcCheckSumByEvenDivision, br))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func CalcSpiralLayerAndPosition(input int) (position int, layer int) {
	if input <= 0 {
		return 0, 0
	}
	if input == 1 {
		return 1, 1
	}

	currentIteration := 1
	currentPosition := input - 1
	for {
		if currentPosition <= 8*currentIteration {
			return currentPosition, currentIteration + 1
		}
		currentPosition -= 8 * currentIteration
		currentIteration++
	}
}

func CalcCoordinates(position int, layer int) (x int, y int) {
	sideSize := layer + (layer - 1)
	if position >= 1 && position <= sideSize-1 {
		return sideSize, position + 1
	} else if position > sideSize-1 && position <= 2*sideSize-2 {
		return (2*sideSize - 2) - position + 1, sideSize
	} else if position > 2*sideSize-2 && position <= 3*sideSize-3 {
		return 1, (3*sideSize - 3) - position + 1
	} else if position > 3*sideSize-3 && position <= 4*sideSize-4 {
		return position - (3*sideSize - 2) + 2, 1
	} else {
		panic("Invalid position and layer")
	}
}

func calcSteps(x int, y int, layer int) int {
	centerX := layer
	centerY := layer
	return utils.Abs(centerX-x) + utils.Abs(centerY-y)
}

//TODO: solve B
// func solveB(target int) int {
// 	directions := map[string]string{"RIGHT": "UP", "UP": "LEFT", "LEFT": "DOWN", "DOWN": "RIGHT"}
// 	m := make(map[string]int)
// 	m["00"] = 1
// 	startingDirection := "RIGHT"

// 	return 0
// }

// func generateNeighboors( x int, y int, direction string ) []string {
// 	switch direction
// }
