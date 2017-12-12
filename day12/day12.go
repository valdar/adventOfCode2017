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

	pipeChart := parse(br)

	switch {
	case caseSelection == "A":
		fmt.Printf("The number of reachable programs is %d\n", len(FindReachableFromStarting(0, pipeChart)))
	case caseSelection == "B":
		fmt.Printf("The number of programs groups is %d\n", FindReachableGroups(pipeChart))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func FindReachableFromStarting(starting int, pipeChart map[int][]int) []int {
	reachable := map[int]bool{}
	toBeVisited := []int{}

	if startingReachable, present := pipeChart[starting]; present {
		reachable[starting] = true
		toBeVisited = append(toBeVisited, startingReachable...)
		for len(toBeVisited) > 0 {
			currInspectedElement := toBeVisited[0]
			toBeVisited = append([]int{}, toBeVisited[1:]...)
			if !reachable[currInspectedElement] {
				reachable[currInspectedElement] = true
				toBeVisited = append(toBeVisited, pipeChart[currInspectedElement]...)
			}
		}
	}

	result := []int{}
	for k := range reachable {
		result = append(result, k)
	}
	return result
}

func FindReachableGroups(pipeChart map[int][]int) int {
	partOfaGroup := map[int]bool{}
	var numberOfGroups int

	for k := range pipeChart {
		if !partOfaGroup[k] {
			numberOfGroups++
			kGroup := FindReachableFromStarting(k, pipeChart)
			for _, kGroupMember := range kGroup {
				partOfaGroup[kGroupMember] = true
			}
		}
	}

	return numberOfGroups
}

func parse(br *bufio.Reader) map[int][]int {
	pipeChart := map[int][]int{}
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

		parts := strings.Split(input, "<->")

		if len(parts) > 2 {
			panic("Not parsable input!")
		}

		key, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		utils.Check(err)

		value := []int{}
		for _, currInt := range strings.Split(parts[1], ",") {
			currValue, err := strconv.Atoi(strings.TrimSpace(currInt))
			utils.Check(err)
			value = append(value, currValue)
		}
		pipeChart[key] = value
	}
	return pipeChart
}
