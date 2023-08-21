// https://exercism.org/tracks/go/exercises/hamming

package main

import (
	"errors"
	"fmt"
)

func main() {
	a := "GAGCCTACTAACGGGAT"
	b := "CATCGTAATGACGGCCT"
	fmt.Println(Distance(a, b))
}

func Distance(a, b string) (int, error) {
	hammingDistance := 0
	if len(a) != len(b) {
		return -1, errors.New("DNA strands not of equal")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDistance += 1
		}
	}
	return hammingDistance, nil

}
