// https://exercism.org/tracks/go/exercises/acronym

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Abbreviate("Halley's Comet"))
}

func Abbreviate(s string) string {
	var abbv string
	abbv += string(s[0])

	for i := 1; i < len(s); i++ {
		if s[i] == 39 {
			i++
			i++
		}
		if !unicode.IsLetter(rune(s[i-1])) && unicode.IsLetter(rune(s[i])) {
			abbv += string(s[i])

		}
	}
	return strings.ToUpper(abbv)
}
