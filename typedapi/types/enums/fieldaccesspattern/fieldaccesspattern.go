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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

// Package fieldaccesspattern
package fieldaccesspattern

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6/specification/ingest/_types/Pipeline.ts#L115-L118
type FieldAccessPattern struct {
	Name string
}

var (
	Classic = FieldAccessPattern{"classic"}

	Flexible = FieldAccessPattern{"flexible"}
)

func (f FieldAccessPattern) MarshalText() (text []byte, err error) {
	return []byte(f.String()), nil
}

func (f *FieldAccessPattern) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "classic":
		*f = Classic
	case "flexible":
		*f = Flexible
	default:
		*f = FieldAccessPattern{string(text)}
	}

	return nil
}

func (f FieldAccessPattern) String() string {
	return f.Name
}
