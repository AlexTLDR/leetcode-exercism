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

type team struct {
	name string
	MP   int
	W    int
	D    int
	L    int
	P    int
}

type ByPoints []team

func (a ByPoints) Len() int      { return len(a) }
func (a ByPoints) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPoints) Less(i, j int) bool {
	if a[i].P > a[j].P {
		return true
	} else if a[i].P == a[j].P {
		if a[i].name > a[j].name {
			return true
		}
	}
	return false
}

func (t team) String() string {
	return fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", t.name, t.MP, t.W, t.D, t.L, t.P)
}

func main() {
	ReadFile()
}

func ReadFile() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	//file output
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	Tally(readFile, f)

}

func Tally(reader io.Reader, writer io.Writer) error {

	teams := map[string]team{}

	fileScanner := bufio.NewScanner(reader)
	fileScanner.Split(bufio.ScanLines)
	fileLines := []string{}
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	for _, line := range fileLines {
		if line == "" {
			continue
		}
		if line[0] == '#' {
			continue
		}
		slice := strings.Split(line, ";")
		home := teams[slice[0]]
		away := teams[slice[1]]
		switch slice[2] {
		case "win":
			home.MP += 1
			home.W += 1
			home.P += 3
			away.MP += 1
			away.L += 1
		case "draw":
			home.MP += 1
			away.MP += 1
			home.D += 1
			away.D += 1
			home.P += 1
			away.P += 1
		case "loss":
			home.MP += 1
			away.MP += 1
			home.L += 1
			away.W += 1
			away.P += 3
		default:
			return fmt.Errorf("invalid")

		}
		teams[slice[0]] = home
		teams[slice[1]] = away
	}

	// sorting
	keys := make([]string, 0, len(teams))
	for k := range teams {
		keys = append(keys, k)
	}
	//writer.WriteString("Team                           | MP |  W |  D |  L |  P\n")
	writer.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	teamSlice := []team{}
	for names, v := range teams {
		v.name = names
		teamSlice = append(teamSlice, v)
	}
	sort.Sort(ByPoints(teamSlice))
	for _, v := range teamSlice {
		writer.Write([]byte(v.String()))
	}
	return nil
}

/*
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
Allegoric Alaskans             |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3
Courageous Californians        |  3 |  0 |  1 |  2 |  1
*/
