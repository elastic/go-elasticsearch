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

// GeoLine type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Geo.ts#L59-L65
type GeoLine struct {
	// Coordinates Array of `[lon, lat]` coordinates
	Coordinates [][]float64 `json:"coordinates"`
	// Type Always `"LineString"`
	Type string `json:"type"`
}

// GeoLineBuilder holds GeoLine struct and provides a builder API.
type GeoLineBuilder struct {
	v *GeoLine
}

// NewGeoLine provides a builder for the GeoLine struct.
func NewGeoLineBuilder() *GeoLineBuilder {
	r := GeoLineBuilder{
		&GeoLine{},
	}

	return &r
}

// Build finalize the chain and returns the GeoLine struct
func (rb *GeoLineBuilder) Build() GeoLine {
	return *rb.v
}

// Coordinates Array of `[lon, lat]` coordinates

func (rb *GeoLineBuilder) Coordinates(coordinates ...[]float64) *GeoLineBuilder {
	rb.v.Coordinates = coordinates
	return rb
}

// Type Always `"LineString"`

func (rb *GeoLineBuilder) Type_(type_ string) *GeoLineBuilder {
	rb.v.Type = type_
	return rb
}
