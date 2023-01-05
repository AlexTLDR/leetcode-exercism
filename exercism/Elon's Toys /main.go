// https://exercism.org/tracks/go/exercises/elons-toys

/*
Note: This exercise is a continuation of the need-for-speed exercise. Until line 65 I have copied the need-for-speed exercise.
*/

package main

import "fmt"

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

func main() {
	//Car.battery := 100
	batteryDrain := 3
	speed := 5

	car := NewCar(speed, batteryDrain)
	fmt.Println("car:", car)

	distance := 800
	track := NewTrack(distance)
	fmt.Println("track:", track)

	fmt.Println("drive:", Drive(car))
	fmt.Println("can finish:", CanFinish(car, track))
}

func NewCar(speed, batteryDrain int) Car {
	return Car{battery: 100,
		batteryDrain: batteryDrain,
		speed:        speed,
	}
}

func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	td := car.battery / car.batteryDrain
	maxDistance := car.speed * td
	return maxDistance >= track.distance
}

// Elon's Toys exercise start

// TODO: define the 'Drive()' method

func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// TODO: define the 'DisplayDistance() string' method

// TODO: define the 'DisplayBattery() string' method

// TODO: define the 'CanFinish(trackDistance int) bool' method
