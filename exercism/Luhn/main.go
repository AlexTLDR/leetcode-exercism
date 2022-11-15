// https://exercism.org/tracks/go/exercises/luhn

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Valid("4539 3195 0343 6468"))
}

func Valid(id string) bool {
	trimmed := strings.Replace(id, " ", "", -1)
	var sum int
	if len(trimmed) < 2 {
		return false
	}
	for i, v := range trimmed {
		if !unicode.IsDigit(v) {
			return false
		}
		digit := int(v - '0')
		switch {
		case i%2 == len(trimmed)%2:
			sum += doubleDigitValue(digit)
		default:
			sum += digit
		}
	}
	return sum%10 == 0
}

func doubleDigitValue(digit int) (value int) {
	doubleDigit := digit * 2

	if doubleDigit > 9 {
		value = doubleDigit - 9
	} else {
		value = doubleDigit
	}
	return value
}
