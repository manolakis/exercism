package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToUpper(word)

	for index, character := range word {
		if unicode.IsLetter(character) {
			if strings.IndexRune(word, character) < index {
				return false
			}
		}
	}

	return true
}
