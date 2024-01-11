package phonenumber

import (
	"errors"
	"fmt"
	"unicode"
)

type NorthAmericanNumberingPlan string

func (nanp NorthAmericanNumberingPlan) String() string {
	return string(nanp)
}

func (nanp NorthAmericanNumberingPlan) AreaCode() string {
	return string(nanp[:3])
}

func (nanp NorthAmericanNumberingPlan) Format() string {
	return fmt.Sprintf("(%s) %s-%s", nanp[:3], nanp[3:6], nanp[6:])
}

func NewNorthAmericanNumberingPlan(phoneNumber string) (NorthAmericanNumberingPlan, error) {
	var digits []rune
	for _, digit := range phoneNumber {
		if unicode.IsNumber(digit) {
			digits = append(digits, digit)
		}
	}

	if digits[0] == '1' {
		digits = digits[1:]
	}

	if len(digits) != 10 || (digits[0]-'0') < 2 || (digits[3]-'0') < 2 {
		return "", errors.New("invalid error")
	}

	return NorthAmericanNumberingPlan(digits), nil
}

func Number(phoneNumber string) (string, error) {
	number, err := NewNorthAmericanNumberingPlan(phoneNumber)
	if err != nil {
		return "", err
	}

	return number.String(), nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := NewNorthAmericanNumberingPlan(phoneNumber)
	if err != nil {
		return "", err
	}

	return number.AreaCode(), nil
}

func Format(phoneNumber string) (string, error) {
	number, err := NewNorthAmericanNumberingPlan(phoneNumber)
	if err != nil {
		return "", err
	}

	return number.Format(), nil
}
