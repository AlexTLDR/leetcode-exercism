package main

import "fmt"

func main() {

	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	var roman = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	if s == "" {
		return 0
	}

	var number, lastint, total int
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		number = roman[char]
		if number < lastint {
			total = total - number
		} else {
			total = total + number
		}
		lastint = number
	}
	return total
}
