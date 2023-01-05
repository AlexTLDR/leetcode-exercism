// https://exercism.org/tracks/go/exercises/animal-magic

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	SeedWithTime()
	fmt.Println(RollADie())
	fmt.Println(GenerateWandEnergy())
}

// SeedWithTime seeds math/rand with the current computer time.
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	//SeedWithTime()
	min, max := 1, 20
	n := min + rand.Intn(max-min+1)
	return n
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	min, max := 0.0, 12.0
	return min + rand.Float64()*(max-min)
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	panic("Please implement the ShuffleAnimals function")
}
