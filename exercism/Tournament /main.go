// https://exercism.org/tracks/go/exercises/tournament

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	win  = 3
	draw = 1
	loss = 0
)

func main() {
	ReadFile()
}

func ReadFile() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	readFile.Close()
}

func Tally(reader io.Reader, writer io.Writer) error {
	panic("Please implement the Tally function")
}
