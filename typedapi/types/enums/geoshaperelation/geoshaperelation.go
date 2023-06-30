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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Package geoshaperelation
package geoshaperelation

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/Geo.ts#L67-L72
type GeoShapeRelation struct {
	Name string
}

var (
	Intersects = GeoShapeRelation{"intersects"}

	Disjoint = GeoShapeRelation{"disjoint"}

	Within = GeoShapeRelation{"within"}

	Contains = GeoShapeRelation{"contains"}
)

func (g GeoShapeRelation) MarshalText() (text []byte, err error) {
	return []byte(g.String()), nil
}

func (g *GeoShapeRelation) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "intersects":
		*g = Intersects
	case "disjoint":
		*g = Disjoint
	case "within":
		*g = Within
	case "contains":
		*g = Contains
	default:
		*g = GeoShapeRelation{string(text)}
	}

	return nil
}

func (g GeoShapeRelation) String() string {
	return g.Name
}
