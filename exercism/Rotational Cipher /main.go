// https://exercism.org/tracks/go/exercises/rotational-cipher

package main

import "fmt"

func main() {
	fmt.Println(RotationalCipher("abcdefghijklmnopqrstuvwxyz", 13))
}

func RotationalCipher(plain string, shiftKey int) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	shiftedAlphabet := make(map[string]string)
	for i, letter := range alphabet {
		shiftedAlphabet[string(letter)] = string(alphabet[(i+shiftKey)%26])
	}

	encriptedMessage := ""
	for _, letter := range plain {
		if string(letter) == " " {
			encriptedMessage += " "
		} else {
			encriptedMessage += shiftedAlphabet[string(letter)]
		}
	}
	return encriptedMessage
}
