// https://exercism.org/tracks/go/exercises/tournament

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

	defer readFile.Close()
	Tally(readFile, nil)

}

func Tally(reader io.Reader, writer io.Writer) error {

	teams := map[string]int{
		"Allegoric Alaskans":      0,
		"Blithering Badgers":      0,
		"Devastating Donkeys":     0,
		"Courageous Californians": 0,
	}
	fileScanner := bufio.NewScanner(reader)
	fileScanner.Split(bufio.ScanLines)
	fileLines := []string{}
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	for _, line := range fileLines {

		slice := strings.Split(line, ";")
		//fmt.Println("line number: ", slice)
		switch slice[2] {
		case "win":
			//fmt.Println(slice[0])
			teams[slice[0]] += 3
		case "draw":
			teams[slice[0]] += 1
			teams[slice[1]] += 1
		case "loss":
			teams[slice[1]] += 3
		}
	}
	fmt.Println(teams)
	fmt.Println(fileLines)
	return nil
}
