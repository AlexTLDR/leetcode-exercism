package lasagna

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
