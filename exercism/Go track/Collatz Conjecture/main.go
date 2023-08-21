//https://exercism.org/tracks/go/exercises/collatz-conjecture

package main

import (
	"errors"
	"fmt"
)

var count = 0

func main() {
	fmt.Println(CollatzConjecture(12))
	fmt.Println(CollatzConjecture2(12))
}

// since exercism is failing my recursive function I created CollatzConjecture2 without recursion to pass the exercism tests
func CollatzConjecture(n int) (int, error) {
	switch {
	case n < 1:
		return count, errors.New("n must be greater or equal to 1")
	case n == 1:
		return count, nil
	case n%2 == 0:
		n /= 2
	default:
		n = n*3 + 1
	}
	count++
	CollatzConjecture(n)
	return count, nil
}

func CollatzConjecture2(n int) (int, error) {
	steps := 0
	if n < 1 {
		return steps, errors.New("n must be greater or equal to 1")
	}

	for n != 1 {
		steps++
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return steps, nil
}
