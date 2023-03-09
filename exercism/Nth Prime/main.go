package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Nth(600))
}

func Nth(n int) (int, error) {
	counter := 0
	i := 1
	for {
		i++
		if IsPrime(i) {
			counter++
		}
		if counter == n {
			return i, fmt.Errorf("")
		}
	}
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
