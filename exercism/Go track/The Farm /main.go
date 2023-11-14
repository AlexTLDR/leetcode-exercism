// https://exercism.org/tracks/go/exercises/the-farm

package thefarm

import (
	"fmt"
)

type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}

func DivideFood(fodderCalculator FodderCalculator, numCows int) (float64, error) {
	totalFodder, err := fodderCalculator.FodderAmount(numCows)
	if err != nil {
		return 0, err
	}

	factor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	foodPerCow := totalFodder / float64(numCows) * factor
	return foodPerCow, nil
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0, fmt.Errorf("invalid number of cows")
	}

	return DivideFood(fodderCalculator, numCows)
}
