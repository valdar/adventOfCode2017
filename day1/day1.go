package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

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

	switch {
	case caseSelection == "A":
		fmt.Printf("The sum is %d\n", DoTheSum(input, 1))
	case caseSelection == "B":
		fmt.Printf("The sum is %d\n", DoTheSum(input, len(input)/2))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}

}

func DoTheSum(input string, pass int) int {
	sum := 0
	for position, char := range input {
		current, _ := strconv.Atoi(string(char))
		compareTo, _ := strconv.Atoi(string(input[(position+pass)%len(input)]))
		if current == compareTo {
			sum += current
		}
	}
	return sum
}
