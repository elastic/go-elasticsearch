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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// DecayPlacementDateMathTime type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/query_dsl/compound.ts#L77-L82
type DecayPlacementDateMathTime struct {
	Decay  *float64  `json:"decay,omitempty"`
	Offset *Time     `json:"offset,omitempty"`
	Origin *DateMath `json:"origin,omitempty"`
	Scale  *Time     `json:"scale,omitempty"`
}

// DecayPlacementDateMathTimeBuilder holds DecayPlacementDateMathTime struct and provides a builder API.
type DecayPlacementDateMathTimeBuilder struct {
	v *DecayPlacementDateMathTime
}

// NewDecayPlacementDateMathTime provides a builder for the DecayPlacementDateMathTime struct.
func NewDecayPlacementDateMathTimeBuilder() *DecayPlacementDateMathTimeBuilder {
	r := DecayPlacementDateMathTimeBuilder{
		&DecayPlacementDateMathTime{},
	}

	return &r
}

// Build finalize the chain and returns the DecayPlacementDateMathTime struct
func (rb *DecayPlacementDateMathTimeBuilder) Build() DecayPlacementDateMathTime {
	return *rb.v
}

func (rb *DecayPlacementDateMathTimeBuilder) Decay(decay float64) *DecayPlacementDateMathTimeBuilder {
	rb.v.Decay = &decay
	return rb
}

func (rb *DecayPlacementDateMathTimeBuilder) Offset(offset *TimeBuilder) *DecayPlacementDateMathTimeBuilder {
	v := offset.Build()
	rb.v.Offset = &v
	return rb
}

func (rb *DecayPlacementDateMathTimeBuilder) Origin(origin DateMath) *DecayPlacementDateMathTimeBuilder {
	rb.v.Origin = &origin
	return rb
}

func (rb *DecayPlacementDateMathTimeBuilder) Scale(scale *TimeBuilder) *DecayPlacementDateMathTimeBuilder {
	v := scale.Build()
	rb.v.Scale = &v
	return rb
}
