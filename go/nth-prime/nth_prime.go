package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("the nth number can't be computed")
	}

	primeNumbers := []int{2, 3}

	for i := 5; len(primeNumbers) < n; i++ {
		for j, primeNumber := range primeNumbers {
			if i%primeNumber == 0 {
				break
			}

			if j == len(primeNumbers)-1 {
				primeNumbers = append(primeNumbers, i)
			}
		}
	}

	return primeNumbers[n-1], nil
}
