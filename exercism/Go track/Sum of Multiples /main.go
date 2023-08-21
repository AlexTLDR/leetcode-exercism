// https://exercism.org/tracks/go/exercises/sum-of-multiples

package main

import "fmt"

func main() {
	fmt.Println(SumMultiples(20, 3, 5))
}

func SumMultiples(limit int, divisors ...int) int {
	var sum int
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if divisor == 0 {
				continue
			}
			if i%divisor == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
