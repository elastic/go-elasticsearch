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

// Package dynamicmapping
package dynamicmapping

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/3ea9ce260df22d3244bff5bace485dd97ff4046d/specification/_types/mapping/dynamic-template.ts#L50-L59
type DynamicMapping struct {
	Name string
}

var (
	Strict = DynamicMapping{"strict"}

	Runtime = DynamicMapping{"runtime"}

	True = DynamicMapping{"true"}

	False = DynamicMapping{"false"}
)

func (d *DynamicMapping) UnmarshalJSON(data []byte) error {
	return d.UnmarshalText(data)
}

func (d DynamicMapping) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DynamicMapping) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "strict":
		*d = Strict
	case "runtime":
		*d = Runtime
	case "true":
		*d = True
	case "false":
		*d = False
	default:
		*d = DynamicMapping{string(text)}
	}

	return nil
}

func (d DynamicMapping) String() string {
	return d.Name
}
