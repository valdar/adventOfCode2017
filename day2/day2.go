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
		fmt.Printf("The checksum is %d\n", calcCS(CalcCheckSum, br))
	case caseSelection == "B":
		fmt.Printf("The checksum is %d\n", calcCS(CalcCheckSumByEvenDivision, br))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}

}

func calcCS(csFunction func(input []string) int, br *bufio.Reader) int {
	checkSum := 0
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

		row := strings.Fields(input)

		checkSum += csFunction(row)
	}
	return checkSum
}

func CalcCheckSum(input []string) int {
	max := 0
	min := 0
	for _, str := range input {
		current, _ := strconv.Atoi(str)
		if current < min || min == 0 {
			min = current
		}
		if current > max {
			max = current
		}
	}
	return max - min
}

func CalcCheckSumByEvenDivision(input []string) int {
	num := 0
	dividedBy := 0
entireLoop:
	for position, str := range input[:len(input)-1] {
		current, _ := strconv.Atoi(str)

		for _, str2 := range input[position+1:] {
			compared, _ := strconv.Atoi(str2)

			if (current % compared) == 0 {
				num = current
				dividedBy = compared
				break entireLoop
			}

			if (compared % current) == 0 {
				num = compared
				dividedBy = current
				break entireLoop
			}
		}

	}
	return num / dividedBy
}
