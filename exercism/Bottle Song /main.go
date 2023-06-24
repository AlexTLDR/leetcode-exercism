// https://exercism.org/tracks/go/exercises/bottle-song

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Recite(3, 3))
}

func Recite(startBottles, takeDown int) []string {
	var song []string
	stopBottles := startBottles - takeDown
	for i := startBottles; i > stopBottles; i-- {
		if i < startBottles {
			song = append(song, "")
		}
		song = Verse(song, i, takeDown)
	}
	return song
}

func Verse(song []string, startBottles, takeDown int) []string {
	numbers := []string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	plural1 := "s"
	plural3 := "s"
	if startBottles == 1 {
		plural1 = ""
	}
	if numbers[startBottles-1] == "one" {
		plural3 = ""
	}

	line1 := fmt.Sprintf("%s green bottle%s hanging on the wall,", strings.Title(numbers[startBottles]), plural1)
	line2 := fmt.Sprintf("And if %s green bottle should accidentally fall,", "one")
	line3 := fmt.Sprintf("There'll be %s green bottle%s hanging on the wall.", numbers[startBottles-1], plural3)
	song = append(song, line1, line1, line2, line3)
	return song
}
