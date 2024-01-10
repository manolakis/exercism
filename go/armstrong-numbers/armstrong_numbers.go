package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	sum := 0
	numbers := strconv.Itoa(n)
	length := float64(len(numbers))

	for _, digit := range numbers {
		sum += int(math.Pow(float64(digit-'0'), length))
	}

	return sum == n
}
