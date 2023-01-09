// https://exercism.org/tracks/go/exercises/roman-numerals

package main

import (
	"fmt"
	"strconv"
)

func main() {

}

// I, II, III, IV, V, VI, VII, VIII, IX, X, XI, XII

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", fmt.Errorf("The number %d is undefined", input)
	}

	// mappings := map[int]string{
	// 	1000: "M",
	// 	900:  "CM",
	// 	500:  "D",
	// 	400:  "CD",
	// 	100:  "C",
	// 	90:   "XC",
	// 	50:   "L",
	// 	40:   "XL",
	// 	10:   "X",
	// 	9:    "IX",
	// 	5:    "V",
	// 	4:    "IV",
	// 	1:    "I",
	// }

	stringInput := strconv.Itoa(input)
	for i, v := range stringInput {
		switch i {
		case 0:
		case 1:
		case 2:
		case 3:
		}

	}
	return "", fmt.Errorf("The number %d is undefined", input)
}
