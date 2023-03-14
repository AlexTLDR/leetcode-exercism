package main

import "fmt"

var scheme = map[string]int{
	"Black":  0,
	"Brown":  1,
	"Red":    2,
	"Orange": 3,
	"Yellow": 4,
	"Green":  5,
	"Blue":   6,
	"Violet": 7,
	"Grey":   8,
	"White":  9,
}

func main() {
	fmt.Println(Colors())
	fmt.Println(ColorCode("Green"))
}

// Colors should return the list of all colors.
func Colors() []string {
	colors := []string{}
	for v := range scheme {
		colors = append(colors, v)
	}
	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for s, i := range scheme {
		if s == color {
			return i
		}
	}
	return -1
}
