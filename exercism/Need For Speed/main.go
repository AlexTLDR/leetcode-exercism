// https://exercism.org/tracks/go/exercises/need-for-speed

package main

import "fmt"

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	fmt.Println(car)
}

func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain}
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
// func NewTrack(distance int) Track {
// 	panic("Please implement the NewTrack function")
// }

// // Drive drives the car one time. If there is not enough battery to drive one more time,
// // the car will not move.
// func Drive(car Car) Car {
// 	panic("Please implement the Drive function")
// }

// // CanFinish checks if a car is able to finish a certain track.
// func CanFinish(car Car, track Track) bool {
// 	panic("Please implement the CanFinish function")
// }
