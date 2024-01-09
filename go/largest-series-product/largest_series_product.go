package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var largest int64
	runes := []rune(digits)

	if span <= 0 {
		return 0, errors.New("span must be positive")
	}

	if len(runes) < span {
		return 0, errors.New("span longer that string")
	}

	for i := 0; i <= len(runes)-span; i++ {
		var value int64 = 1

		for j := i; j < i+span; j++ {
			digit, err := strconv.Atoi(string(runes[j]))
			if err != nil {
				return 0, err
			}

			value *= int64(digit)
		}

		if value > largest {
			largest = value
		}
	}

	return largest, nil
}
