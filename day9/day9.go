package main

import (
	"bufio"
	"fmt"
	"os"
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

	inputSlice := make([]int, 256, 256)
	for i := 0; i < 256; i++ {
		inputSlice[i] = i
	}

	switch {
	case caseSelection == "A":
		result, _ := CaclScoreAndGarbage(input)
		fmt.Printf("The score is %d\n", result)
	case caseSelection == "B":
		_, result := CaclScoreAndGarbage(input)
		fmt.Printf("The number of garbage char is %d\n", result)
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func CaclScoreAndGarbage(input string) (int, int) {
	const openCurl = '{'
	const closedCurl = '}'
	const openAngular = '<'
	const closedAngular = '>'
	const bang = '!'

	const INIT = "INIT"
	const OPEN_CURL = "OPEN_CURL"
	const OPEN_ANGULAR = "OPEN_ANGULAR"

	score := 0
	garbageChar := 0
	currPoints := 0
	state := INIT
	var remainingChar []rune
	for _, str := range strings.Split(input, "") {
		remainingChar = append(remainingChar, []rune(str)...)
	}
	for len(remainingChar) > 0 {
		currChar := remainingChar[0]
		if currChar == bang {
			//skip next char
			remainingChar = append([]rune{}, remainingChar[1:]...)
		} else {
			switch state {
			case INIT:
				if currChar == openCurl {
					state = OPEN_CURL
					currPoints++
				}
			case OPEN_CURL:
				if currChar == openCurl {
					currPoints++
				} else if currChar == closedCurl {
					score += currPoints
					currPoints--
					if currPoints == 0 {
						state = INIT
					}
				} else if currChar == openAngular {
					state = OPEN_ANGULAR
				}
			case OPEN_ANGULAR:
				if currChar == closedAngular {
					state = OPEN_CURL
				} else {
					garbageChar++
				}
			default:
				panic("Not recognized state!")
			}
		}
		//remove currChar
		remainingChar = append([]rune{}, remainingChar[1:]...)
	}
	return score, garbageChar
}
