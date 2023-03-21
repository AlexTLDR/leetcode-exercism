package main

import (
	"fmt"
	"strconv"
)

var codes = []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}

func main() {
	fmt.Println(Value([]string{"brown", "green", "violet"}))
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	var resistance string
	for i := 0; i < 2; i++ {
		fmt.Println(codes[i])
	}
	// for _, color := range colors {
	// 	for code, value := range codes {
	// 		if color == code {
	// 			resistance = resistance + strconv.Itoa(value)
	// 		}
	// 	}
	// }
	result, _ := strconv.Atoi(resistance)
	return result
}
