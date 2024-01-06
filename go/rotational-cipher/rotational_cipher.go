package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	cyphered := make([]rune, len(plain))

	for index, char := range plain {
		if unicode.IsLetter(char) {
			root := 'a'
			if unicode.IsUpper(char) {
				root = 'A'
			}

			cyphered[index] = (char-root+rune(shiftKey))%26 + root
		} else {
			cyphered[index] = char
		}
	}

	return string(cyphered)
}
