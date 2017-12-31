package instructions

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/valdar/adventOfCode2017/utils"
)

type Instruction struct {
	Operation             string
	FirstOperandNumeric   int
	FirstOperandRegistry  string
	SecondOperandNumeric  int
	SecondOperandRegistry string
	OneOperand            bool
}

func Parse(br *bufio.Reader) []Instruction {
	program := []Instruction{}
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

		if len(parts) == 2 {
			firstOperandNum, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err == nil {
				program = append(program, Instruction{strings.TrimSpace(parts[0]), firstOperandNum, "", 0, "", true})
			} else {
				program = append(program, Instruction{strings.TrimSpace(parts[0]), 0, strings.TrimSpace(parts[1]), 0, "", true})
			}
		} else if len(parts) == 3 {
			firstOperandNum, err1 := strconv.Atoi(strings.TrimSpace(parts[1]))
			secondOperandNum, err2 := strconv.Atoi(strings.TrimSpace(parts[2]))

			if err1 == nil && err2 == nil {
				program = append(program, Instruction{strings.TrimSpace(parts[0]), firstOperandNum, "", secondOperandNum, "", false})
			} else if err1 != nil && err2 == nil {
				program = append(program, Instruction{strings.TrimSpace(parts[0]), 0, strings.TrimSpace(parts[1]), secondOperandNum, "", false})
			} else if err1 == nil && err2 != nil {
				program = append(program, Instruction{strings.TrimSpace(parts[0]), firstOperandNum, "", 0, strings.TrimSpace(parts[2]), false})
			} else if err1 != nil && err2 != nil {
				program = append(program, Instruction{strings.TrimSpace(parts[0]), 0, strings.TrimSpace(parts[1]), 0, strings.TrimSpace(parts[2]), false})
			}
		} else {
			panic("Not parsable input!")
		}
	}
	return program
}
