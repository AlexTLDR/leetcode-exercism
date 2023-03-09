package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Nth(6))
}

func Nth(n int) (int, error) {
	counter := 0
	for i := 2; i < int(math.Pow(2, 64)); i++ {
		sqRoot := int(math.Sqrt(float64(i)))
		for j := 2; j <= sqRoot; j++ {
			if i%j != 0 {
				counter++
			}
			if n == counter {
				return i, fmt.Errorf("")
			}
		}

	}
	return 0, fmt.Errorf("no prime number")
}
