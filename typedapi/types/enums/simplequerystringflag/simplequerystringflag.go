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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package simplequerystringflag
package simplequerystringflag

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/_types/query_dsl/fulltext.ts#L781-L836
type SimpleQueryStringFlag struct {
	Name string
}

var (

	// NONE Disables all operators.
	NONE = SimpleQueryStringFlag{"NONE"}

	// AND Enables the `+` AND operator.
	AND = SimpleQueryStringFlag{"AND"}

	// NOT Enables the `-` NOT operator.
	NOT = SimpleQueryStringFlag{"NOT"}

	// OR Enables the `\|` OR operator.
	OR = SimpleQueryStringFlag{"OR"}

	// PREFIX Enables the `*` prefix operator.
	PREFIX = SimpleQueryStringFlag{"PREFIX"}

	// PHRASE Enables the `"` quotes operator used to search for phrases.
	PHRASE = SimpleQueryStringFlag{"PHRASE"}

	// PRECEDENCE Enables the `(` and `)` operators to control operator precedence.
	PRECEDENCE = SimpleQueryStringFlag{"PRECEDENCE"}

	// ESCAPE Enables `\` as an escape character.
	ESCAPE = SimpleQueryStringFlag{"ESCAPE"}

	// WHITESPACE Enables whitespace as split characters.
	WHITESPACE = SimpleQueryStringFlag{"WHITESPACE"}

	// FUZZY Enables the `~N` operator after a word, where `N` is an integer denoting the
	// allowed edit distance for matching.
	FUZZY = SimpleQueryStringFlag{"FUZZY"}

	// NEAR Enables the `~N` operator, after a phrase where `N` is the maximum number of
	// positions allowed between matching tokens. Synonymous to `SLOP`.
	NEAR = SimpleQueryStringFlag{"NEAR"}

	// SLOP Enables the `~N` operator, after a phrase where `N` is maximum number of
	// positions allowed between matching tokens. Synonymous to `NEAR`.
	SLOP = SimpleQueryStringFlag{"SLOP"}

	// ALL Enables all optional operators.
	ALL = SimpleQueryStringFlag{"ALL"}
)

func (s SimpleQueryStringFlag) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *SimpleQueryStringFlag) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "none":
		*s = NONE
	case "and":
		*s = AND
	case "not":
		*s = NOT
	case "or":
		*s = OR
	case "prefix":
		*s = PREFIX
	case "phrase":
		*s = PHRASE
	case "precedence":
		*s = PRECEDENCE
	case "escape":
		*s = ESCAPE
	case "whitespace":
		*s = WHITESPACE
	case "fuzzy":
		*s = FUZZY
	case "near":
		*s = NEAR
	case "slop":
		*s = SLOP
	case "all":
		*s = ALL
	default:
		*s = SimpleQueryStringFlag{string(text)}
	}

	return nil
}

func (s SimpleQueryStringFlag) String() string {
	return s.Name
}
