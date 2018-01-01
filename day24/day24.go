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

type port struct {
	A        int
	B        int
	freePort int
}

type inProgressBridge struct {
	buildPart      []port
	remainingParts []port
}

func main() {
	caseSelection := os.Args[1]
	f, err := os.Open(os.Args[2])
	defer f.Close()
	utils.Check(err)
	br := bufio.NewReader(f)
	ports := parse(br)

	switch {
	case caseSelection == "A":
		fmt.Printf("Solving... please wait upp to 20 minutes on a modern hardware (if you want to see the progress decomment line 129)\n")
		fmt.Printf("The strongest bridge available is %d\n", SolveA(ports))
	case caseSelection == "B":
		fmt.Printf("Solving... please wait upp to 20 minutes on a modern hardware (if you want to see the progress decomment line 129)\n")
		fmt.Printf("The strongest longest bridge available is %d\n", SolveB(ports))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func SolveA(parts []port) int {
	completeBridges := buildBridges(parts)
	maxBridge := 0
	for _, bridge := range completeBridges {
		bridgeTotal := 0
		for _, part := range bridge {
			bridgeTotal += part.A
			bridgeTotal += part.B
		}
		if bridgeTotal > maxBridge {
			maxBridge = bridgeTotal
		}
	}
	return maxBridge
}

func SolveB(parts []port) int {
	completeBridges := buildBridges(parts)

	longestBridgeLenght := 0
	for _, bridge := range completeBridges {
		if len(bridge) > longestBridgeLenght {
			longestBridgeLenght = len(bridge)
		}
	}

	longestBridges := [][]port{}
	for _, bridge := range completeBridges {
		if len(bridge) == longestBridgeLenght {
			longestBridges = append(longestBridges, bridge)
		}
	}

	maxBridge := 0
	for _, bridge := range longestBridges {
		bridgeTotal := 0
		for _, part := range bridge {
			bridgeTotal += part.A
			bridgeTotal += part.B
		}
		if bridgeTotal > maxBridge {
			maxBridge = bridgeTotal
		}
	}
	return maxBridge
}

func buildBridges(parts []port) [][]port {
	completeBridges := [][]port{}
	inprogressBridges := []inProgressBridge{}

	//find initial bridges
	for i, part := range parts {
		if part.A == 0 {
			remaningParts := append([]port{}, parts[:i]...)
			remaningParts = append(remaningParts, parts[i+1:]...)
			inprogressBridges = append(inprogressBridges, inProgressBridge{[]port{port{part.A, part.B, part.B}}, remaningParts})
		} else if part.B == 0 {
			remaningParts := append([]port{}, parts[:i]...)
			remaningParts = append(remaningParts, parts[i+1:]...)
			inprogressBridges = append(inprogressBridges, inProgressBridge{[]port{port{part.A, part.B, part.A}}, remaningParts})
		}
	}

	for len(inprogressBridges) > 0 {
		currBridge := inprogressBridges[0]
		inprogressBridges = append([]inProgressBridge{}, inprogressBridges[1:]...)
		initialLenght := len(inprogressBridges)
		//find next port
		for i, part := range currBridge.remainingParts {
			if part.A == currBridge.buildPart[0].freePort {
				remaningParts := append([]port{}, currBridge.remainingParts[:i]...)
				remaningParts = append(remaningParts, currBridge.remainingParts[i+1:]...)
				inprogressBridges = append(inprogressBridges, inProgressBridge{append([]port{port{part.A, part.B, part.B}}, currBridge.buildPart...), remaningParts})
			} else if part.B == currBridge.buildPart[0].freePort {
				remaningParts := append([]port{}, currBridge.remainingParts[:i]...)
				remaningParts = append(remaningParts, currBridge.remainingParts[i+1:]...)
				inprogressBridges = append(inprogressBridges, inProgressBridge{append([]port{port{part.A, part.B, part.A}}, currBridge.buildPart...), remaningParts})
			}
		}

		if initialLenght == len(inprogressBridges) {
			completeBridges = append(completeBridges, currBridge.buildPart)
		}
		//fmt.Printf("\rExploring [%d] options", len(inprogressBridges))
	}
	return completeBridges
}

func parse(br *bufio.Reader) []port {
	result := []port{}
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

		terminals := strings.Split(input, "/")
		a, errA := strconv.Atoi(terminals[0])
		utils.Check(errA)
		b, errB := strconv.Atoi(terminals[1])
		utils.Check(errB)
		result = append(result, port{a, b, -1})
	}
	return result
}
