package main

import "fmt"

func main() {
	fmt.Println(CollatzConjecture(12))
}

func CollatzConjecture(n int) (int, error) {
	steps := 0
	fmt.Println(n)
	switch {
	case n == 1:
		return steps, nil
	case n%2 == 0:
		steps++
		n /= 2
	default:
		steps++
		n = n*3 + 1
	}
	CollatzConjecture(n)
	return steps, nil
}
