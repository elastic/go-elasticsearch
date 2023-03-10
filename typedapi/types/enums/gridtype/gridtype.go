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

// Package gridtype
package gridtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_global/search_mvt/_types/GridType.ts#L20-L25
type GridType struct {
	Name string
}

var (
	Grid = GridType{"grid"}

	Point = GridType{"point"}

	Centroid = GridType{"centroid"}
)

func (g GridType) MarshalText() (text []byte, err error) {
	return []byte(g.String()), nil
}

func (g *GridType) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "grid":
		*g = Grid
	case "point":
		*g = Point
	case "centroid":
		*g = Centroid
	default:
		*g = GridType{string(text)}
	}

	return nil
}

func (g GridType) String() string {
	return g.Name
}
