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

type Program struct {
	name       string
	weight     int
	successors []string
}

type programTreeWeight struct {
	program    Program
	treeWeight int
}

func main() {
	caseSelection := os.Args[1]
	f, err := os.Open(os.Args[2])
	defer f.Close()
	utils.Check(err)
	br := bufio.NewReader(f)

	programsIndex := parse(br)

	switch {
	case caseSelection == "A":
		fmt.Printf("The starting program name is %v\n", FindStartProgram(programsIndex))
	case caseSelection == "B":
		fmt.Printf("The weight of the unbalancing program needs to be %d\n", FindTheImbalanceDifference(programsIndex))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func FindTheImbalanceDifference(programsIndex map[string]Program) int {
	for k := range programsIndex {
		currProgram := programsIndex[k]
		treeWeights := []programTreeWeight{}
		for _, currSuccessorName := range currProgram.successors {
			treeWeights = append(treeWeights, programTreeWeight{programsIndex[currSuccessorName], calcWeightOfTreeStartingAt(currSuccessorName, programsIndex)})
		}

		if len(treeWeights) > 0 {
			groupA := []programTreeWeight{}
			groupB := []programTreeWeight{}

			firstElement := treeWeights[0]
			groupA = append(groupA, firstElement)
			firstElementTotal := firstElement.program.weight + firstElement.treeWeight

			for _, currTreeWeight := range treeWeights[1:] {
				if firstElementTotal != currTreeWeight.program.weight+currTreeWeight.treeWeight {
					groupB = append(groupB, currTreeWeight)
				} else {
					groupA = append(groupA, currTreeWeight)
				}
			}

			if len(groupB) > 0 {
				if len(groupA) > len(groupB) {
					return firstElementTotal - groupB[0].treeWeight
				} else {
					return groupB[0].treeWeight + groupB[0].program.weight - groupA[0].treeWeight
				}
			}

		}
	}
	panic("Inbalance not found!")
}

func FindStartProgram(programsIndex map[string]Program) string {
	successors := map[string]bool{}
	for k := range programsIndex {
		currProgram := programsIndex[k]

		for _, currSuccessorName := range currProgram.successors {
			successors[currSuccessorName] = true
		}
	}

	programRoot := []string{}
	for k := range programsIndex {
		if !successors[k] {
			programRoot = append(programRoot, k)
		}
	}

	if len(programRoot) != 1 {
		panic("There should only be one root Program!")
	}
	return programRoot[0]
}

func calcWeightOfTreeStartingAt(start string, programsIndex map[string]Program) int {
	treeWeight := 0
	toBeInspected := programsIndex[start].successors
	for len(toBeInspected) > 0 {
		currProgram := programsIndex[toBeInspected[0]]
		treeWeight += currProgram.weight
		toBeInspected = append(toBeInspected[1:], currProgram.successors...)
	}
	return treeWeight
}

func parse(br *bufio.Reader) map[string]Program {
	result := make(map[string]Program)
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

		parts := strings.Split(input, "->")

		newProgram := new(Program)
		if len(parts) >= 1 {
			//fill in name and weight
			subParts := strings.Split(strings.TrimSpace(parts[0]), " ")
			newProgram.name = strings.TrimSpace(subParts[0])
			weigthWithBrackets := strings.TrimSpace(subParts[1])
			strconv.Atoi(weigthWithBrackets[1 : len(weigthWithBrackets)-1])
			weight, err := strconv.Atoi(weigthWithBrackets[1 : len(weigthWithBrackets)-1])
			utils.Check(err)
			newProgram.weight = weight
			newProgram.successors = []string{}
		}
		if len(parts) == 2 {
			//fill in successors
			subParts := strings.Split(strings.TrimSpace(parts[1]), ",")
			for _, currSuccessorName := range subParts {
				newProgram.successors = append(newProgram.successors, strings.TrimSpace(currSuccessorName))
			}
		}
		if len(parts) > 2 {
			panic("Not parable input!")
		}
		result[newProgram.name] = *newProgram
	}
	return result
}
