package grains

import (
    "math"
    "errors"
    )

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		err := errors.New("not a valid chess square")
		return uint64(math.Pow(2, float64(number-1))), err
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
var grains uint64 = 1
	return (grains << 64) - 1
}

