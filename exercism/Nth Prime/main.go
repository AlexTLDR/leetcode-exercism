package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Nth(6))
}

func Nth(n int) (int, error) {
	sqRoot := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqRoot; i++ {
		if n%i == 0 {
			fmt.Println("Non Prime Number")
			return 0, fmt.Errorf("Non Prime Number")
		}
	}
	return n, fmt.Errorf("")
}
