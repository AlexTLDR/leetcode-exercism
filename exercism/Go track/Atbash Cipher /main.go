// https://exercism.org/tracks/go/exercises/atbash-cipher

package main

import (
	"fmt"
	"strings"
	"unicode"
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

func Atbash2(s string) string {

	var alphabetMap = map[rune]rune{
		'a': 'z',
		'b': 'y',
		'c': 'x',
		'd': 'w',
		'e': 'v',
		'f': 'u',
		'g': 't',
		'h': 's',
		'i': 'r',
		'j': 'q',
		'k': 'p',
		'l': 'o',
		'm': 'n',
		'n': 'm',
		'o': 'l',
		'p': 'k',
		'q': 'j',
		'r': 'i',
		's': 'h',
		't': 'g',
		'u': 'f',
		'v': 'e',
		'w': 'd',
		'x': 'c',
		'y': 'b',
		'z': 'a',
		'0': '0',
		'1': '1',
		'2': '2',
		'3': '3',
		'4': '4',
		'5': '5',
		'6': '6',
		'7': '7',
		'8': '8',
		'9': '9',
	}

	s = strings.ToLower(s)
	cipherText := strings.Map(func(r rune) rune {
		if val, ok := alphabetMap[r]; ok {
			return val
		}
		return -1
	}, s)

	groupSize := 5
	groupedText := ""
	for i, char := range cipherText {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			continue
		}
		groupedText += string(char)
		if (i+1)%groupSize == 0 && i != len(cipherText)-1 {
			groupedText += " "
		}
	}

	return groupedText
}
