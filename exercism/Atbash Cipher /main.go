// https://exercism.org/tracks/go/exercises/atbash-cipher

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Atbash("x123 yes"))
}

// func Atbash(s string) string {
// 	var alphabetMap = map[string]string{
// 		"a": "z",
// 		"b": "y",
// 		"c": "x",
// 		"d": "w",
// 		"e": "v",
// 		"f": "u",
// 		"g": "t",
// 		"h": "s",
// 		"i": "r",
// 		"j": "q",
// 		"k": "p",
// 		"l": "o",
// 		"m": "n",
// 		"n": "m",
// 		"o": "l",
// 		"p": "k",
// 		"q": "j",
// 		"r": "i",
// 		"s": "h",
// 		"t": "g",
// 		"u": "f",
// 		"v": "e",
// 		"w": "d",
// 		"x": "c",
// 		"y": "b",
// 		"z": "a",
// 	}

// 	for _, v := range s {
// 		for w := range alphabetMap {
// 			if strings.ToLower(string(v)) == w {
// 				s = strings.Replace(s, w, alphabetMap[w], 1)
// 			}
// 		}
// 	}
// 	return s

// }

func Atbash(s string) string {
	s = strings.ToLower(s)
	cipherText := ""
	groupedText := ""
	count := 0

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			cipherChar := 'z' - (char - 'a')
			cipherText += string(cipherChar)
		} else {
			cipherText += string(char)
		}

		if count == 4 {
			groupedText += cipherText + " "
			cipherText = ""
			count = 0
		} else {
			count++
		}
	}

	if cipherText != "" {
		groupedText += cipherText
	}

	return groupedText

}
