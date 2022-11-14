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
	/* brute force
	var chessBoard []uint64
	total := 0.0
	for i := 0; i < 64; i++ {
		chessBoard = append(chessBoard, uint64(i))
	}

	for i, v := range chessBoard {
		total += math.Pow(float64(v), float64(i))
	}
	return uint64(total)
	*/

	// Mathematical Solution
	return uint64(math.Pow(2, 64) - 1)
}
