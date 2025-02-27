package utilities

import (
	"regexp"
)

func AlphaNumericCount(str string) int64 {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	stripped := reg.ReplaceAllString(str, "")
	return int64(len(stripped))
}
