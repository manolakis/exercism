package sieve

func Sieve(limit int) (output []int) {
	list := make([]bool, limit+1)

	for i := 2; i <= limit; i++ {
		if !list[i] {
			output = append(output, i)

			for j := i * i; j <= limit; j += i {
				list[j] = true
			}
		}
	}

	return output
}
