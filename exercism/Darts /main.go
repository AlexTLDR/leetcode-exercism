// https://exercism.org/tracks/go/exercises/darts

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Score(1, 1))
}

func Score(x, y float64) int {
	score := math.Pow(x, 2) + math.Pow(y, 2)
	switch {
	case score > 100:
		return 0
	case score > 25:
		return 1
	case score > 1:
		return 5
	}

	return 10
}
