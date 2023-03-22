package main

import "fmt"

var count = 0

func main() {
	fmt.Println(CollatzConjecture(12))
}

func CollatzConjecture(n int) (int, error) {
	switch {
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
