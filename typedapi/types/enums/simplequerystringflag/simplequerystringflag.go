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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Package simplequerystringflag
package simplequerystringflag

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/query_dsl/fulltext.ts#L278-L292
type SimpleQueryStringFlag struct {
	Name string
}

var (
	NONE = SimpleQueryStringFlag{"NONE"}

	AND = SimpleQueryStringFlag{"AND"}

	OR = SimpleQueryStringFlag{"OR"}

	NOT = SimpleQueryStringFlag{"NOT"}

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
	switch strings.ToLower(string(text)) {

	case "NONE":
		*s = NONE
	case "AND":
		*s = AND
	case "OR":
		*s = OR
	case "NOT":
		*s = NOT
	case "PREFIX":
		*s = PREFIX
	case "PHRASE":
		*s = PHRASE
	case "PRECEDENCE":
		*s = PRECEDENCE
	case "ESCAPE":
		*s = ESCAPE
	case "WHITESPACE":
		*s = WHITESPACE
	case "FUZZY":
		*s = FUZZY
	case "NEAR":
		*s = NEAR
	case "SLOP":
		*s = SLOP
	case "ALL":
		*s = ALL
	default:
		*s = SimpleQueryStringFlag{string(text)}
	}

	return nil
}

func (s SimpleQueryStringFlag) String() string {
	return s.Name
}
