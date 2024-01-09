package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) (output []string) {
	lowerSubject := strings.ToLower(subject)
	sortedSubject := sortString(lowerSubject)

	for _, candidate := range candidates {
		lowerCandidate := strings.ToLower(candidate)

		if lowerSubject != lowerCandidate && sortedSubject == sortString(lowerCandidate) {
			output = append(output, candidate)
		}
	}

	return output
}

func sortString(input string) string {
	slice := []rune(input)
	slices.Sort(slice)

	return string(slice)
}
