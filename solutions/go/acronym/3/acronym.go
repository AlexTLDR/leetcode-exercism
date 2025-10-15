package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	abbv := string(s[0])

	for i := 1; i < len(s); i++ {
		if s[i] == '\'' {
			i += 2
		}

		prevCharIsLetter := unicode.IsLetter(rune(s[i-1]))
		currentCharIsLetter := unicode.IsLetter(rune(s[i]))

		if !prevCharIsLetter && currentCharIsLetter {
			abbv += string(s[i])
		}
	}

	return strings.ToUpper(abbv)
}
