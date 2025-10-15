package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	var abbv string
	abbv += string(s[0])

	for i := 1; i < len(s); i++ {
		if s[i] == 39 {
			i+=2
		}
		if !unicode.IsLetter(rune(s[i-1])) && unicode.IsLetter(rune(s[i])) {
			abbv += string(s[i])

		}
	}
	return strings.ToUpper(abbv)
}
