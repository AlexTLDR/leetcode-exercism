package armstrong

import (
	"math"
)

func IsNumber(n int) bool {
		numOfDigits := 0
	temp := n
	for temp > 0 {
		temp /= 10
		numOfDigits++
	}
	sum := 0
	temp = n
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numOfDigits)))
		temp /= 10
	}
	return sum == n
}