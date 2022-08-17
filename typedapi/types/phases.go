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

// Phases type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ilm/_types/Phase.ts#L35-L40
type Phases struct {
	Cold   *Phase `json:"cold,omitempty"`
	Delete *Phase `json:"delete,omitempty"`
	Hot    *Phase `json:"hot,omitempty"`
	Warm   *Phase `json:"warm,omitempty"`
}

// PhasesBuilder holds Phases struct and provides a builder API.
type PhasesBuilder struct {
	v *Phases
}

// NewPhases provides a builder for the Phases struct.
func NewPhasesBuilder() *PhasesBuilder {
	r := PhasesBuilder{
		&Phases{},
	}

	return &r
}

// Build finalize the chain and returns the Phases struct
func (rb *PhasesBuilder) Build() Phases {
	return *rb.v
}

func (rb *PhasesBuilder) Cold(cold *PhaseBuilder) *PhasesBuilder {
	v := cold.Build()
	rb.v.Cold = &v
	return rb
}

func (rb *PhasesBuilder) Delete(delete *PhaseBuilder) *PhasesBuilder {
	v := delete.Build()
	rb.v.Delete = &v
	return rb
}

func (rb *PhasesBuilder) Hot(hot *PhaseBuilder) *PhasesBuilder {
	v := hot.Build()
	rb.v.Hot = &v
	return rb
}

func (rb *PhasesBuilder) Warm(warm *PhaseBuilder) *PhasesBuilder {
	v := warm.Build()
	rb.v.Warm = &v
	return rb
}
