package luhn

import (
	"errors"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	digits, err := normalize(id)
	if err != nil {
		return false
	}

	secondDigitsDoubled := doubleEverySecondDigit(digits)
	sumOfDigits := sumDigits(secondDigitsDoubled)

	return isDivisibleByTen(sumOfDigits)
}

func normalize(id string) ([]int, error) {
	value := strings.ReplaceAll(id, " ", "")

	if containsInvalidCharacters(value) {
		return nil, errors.New("the input contains invalid characters")
	}

	return toIntSlice(value), nil
}

func containsInvalidCharacters(id string) bool {
	for _, value := range id {
		if !unicode.IsDigit(value) {
			return true
		}
	}

	return len(id) <= 1
}

func toIntSlice(id string) []int {
	return mapSlice[rune, int]([]rune(id), func(r rune, index int) int {
		return int(r - '0')
	})
}

func doubleEverySecondDigit(input []int) []int {
	isSecondDigit := len(input)%2 != 0

	return mapSlice[int, int](input, func(digit int, index int) int {
		isSecondDigit = !isSecondDigit

		if isSecondDigit {
			value := digit * 2
			if value > 9 {
				value -= 9
			}

			return value
		}

		return digit
	})
}

func mapSlice[T any, M any](input []T, f func(T, int) M) []M {
	output := make([]M, len(input))

	for index, value := range input {
		output[index] = f(value, index)
	}

	return output
}

func sumDigits(input []int) int {
	sum := 0
	for _, digit := range input {
		sum += digit
	}

	return sum
}

func isDivisibleByTen(input int) bool {
	return input%10 == 0
}
