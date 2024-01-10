package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var output []rune

	j := 0
	for _, char := range strings.ToLower(s) {
		if 'a' <= char && char <= 'z' {
			char = 'z' - char + 'a'
		}

		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			if j == 5 {
				output = append(output, ' ')
				j = 0
			}

			j++
			output = append(output, char)
		}
	}

	return string(output)
}
