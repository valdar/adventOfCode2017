package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/valdar/adventOfCode2017/utils"
)

type coordinates struct {
	i int
	j int
}

func main() {
	caseSelection := os.Args[1]
	startValueA, err1 := strconv.Atoi(os.Args[2])
	utils.Check(err1)
	startValueB, err2 := strconv.Atoi(os.Args[3])
	utils.Check(err2)

	switch {
	case caseSelection == "A":
		fmt.Printf("Number of matches found by judge are %d\n", CalcJudgeMatches(startValueA, startValueB))
	case caseSelection == "B":
		fmt.Printf("Number of matches found by judge for the second part are %d\n", CalcJudgeMatchesSecondPart(startValueA, startValueB))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func CalcJudgeMatches(startValueA int, startValueB int) int {
	const factorA int64 = 16807
	const factorB int64 = 48271
	const bitMsk = 0x000000000000FFFF
	result := 0
	prevValueA := int64(startValueA)
	prevValueB := int64(startValueB)
	for i := 0; i < 40000000; i++ {
		prevValueA = calcNextCounterValue(prevValueA, factorA)
		prevValueB = calcNextCounterValue(prevValueB, factorB)
		if prevValueA&bitMsk == prevValueB&bitMsk {
			result++
		}
	}
	return result
}

func calcNextCounterValue(prevValue int64, factor int64) int64 {
	const divisor int64 = 2147483647
	return (prevValue * factor) % divisor
}

func CalcJudgeMatchesSecondPart(startValueA int, startValueB int) int {
	const factorA int64 = 16807
	const factorB int64 = 48271
	const exactDivedendA int64 = 4
	const exactDivedendB int64 = 8
	const bitMsk = 0x000000000000FFFF
	result := 0
	prevValueA := int64(startValueA)
	prevValueB := int64(startValueB)
	for i := 0; i < 5000000; i++ {
		prevValueA = calcNextCounterValueSecondPart(prevValueA, factorA, exactDivedendA)
		prevValueB = calcNextCounterValueSecondPart(prevValueB, factorB, exactDivedendB)
		if prevValueA&bitMsk == prevValueB&bitMsk {
			result++
		}
	}
	return result
}

func calcNextCounterValueSecondPart(prevValue int64, factor int64, exactDividend int64) int64 {
	currValue := prevValue
	for {
		currValue = calcNextCounterValue(currValue, factor)
		if currValue%exactDividend == 0 {
			return currValue
		}
	}
}
