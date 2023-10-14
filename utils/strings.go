package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func Trim(str string) string {
	trimmed := strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(str, " "))
	trimmed = strings.Replace(trimmed, " <", "<", -1)
	trimmed = strings.Replace(trimmed, "> ", ">", -1)
	return trimmed
}

func TransformToHex(text string) string {
	var output = fmt.Sprintf("#%s", strings.Trim(text, "\""))
	return output
}

func TrimQuotes(text string) string {
	return strings.Trim(text, "\"")
}
