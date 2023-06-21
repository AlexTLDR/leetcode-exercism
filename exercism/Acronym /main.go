// https://exercism.org/tracks/go/exercises/acronym

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Abbreviate("Liquid-crystal display"))
}

func Abbreviate(s string) string {
	var abbv string
	abbv += string(s[0])
	for i := 1; i < len(s); i++ {
		if !unicode.IsLetter(rune(s[i])) {
			abbv += string(s[i+1])
		}
	}
	return strings.ToUpper(abbv)
}
