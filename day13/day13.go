package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	firewallStructure := parse(br)

	switch {
	case caseSelection == "A":
		fmt.Printf("The severity of the firewall passthrough is %d\n", CalcSeverity(firewallStructure))
	case caseSelection == "B":
		fmt.Printf("The minimum delay to passthrough the firewall with severity 0 (i.e. uncought) is %d\n", CalcDelayToSeverity0(firewallStructure))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func CalcScannerPosition(time int, layerRange int) int {
	// this is the Triangle wave (https://en.wikipedia.org/wiki/Triangle_wave)
	// parametrized to the layer range it represent the position of the scanner at a given time
	pass := float64(2 * layerRange)
	timeOnPass := float64(time) / pass
	return utils.Abs(utils.Round((pass * (timeOnPass - math.Floor(timeOnPass+0.5)))))
}

func CalcSeverity(firewallStructure map[int]int) int {
	severity := 0
	for k := range firewallStructure {
		if CalcScannerPosition(k, firewallStructure[k]-1) == 0 {
			severity += k * firewallStructure[k]
		}
	}
	return severity
}

func CalcDelayToSeverity0(firewallStructure map[int]int) int {
	delay := 0
	for {
		severity := 0
		for k := range firewallStructure {
			if CalcScannerPosition(k+delay, firewallStructure[k]-1) == 0 {
				severity++
			}
		}
		if severity == 0 {
			break
		}
		delay++
	}
	return delay
}

func parse(br *bufio.Reader) map[int]int {
	firewallStructure := map[int]int{}
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

		parts := strings.Split(input, ":")

		if len(parts) > 2 {
			panic("Not parsable input!")
		}

		key, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		utils.Check(err)

		value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		utils.Check(err)

		firewallStructure[key] = value
	}
	return firewallStructure
}
