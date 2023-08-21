// https://exercism.org/tracks/go/exercises/lasagna-master

package main

import "fmt"

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	quantitiesStrings := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	quantitiesFloats := []float64{1.2, 3.6, 10.5}
	scaledQuantities := ScaleRecipe(quantitiesFloats, 4)
	fmt.Println(PreparationTime(layers, 0))
	fmt.Println(Quantities(quantitiesStrings))
	AddSecretIngredient(friendsList, myList)
	fmt.Println(myList)
	fmt.Println(scaledQuantities)
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

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) (scaledQuantities []float64) {
	for _, v := range quantities {
		scaledQuantities = append(scaledQuantities, v*float64(portions)/2)
	}
	return
}
