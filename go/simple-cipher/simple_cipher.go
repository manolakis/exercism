package cipher

import (
	"math"
	"strings"
	"unicode"
)

type shift int
type vigenere string

func NewCaesar() Cipher {
	return shift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || int(math.Abs(float64(distance))) >= 26 {
		return nil
	}

	return shift(distance)
}

func (c shift) Encode(input string) string {
	ciphered := make([]rune, 0)

	for _, char := range strings.ToLower(input) {
		if unicode.IsLetter(char) {
			ciphered = append(ciphered, toRune(int(26+toShift(char)+c)%26))
		}
	}

	return string(ciphered)
}

func (c shift) Decode(input string) string {
	plain := make([]rune, 0)

	for _, char := range strings.ToLower(input) {
		if unicode.IsLetter(char) {
			plain = append(plain, toRune(int(26+toShift(char)-c)%26))
		}
	}

	return string(plain)
}

func NewVigenere(key string) Cipher {
	sum := 0
	for _, key := range key {
		if !unicode.IsLetter(key) || unicode.IsUpper(key) {
			return nil
		} else {
			sum += int(toShift(key))
		}
	}

	if sum == 0 {
		return nil
	}

	return vigenere(key)
}

func (v vigenere) Encode(input string) string {
	key := []rune(v)
	ciphered := make([]rune, 0)
	index := 0

	for _, char := range strings.ToLower(input) {
		if unicode.IsLetter(char) {
			ciphered = append(ciphered, toRune(int(26+toShift(char)+toShift(key[index%len(key)]))%26))
			index++
		}
	}

	return string(ciphered)
}

func (v vigenere) Decode(input string) string {
	key := []rune(v)
	ciphered := make([]rune, 0)
	index := 0

	for _, char := range strings.ToLower(input) {
		if unicode.IsLetter(char) {
			ciphered = append(ciphered, toRune(int(26+toShift(char)-toShift(key[index%len(key)]))%26))
			index++
		}
	}

	return string(ciphered)
}

func toShift(letter rune) shift {
	return shift(letter - 'a')
}

func toRune(distance int) rune {
	return rune(distance + 'a')
}
