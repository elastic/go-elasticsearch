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


package types

// GeoPolygonPoints type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/geo.ts#L59-L61
type GeoPolygonPoints struct {
	Points []GeoLocation `json:"points"`
}

// GeoPolygonPointsBuilder holds GeoPolygonPoints struct and provides a builder API.
type GeoPolygonPointsBuilder struct {
	v *GeoPolygonPoints
}

// NewGeoPolygonPoints provides a builder for the GeoPolygonPoints struct.
func NewGeoPolygonPointsBuilder() *GeoPolygonPointsBuilder {
	r := GeoPolygonPointsBuilder{
		&GeoPolygonPoints{},
	}

	return &r
}

// Build finalize the chain and returns the GeoPolygonPoints struct
func (rb *GeoPolygonPointsBuilder) Build() GeoPolygonPoints {
	return *rb.v
}

func (rb *GeoPolygonPointsBuilder) Points(points ...GeoLocation) *GeoPolygonPointsBuilder {
	rb.v.Points = points
	return rb
}
