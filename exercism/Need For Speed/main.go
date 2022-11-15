// https://exercism.org/tracks/go/exercises/need-for-speed

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
	batteryDrain := 2
	speed := 5

	car := NewCar(speed, batteryDrain)
	fmt.Println(car)

	distance := 800
	track := NewTrack(distance)
	fmt.Println(track)

	fmt.Println(Drive(car))
	fmt.Println(CanFinish(car, track))
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
	if car.battery/car.batteryDrain >= 1 {
		return true
	}
	return false
}
