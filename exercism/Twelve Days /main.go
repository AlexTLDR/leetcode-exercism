// https://exercism.org/tracks/go/exercises/twelve-days

package main

import (
	"fmt"
	"strings"
)

func main() {

}

var days = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var gifts = []string{"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens", "four Calling Birds", "five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}

const repeat = "On the %s day of Christmas my true love gave to me: %s"

func Verse(i int) string {
	if i < 1 || i > 12 {
		return ""
	}

	return fmt.Sprintf(repeat, days[i-1], strings.Join(gifts[0:i], ", "))
}

func Song() string {
	panic("Please implement the Song function")
}
