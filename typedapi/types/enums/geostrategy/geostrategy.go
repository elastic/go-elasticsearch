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
// https://github.com/elastic/elasticsearch-specification/tree/33e8a1c9cad22a5946ac735c4fba31af2da2cec2

// Package geostrategy
package geostrategy

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/33e8a1c9cad22a5946ac735c4fba31af2da2cec2/specification/_types/mapping/geo.ts#L52-L55
type GeoStrategy struct {
	Name string
}

var (
	Recursive = GeoStrategy{"recursive"}

	Term = GeoStrategy{"term"}
)

func (g GeoStrategy) MarshalText() (text []byte, err error) {
	return []byte(g.String()), nil
}

func (g *GeoStrategy) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "recursive":
		*g = Recursive
	case "term":
		*g = Term
	default:
		*g = GeoStrategy{string(text)}
	}

	return nil
}

func (g GeoStrategy) String() string {
	return g.Name
}
