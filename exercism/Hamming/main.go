// https://exercism.org/tracks/go/exercises/hamming

package main

import "fmt"

func main() {
	a := "GAGCCTACTAACGGGAT"
	b := "CATCGTAATGACGGCCT"
	fmt.Println(Distance(a, b))
}

func Distance(a, b string) (int, error) {
	return 0, nil
}
