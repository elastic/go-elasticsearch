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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// DecayPlacementdoubledouble type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/query_dsl/compound.ts#L77-L82
type DecayPlacementdoubledouble struct {
	Decay  *float64 `json:"decay,omitempty"`
	Offset *float64 `json:"offset,omitempty"`
	Origin *float64 `json:"origin,omitempty"`
	Scale  *float64 `json:"scale,omitempty"`
}

// DecayPlacementdoubledoubleBuilder holds DecayPlacementdoubledouble struct and provides a builder API.
type DecayPlacementdoubledoubleBuilder struct {
	v *DecayPlacementdoubledouble
}

// NewDecayPlacementdoubledouble provides a builder for the DecayPlacementdoubledouble struct.
func NewDecayPlacementdoubledoubleBuilder() *DecayPlacementdoubledoubleBuilder {
	r := DecayPlacementdoubledoubleBuilder{
		&DecayPlacementdoubledouble{},
	}

	return &r
}

// Build finalize the chain and returns the DecayPlacementdoubledouble struct
func (rb *DecayPlacementdoubledoubleBuilder) Build() DecayPlacementdoubledouble {
	return *rb.v
}

func (rb *DecayPlacementdoubledoubleBuilder) Decay(decay float64) *DecayPlacementdoubledoubleBuilder {
	rb.v.Decay = &decay
	return rb
}

func (rb *DecayPlacementdoubledoubleBuilder) Offset(offset float64) *DecayPlacementdoubledoubleBuilder {
	rb.v.Offset = &offset
	return rb
}

func (rb *DecayPlacementdoubledoubleBuilder) Origin(origin float64) *DecayPlacementdoubledoubleBuilder {
	rb.v.Origin = &origin
	return rb
}

func (rb *DecayPlacementdoubledoubleBuilder) Scale(scale float64) *DecayPlacementdoubledoubleBuilder {
	rb.v.Scale = &scale
	return rb
}
