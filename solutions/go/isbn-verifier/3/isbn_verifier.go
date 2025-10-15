package isbn

import "strings"

func IsValidISBN(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	if len(isbn) != 10 {
		return false
	}
	var checkDigit int
	if isbn[9] == 'X' {
		checkDigit = 10
	} else {
		checkDigit = int(isbn[9] - '0')
	}

	isbn = isbn[:9]

	var ISBN int
	for i := 0; i < len(isbn); i++ {
		if int(isbn[i])-'0' < 0 || int(isbn[i])-'0' > 9 || checkDigit < 0 || checkDigit > 10 {
			return false
		}
		ISBN += (int(isbn[i]) - '0') * (10 - i)
	}
	ISBN += checkDigit
	if ISBN%11 == 0 {
		return true
	}
	return false
}