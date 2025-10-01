package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	workingCarsPerHour := float64(productionRate) * successRate / 100
	return workingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerMinute := float64(productionRate) * successRate / 100 / 60
	return int(workingCarsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
		tenGroup := carsCount / 10
	remainingGroup := carsCount % 10
	return uint(tenGroup)*95000 + uint(remainingGroup)*10000
}
