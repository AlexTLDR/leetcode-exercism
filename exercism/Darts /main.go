// https://exercism.org/tracks/go/exercises/darts

package main

import "fmt"

func main() {
	fmt.Println(Score(1, 1))
}

func Score(x, y float64) int {
	score := 0.0
	switch {
	case x >= y:
		score = x
	default:
		score = y
	}

	switch {
	case score > 10:
		score = 0
	case score > 5 && score <= 10:
		score = 1
	case score > 1 && score <= 5:
		score = 5
	default:
		score = 10
	}
	return int(score)
}
