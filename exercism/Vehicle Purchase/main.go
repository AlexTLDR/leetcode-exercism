// https://exercism.org/tracks/go/exercises/vehicle-purchase

package main

import (
	"fmt"
)

func main() {
	fmt.Println(NeedsLicense("car"))
	fmt.Println(ChooseVehicle("Wuling Hongguang", "Toyota Corolla"))
	fmt.Println(CalculateResellPrice(1000, 1))
}

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"

}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	phrase := " is clearly the better choice."
	if option1 < option2 {
		return option1 + phrase
	}
	return option2 + phrase
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var reselPrice float64
	if age < 3 {
		reselPrice = originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		reselPrice = originalPrice * 0.7
	} else {
		reselPrice = originalPrice * 0.5
	}

	return reselPrice
}
