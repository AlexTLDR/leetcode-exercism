// https://exercism.org/tracks/go/exercises/armstrong-numbers

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(IsNumber(153))
	fmt.Println(isArmstrong(153))
}

func IsNumber(n int) bool {
	var sum float64
	strNum := strconv.Itoa(n)
	for _, v := range strNum {
		sum += math.Pow(float64(int(v-48)), float64(len(strNum)))
	}
	if n == int(sum) {
		return true
	}
	return false
}

func isArmstrong(num int) bool {
	numOfDigits := 0
	temp := num
	for temp > 0 {
		temp /= 10
		numOfDigits++
	}
	sum := 0
	temp = num
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numOfDigits)))
		temp /= 10
	}
	return sum == num
}
