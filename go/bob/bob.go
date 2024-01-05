package bob

import (
	"strings"
	"unicode"
)

// Hey reply to someone when they say something
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	isQuestion, isYelling := checkIsQuestion(remark), checkIsYelling(remark)

	switch {
	case isQuestion && isYelling:
		return "Calm down, I know what I'm doing!"
	case isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}

func checkIsQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func checkIsYelling(remark string) bool {
	var hasLetters bool

	for _, character := range remark {
		if unicode.IsLetter(character) {
			hasLetters = true

			if unicode.IsLower(character) {
				return false
			}
		}
	}

	return hasLetters
}
