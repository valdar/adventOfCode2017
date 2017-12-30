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
	iterations, err := strconv.Atoi(os.Args[1])
	utils.Check(err)

	f, err := os.Open(os.Args[2])
	defer f.Close()
	utils.Check(err)
	br := bufio.NewReader(f)

	mapping := parse(br)

	startPixls := make([][]bool, 3)
	startPixls[0] = []bool{false, true, false}
	startPixls[1] = []bool{false, false, true}
	startPixls[2] = []bool{true, true, true}

	fmt.Printf("The number of ON pixels after %d iterations is %d\n", iterations, Solve(mapping, startPixls, iterations))
}

func Solve(mapping map[string][][]bool, startPixels [][]bool, iterations int) int {
	currentStatePixels := startPixels
	fmt.Println("STARTING PIXEL MATRIX:")
	printBlock(currentStatePixels)
	fmt.Println("")
	for i := 0; i < iterations; i++ {
		if len(currentStatePixels)%2 == 0 {
			currentStatePixels = expandPixelMatrix(mapping, currentStatePixels, 2, 3)
		} else if len(currentStatePixels)%3 == 0 {
			currentStatePixels = expandPixelMatrix(mapping, currentStatePixels, 3, 4)
		} else {
			panic("Error, pixel matrix should be multiple of 2 or 3!")
		}
	}
	fmt.Println("FINISHED PIXEL MATRIX:")
	printBlock(currentStatePixels)
	fmt.Println("")

	return countOnPixels(currentStatePixels)
}

func expandPixelMatrix(mapping map[string][][]bool, pixelMatrix [][]bool, divedBlockBy int, mappedBlockSize int) [][]bool {
	numberOfYsubBlocks := len(pixelMatrix) / divedBlockBy
	numberOfXsubBlocks := len(pixelMatrix[0]) / divedBlockBy

	result := make([][]bool, numberOfYsubBlocks*mappedBlockSize)
	for i, _ := range result {
		result[i] = make([]bool, numberOfXsubBlocks*mappedBlockSize)
	}

	for i := 0; i < numberOfYsubBlocks; i++ {
		for j := 0; j < numberOfXsubBlocks; j++ {
			startY := i * divedBlockBy
			startX := j * divedBlockBy
			endY := startY + divedBlockBy - 1
			endX := startX + divedBlockBy - 1
			remappedBlock := mapping[getKey(extractSubBlock(pixelMatrix, startY, startX, endY, endX))]

			destSartY := i * mappedBlockSize
			destSartX := j * mappedBlockSize
			copySubBlock(result, remappedBlock, destSartY, destSartX)
		}
	}

	return result
}

func copySubBlock(destBlock [][]bool, sourceBlock [][]bool, startY int, startX int) {
	for i, _ := range sourceBlock {
		for j, _ := range sourceBlock[i] {
			destBlock[startY+i][startX+j] = sourceBlock[i][j]
		}
	}
}

func extractSubBlock(sourceBlock [][]bool, startY int, startX int, endY int, endX int) [][]bool {
	ysize := endY - startY + 1
	xsize := endX - startX + 1
	result := make([][]bool, ysize)
	for i, _ := range result {
		result[i] = make([]bool, xsize)
	}

	for i, _ := range result {
		for j, _ := range result[i] {
			result[i][j] = sourceBlock[startY+i][startX+j]
		}
	}

	return result
}

func countOnPixels(block [][]bool) int {
	result := 0
	for i, _ := range block {
		for j, _ := range block[i] {
			if block[i][j] {
				result++
			}
		}
	}
	return result
}

func printBlock(block [][]bool) {
	for i, _ := range block {
		for j, _ := range block[i] {
			if block[i][j] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func parse(br *bufio.Reader) map[string][][]bool {
	mapping := map[string][][]bool{}
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

		parts := strings.Split(input, " => ")
		fromBlock := parseBlock(parts[0])
		rotate90Block := rotate90degreesClockwiseNxNblock(fromBlock)
		rotate180Block := rotate90degreesClockwiseNxNblock(rotate90Block)
		rotate270Block := rotate90degreesClockwiseNxNblock(rotate180Block)

		//we just flip horizontally since all the rotations account also for the vertical flipp plus its rotations
		flippedBlock := flipHorizNxNblock(fromBlock)
		rotate90FlippedBlock := rotate90degreesClockwiseNxNblock(flippedBlock)
		rotate180FlippedBlock := rotate90degreesClockwiseNxNblock(rotate90FlippedBlock)
		rotate270FlippedBlock := rotate90degreesClockwiseNxNblock(rotate180FlippedBlock)

		toBlock := parseBlock(parts[1])

		mapping[getKey(fromBlock)] = toBlock
		mapping[getKey(rotate90Block)] = toBlock
		mapping[getKey(rotate180Block)] = toBlock
		mapping[getKey(rotate270Block)] = toBlock

		mapping[getKey(flippedBlock)] = toBlock
		mapping[getKey(rotate90FlippedBlock)] = toBlock
		mapping[getKey(rotate180FlippedBlock)] = toBlock
		mapping[getKey(rotate270FlippedBlock)] = toBlock
	}
	return mapping
}

func rotate90degreesClockwiseNxNblock(block [][]bool) [][]bool {
	result := make([][]bool, len(block))
	for i, _ := range block {
		result[i] = make([]bool, len(block[0]))
	}

	for i, row := range block {
		for j, pixel := range row {
			result[j][len(block)-1-i] = pixel
		}
	}

	return result
}

func flipHorizNxNblock(block [][]bool) [][]bool {
	result := make([][]bool, len(block))
	for i, _ := range block {
		result[i] = make([]bool, len(block[0]))
	}

	for i, _ := range block {
		swapIndex := len(block) - 1 - i
		for j := range block {
			result[swapIndex][j] = block[i][j]
		}
	}

	return result
}

func getKey(block [][]bool) string {
	result := ""
	for i, row := range block {
		for j, _ := range row {
			if block[i][j] {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "/"
	}
	return result[:len(result)-1]
}

func parseBlock(key string) [][]bool {
	rows := strings.Split(key, "/")
	block := make([][]bool, len(rows))
	for i, row := range rows {
		pixels := strings.Split(row, "")
		blockRow := make([]bool, len(pixels))
		for j, pixel := range pixels {
			if pixel == "#" {
				blockRow[j] = true
			} else {
				blockRow[j] = false
			}
		}
		block[i] = blockRow
	}
	return block
}
