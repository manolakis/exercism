package accumulate

type unaryFunc func(data string) string

func Accumulate(slice []string, f unaryFunc) []string {
	var newSlice = make([]string, len(slice))

	for idx, item := range slice {
		newSlice[idx] = f(item)
	}

	return newSlice
}
