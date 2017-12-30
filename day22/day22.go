package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/valdar/adventOfCode2017/utils"
)

const UP string = "UP"
const DOWN string = "DOWN"
const LEFT string = "LEFT"
const RIGHT string = "RIGHT"

var turnLeftMap map[string]string = map[string]string{
	UP:    LEFT,
	LEFT:  DOWN,
	DOWN:  RIGHT,
	RIGHT: UP,
}
var turnRightMap map[string]string = map[string]string{
	UP:    RIGHT,
	RIGHT: DOWN,
	DOWN:  LEFT,
	LEFT:  UP,
}
var reverseirectionMap map[string]string = map[string]string{
	UP:    DOWN,
	RIGHT: LEFT,
	DOWN:  UP,
	LEFT:  RIGHT,
}

const INFECTED string = "INFECTED"
const CLEAN string = "CLEAN"
const WEAKENED string = "WEAKENED"
const FLAGGED string = "FLAGGED"

var stateChangeMap map[string]string = map[string]string{
	INFECTED: FLAGGED,
	CLEAN:    WEAKENED,
	WEAKENED: INFECTED,
	FLAGGED:  CLEAN,
}

func main() {
	caseSelection := os.Args[1]
	f, err := os.Open(os.Args[2])
	defer f.Close()
	utils.Check(err)
	br := bufio.NewReader(f)

	switch {
	case caseSelection == "A":
		gridCenter, gridRows := parseA(br)
		fmt.Printf("The number of new infected nodes after %d bursts is %d\n", 10000, SolveA(gridCenter, gridRows/2, gridRows/2, 10000))
	case caseSelection == "B":
		gridCenter, gridRows := parseB(br)
		fmt.Printf("The number of new infected nodes after %d bursts is %d\n", 10000000, SolveB(gridCenter, gridRows/2, gridRows/2, 10000000))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func SolveA(startingGrid map[string]bool, startX int, startY int, numberOfBurst int) int {
	carrierCurrX := startX
	carrierCurrY := startY
	carrierCurrOrientation := UP
	result := 0

	for i := 0; i < numberOfBurst; i++ {
		if startingGrid[getKey(carrierCurrX, carrierCurrY)] {
			carrierCurrOrientation = turnRightMap[carrierCurrOrientation]
			startingGrid[getKey(carrierCurrX, carrierCurrY)] = false
		} else {
			carrierCurrOrientation = turnLeftMap[carrierCurrOrientation]
			startingGrid[getKey(carrierCurrX, carrierCurrY)] = true
			result++
		}

		carrierCurrX, carrierCurrY = getNextPosition(carrierCurrOrientation, carrierCurrX, carrierCurrY)
	}

	return result
}

func SolveB(startingGrid map[string]string, startX int, startY int, numberOfBurst int) int {
	carrierCurrX := startX
	carrierCurrY := startY
	carrierCurrOrientation := UP
	result := 0

	for i := 0; i < numberOfBurst; i++ {
		curState, ok := startingGrid[getKey(carrierCurrX, carrierCurrY)]
		if !ok {
			startingGrid[getKey(carrierCurrX, carrierCurrY)] = CLEAN
			curState = CLEAN
		}

		switch curState {
		case CLEAN:
			carrierCurrOrientation = turnLeftMap[carrierCurrOrientation]
		case WEAKENED:
			//nothing to do
		case INFECTED:
			carrierCurrOrientation = turnRightMap[carrierCurrOrientation]
		case FLAGGED:
			carrierCurrOrientation = reverseirectionMap[carrierCurrOrientation]
		}

		startingGrid[getKey(carrierCurrX, carrierCurrY)] = stateChangeMap[startingGrid[getKey(carrierCurrX, carrierCurrY)]]

		if startingGrid[getKey(carrierCurrX, carrierCurrY)] == INFECTED {
			result++
		}

		carrierCurrX, carrierCurrY = getNextPosition(carrierCurrOrientation, carrierCurrX, carrierCurrY)
	}

	return result
}

func parseA(br *bufio.Reader) (map[string]bool, int) {
	grid := map[string]bool{}
	i := 0
	for {
		line, err := utils.ReadLine(br)
		if err != nil {
			if err == io.EOF {
				//file is ended
				break
			}
			panic(err)
		}
		//discard endline
		input := line[:len(line)-1]

		nodes := strings.Split(input, "")
		for j, node := range nodes {
			if node == "#" {
				grid[getKey(i, j)] = true
			} else {
				grid[getKey(i, j)] = false
			}
		}
		i++
	}
	return grid, i
}

func getNextPosition(carrierCurrOrientation string, x int, y int) (int, int) {
	switch carrierCurrOrientation {
	case UP:
		return x - 1, y
	case DOWN:
		return x + 1, y
	case LEFT:
		return x, y - 1
	case RIGHT:
		return x, y + 1
	default:
		panic("Not Recognized direction!")
	}
}

func parseB(br *bufio.Reader) (map[string]string, int) {
	grid := map[string]string{}
	i := 0
	for {
		line, err := utils.ReadLine(br)
		if err != nil {
			if err == io.EOF {
				//file is ended
				break
			}
			panic(err)
		}
		//discard endline
		input := line[:len(line)-1]

		nodes := strings.Split(input, "")
		for j, node := range nodes {
			if node == "#" {
				grid[getKey(i, j)] = INFECTED
			} else {
				grid[getKey(i, j)] = CLEAN
			}
		}
		i++
	}
	return grid, i
}

func getKey(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}
