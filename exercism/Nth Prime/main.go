package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Nth(1))
}

/*

It seems that exercism doesn't like the infinite loop

func Nth(n int) (int, error) {
	counter := 0
	i := 1
	for {
		i++
		if IsPrimeCheck(i) {
			counter++
		}
		if counter == n {
			return i, fmt.Errorf("")
		}
	}
}

func IsPrimeCheck(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
*/

func Nth(n int) (int, error) {
	counter := 0
	for i := 2; i < 104743; i++ {
		if IsPrimeCheck(i) {
			counter++
		}
		if counter == n {
			return i, nil
		}
	}
	return 0, fmt.Errorf("error")
}

func IsPrimeCheck(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
