//https://exercism.org/tracks/go/exercises/resistor-color-trio/edit

package main

import (
	"fmt"
	"strconv"
)

var codes = []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}

func main() {

	fmt.Println(Label([]string{"orange", "orange", "black"}))

}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	if len(colors) != 3 {
		return "Please input only 3 colors"
	}
	var resistance string
	for i := 0; i < 2; i++ {
		for value, code := range codes {
			if colors[i] == code {
				resistance = resistance + strconv.Itoa(value)
			}
		}
	}
	return resistance
}
