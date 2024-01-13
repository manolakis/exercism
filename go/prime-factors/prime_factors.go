package prime

func Factors(number int64) []int64 {
	var factors []int64

	for i := 2; number != 1; {
		n := int64(i)

		if number%n == 0 {
			number /= n

			factors = append(factors, n)
		} else {
			i++
		}
	}

	return factors
}
