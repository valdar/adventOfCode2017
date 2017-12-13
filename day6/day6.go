package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/valdar/adventOfCode2017/utils"
)

func main() {
	caseSelection := os.Args[1]
	f, err := os.Open(os.Args[2])
	defer f.Close()
	utils.Check(err)

	line, err := utils.ReadLine(bufio.NewReader(f))
	utils.Check(err)

	//discard endline
	input := line[:len(line)-1]
	parts := strings.Split(input, " ")
	initialConfig := []int{}
	for _, part := range parts {
		stringPart, err := strconv.Atoi(part)
		utils.Check(err)
		initialConfig = append(initialConfig, stringPart)
	}

	switch {
	case caseSelection == "A":
		result, _ := CalcStepsSameConfig(initialConfig)
		fmt.Printf("The number of steps to see an already seen config are %d\n", result)
	case caseSelection == "B":
		_, alreadySeen := CalcStepsSameConfig(initialConfig)
		result, _ := CalcStepsSameConfig(alreadySeen)
		fmt.Printf("The number of steps in the infinite loop are %d\n", result)
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func CalcStepsSameConfig(input []int) (int, []int) {
	steps := 0

	seenPositions := map[string]bool{}

	currPosition := input

	for !seenPositions[positionToString(currPosition)] {
		seenPositions[positionToString(currPosition)] = true
		//find block to redistribute
		maxMemBlocks := 0
		for _, memBlocks := range currPosition {
			if memBlocks > maxMemBlocks {
				maxMemBlocks = memBlocks
			}
		}
		minIndexWithMaxValue := len(currPosition)
		for index, memBlocks := range currPosition {
			if memBlocks == maxMemBlocks && index < minIndexWithMaxValue {
				minIndexWithMaxValue = index
			}
		}

		//redstribute in currPosition
		currPosition[minIndexWithMaxValue] = 0
		for i := 1; i <= maxMemBlocks; i++ {
			currPosition[(minIndexWithMaxValue+i)%len(currPosition)]++
		}
		steps++
	}
	return steps, currPosition
}

func positionToString(position []int) string {
	result := ""
	for _, curr := range position {
		result += strconv.Itoa(curr)
	}
	return result
}
