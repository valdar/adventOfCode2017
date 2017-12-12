package main

import (
	"bufio"
	"fmt"
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
		fmt.Printf("The hash is %d\n", CalcHash(inputSlice, strings.Split(input, ",")))
	case caseSelection == "B":
		fmt.Printf("The hash knot is %v\n", CalcHashKnot(inputSlice, input))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func CalcHashKnot(inputSlice []int, lenghtsString string) string {
	lenghts := []int{}
	for _, currRune := range lenghtsString {
		lenghts = append(lenghts, int(currRune))
	}
	lenghts = append(lenghts, []int{17, 31, 73, 47, 23}...)

	curPosition := 0
	currSkip := 0

	for i := 0; i < 64; i++ {
		for _, currLength := range lenghts {
			curPosition = PerformStep(inputSlice, curPosition, currSkip, currLength, cap(inputSlice))
			currSkip++
		}
	}

	digit1 := convertToHexString(xor(inputSlice[:16]))
	digit2 := convertToHexString(xor(inputSlice[16:32]))
	digit3 := convertToHexString(xor(inputSlice[32:48]))
	digit4 := convertToHexString(xor(inputSlice[48:64]))
	digit5 := convertToHexString(xor(inputSlice[64:80]))
	digit6 := convertToHexString(xor(inputSlice[80:96]))
	digit7 := convertToHexString(xor(inputSlice[96:112]))
	digit8 := convertToHexString(xor(inputSlice[112:128]))
	digit9 := convertToHexString(xor(inputSlice[128:144]))
	digit10 := convertToHexString(xor(inputSlice[144:160]))
	digit11 := convertToHexString(xor(inputSlice[160:176]))
	digit12 := convertToHexString(xor(inputSlice[176:192]))
	digit13 := convertToHexString(xor(inputSlice[192:208]))
	digit14 := convertToHexString(xor(inputSlice[208:224]))
	digit15 := convertToHexString(xor(inputSlice[224:240]))
	digit16 := convertToHexString(xor(inputSlice[240:]))

	return digit1 + digit2 + digit3 + digit4 + digit5 + digit6 + digit7 + digit8 + digit9 + digit10 + digit11 + digit12 + digit13 + digit14 + digit15 + digit16
}

func xor(input []int) int {
	result := 0
	for _, currElement := range input {
		result ^= currElement
	}
	return result
}

func convertToHexString(input int) string {
	hexDigit := fmt.Sprintf("%x", input)
	if len(hexDigit) == 1 {
		hexDigit = "0" + hexDigit
	}
	return hexDigit
}

func CalcHash(inputSlice []int, lenghts []string) int {
	curPosition := 0
	currSkip := 0
	for _, currLength := range lenghts {
		currIntLength, err := strconv.Atoi(currLength)
		utils.Check(err)
		curPosition = PerformStep(inputSlice, curPosition, currSkip, currIntLength, cap(inputSlice))
		currSkip++
	}
	return inputSlice[0] * inputSlice[1]
}

func PerformStep(slice []int, startingPosition int, skipSize int, length int, sliceSize int) int {
	if length > sliceSize {
		panic("Lengths larger than the size of the list are invalid.")
	}

	endOfreverseSection := startingPosition + length

	var toBeReversed []int
	if endOfreverseSection > sliceSize {
		toBeReversed = append(slice[startingPosition:], slice[:endOfreverseSection%sliceSize]...)
	} else {
		toBeReversed = append([]int{}, slice[startingPosition:endOfreverseSection]...)
	}

	utils.Reverse(toBeReversed)

	if endOfreverseSection > sliceSize {
		copy(slice[startingPosition:], toBeReversed[:sliceSize-startingPosition])
		copy(slice[:endOfreverseSection%sliceSize], toBeReversed[sliceSize-startingPosition:])
	} else {
		copy(slice[startingPosition:endOfreverseSection], toBeReversed)
	}

	return (startingPosition + length + skipSize) % sliceSize
}
