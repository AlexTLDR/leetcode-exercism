// https://exercism.org/tracks/go/exercises/twelve-days

package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(Song())
	fmt.Println(Verse(2))
}

var days = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var gifts = []string{"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens", "four Calling Birds", "five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}

const repeat = "On the %s day of Christmas my true love gave to me: %s"

func Verse(i int) string {
	if i < 1 || i > 12 {
		return ""
	}
	return fmt.Sprintf(repeat, days[i-1], giftsForDay(i)) + "."
}

func Song() string {
	var verses []string
	for i := 1; i <= 12; i++ {
		if i == 1 {
			verses = append(verses, "and", Verse(i))
		} else {
			verses = append(verses, Verse(i))
		}
	}

	return strings.Join(verses, "\n")
}

func giftsForDay(i int) string {
	if i == 1 {
		return gifts[0]
	}

	reversedGifts := make([]string, i)
	for j := 0; j < i; j++ {
		reversedGifts[j] = gifts[i-j-1]
	}

	reversedGifts[len(reversedGifts)-1] = "and " + reversedGifts[len(reversedGifts)-1]

	return strings.Join(reversedGifts, ", ")
}
