package isbn

import (
	"regexp"
	"strings"
)

func IsValidISBN(isbn string) bool {
	reg := regexp.MustCompile("^[0-9]{9}[0-9X]$")
	isbn = strings.Replace(isbn, "-", "", -1)
	if !reg.MatchString(isbn) {
		return false
	}

	digit := 0
	ISBN := 0

	for i := 0; i < len(isbn); i++ {
		if isbn[i] == 'X' {
			digit = 10
		} else {
			digit = int(isbn[i] - '0')
		}
		ISBN += digit * (len(isbn) - i)
	}
	return ISBN%11 == 0
}