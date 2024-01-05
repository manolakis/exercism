package romannumerals

import "errors"

var romanNumbers = map[int]string{
	0:    "",
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("the number is too small or too big")
	}

	return handle(&input, 1000) +
		handle(&input, 100) +
		handle(&input, 10) +
		handle(&input, 1), nil
}

func handle(input *int, unit int) string {
	var romanValue string

	for value := *input / unit; value > 0; value = *input / unit {
		switch {
		case value == 9:
			romanValue += romanNumbers[unit] + romanNumbers[unit*10]
			*input -= unit * 9
		case value >= 5:
			romanValue += romanNumbers[unit*5]
			*input -= unit * 5
		case value == 4:
			romanValue += romanNumbers[unit] + romanNumbers[unit*5]
			*input -= unit * 4
		default:
			romanValue += romanNumbers[unit]
			*input -= unit
		}
	}

	return romanValue
}
