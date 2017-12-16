package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/valdar/adventOfCode2017/utils"
)

type coordinates struct {
	i int
	j int
}

func main() {
	caseSelection := os.Args[1]
	f, err := os.Open(os.Args[2])
	defer f.Close()
	utils.Check(err)

	line, err := utils.ReadLine(bufio.NewReader(f))
	utils.Check(err)

	//discard endline
	input := line[:len(line)-1]

	dancePosition := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}
	switch {
	case caseSelection == "A":
		ApplyDanceSteps(Parse(input), dancePosition)
		fmt.Printf("Ths final position after the dance steps in input is %v\n", strings.Join(dancePosition, ""))
	case caseSelection == "B":
		fmt.Printf("Ths final position after a bilion time the dance steps in input is %v\n", ApplyDanceStepsConecutivelyAbilionTime(Parse(input), dancePosition))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func ApplyDanceStepsConecutivelyAbilionTime(steps []func([]string), startingPosition []string) string {
	seen := map[string]int{}
	for i := 0; i < 1000000000; i++ {
		if seenIndex, ok := seen[strings.Join(startingPosition, "")]; ok {
			fmt.Printf("\nFound already seen end sequence [of step %d] at step %d\n", seenIndex, i)
			indexTosearch := seenIndex + (1000000000 % (i - seenIndex))
			for k := range seen {
				if seen[k] == indexTosearch {
					return k
				}
			}
		}
		seen[strings.Join(startingPosition, "")] = i
		fmt.Printf("Performing dance number: [%d] of [1000000000]\r", i)
		for _, step := range steps {
			step(startingPosition)
		}
	}
	return strings.Join(startingPosition, "")
}

func ApplyDanceSteps(steps []func([]string), startingPosition []string) {
	for _, step := range steps {
		step(startingPosition)
	}
}

func Parse(input string) []func(input []string) {
	parts := strings.Split(input, ",")
	result := []func([]string){}
	for _, danceStep := range parts {
		subParts := strings.Split(danceStep, "")
		switch subParts[0] {
		case "s":
			spinnSize, err := strconv.Atoi(strings.Join(subParts[1:], ""))
			utils.Check(err)
			stepFunction := func(input []string) {
				lastSpinnsizeElements := append([]string{}, input[len(input)-spinnSize:]...)
				others := append([]string{}, input[:len(input)-spinnSize]...)
				copy(input[:len(lastSpinnsizeElements)], lastSpinnsizeElements)
				copy(input[len(lastSpinnsizeElements):], others)
			}
			result = append(result, stepFunction)
		case "x":
			operands := strings.Split(strings.Join(subParts[1:], ""), "/")
			swapIndex1, err1 := strconv.Atoi(operands[0])
			utils.Check(err1)
			swapIndex2, err2 := strconv.Atoi(operands[1])
			utils.Check(err2)
			stepFunction := func(input []string) {
				input[swapIndex1], input[swapIndex2] = input[swapIndex2], input[swapIndex1]
			}
			result = append(result, stepFunction)
		case "p":
			operands := strings.Split(strings.Join(subParts[1:], ""), "/")
			nameToSwap1 := operands[0]
			nameToSwap2 := operands[1]
			stepFunction := func(input []string) {
				// fmt.Printf("XXX: %v\n", input)
				// fmt.Printf("XXX: %v\n", nameToSwap1)
				// fmt.Printf("XXX: %v\n", nameToSwap2)

				index1 := -1
				index2 := -1
				for index, value := range input {

					if value == nameToSwap1 {
						index1 = index
					}
					if value == nameToSwap2 {
						index2 = index
					}
				}
				input[index1], input[index2] = input[index2], input[index1]
			}
			result = append(result, stepFunction)
		default:
			panic("Dance step not recognised!")
		}
	}
	return result
}
