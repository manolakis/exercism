package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)

	for value, stringValues := range input {
		for _, letter := range stringValues {
			output[strings.ToLower(letter)] = value
		}
	}

	return output
}
