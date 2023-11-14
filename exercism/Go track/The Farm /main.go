// https://exercism.org/tracks/go/exercises/the-farm

package thefarm

import (
	"fmt"
)

type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}

func DivideFood(calculator FodderCalculator, numCows int) (float64, error) {
	total, err := calculator.FodderAmount(numCows)
	if err != nil {
		return 0, fmt.Errorf("error calculating total fodder: %w", err)
	}

	factor, err := calculator.FatteningFactor()
	if err != nil {
		return 0, fmt.Errorf("error retrieving fattening factor: %w", err)
	}

	return total / float64(numCows) * factor, nil
}

func ValidateInputAndDivideFood(calculator FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0, NewInvalidCowsError(numCows, "invalid number of cows")
	}

	return DivideFood(calculator, numCows)
}

type InvalidCowsError struct {
	num int
	msg string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.num, err.msg)
}

func NewInvalidCowsError(num int, msg string) *InvalidCowsError {
	return &InvalidCowsError{num, msg}
}

func ValidateNumberOfCows(num int) error {
	switch {
	case num < 0:
		return NewInvalidCowsError(num, "there are no negative cows")
	case num == 0:
		return NewInvalidCowsError(num, "no cows don't need food")
	default:
		return nil
	}
}
