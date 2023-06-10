// https://exercism.org/tracks/go/exercises/isbn-verifier

// ISBN formula => (d₁ * 10 + d₂ * 9 + d₃ * 8 + d₄ * 7 + d₅ * 6 + d₆ * 5 + d₇ * 4 + d₈ * 3 + d₉ * 2 + d₁₀ * 1) mod 11 == 0

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsValidISBN("3-598-21508-8"))
}

func IsValidISBN(isbn string) bool {
	trimmed := strings.Replace(isbn, "-", "", -1)
	if len(trimmed) != 10 {
		return false
	}
	var checkDigit int
	if trimmed[9] == 88 {
		checkDigit = 10
	} else {
		checkDigit = int(trimmed[9] - 48)
	}

	trimmed = trimmed[:9]

	var ISBN int
	for i := 0; i < len(trimmed); i++ {
		if int(trimmed[i])-48 < 0 || int(trimmed[i])-48 > 9 || checkDigit < 0 || checkDigit > 10 {
			return false
		}
		ISBN += (int(trimmed[0]) - 48) * (len(trimmed) - i)
	}
	ISBN += checkDigit
	if ISBN%11 == 0 {
		return true
	}
	return false
}
