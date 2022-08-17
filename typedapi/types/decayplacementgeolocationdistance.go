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

// DecayPlacementGeoLocationDistance type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/compound.ts#L77-L82
type DecayPlacementGeoLocationDistance struct {
	Decay  *float64     `json:"decay,omitempty"`
	Offset *Distance    `json:"offset,omitempty"`
	Origin *GeoLocation `json:"origin,omitempty"`
	Scale  *Distance    `json:"scale,omitempty"`
}

// DecayPlacementGeoLocationDistanceBuilder holds DecayPlacementGeoLocationDistance struct and provides a builder API.
type DecayPlacementGeoLocationDistanceBuilder struct {
	v *DecayPlacementGeoLocationDistance
}

// NewDecayPlacementGeoLocationDistance provides a builder for the DecayPlacementGeoLocationDistance struct.
func NewDecayPlacementGeoLocationDistanceBuilder() *DecayPlacementGeoLocationDistanceBuilder {
	r := DecayPlacementGeoLocationDistanceBuilder{
		&DecayPlacementGeoLocationDistance{},
	}

	return &r
}

// Build finalize the chain and returns the DecayPlacementGeoLocationDistance struct
func (rb *DecayPlacementGeoLocationDistanceBuilder) Build() DecayPlacementGeoLocationDistance {
	return *rb.v
}

func (rb *DecayPlacementGeoLocationDistanceBuilder) Decay(decay float64) *DecayPlacementGeoLocationDistanceBuilder {
	rb.v.Decay = &decay
	return rb
}

func (rb *DecayPlacementGeoLocationDistanceBuilder) Offset(offset Distance) *DecayPlacementGeoLocationDistanceBuilder {
	rb.v.Offset = &offset
	return rb
}

func (rb *DecayPlacementGeoLocationDistanceBuilder) Origin(origin *GeoLocationBuilder) *DecayPlacementGeoLocationDistanceBuilder {
	v := origin.Build()
	rb.v.Origin = &v
	return rb
}

func (rb *DecayPlacementGeoLocationDistanceBuilder) Scale(scale Distance) *DecayPlacementGeoLocationDistanceBuilder {
	rb.v.Scale = &scale
	return rb
}
