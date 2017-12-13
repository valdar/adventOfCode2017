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

func main() {
	caseSelection := os.Args[1]
	f, err := os.Open(os.Args[2])
	defer f.Close()
	utils.Check(err)
	br := bufio.NewReader(f)

	switch {
	case caseSelection == "A":
		fmt.Printf("Number of steps to exit is %d\n", CalcStepsOutA(parse(br)))
	case caseSelection == "B":
		fmt.Printf("Number of steps to exit is %d\n", CalcStepsOutB(parse(br)))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func CalcStepsOutA(input []int) int {
	steps := 0
	position := 0
	for position < len(input) {
		currPosition := position
		position += input[position]
		input[currPosition]++
		steps++
	}
	return steps
}

func CalcStepsOutB(input []int) int {
	steps := 0
	position := 0
	for position < len(input) {
		currPosition := position
		position += input[position]
		if input[currPosition] >= 3 {
			input[currPosition]--
		} else {
			input[currPosition]++
		}
		steps++
	}
	return steps
}

func parse(br *bufio.Reader) []int {
	result := []int{}
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

		elemnt, err := strconv.Atoi(strings.TrimSpace(input))
		utils.Check(err)

		result = append(result, elemnt)
	}
	return result
}
