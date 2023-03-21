package main

import "fmt"

var colors = map[string]int{
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
	panic("Implement the Value function")
}
