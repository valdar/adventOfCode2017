package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/valdar/adventOfCode2017/utils"
)

type instruction struct {
	operation             string
	firstOperandNumeric   int
	firstOperandRegistry  string
	secondOperandNumeric  int
	secondOperandRegistry string
	oneOperand            bool
}

func main() {
	caseSelection := os.Args[1]
	f, err := os.Open(os.Args[2])
	defer f.Close()
	utils.Check(err)
	br := bufio.NewReader(f)

	maze := parse(br)

	switch {
	case caseSelection == "A":
		result, _ := WalkTheMaze(maze)
		fmt.Printf("The letters collected walking the maze are %v\n", result)
	case caseSelection == "B":
		_, result := WalkTheMaze(maze)
		fmt.Printf("The number of steps walking the maze are %d\n", result)
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func WalkTheMaze(maze [][]string) (string, int) {
	const VERT_ROAD = "|"
	const ORIZ_ROAD = "-"
	const JUNCTION = "+"
	const EMPTY = " "
	const DOWN = "DOWN"
	const UP = "UP"
	const LEFT = "LEFT"
	const RIGHT = "RIGHT"

	collectedLetters := ""
	steps := 0

	//find starting position
	currY := 0
	var currX int
	for x, content := range maze[0] {
		if content == VERT_ROAD {
			currX = x
		}
	}
	currDirection := DOWN

	stop := false

	for !stop {
		//fmt.Printf("Current position [%d,%d] facing %v, current tile: [%v]\n", currX, currY, currDirection, maze[currY][currX])
		steps++
		switch currDirection {
		case DOWN:
			if currY+1 >= len(maze) {
				//we have nowhere to go
				stop = true
			} else {
				nextContent := maze[currY+1][currX]
				if nextContent == VERT_ROAD || nextContent == ORIZ_ROAD {
					currY++
				} else if nextContent == JUNCTION {
					stop, currX, currY, currDirection = exploreJunction(maze, currX, currY+1, currDirection, &collectedLetters, &steps)
				} else if nextContent != EMPTY {
					//A letter!
					currY++
					collectedLetters += nextContent
				} else {
					//we have nowhere to go
					stop = true
				}
			}
		case UP:
			if currY-1 < 0 {
				//we have nowhere to go
				stop = true
			} else {
				nextContent := maze[currY-1][currX]
				if nextContent == VERT_ROAD || nextContent == ORIZ_ROAD {
					currY--
				} else if nextContent == JUNCTION {
					stop, currX, currY, currDirection = exploreJunction(maze, currX, currY-1, currDirection, &collectedLetters, &steps)
				} else if nextContent != EMPTY {
					//A letter!
					currY--
					collectedLetters += nextContent
				} else {
					//we have nowhere to go
					stop = true
				}
			}
		case RIGHT:
			if currX+1 >= len(maze[currY]) {
				//we have nowhere to go
				stop = true
			} else {
				nextContent := maze[currY][currX+1]
				if nextContent == ORIZ_ROAD || nextContent == VERT_ROAD {
					currX++
				} else if nextContent == JUNCTION {
					stop, currX, currY, currDirection = exploreJunction(maze, currX+1, currY, currDirection, &collectedLetters, &steps)
				} else if nextContent != EMPTY {
					//A letter!
					currX++
					collectedLetters += nextContent
				} else {
					//we have nowhere to go
					stop = true
				}
			}
		case LEFT:
			if currX-1 < 0 {
				//we have nowhere to go
				stop = true
			} else {
				nextContent := maze[currY][currX-1]
				if nextContent == ORIZ_ROAD || nextContent == VERT_ROAD {
					currX--
				} else if nextContent == JUNCTION {
					stop, currX, currY, currDirection = exploreJunction(maze, currX-1, currY, currDirection, &collectedLetters, &steps)
				} else if nextContent != EMPTY {
					//A letter!
					currX--
					collectedLetters += nextContent
				} else {
					//we have nowhere to go
					stop = true
				}
			}
		default:
			panic("Unrecognized direction!")
		}
	}
	return collectedLetters, steps
}

func exploreJunction(maze [][]string, currX int, currY int, currDirection string, collectedLetters *string, steps *int) (bool, int, int, string) {
	//fmt.Printf("Explore JUNCTION at [%d,%d] facing %v, junction tile: [%v]\n", currX, currY, currDirection, maze[currY][currX])
	*steps += 1

	const VERT_ROAD = "|"
	const ORIZ_ROAD = "-"
	const JUNCTION = "+"
	const EMPTY = " "
	const DOWN = "DOWN"
	const UP = "UP"
	const LEFT = "LEFT"
	const RIGHT = "RIGHT"

	viableStepCount := 0

	var nextX, nextY int
	var nextDir string
	switch currDirection {
	case DOWN, UP:
		if currX-1 >= 0 && maze[currY][currX-1] == ORIZ_ROAD {
			viableStepCount++
			nextX = currX - 1
			nextY = currY
			nextDir = LEFT
		}
		if currX+1 < len(maze[currY]) && maze[currY][currX+1] == ORIZ_ROAD {
			viableStepCount++
			nextX = currX + 1
			nextY = currY
			nextDir = RIGHT
		}
		if currX-1 >= 0 && maze[currY][currX-1] == JUNCTION {
			return exploreJunction(maze, currX-1, currY, LEFT, collectedLetters, steps)
		}
		if currX+1 < len(maze[currY]) && maze[currY][currX+1] == JUNCTION {
			return exploreJunction(maze, currX+1, currY, RIGHT, collectedLetters, steps)
		}
		if currX-1 >= 0 && maze[currY][currX-1] != EMPTY && maze[currY][currX-1] != ORIZ_ROAD && maze[currY][currX-1] != VERT_ROAD {
			//A letter!
			*collectedLetters += maze[currY][currX-1]
			viableStepCount++
			nextX = currX - 1
			nextY = currY
			nextDir = LEFT
		}
		if currX+1 < len(maze[currY]) && maze[currY][currX+1] != EMPTY && maze[currY][currX+1] != ORIZ_ROAD && maze[currY][currX+1] != VERT_ROAD {
			//A letter!
			*collectedLetters += maze[currY][currX+1]
			viableStepCount++
			nextX = currX + 1
			nextY = currY
			nextDir = RIGHT
		}
	case LEFT, RIGHT:
		if currY-1 >= 0 && maze[currY-1][currX] == VERT_ROAD {
			viableStepCount++
			nextX = currX
			nextY = currY - 1
			nextDir = UP
		}
		if currY+1 < len(maze) && maze[currY+1][currX] == VERT_ROAD {
			viableStepCount++
			nextX = currX
			nextY = currY + 1
			nextDir = DOWN
		}
		if currY-1 >= 0 && maze[currY-1][currX] == JUNCTION {
			return exploreJunction(maze, currX, currY-1, UP, collectedLetters, steps)
		}
		if currY+1 < len(maze) && maze[currY+1][currX] == JUNCTION {
			return exploreJunction(maze, currX, currY+1, DOWN, collectedLetters, steps)
		}
		if currY-1 >= 0 && maze[currY-1][currX] != ORIZ_ROAD && maze[currY-1][currX] != VERT_ROAD && maze[currY-1][currX] != EMPTY {
			//A letter!
			*collectedLetters += maze[currY-1][currX]
			viableStepCount++
			nextX = currX
			nextY = currY - 1
			nextDir = UP
		}
		if currY+1 < len(maze) && maze[currY+1][currX] != ORIZ_ROAD && maze[currY+1][currX] != VERT_ROAD && maze[currY+1][currX] != EMPTY {
			//A letter!
			*collectedLetters += maze[currY+1][currX]
			viableStepCount++
			nextX = currX
			nextY = currY + 1
			nextDir = DOWN
		}
	default:
		panic("Unrecognized direction!")
	}

	if viableStepCount == 0 {
		//we have nowhere to go
		return true, currX, currY, currDirection
	} else if viableStepCount == 1 {
		return false, nextX, nextY, nextDir
	} else {
		panic("Undecidible junction!")
	}
}

func parse(br *bufio.Reader) [][]string {
	maze := [][]string{}
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

		parts := strings.Split(input, "")
		maze = append(maze, parts)
	}
	return maze
}
