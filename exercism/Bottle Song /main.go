// https://exercism.org/tracks/go/exercises/bottle-song

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Recite(10, 1))
}

func Recite(startBottles, takeDown int) []string {
	var song []string
	for i := startBottles; i > 0; i-- {
		song = Verse(song, i, takeDown)
	}
	return song
}

func Verse(song []string, startBottles, takeDown int) []string {
	numbers := []string{"no more", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	if startBottles < 1 {
		return song
	}
	if startBottles == 1 {
		song = append(song, numbers[startBottles], "green bottle hanging on the wall,\n")
		song = append(song, numbers[startBottles], "green bottle hanging on the wall,\n")
		song = append(song, "And if", numbers[takeDown], "green bottle should accidentally fall,\n")
		song = append(song, "There'll be no more bottles hanging on the wall.\n")
		return song
	}
	song = append(song, numbers[startBottles], "green bottles hanging on the wall,\n")
	song = append(song, numbers[startBottles], "green bottles hanging on the wall,\n")
	song = append(song, "And if", numbers[takeDown], "green bottle should accidentally fall,\n")
	song = append(song, "There'll be", numbers[startBottles-takeDown], "green bottles hanging on the wall.\n\n")
	return song

}
