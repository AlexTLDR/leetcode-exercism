package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
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

