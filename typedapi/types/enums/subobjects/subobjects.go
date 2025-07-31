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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package subobjects
package subobjects

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/mapping/TypeMapping.ts#L63-L74
type Subobjects struct {
	Name string
}

var (
	True = Subobjects{"true"}

	False = Subobjects{"false"}

	Auto = Subobjects{"auto"}
)

func (s *Subobjects) UnmarshalJSON(data []byte) error {
	return s.UnmarshalText(data)
}

func (s Subobjects) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *Subobjects) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "true":
		*s = True
	case "false":
		*s = False
	case "auto":
		*s = Auto
	default:
		*s = Subobjects{string(text)}
	}

	return nil
}

func (s Subobjects) String() string {
	return s.Name
}
