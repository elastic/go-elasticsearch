package utils

import (
	"regexp"
)

var (
	reIDString = regexp.MustCompile(`\bid\b`)
)

// IDToUpper returns a string with all occurrences of "id" capitalized.
//
func IDToUpper(s string) string {
	return reIDString.ReplaceAllLiteralString(s, "ID")
}
