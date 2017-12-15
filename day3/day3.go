package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/valdar/adventOfCode2017/utils"
)

type position struct {
	x    int
	y    int
	hash string
}

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
		fmt.Printf("The first greatest of %d value encoutered is %d\n", input, SolveB(input))
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

func SolveB(target int) int {
	const UP = "UP"
	const DOWN = "DOWN"
	const RIGHT = "RIGHT"
	const LEFT = "LEFT"

	m := make(map[string]int)

	currentPosition := position{0, 0, "0,0"}
	m[currentPosition.hash] = 1
	currDirection := DOWN
	currValue := 1

	for currValue <= target {
		var positionOnTheLeft position
		var newPositionOnTheLeft string
		switch currDirection {
		case UP:
			positionOnTheLeft = newPosition(currentPosition.x-1, currentPosition.y)
			newPositionOnTheLeft = LEFT
		case DOWN:
			positionOnTheLeft = newPosition(currentPosition.x+1, currentPosition.y)
			newPositionOnTheLeft = RIGHT
		case RIGHT:
			positionOnTheLeft = newPosition(currentPosition.x, currentPosition.y+1)
			newPositionOnTheLeft = UP
		case LEFT:
			positionOnTheLeft = newPosition(currentPosition.x, currentPosition.y-1)
			newPositionOnTheLeft = DOWN
		default:
			panic("Unrecognized position!")

		}

		var nextPosition position
		if m[positionOnTheLeft.hash] == 0 {
			nextPosition = positionOnTheLeft
			currDirection = newPositionOnTheLeft
		} else {
			switch currDirection {
			case UP:
				nextPosition = newPosition(currentPosition.x, currentPosition.y+1)
				currDirection = UP
			case DOWN:
				nextPosition = newPosition(currentPosition.x, currentPosition.y-1)
				currDirection = DOWN
			case RIGHT:
				nextPosition = newPosition(currentPosition.x+1, currentPosition.y)
				currDirection = RIGHT
			case LEFT:
				nextPosition = newPosition(currentPosition.x-1, currentPosition.y)
				currDirection = LEFT
			default:
				panic("Unrecognized position!")

			}
		}

		currValue = 0
		for _, currNeighbor := range generateNeighbors(nextPosition) {
			currValue += m[currNeighbor.hash]
		}
		m[nextPosition.hash] = currValue
		currentPosition = nextPosition
	}
	return currValue
}

func generateNeighbors(p position) []position {
	result := []position{}
	result = append(result, newPosition(p.x, p.y-1))
	result = append(result, newPosition(p.x, p.y+1))
	result = append(result, newPosition(p.x-1, p.y))
	result = append(result, newPosition(p.x+1, p.y))

	result = append(result, newPosition(p.x+1, p.y+1))
	result = append(result, newPosition(p.x-1, p.y-1))
	result = append(result, newPosition(p.x-1, p.y+1))
	result = append(result, newPosition(p.x+1, p.y-1))

	return result
}

func newPosition(x int, y int) position {
	return position{x, y, strconv.Itoa(x) + "," + strconv.Itoa(y)}
}
