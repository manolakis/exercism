package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	fodderAmount, err := fodderCalculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodderAmount / float64(cows) * fatteningFactor, nil
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fodderCalculator, cows)
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (invalidCowsError InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", invalidCowsError.cows, invalidCowsError.message)
}

func NewInvalidCowsError(cows int, message string) *InvalidCowsError {
	return &InvalidCowsError{cows: cows, message: message}
}

func ValidateNumberOfCows(cows int) error {
	switch {
	case cows < 0:
		return NewInvalidCowsError(cows, "there are no negative cows")
	case cows == 0:
		return NewInvalidCowsError(cows, "no cows don't need food")
	default:
		return nil
	}
}
