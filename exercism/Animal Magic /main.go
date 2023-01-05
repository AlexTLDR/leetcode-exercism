// https://exercism.org/tracks/go/exercises/animal-magic

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(RollADie())

}

// SeedWithTime seeds math/rand with the current computer time.
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	SeedWithTime()
	a, b := 1, 20
	n := a + rand.Intn(b-a+1)
	return n
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	panic("Please implement the GenerateWandEnergy function")
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	panic("Please implement the ShuffleAnimals function")
}
