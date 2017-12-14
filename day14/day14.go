package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/valdar/adventOfCode2017/day10/hashKnot"
)

type coordinates struct {
	i int
	j int
}

func main() {
	caseSelection := os.Args[1]
	input := os.Args[2]

	switch {
	case caseSelection == "A":
		fmt.Printf("Number of Full blocks on disk are %d\n", CalcFullBlocks(input))
	case caseSelection == "B":
		fmt.Printf("Number of Full regions on disk are %d\n", CalcRegions(input))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func CalcFullBlocks(key string) int {
	fullBlocks := 0
	for i := 0; i < 128; i++ {
		suffix := strconv.Itoa(i)

		hashBits := calcBinayString(key + "-" + suffix)
		for _, bit := range strings.Split(hashBits, "") {
			if bit == "1" {
				fullBlocks++
			}
		}
	}
	return fullBlocks
}

func CalcRegions(key string) int {
	var reagionMask [128][128]int
	var blocks [128][128]int

	for i := 0; i < 128; i++ {
		suffix := strconv.Itoa(i)

		hashBits := calcBinayString(key + "-" + suffix)
		for j, bit := range strings.Split(hashBits, "") {
			if bit == "1" {
				blocks[i][j] = 1
			}
		}
	}

	regionCounter := 0
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if blocks[i][j] == 0 {
				//empty block
				reagionMask[i][j] = -1
			} else if blocks[i][j] == 1 && reagionMask[i][j] == 0 {
				//full block not yet explored starting point of a new region
				regionCounter++
				regionCoordinates := []coordinates{{i, j}}
				for len(regionCoordinates) > 0 {
					currCoordinate := regionCoordinates[0]
					regionCoordinates = regionCoordinates[1:]
					if reagionMask[currCoordinate.i][currCoordinate.j] == 0 && blocks[currCoordinate.i][currCoordinate.j] == 1 {
						//the full block at coordinates is part of the current region
						reagionMask[currCoordinate.i][currCoordinate.j] = regionCounter
						//add adiacent block for exploration
						regionCoordinates = append(regionCoordinates, findAdiacentBlocks(currCoordinate, 128, 128)...)
					}
				}
			}
		}
	}

	max := 0
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if reagionMask[i][j] > max {
				max = reagionMask[i][j]
			}
		}
	}
	return max
}

func findAdiacentBlocks(block coordinates, iMaxSize int, jMaxSize int) []coordinates {
	var canditadeCoordinates [4]coordinates
	canditadeCoordinates[0] = coordinates{block.i - 1, block.j}
	canditadeCoordinates[1] = coordinates{block.i + 1, block.j}
	canditadeCoordinates[2] = coordinates{block.i, block.j - 1}
	canditadeCoordinates[3] = coordinates{block.i, block.j + 1}

	result := []coordinates{}
	for _, currCoordinate := range canditadeCoordinates {
		if currCoordinate.i >= 0 && currCoordinate.i < iMaxSize && currCoordinate.j >= 0 && currCoordinate.j < jMaxSize {
			result = append(result, currCoordinate)
		}
	}
	return result
}

func calcBinayString(stringToHash string) string {
	hashKnot := hashKnot.HashKnot(stringToHash)

	conversionTable := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"a": "1010",
		"b": "1011",
		"c": "1100",
		"d": "1101",
		"e": "1110",
		"f": "1111",
	}

	result := ""
	for _, hexCar := range strings.Split(hashKnot, "") {
		if convertedChar, ok := conversionTable[hexCar]; ok {
			result += convertedChar
		} else {
			panic("Unrecognized hex char!")
		}
	}
	return result
}
