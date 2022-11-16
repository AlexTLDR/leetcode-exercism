// https://exercism.org/tracks/go/exercises/lasagna-master

package main

import "fmt"

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println(PreparationTime(layers, 0))
}

func PreparationTime(layers []string, preparation int) int {
	if preparation == 0 {
		preparation = 2
	}
	return len(layers) * preparation
}
