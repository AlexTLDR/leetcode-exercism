package isbn

import "strings"

func IsValidISBN(isbn string) bool {
	trimmed := strings.Replace(isbn, "-", "", -1)
	if len(trimmed) != 10 {
		return false
	}
	var checkDigit int
	if trimmed[9] == 'X' {
		checkDigit = 10
	} else {
		checkDigit = int(trimmed[9] - '0')
	}

	trimmed = trimmed[:9]

	var ISBN int
	for i := 0; i < len(trimmed); i++ {
		if int(trimmed[i])-'0' < 0 || int(trimmed[i])-'0' > 9 || checkDigit < 0 || checkDigit > 10 {
			return false
		}
		ISBN += (int(trimmed[i]) - '0') * (10 - i)
	}
	ISBN += checkDigit
	if ISBN%11 == 0 {
		return true
	}
	return false
}