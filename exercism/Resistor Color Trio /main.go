//https://exercism.org/tracks/go/exercises/resistor-color-trio/edit

package main

import (
	"fmt"
	"math"
)

var codes = []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}

func main() {

	fmt.Println(Label([]string{"orange", "red", "white"}))

}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	if len(colors) != 3 {
		return "Please input only 3 colors"
	}
	var resistance int
	//var unit string
	//var zeroes string

	for value, code := range codes {
		if colors[0] == code {
			resistance = value * 10
		}

		// if colors[2] == code {
		// 	zeroes = strings.Repeat("0", value)
		// }
	}

	for value, code := range codes {
		if colors[1] == code {
			resistance += value
		}
	}

	for value, code := range codes {
		if colors[2] == code {
			resistance *= int(math.Pow(10, float64(value)))
		}
	}

	//intResistance, _ := strconv.Atoi(resistance + zeroes)
	fmt.Println(resistance)
	return ""
}
