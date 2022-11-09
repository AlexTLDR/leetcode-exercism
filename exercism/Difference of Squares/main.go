// https://exercism.org/tracks/go/exercises/difference-of-squares

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SquareOfSum(10))
	fmt.Println(SumOfSquares(10))
	fmt.Println(Difference(10))
}

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

/*
A better approach is to use Gauss' formula to calculate the sum of the numbers;

	    n(n+1)
	    ------
	S =   2

	and the square pyramidal number is the sum of the squares

			n(n + 1)(2n + 1)
			---------------
	S = 			 6

	Which yields O(1) time requirement and O(1) space requirement.


func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

*/
