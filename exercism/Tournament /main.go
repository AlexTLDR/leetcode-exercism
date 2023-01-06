// https://exercism.org/tracks/go/exercises/tournament

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func WriteFile(key string, teams map[string]int) {
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
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
		switch slice[2] {
		case "win":
			teams[slice[0]] += 3
		case "draw":
			teams[slice[0]] += 1
			teams[slice[1]] += 1
		case "loss":
			teams[slice[1]] += 3
		}
	}

	// file output
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// map sorting
	keys := make([]string, 0, len(teams))
	for k := range teams {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	for _, k := range keys {
		fmt.Println(k, teams[k])
		_, err := f.WriteString(k)
		if err != nil {
			fmt.Println(err)
		}

	}

	//fmt.Println(teams)
	// fmt.Println(fileLines)

	return nil
}
