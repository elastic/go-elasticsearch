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
// https://github.com/elastic/elasticsearch-specification/tree/f1932ce6b46a53a8342db522b1a7883bcc9e0996

// Package simplequerystringflag
package simplequerystringflag

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/f1932ce6b46a53a8342db522b1a7883bcc9e0996/specification/_types/query_dsl/fulltext.ts#L729-L784
type SimpleQueryStringFlag struct {
	Name string
}

var (
	NONE = SimpleQueryStringFlag{"NONE"}

	AND = SimpleQueryStringFlag{"AND"}

	NOT = SimpleQueryStringFlag{"NOT"}

	OR = SimpleQueryStringFlag{"OR"}

	PREFIX = SimpleQueryStringFlag{"PREFIX"}

	PHRASE = SimpleQueryStringFlag{"PHRASE"}

	PRECEDENCE = SimpleQueryStringFlag{"PRECEDENCE"}

	ESCAPE = SimpleQueryStringFlag{"ESCAPE"}

	WHITESPACE = SimpleQueryStringFlag{"WHITESPACE"}

	FUZZY = SimpleQueryStringFlag{"FUZZY"}

	NEAR = SimpleQueryStringFlag{"NEAR"}

	SLOP = SimpleQueryStringFlag{"SLOP"}

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
