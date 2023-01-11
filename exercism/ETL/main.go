// https://exercism.org/tracks/go/exercises/etl

package main

import (
	"fmt"
	"strings"
)

func main() {
	old := map[int][]string{
		1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		2:  {"D", "G"},
		3:  {"B", "C", "M", "P"},
		4:  {"F", "H", "V", "W", "Y"},
		5:  {"K"},
		8:  {"J", "X"},
		10: {"Q", "Z"},
	}

	new := Transform(old)
	fmt.Println(new)

}

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)

	for i, v := range in {

		for _, s := range v {
			out[strings.ToLower(s)] = i
		}

	}
	fmt.Println(out)
	return out
}
