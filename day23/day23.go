package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/valdar/adventOfCode2017/day18/instructions"
	"github.com/valdar/adventOfCode2017/utils"
)

func main() {
	caseSelection := os.Args[1]

	switch {
	case caseSelection == "A":
		f, err := os.Open(os.Args[2])
		defer f.Close()
		utils.Check(err)
		br := bufio.NewReader(f)
		program := instructions.Parse(br)
		fmt.Printf("The number of tims a mul is executed is %d\n", SolveA(program, map[string]int{}))
	case caseSelection == "B":
		fmt.Printf("The value of registry h is %d\n", SolveB())
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func SolveA(program []instructions.Instruction, registries map[string]int) int {
	result := 0
	pc := 0

	for pc >= 0 && pc < len(program) {
		currInstruction := program[pc]
		secondOperandResolved := currInstruction.SecondOperandNumeric
		if currInstruction.SecondOperandRegistry != "" {
			secondOperandResolved = registries[currInstruction.SecondOperandRegistry]
		}

		switch currInstruction.Operation {
		case "set":
			registries[currInstruction.FirstOperandRegistry] = secondOperandResolved
			pc++
		case "sub":
			registries[currInstruction.FirstOperandRegistry] -= secondOperandResolved
			pc++
		case "mul":
			registries[currInstruction.FirstOperandRegistry] *= secondOperandResolved
			pc++
			result++
		case "jnz":
			firstOperandResolved := currInstruction.FirstOperandNumeric
			if currInstruction.FirstOperandRegistry != "" {
				firstOperandResolved = registries[currInstruction.FirstOperandRegistry]
			}
			if firstOperandResolved != 0 {
				pc += secondOperandResolved
			} else {
				pc++
			}
		}
	}

	return result
}

func SolveB() int {
	result := 0
	for i := (65 * 100) + 100000; i <= (65*100)+100000+17000; i += 17 {
		prime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if !prime {
			result++
		}
	}
	return result
}
