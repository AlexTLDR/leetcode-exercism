//https://exercism.org/tracks/go/exercises/resistor-color-trio/edit

package main

import (
	"fmt"
	"strconv"
)

var codes = []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}

func main() {

	fmt.Println(Label([]string{"red", "black", "red"}))

}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	if len(colors) != 3 {
		return "Please input only 3 colors"
	}
	var resistance int
	var unit string
	//var zeroes string

	for value, code := range codes {
		if colors[0] == code {
			resistance = value * 10
		}
	}

	for value, code := range codes {
		if colors[1] == code {
			resistance += value
		}
	}

	for value, code := range codes {
		if colors[2] == code {
			switch {
			case value == 0:
				unit = strconv.Itoa(resistance) + " ohms"
			case value == 1 || value == 2:
				unit = strconv.Itoa(resistance/10) + " kiloohms"
			case value >= 3 && value < 6:
				unit = strconv.Itoa(resistance) + " kiloohms"
			case value >= 6 && value < 9:
				unit = strconv.Itoa(resistance) + " megaohms"
			case value == 9:
				unit = strconv.Itoa(resistance) + " gigaohms"
			}
		}

	}

	return unit
}
