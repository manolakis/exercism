package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var re = regexp.MustCompile(`\w+('\w+)?`)

func WordCount(phrase string) Frequency {
	frequency := make(Frequency)

	for _, word := range re.FindAllString(strings.ToLower(phrase), -1) {
		frequency[word]++
	}

	return frequency
}
