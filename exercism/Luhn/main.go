// https://exercism.org/tracks/go/exercises/luhn

package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(Valid("4539 3195 0343 6467"))
	fmt.Println(Valid(strings.Replace("4539 3195 0343 6467", " ", "", -1)))
}

func Valid(id string) bool {
	//strings.Replace(id, " ", "", -1)
	var idArr []int
	fmt.Println(id)
	for _, v := range id {
		idArr = append(idArr, int(v-48))
	}

	fmt.Println(idArr)
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
