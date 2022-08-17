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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


// Package geodistancetype
package geodistancetype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Geo.ts#L51-L54
type GeoDistanceType struct {
	name string
}

var (
	Arc = GeoDistanceType{"arc"}

	Plane = GeoDistanceType{"plane"}
)

func (g GeoDistanceType) MarshalText() (text []byte, err error) {
	return []byte(g.String()), nil
}

func (g *GeoDistanceType) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "arc":
		*g = Arc
	case "plane":
		*g = Plane
	default:
		*g = GeoDistanceType{string(text)}
	}

	return nil
}

func (g GeoDistanceType) String() string {
	return g.name
}
