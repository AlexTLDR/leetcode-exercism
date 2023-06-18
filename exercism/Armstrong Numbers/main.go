// https://exercism.org/tracks/go/exercises/armstrong-numbers

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(IsNumber(153))
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
