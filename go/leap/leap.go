package leap

// IsLeapYear should return true when is a leap year
func IsLeapYear(year int) bool {
	return isDivisibleBy(year, 400) || (isDivisibleBy(year, 4) && !isDivisibleBy(year, 100))
}

func isDivisibleBy(year int, number int) bool {
	return year%number == 0
}
