package rotationalcipher

import (
	"strings"
)


func RotationalCipher(plain string, shiftKey int) string {
	shiftedCapitalAlphabet, shiftedLowerAlphabet := shift("abcdefghijklmnopqrstuvwxyz", shiftKey)

	encriptedMessage := ""
	for _, letter := range plain {
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

func shift(alphabet string, shiftKey int) (map[string]string, map[string]string) {
	upperShifted := make(map[string]string)
	lowerShifted := make(map[string]string)
	for i, letter := range alphabet {
		upperShifted[strings.ToUpper(string(letter))] = strings.ToUpper(string(alphabet[(i+shiftKey)%26]))
		lowerShifted[string(letter)] = string(alphabet[(i+shiftKey)%26])
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
