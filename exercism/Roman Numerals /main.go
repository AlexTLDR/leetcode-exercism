// https://exercism.org/tracks/go/exercises/roman-numerals

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(ToRomanNumeral(3145))
}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", fmt.Errorf("The number %d is undefined", input)
	}
	hundreds := map[int]string{
		1: "C",
		2: "CC",
		3: "CCC",
		4: "CD",
		5: "D",
		6: "DC",
		7: "DCC",
		8: "DCCC",
		9: "CM",
	}
	tens := map[int]string{
		1: "X",
		2: "XX",
		3: "XXX",
		4: "XL",
		5: "L",
		6: "LX",
		7: "LXX",
		8: "LXXX",
		9: "XC",
	}
	numbers := map[int]string{
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
		6: "VI",
		7: "VII",
		8: "VIII",
		9: "IX",
	}
	roman := ""
	stringInput := strconv.Itoa(input)
	switch len(stringInput) {
	case 3:
		stringInput = "0" + stringInput
	case 2:
		stringInput = "00" + stringInput
	case 1:
		stringInput = "000" + stringInput
	}
	for i, v := range stringInput {
		s := string(v)
		si, _ := strconv.Atoi(s)
		switch i {
		case 0:
			roman = strings.Repeat("M", si)
		case 1:
			roman = roman + hundreds[si]
		case 2:
			roman = roman + tens[si]
		case 3:
			roman = roman + numbers[si]
		}

	}
	return roman, nil
}
