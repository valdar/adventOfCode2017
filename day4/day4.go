package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
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
		fmt.Printf("Number of vali passphrases is %d\n", checkValisPassphrase(checkDuplicates, br))
	case caseSelection == "B":
		fmt.Printf("Number of vali passphrases is %d\n", checkValisPassphrase(CheckNotAnagram, br))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func CalcOccurencies(input []string) map[string]int {
	m := make(map[string]int)
	for _, str := range input {
		m[str] += 1
	}
	return m
}

func CountOccurenciesGreatherEqualThan(m map[string]int, trashold int) int {
	greatherThanTrashold := 0
	for _, value := range m {
		if value >= trashold {
			greatherThanTrashold++
		}
	}
	return greatherThanTrashold
}

func checkDuplicates(input []string) bool {
	return CountOccurenciesGreatherEqualThan(CalcOccurencies(input), 2) == 0
}

func CheckAnagram(input []string) bool {
	letterOccurenciesPerWord := make([]map[string]int, len(input))
	for position, str := range input {
		letterOccurenciesPerWord[position] = CalcOccurencies(strings.Split(str, ""))
	}
	for pos, currLetterMaprange := range letterOccurenciesPerWord[:len(letterOccurenciesPerWord)-1] {
		for _, compareLetterMap := range letterOccurenciesPerWord[pos+1:] {
			if reflect.DeepEqual(currLetterMaprange, compareLetterMap) {
				return true
			}
		}
	}
	return false
}

func CheckNotAnagram(input []string) bool {
	return !CheckAnagram(input)
}

func checkValisPassphrase(checkFunction func(input []string) bool, br *bufio.Reader) int {
	validPass := 0
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

		if checkFunction(row) {
			validPass++
		}
	}
	return validPass
}
