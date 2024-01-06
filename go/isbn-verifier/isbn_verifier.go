package isbn

import (
	"regexp"
	"strconv"
	"strings"
)

var ValidatorRegex = regexp.MustCompile(`^\d{9}(\d|X)$`)

func IsValidISBN(isbn string) bool {
	value := strings.ReplaceAll(strings.ToUpper(strings.TrimSpace(isbn)), "-", "")

	if !ValidatorRegex.MatchString(value) {
		return false
	}

	var sum int
	for index, digit := range value {
		if number, err := strconv.Atoi(string(digit)); err == nil {
			sum += number * (10 - index)
		} else {
			sum += 10
		}
	}

	return sum%11 == 0
}
