package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/valdar/adventOfCode2017/day10/hashKnot"
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
		fmt.Printf("The hash is %d\n", hashKnot.CalcHash(inputSlice, strings.Split(input, ",")))
	case caseSelection == "B":
		fmt.Printf("The hash knot is %v\n", hashKnot.CalcHashKnot(inputSlice, input))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}
