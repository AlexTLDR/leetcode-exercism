package main

import (
	"fmt"
	"strconv"
)

var codes = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func main() {
	fmt.Println(Value([]string{"brown", "green"}))
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	var resistance string
	for _, color := range colors {
		for code, value := range codes {
			if color == code {
				resistance = resistance + strconv.Itoa(value)
			}
		}
	}
	result, _ := strconv.Atoi(resistance)
	return result
}
