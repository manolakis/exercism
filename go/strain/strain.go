package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](input []T, predicate func(T) bool) []T {
	var output []T

	for _, value := range input {
		if predicate(value) {
			output = append(output, value)
		}
	}

	return output
}

func Discard[T any](input []T, predicate func(T) bool) []T {
	return Keep(input, func(value T) bool {
		return !predicate(value)
	})
}
