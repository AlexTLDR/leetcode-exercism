package atbash

import (
	"strings"
	"unicode"
)



func Atbash(s string) string {

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



