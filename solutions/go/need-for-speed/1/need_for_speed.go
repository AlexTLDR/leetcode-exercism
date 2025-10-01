package speed

// TODO: define the 'Car' type struct

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
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
