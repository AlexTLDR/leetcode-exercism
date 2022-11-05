// https://exercism.org/tracks/go/exercises/cars-assemble

package main

import "fmt"

func main() {
	fmt.Println(CalculateWorkingCarsPerHour(1547, 90))
	fmt.Println(CalculateWorkingCarsPerMinute(1105, 90))
	fmt.Println(CalculateCost(37))
	fmt.Println(CalculateCost(21))
}

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// workingCarsPerHour := float64(productionRate) * successRate / 100
	// return workingCarsPerHour
	return float64(productionRate) * successRate / 100
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// workingCarsPerMinute := float64(productionRate) * successRate / 100 / 60
	// return int(workingCarsPerMinute)
	return (int)(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

func CalculateCost(carsCount int) uint {
	tenGroup := carsCount / 10
	remainingGroup := carsCount % 10
	return uint(tenGroup)*95000 + uint(remainingGroup)*10000
}
