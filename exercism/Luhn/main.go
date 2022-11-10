// https://exercism.org/tracks/go/exercises/luhn

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Valid("4539 3195 0343 6467"))
}

func Valid(id string) bool {
	if len(strTrimmer(id)) < 2 {
		return false
	}
	var sum int
	var idArr []int

	for _, v := range strTrimmer(id) {
		idArr = append(idArr, int(v-48))
	}
	for i, v := range idArr {
		if i == 0 || i%2 == 0 {
			v = doubleDigitValue(v)
		}
		sum += v
	}
	if sum%10 == 0 {
		return true
	}
	return false
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

func strTrimmer(str string) (trimmed string) {
	trimmed = strings.Replace(str, " ", "", -1)
	return trimmed
}
