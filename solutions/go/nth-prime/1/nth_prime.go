package prime

import (
	"math"
    "fmt"
)


func Nth(n int) (int, error) {
	if n == 0{
		return 0, fmt.Errorf("There is no 0th prime")
	}
	counter := 0
	for i := 2; i <= 104743; i++ {
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
