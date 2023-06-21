// https://exercism.org/tracks/go/exercises/atbash-cipher

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Atbash("x123 yes"))
}

func Atbash(s string) string {
	s = strings.ToLower(s)
	cipherText := strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return 'z' - (r - 'a')
		case r >= '0' && r <= '9':
			return r
		default:
			return -1
		}
	}, s)

	groupSize := 5
	groupedText := ""
	for i, char := range cipherText {
		groupedText += string(char)
		if (i+1)%groupSize == 0 && i != len(cipherText)-1 {
			groupedText += " "
		}
	}

	return groupedText
}
