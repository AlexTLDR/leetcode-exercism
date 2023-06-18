// https://exercism.org/tracks/go/exercises/phone-number

package main

import "fmt"

func main() {
	phoneNumber := "+1 (613)-995-0253"
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
	return cleanNo, nil
}

func AreaCode(phoneNumber string) (string, error) {
	panic("Please implement the AreaCode function")

}

func Format(phoneNumber string) (string, error) {
	panic("Please implement the Format function")
}
