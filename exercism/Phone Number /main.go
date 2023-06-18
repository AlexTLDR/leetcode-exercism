// https://exercism.org/tracks/go/exercises/phone-number

package main

import "fmt"

func main() {
	phoneNumber := " 223 456   7890   "
	fmt.Println(Number(phoneNumber))
	fmt.Println(AreaCode(phoneNumber))
	fmt.Println(Format(phoneNumber))

}

func Number(phoneNumber string) (string, error) {
	var cleanNo string
	for _, char := range phoneNumber {
		if char >= 48 && char <= 57 {
			cleanNo += string(char)
		}
	}
	if len(cleanNo) > 11 || len(cleanNo) < 10 {
		return "", fmt.Errorf("Invalid phone number")
	}
	if len(cleanNo) == 11 && cleanNo[0] != 49 {
		return "", fmt.Errorf("Invalid phone number")
	}
	if len(cleanNo) == 11 {
		cleanNo = cleanNo[1:]
	}
	if cleanNo[0] < 50 || cleanNo[3] < 50 {
		return "", fmt.Errorf("Invalid phone number")
	}
	return cleanNo, nil
}

func AreaCode(phoneNumber string) (string, error) {
	areaCode, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	areaCode = areaCode[:3]
	return areaCode, err

}

func Format(phoneNumber string) (string, error) {
	format, err := Number(phoneNumber)
	format = "(" + format[:3] + ")" + " " + format[3:6] + "-" + format[6:]
	return format, err
}
