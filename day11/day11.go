package main

import (
	"bufio"
	"fmt"
	"os"
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

	switch {
	case caseSelection == "A":
		endX, endy, _ := CalcEndingHexAndMax(input)
		fmt.Printf("The distance is %d\n", CalcCellDistance(0, 0, endX, endy))
	case caseSelection == "B":
		_, _, maxDistance := CalcEndingHexAndMax(input)
		fmt.Printf("The MAX distance is %d\n", maxDistance)
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func CalcEndingHexAndMax(input string) (int, int, int) {
	var x, y, max int
	for _, currDir := range strings.Split(input, ",") {
		x, y = CalcNextHex(x, y, currDir)
		currDistance := CalcCellDistance(0, 0, x, y)
		if currDistance > max {
			max = currDistance
		}
	}
	return x, y, max
}

func CalcNextHex(x int, y int, direction string) (int, int) {
	switch direction {
	case "n":
		return x, y + 1
	case "ne":
		return x + 1, y
	case "se":
		return x + 1, y - 1
	case "s":
		return x, y - 1
	case "sw":
		return x - 1, y
	case "nw":
		return x - 1, y + 1
	default:
		panic("Direction not recognised!")
	}
}

func CalcCellDistance(x0 int, y0 int, x1 int, y1 int) int {
	if x1 < x0 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}
	if y1 > y0 {
		return x1 - x0 + y1 - y0
	} else if x0+y0 > x1+y1 {
		return y0 - y1
	} else {
		return x1 - x0
	}
}
