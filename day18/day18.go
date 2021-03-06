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
	f, err := os.Open(os.Args[2])
	defer f.Close()
	utils.Check(err)
	br := bufio.NewReader(f)

	program := instructions.Parse(br)

	switch {
	case caseSelection == "A":
		registryStatus := map[string]int{}
		const recoveryRegistryName = "RECOVERY"
		fmt.Printf("The frequency of the recovere sound is %d\n", SolveA(program, registryStatus, recoveryRegistryName))
	case caseSelection == "B":
		fmt.Printf("The number of time program 1 send a value are %d\n", SolveB(program))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func SolveA(program []instructions.Instruction, registryStatus map[string]int, recoveryregistryname string) int {
	pc := 0

	for pc >= 0 && pc < len(program) {
		currInstruction := program[pc]
		executed, offsetNextInstruction := execInstruction(currInstruction, registryStatus, recoveryregistryname)
		if executed && currInstruction.Operation == "rcv" {
			return registryStatus[recoveryregistryname]
		}
		pc += offsetNextInstruction
	}
	return registryStatus[recoveryregistryname]
}

func SolveB(program []instructions.Instruction) int {
	pc0 := 0
	pc1 := 0

	registryStatus0 := map[string]int{"p": 0}
	registryStatus1 := map[string]int{"p": 1}

	queueTo0 := []int{}
	queueTo1 := []int{}

	consecutiveStops := 0

	hasProgressed0 := true
	hasProgressed1 := true

	program1EmittedValues := 0

	for {
		fmt.Printf("hasProgressed0: %t, hasProgressed1 %t, queueTo0: %d, queueTo1: %d\n", hasProgressed0, hasProgressed1, len(queueTo0), len(queueTo1))
		if !hasProgressed0 && !hasProgressed1 {
			consecutiveStops++
		} else {
			consecutiveStops = 0
		}

		hasProgressed0 = false
		hasProgressed1 = false

		if consecutiveStops >= 2 {
			return program1EmittedValues
		}

		running0 := true
		//run program 0 untill posible
		for pc0 >= 0 && pc0 < len(program) && running0 {
			currInstruction0 := program[pc0]
			var offsetNextInstruction0 int
			running0, offsetNextInstruction0 = execInstructionB(currInstruction0, registryStatus0, &queueTo1, &queueTo0)
			if offsetNextInstruction0 != 0 {
				hasProgressed0 = true
			}
			pc0 += offsetNextInstruction0
		}

		running1 := true
		//run program 1 untill posible
		for pc1 >= 0 && pc1 < len(program) && running1 {
			currInstruction1 := program[pc1]
			var offsetNextInstruction1 int
			running1, offsetNextInstruction1 = execInstructionB(currInstruction1, registryStatus1, &queueTo0, &queueTo1)
			if currInstruction1.Operation == "snd" {
				program1EmittedValues++
			}
			if offsetNextInstruction1 != 0 {
				hasProgressed1 = true
			}
			pc1 += offsetNextInstruction1
		}
	}
}

func execInstruction(currInstruction instructions.Instruction, registryStatus map[string]int, recoveryregistryname string) (bool, int) {
	if currInstruction.OneOperand {
		switch currInstruction.Operation {
		case "snd":
			var firstOperand int
			if currInstruction.FirstOperandRegistry == "" {
				firstOperand = currInstruction.FirstOperandNumeric
			} else {
				firstOperand = registryStatus[currInstruction.FirstOperandRegistry]
			}
			registryStatus[recoveryregistryname] = firstOperand
			return true, 1
		case "rcv":
			if currInstruction.FirstOperandRegistry == "" {
				panic("Malformed rcv operation!")
			} else if registryStatus[currInstruction.FirstOperandRegistry] != 0 {
				return true, 1
			} else {
				return false, 1
			}
		default:
			panic("Not supported single operand operation!")
		}
	} else {
		if currInstruction.Operation != "jgz" && currInstruction.FirstOperandRegistry == "" {
			panic("Malformed double operand set. add, mul, mod operation!")
		}

		firstOperandNum := currInstruction.FirstOperandNumeric
		firstOperandReg := currInstruction.FirstOperandRegistry
		var secondOperand int

		if currInstruction.SecondOperandRegistry == "" {
			secondOperand = currInstruction.SecondOperandNumeric
		} else {
			secondOperand = registryStatus[currInstruction.SecondOperandRegistry]
		}

		switch currInstruction.Operation {
		case "set":
			registryStatus[firstOperandReg] = secondOperand
			return true, 1
		case "add":
			registryStatus[firstOperandReg] += secondOperand
			return true, 1
		case "mul":
			registryStatus[firstOperandReg] *= secondOperand
			return true, 1
		case "mod":
			registryStatus[firstOperandReg] %= secondOperand
			return true, 1
		case "jgz":
			var firstOperandJgz int
			if firstOperandReg == "" {
				firstOperandJgz = firstOperandNum
			} else {
				firstOperandJgz = registryStatus[firstOperandReg]
			}
			if firstOperandJgz > 0 {
				return true, secondOperand
			} else {
				return false, 1
			}
		default:
			panic("Not supported two operand operation!")
		}
	}
}

func execInstructionB(currInstruction instructions.Instruction, registryStatus map[string]int, queueTo *[]int, queueFrom *[]int) (bool, int) {
	if currInstruction.OneOperand {
		switch currInstruction.Operation {
		case "snd":
			var firstOperand int
			if currInstruction.FirstOperandRegistry == "" {
				firstOperand = currInstruction.FirstOperandNumeric
			} else {
				firstOperand = registryStatus[currInstruction.FirstOperandRegistry]
			}
			*queueTo = append(*queueTo, firstOperand)
			return true, 1
		case "rcv":
			if currInstruction.FirstOperandRegistry == "" {
				panic("Malformed rcv operation!")
			} else if len(*queueFrom) > 0 {
				registryStatus[currInstruction.FirstOperandRegistry] = (*queueFrom)[0]
				*queueFrom = (*queueFrom)[1:]
				return true, 1
			} else {
				return false, 0
			}
		default:
			panic("Not supported single operand operation!")
		}
	} else {
		if currInstruction.Operation != "jgz" && currInstruction.FirstOperandRegistry == "" {
			panic("Malformed double operand set. add, mul, mod operation!")
		}

		firstOperandNum := currInstruction.FirstOperandNumeric
		firstOperandReg := currInstruction.FirstOperandRegistry
		var secondOperand int

		if currInstruction.SecondOperandRegistry == "" {
			secondOperand = currInstruction.SecondOperandNumeric
		} else {
			secondOperand = registryStatus[currInstruction.SecondOperandRegistry]
		}

		switch currInstruction.Operation {
		case "set":
			registryStatus[firstOperandReg] = secondOperand
			return true, 1
		case "add":
			registryStatus[firstOperandReg] += secondOperand
			return true, 1
		case "mul":
			registryStatus[firstOperandReg] *= secondOperand
			return true, 1
		case "mod":
			registryStatus[firstOperandReg] %= secondOperand
			return true, 1
		case "jgz":
			var firstOperandJgz int
			if firstOperandReg == "" {
				firstOperandJgz = firstOperandNum
			} else {
				firstOperandJgz = registryStatus[firstOperandReg]
			}
			if firstOperandJgz > 0 {
				return true, secondOperand
			} else {
				return true, 1
			}
		default:
			panic("Not supported two operand operation!")
		}
	}
}
