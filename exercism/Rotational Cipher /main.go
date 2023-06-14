// https://exercism.org/tracks/go/exercises/rotational-cipher

package main

import "fmt"

func main() {
	fmt.Println(RotationalCipher("abcdefghijklmnopqrstuvwxyz", 13))
}

func RotationalCipher(plain string, shiftKey int) string {
	capitalAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerAlphabet := "abcdefghijklmnopqrstuvwxyz"
	shiftedCapitalAlphabet := make(map[string]string)
	shiftedLowerAlphabet := make(map[string]string)
	for i, letter := range capitalAlphabet {
		shiftedCapitalAlphabet[string(letter)] = string(capitalAlphabet[(i+shiftKey)%26])
	}

	for i, letter := range lowerAlphabet {
		shiftedLowerAlphabet[string(letter)] = string(lowerAlphabet[(i+shiftKey)%26])
	}

	encriptedMessage := ""
	for _, letter := range plain {
		switch {
		case string(letter) == " ":
			encriptedMessage += " "
		case letter >= 'A' && letter <= 'Z':
			encriptedMessage += shiftedCapitalAlphabet[string(letter)]
		case letter >= 'a' && letter <= 'z':
			encriptedMessage += shiftedLowerAlphabet[string(letter)]
		}
	}
	return encriptedMessage
}
