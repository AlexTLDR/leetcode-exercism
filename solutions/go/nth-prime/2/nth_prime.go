package prime

import (
	"math"
    "fmt"
)


func Nth(n int) (int, error) {
	if n == 0 {
		return 0, fmt.Errorf("There is no 0th prime")
	}
	counter := 0
	i := 1
	for {
		i++
		if IsPrimeCheck(i) {
			counter++
		}
		if counter == n {
			break

		}
	}
	return i, nil
}

func IsPrimeCheck(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}