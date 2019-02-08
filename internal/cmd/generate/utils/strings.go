package utils

import (
	"regexp"
)

var (
	reIDString = regexp.MustCompile(`\bid\b`)
)

func IDToUpper(s string) string {
	return reIDString.ReplaceAllLiteralString(s, "ID")
}
