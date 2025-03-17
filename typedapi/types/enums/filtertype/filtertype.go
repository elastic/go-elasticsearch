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
// https://github.com/elastic/elasticsearch-specification/tree/3ea9ce260df22d3244bff5bace485dd97ff4046d

// Package filtertype
package filtertype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/3ea9ce260df22d3244bff5bace485dd97ff4046d/specification/ml/_types/Filter.ts#L43-L46
type FilterType struct {
	Name string
}

var (
	Include = FilterType{"include"}

	Exclude = FilterType{"exclude"}
)

func (f FilterType) MarshalText() (text []byte, err error) {
	return []byte(f.String()), nil
}

func (f *FilterType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "include":
		*f = Include
	case "exclude":
		*f = Exclude
	default:
		*f = FilterType{string(text)}
	}

	return nil
}

func (f FilterType) String() string {
	return f.Name
}
