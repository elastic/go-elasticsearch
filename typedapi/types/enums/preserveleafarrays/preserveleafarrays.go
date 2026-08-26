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
// https://github.com/elastic/elasticsearch-specification/tree/964a36594f01c23463551aa7d09d17e514f361d1

// Package preserveleafarrays
package preserveleafarrays

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/964a36594f01c23463551aa7d09d17e514f361d1/specification/_types/mapping/complex.ts#L46-L49
type PreserveLeafArrays struct {
	Name string
}

var (
	Lossy = PreserveLeafArrays{"lossy"}

	Exact = PreserveLeafArrays{"exact"}
)

func (p PreserveLeafArrays) MarshalText() (text []byte, err error) {
	return []byte(p.String()), nil
}

func (p *PreserveLeafArrays) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "lossy":
		*p = Lossy
	case "exact":
		*p = Exact
	default:
		*p = PreserveLeafArrays{string(text)}
	}

	return nil
}

func (p PreserveLeafArrays) String() string {
	return p.Name
}
