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

// DecayPlacementDateMathDuration type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/compound.ts#L77-L82
type DecayPlacementDateMathDuration struct {
	Decay  *float64  `json:"decay,omitempty"`
	Offset *Duration `json:"offset,omitempty"`
	Origin *DateMath `json:"origin,omitempty"`
	Scale  *Duration `json:"scale,omitempty"`
}

// DecayPlacementDateMathDurationBuilder holds DecayPlacementDateMathDuration struct and provides a builder API.
type DecayPlacementDateMathDurationBuilder struct {
	v *DecayPlacementDateMathDuration
}

// NewDecayPlacementDateMathDuration provides a builder for the DecayPlacementDateMathDuration struct.
func NewDecayPlacementDateMathDurationBuilder() *DecayPlacementDateMathDurationBuilder {
	r := DecayPlacementDateMathDurationBuilder{
		&DecayPlacementDateMathDuration{},
	}

	return &r
}

// Build finalize the chain and returns the DecayPlacementDateMathDuration struct
func (rb *DecayPlacementDateMathDurationBuilder) Build() DecayPlacementDateMathDuration {
	return *rb.v
}

func (rb *DecayPlacementDateMathDurationBuilder) Decay(decay float64) *DecayPlacementDateMathDurationBuilder {
	rb.v.Decay = &decay
	return rb
}

func (rb *DecayPlacementDateMathDurationBuilder) Offset(offset *DurationBuilder) *DecayPlacementDateMathDurationBuilder {
	v := offset.Build()
	rb.v.Offset = &v
	return rb
}

func (rb *DecayPlacementDateMathDurationBuilder) Origin(origin DateMath) *DecayPlacementDateMathDurationBuilder {
	rb.v.Origin = &origin
	return rb
}

func (rb *DecayPlacementDateMathDurationBuilder) Scale(scale *DurationBuilder) *DecayPlacementDateMathDurationBuilder {
	v := scale.Build()
	rb.v.Scale = &v
	return rb
}
