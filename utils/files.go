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

func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
