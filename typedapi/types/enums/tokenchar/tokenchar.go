// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

// Package tokenchar
package tokenchar

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/analysis/tokenizers.ts#L46-L53
type TokenChar struct {
	Name string
}

var (
	Letter = TokenChar{"letter"}

	Digit = TokenChar{"digit"}

	Whitespace = TokenChar{"whitespace"}

	Punctuation = TokenChar{"punctuation"}

	Symbol = TokenChar{"symbol"}

	Custom = TokenChar{"custom"}
)

func (t TokenChar) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TokenChar) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "letter":
		*t = Letter
	case "digit":
		*t = Digit
	case "whitespace":
		*t = Whitespace
	case "punctuation":
		*t = Punctuation
	case "symbol":
		*t = Symbol
	case "custom":
		*t = Custom
	default:
		*t = TokenChar{string(text)}
	}

	return nil
}

func (t TokenChar) String() string {
	return t.Name
}
