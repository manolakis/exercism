package collatzconjecture

import (
	"errors"
	"fmt"
)

func CollatzConjecture(number int) (int, error) {
	if number <= 0 {
		return 0, errors.New(fmt.Sprintf("%d is not a valid number", number))
	}

	return collatzConjecture(number), nil
}

func collatzConjecture(number int) int {
	if number == 1 {
		return 0
	}

	if isEven(number) {
		return 1 + collatzConjecture(number/2)
	}

	return 1 + collatzConjecture(number*3+1)
}

func isEven(number int) bool {
	return number%2 == 0
}
