package luhn

import (
    "strings"
    "unicode"
    )

func Valid(id string) bool {
	trimmed := strings.Replace(id, " ", "", -1)

	if len(trimmed) < 2 {
		return false
	}
	for _, v := range trimmed {
		if !unicode.IsDigit(v) {
			return false
		}
	}
	var sum int
	var idArr []int

	for _, v := range trimmed {
		idArr = append(idArr, int(v-48))
	}
	switch len(idArr)%2 == 0 {
	case true:

		for i, v := range idArr {
			if i == 0 || i%2 == 0 {
				v = doubleDigitValue(v)
			}
			sum += v
		}
		if sum%10 == 0 {
			return true
		}
		return false

	case false:
		for i, v := range idArr {
			if i == 1 || i%2 != 0 {
				v = doubleDigitValue(v)
			}
			sum += v
		}
		if sum%10 == 0 {
			return true
		}

	}

	return false
}

func doubleDigitValue(digit int) (value int) {
	doubleDigit := digit * 2

	if doubleDigit > 9 {
		value = doubleDigit - 9
	} else {
		value = doubleDigit
	}
	return value
}