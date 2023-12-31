package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var validLineRegex = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)
var strangeSeparatorsRegex = regexp.MustCompile(`<(~|\*|=|-)*>`)
var passwordRegex = regexp.MustCompile(`".*(?i:password).*"`)
var endOfLineRegex = regexp.MustCompile(`end-of-line\d*`)
var userRegex = regexp.MustCompile(`User\s+(\w+)`)

func IsValidLine(text string) bool {
	return validLineRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	return strangeSeparatorsRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	counter := 0

	for _, line := range lines {
		counter += len(passwordRegex.FindAllString(line, -1))
	}

	return counter
}

func RemoveEndOfLineText(text string) string {
	return endOfLineRegex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	for index, line := range lines {
		user := userRegex.FindStringSubmatch(line)

		if user != nil {
			lines[index] = fmt.Sprintf("[USR] %s %s", user[1], line)
		}
	}

	return lines
}
