package reverse

func Reverse(input string) string {
	var reverse string

	for _, character := range input {
		reverse = string(character) + reverse
	}

	return reverse
}
