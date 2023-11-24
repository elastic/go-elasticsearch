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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Package termvectoroption
package termvectoroption

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/mapping/TermVectorOption.ts#L20-L28
type TermVectorOption struct {
	Name string
}

var (
	No = TermVectorOption{"no"}

	Yes = TermVectorOption{"yes"}

	Withoffsets = TermVectorOption{"with_offsets"}

	Withpositions = TermVectorOption{"with_positions"}

	Withpositionsoffsets = TermVectorOption{"with_positions_offsets"}

	Withpositionsoffsetspayloads = TermVectorOption{"with_positions_offsets_payloads"}

	Withpositionspayloads = TermVectorOption{"with_positions_payloads"}
)

func (t TermVectorOption) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TermVectorOption) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "no":
		*t = No
	case "yes":
		*t = Yes
	case "with_offsets":
		*t = Withoffsets
	case "with_positions":
		*t = Withpositions
	case "with_positions_offsets":
		*t = Withpositionsoffsets
	case "with_positions_offsets_payloads":
		*t = Withpositionsoffsetspayloads
	case "with_positions_payloads":
		*t = Withpositionspayloads
	default:
		*t = TermVectorOption{string(text)}
	}

	return nil
}

func (t TermVectorOption) String() string {
	return t.Name
}
