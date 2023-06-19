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
	shiftedCapitalAlphabet, shiftedLowerAlphabet := shift("abcdefghijklmnopqrstuvwxyz")

	encriptedMessage := ""
	for _, letter := range s {
		switch {
		case nonLetter(letter):
			encriptedMessage += string(letter)
		case capital(letter):
			encriptedMessage += shiftedCapitalAlphabet[string(letter)]
		case lowerCase(letter):
			encriptedMessage += shiftedLowerAlphabet[string(letter)]
		}
	}
	return encriptedMessage
}

func shift(alphabet string) (map[string]string, map[string]string) {
	upperShifted := make(map[string]string)
	lowerShifted := make(map[string]string)
	for i, letter := range alphabet {
		upperShifted[strings.ToUpper(string(letter))] = strings.ToUpper(string(alphabet[len(alphabet)-i]))
		lowerShifted[string(letter)] = string(alphabet[len(alphabet)-i])
	}
	return upperShifted, lowerShifted
}

func capital(letter rune) bool {
	return letter >= 'A' && letter <= 'Z'
}

func lowerCase(letter rune) bool {
	return letter >= 'a' && letter <= 'z'
}

func nonLetter(symbol rune) bool {
	return symbol < 'A' || (symbol > 'Z' && symbol < 'a') || (symbol > 'z')
}
