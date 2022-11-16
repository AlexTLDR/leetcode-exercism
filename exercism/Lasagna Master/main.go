// https://exercism.org/tracks/go/exercises/lasagna-master

package main

import "fmt"

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	quantities := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println(PreparationTime(layers, 0))
	fmt.Println(Quantities(quantities))
}

func PreparationTime(layers []string, preparation int) int {
	if preparation == 0 {
		preparation = 2
	}
	return len(layers) * preparation
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		switch v {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}
