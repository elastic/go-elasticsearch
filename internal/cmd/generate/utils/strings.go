// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

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
