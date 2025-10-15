package diffsquares

import "math"

func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		square := SquareInt(i)
		sum += square
	}
	return sum
}

func SquareInt(n int) int {
	result := math.Pow(float64(n), 2)
	return int(result)
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
