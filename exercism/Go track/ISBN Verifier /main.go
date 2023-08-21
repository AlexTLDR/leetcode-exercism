// https://exercism.org/tracks/go/exercises/isbn-verifier

// ISBN formula => (d₁ * 10 + d₂ * 9 + d₃ * 8 + d₄ * 7 + d₅ * 6 + d₆ * 5 + d₇ * 4 + d₈ * 3 + d₉ * 2 + d₁₀ * 1) mod 11 == 0

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(IsValidISBN("3-598-21507-X"))
	fmt.Println(IsValidISBN2("3-598-21507-X"))
}

func IsValidISBN(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	if len(isbn) != 10 {
		return false
	}
	var checkDigit int
	if isbn[9] == 'X' {
		checkDigit = 10
	} else {
		checkDigit = int(isbn[9] - '0')
	}

	isbn = isbn[:9]

	var ISBN int
	for i := 0; i < len(isbn); i++ {
		if int(isbn[i])-'0' < 0 || int(isbn[i])-'0' > 9 || checkDigit < 0 || checkDigit > 10 {
			return false
		}
		ISBN += (int(isbn[i]) - '0') * (10 - i)
	}
	ISBN += checkDigit
	return ISBN%11 == 0
}

func IsValidISBN2(isbn string) bool {
	reg := regexp.MustCompile("^[0-9]{9}[0-9X]$")
	isbn = strings.Replace(isbn, "-", "", -1)
	if !reg.MatchString(isbn) {
		return false
	}

	digit := 0
	ISBN := 0

	for i := 0; i < len(isbn); i++ {
		if isbn[i] == 'X' {
			digit = 10
		} else {
			digit = int(isbn[i] - '0')
		}
		ISBN += digit * (len(isbn) - i)
	}
	return ISBN%11 == 0
}
