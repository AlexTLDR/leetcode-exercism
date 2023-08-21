// https://exercism.org/tracks/go/exercises/raindrops

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Convert(35))
}

func Convert(n int) string {
	s := ""
	if n%3 == 0 {
		s += "Pling"
	}
	if n%5 == 0 {
		s += "Plang"
	}
	if n%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		return strconv.Itoa(n)
	}
	return s
}
