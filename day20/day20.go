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

type vector struct {
	x int
	y int
	z int
}

type particel struct {
	id  int
	pos vector
	vel vector
	acc vector
}

func main() {
	caseSelection := os.Args[1]
	f, err := os.Open(os.Args[2])
	defer f.Close()
	utils.Check(err)
	br := bufio.NewReader(f)

	universe := parse(br)

	switch {
	case caseSelection == "A":
		fmt.Printf("The the eventually closer to <0,0,0> particel is %d\n", SolveA(universe).id)
	case caseSelection == "B":
		fmt.Printf("The the total of surviving particel is %d\n", SolveB(universe))
	default:
		fmt.Printf("Invalid Selection, possible values: A or B\n")
	}
}

func SolveA(universe []particel) particel {
	const (
		MinUint uint = 0
		MaxUint      = ^MinUint
		MaxInt       = int(MaxUint >> 1)
		MinInt       = ^MaxInt
	)

	var result particel
	minAccParticels := []particel{}
	minAcc := MaxInt

	for _, currParticel := range universe {
		currAcc := calcVectorModule(currParticel.acc)
		if currAcc < minAcc {
			minAcc = currAcc
		}
	}
	for _, currParticel := range universe {
		currAcc := calcVectorModule(currParticel.acc)
		if currAcc == minAcc {
			minAccParticels = append(minAccParticels, currParticel)
		}
	}

	driftingAwayFromCenter := make([]bool, len(minAccParticels), len(minAccParticels))

	for !checkAllDrifting(driftingAwayFromCenter) {
		for index := 0; index < len(minAccParticels); index++ {
			startDistanceFrom000 := calcVectorModule(minAccParticels[index].pos)
			startVelocity := calcVectorModule(minAccParticels[index].vel)

			//update velocity
			minAccParticels[index].vel.x += minAccParticels[index].acc.x
			minAccParticels[index].vel.y += minAccParticels[index].acc.y
			minAccParticels[index].vel.z += minAccParticels[index].acc.z
			//update position
			minAccParticels[index].pos.x += minAccParticels[index].vel.x
			minAccParticels[index].pos.y += minAccParticels[index].vel.y
			minAccParticels[index].pos.z += minAccParticels[index].vel.z

			newDistanceFrom000 := calcVectorModule(minAccParticels[index].pos)
			newVelocity := calcVectorModule(minAccParticels[index].vel)
			if newDistanceFrom000 >= startDistanceFrom000 && newVelocity >= startVelocity {
				driftingAwayFromCenter[index] = true
			}
		}
	}

	minDistance := MaxInt

	for _, currParticel := range minAccParticels {
		currDistance := calcVectorModule(currParticel.pos)
		if currDistance < minDistance {
			minDistance = currDistance
		}
	}
	for _, currParticel := range minAccParticels {
		currDistance := calcVectorModule(currParticel.pos)
		if currDistance == minDistance {
			result = currParticel
		}
	}

	return result
}

func SolveB(universe []particel) int {
	survivedParticelsIndexes := make([]int, len(universe), len(universe))
	for i := 0; i < len(universe); i++ {
		survivedParticelsIndexes[i] = i
	}

	prevPositions := make([]vector, len(universe), len(universe))
	for i, obj := range universe {
		prevPositions[i] = obj.pos
	}
	t := 0

	for {
		t++
		fmt.Printf("Time [%d]\r", t)
		destroyed := map[int]bool{}
		newPositions := make([]vector, len(universe), len(universe))
		for place, i := range survivedParticelsIndexes {
			for _, j := range survivedParticelsIndexes[place+1:] {
				newPositions[i] = calcPositionGivenTime(t, universe[i])
				newPositions[j] = calcPositionGivenTime(t, universe[j])
				if newPositions[i].x == newPositions[j].x && newPositions[i].y == newPositions[j].y && newPositions[i].z == newPositions[j].z {
					destroyed[i] = true
					destroyed[j] = true
				}
			}
		}

		newSurvivors := []int{}
		for _, index := range survivedParticelsIndexes {
			if !destroyed[index] {
				newSurvivors = append(newSurvivors, index)
			}
		}

		gettingCloser := map[int]bool{}
		for place, i := range newSurvivors {
			for _, j := range newSurvivors[place+1:] {
				newDistance := calcSquareOfDistance(newPositions[i], newPositions[j])
				oldDistance := calcSquareOfDistance(prevPositions[i], prevPositions[j])
				newVelocityI := calcVectorModule(calcVelocityGivenTime(t, universe[i]))
				oldVelocityI := calcVectorModule(calcVelocityGivenTime(t-1, universe[i]))
				newVelocityJ := calcVectorModule(calcVelocityGivenTime(t, universe[j]))
				oldVelocityJ := calcVectorModule(calcVelocityGivenTime(t-1, universe[j]))

				if newDistance < oldDistance && newVelocityI >= oldVelocityI && newVelocityJ >= oldVelocityJ {
					gettingCloser[i] = true
					gettingCloser[j] = true
				}
			}
		}

		survivedParticelsIndexes = newSurvivors
		prevPositions = newPositions

		if checkNoOneGettingCloser(gettingCloser, newSurvivors) {
			break
		}
	}

	fmt.Printf("\n")

	return len(survivedParticelsIndexes)
}

func checkNoOneGettingCloser(flags map[int]bool, indexes []int) bool {
	result := true
	for _, currIndex := range indexes {
		result = result && !flags[currIndex]
	}
	return result
}

func checkAllDrifting(flags []bool) bool {
	result := true
	for _, currFlag := range flags {
		result = result && currFlag
	}
	return result
}

func parse(br *bufio.Reader) []particel {
	universe := []particel{}
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

		parts := strings.Split(input, ", ")

		universe = append(universe, particel{len(universe), parseVector(parts[0]), parseVector(parts[1]), parseVector(parts[2])})
	}
	return universe
}

func parseVector(coordinates string) vector {
	subparts := strings.Split(coordinates, "=")

	components := strings.Split(subparts[1][1:len(subparts[1])-1], ",")

	x, err1 := strconv.Atoi(components[0])
	utils.Check(err1)
	y, err2 := strconv.Atoi(components[1])
	utils.Check(err2)
	z, err3 := strconv.Atoi(components[2])
	utils.Check(err3)

	return vector{x, y, z}
}

func calcSquareOfDistance(one vector, two vector) int {
	return (one.x-two.x)*(one.x-two.x) + (one.y-two.y)*(one.y-two.y) + (one.z-two.z)*(one.z-two.z)
}

func calcPositionGivenTime(t int, obj particel) vector {
	x := calcPositiogivenTimeSingleDirection(t, obj.acc.x, obj.vel.x, obj.pos.x)
	y := calcPositiogivenTimeSingleDirection(t, obj.acc.y, obj.vel.y, obj.pos.y)
	z := calcPositiogivenTimeSingleDirection(t, obj.acc.z, obj.vel.z, obj.pos.z)
	return vector{x, y, z}
}

func calcVelocityGivenTime(t int, obj particel) vector {
	x := calcVelocityTimeSingleDirection(t, obj.acc.x, obj.vel.x)
	y := calcVelocityTimeSingleDirection(t, obj.acc.y, obj.vel.y)
	z := calcVelocityTimeSingleDirection(t, obj.acc.z, obj.vel.z)
	return vector{x, y, z}
}

func calcPositiogivenTimeSingleDirection(t int, a int, v int, p int) int {
	return a*t*(t+1)/2 + v*t + p
}

func calcVelocityTimeSingleDirection(t int, a int, v int) int {
	return v + a*t
}

func calcVectorModule(v vector) int {
	return utils.Abs(v.x) + utils.Abs(v.y) + utils.Abs(v.z)
}
