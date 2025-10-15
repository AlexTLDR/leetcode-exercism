package resistorcolor

import "strings"

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


// Colors should return the list of all colors.
func Colors() []string {
	colors := []string{}
	for v := range scheme {
		colors = append(colors, strings.ToLower(v))
	}
	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for s, i := range scheme {
		if strings.ToLower(s) == color {
			return i
		}
	}
	return -1
}
