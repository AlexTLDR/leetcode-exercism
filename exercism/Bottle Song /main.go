// https://exercism.org/tracks/go/exercises/bottle-song

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Recite(10, 1))
}

func Recite(startBottles, takeDown int) []string {
	var song []string
	for i := 0; i < takeDown; i++ {
		song = Verse(song, startBottles, takeDown)

	}
	return song
}

func Verse(song []string, startBottles, takeDown int) []string {
	song = append(song, strconv.Itoa(startBottles), "green bottles hanging on the wall,\n")
	song = append(song, strconv.Itoa(startBottles), "green bottles hanging on the wall,\n")
	song = append(song, "And if", strconv.Itoa(takeDown), "green bottle should accidentally fall,\n")
	return song

}
