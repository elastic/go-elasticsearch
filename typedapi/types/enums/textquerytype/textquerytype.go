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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Package textquerytype
package textquerytype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/query_dsl/fulltext.ts#L219-L226
type TextQueryType struct {
	Name string
}

var (
	Bestfields = TextQueryType{"best_fields"}

	Mostfields = TextQueryType{"most_fields"}

	Crossfields = TextQueryType{"cross_fields"}

	Phrase = TextQueryType{"phrase"}

	Phraseprefix = TextQueryType{"phrase_prefix"}

	Boolprefix = TextQueryType{"bool_prefix"}
)

func (t TextQueryType) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TextQueryType) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "best_fields":
		*t = Bestfields
	case "most_fields":
		*t = Mostfields
	case "cross_fields":
		*t = Crossfields
	case "phrase":
		*t = Phrase
	case "phrase_prefix":
		*t = Phraseprefix
	case "bool_prefix":
		*t = Boolprefix
	default:
		*t = TextQueryType{string(text)}
	}

	return nil
}

func (t TextQueryType) String() string {
	return t.Name
}
