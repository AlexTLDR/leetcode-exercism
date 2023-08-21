// https://exercism.org/tracks/go/exercises/grains

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println(Total())
	fmt.Println(Square(4))
}

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		err := errors.New("not a valid chess square")
		return uint64(math.Pow(2, float64(number-1))), err
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	return (1 << 64) - 1
}
