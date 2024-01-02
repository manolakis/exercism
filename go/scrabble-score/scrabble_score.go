package scrabble

import "strings"

func getLetterValue(letter rune) int {
	var value int

	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		value = 1
	case 'D', 'G':
		value = 2
	case 'B', 'C', 'M', 'P':
		value = 3
	case 'F', 'H', 'V', 'W', 'Y':
		value = 4
	case 'K':
		value = 5
	case 'J', 'X':
		value = 8
	case 'Q', 'Z':
		value = 10
	default:
		value = 0
	}

	return value
}

func Score(word string) int {
	counter := 0

	for _, letter := range strings.ToUpper(word) {
		counter += getLetterValue(letter)
	}

	return counter
}
