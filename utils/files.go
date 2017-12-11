package utils

import (
	"bufio"
)

func ReadLine(fileReader *bufio.Reader) (string, error) {
	return fileReader.ReadString('\n')
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
