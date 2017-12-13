package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	instructions := parse(br)

	switch {
	case caseSelection == "A":
		registryState, _ := ComputeInstructions(instructions)
		fmt.Printf("The max value in registries at the end of instructions executione is %v\n", CalcMax(registryState))
	case caseSelection == "B":
		_, maxOverall := ComputeInstructions(instructions)
		fmt.Printf("The max value reached during execution of instructions is %v\n", maxOverall)
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func ComputeInstructions(instructions [][]string) (map[string]int, int) {
	result := map[string]int{}
	maxDuringCompuitation := math.MinInt32
	for _, currInstr := range instructions {
		condValue, err := strconv.Atoi(strings.TrimSpace(currInstr[6]))
		utils.Check(err)
		switch currInstr[5] {
		case ">":
			if result[currInstr[4]] > condValue {
				performRegistryOperation(result, currInstr)
			}
		case "<":
			if result[currInstr[4]] < condValue {
				performRegistryOperation(result, currInstr)
			}
		case "==":
			if result[currInstr[4]] == condValue {
				performRegistryOperation(result, currInstr)
			}
		case "!=":
			if result[currInstr[4]] != condValue {
				performRegistryOperation(result, currInstr)
			}
		case "<=":
			if result[currInstr[4]] <= condValue {
				performRegistryOperation(result, currInstr)
			}
		case ">=":
			if result[currInstr[4]] >= condValue {
				performRegistryOperation(result, currInstr)
			}
		default:
			panic("Operation not supported!")
		}
		if result[currInstr[0]] > maxDuringCompuitation {
			maxDuringCompuitation = result[currInstr[0]]
		}
	}
	return result, maxDuringCompuitation
}

func performRegistryOperation(registries map[string]int, operation []string) {
	incDecValue, err := strconv.Atoi(strings.TrimSpace(operation[2]))
	utils.Check(err)
	if operation[1] == "dec" {
		registries[operation[0]] -= incDecValue
	} else if operation[1] == "inc" {
		registries[operation[0]] += incDecValue
	} else {
		panic("Unsupported operation!")
	}
}

func CalcMax(registriesState map[string]int) int {
	max := math.MinInt32
	for _, v := range registriesState {
		if v > max {
			max = v
		}
	}
	return max
}

func parse(br *bufio.Reader) [][]string {
	result := [][]string{}
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

		parts := strings.Split(input, " ")

		result = append(result, parts)

	}
	return result
}
