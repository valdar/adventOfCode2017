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
