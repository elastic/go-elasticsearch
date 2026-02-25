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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package geoshaperelation
package geoshaperelation

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/_types/Geo.ts#L64-L82
type GeoShapeRelation struct {
	Name string
}

var (

	// Intersects Return all documents whose `geo_shape` or `geo_point` field intersects the
	// query geometry.
	Intersects = GeoShapeRelation{"intersects"}

	// Disjoint Return all documents whose `geo_shape` or `geo_point` field has nothing in
	// common with the query geometry.
	Disjoint = GeoShapeRelation{"disjoint"}

	// Within Return all documents whose `geo_shape` or `geo_point` field is within the
	// query geometry. Line geometries are not supported.
	Within = GeoShapeRelation{"within"}

	// Contains Return all documents whose `geo_shape` or `geo_point` field contains the
	// query geometry.
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
