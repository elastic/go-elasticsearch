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

// Phase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ilm/_types/Phase.ts#L25-L33
type Phase struct {
	Actions        *Actions        `json:"actions,omitempty"`
	Configurations *Configurations `json:"configurations,omitempty"`
	MinAge         int64           `json:"min_age,omitempty"`
}

// PhaseBuilder holds Phase struct and provides a builder API.
type PhaseBuilder struct {
	v *Phase
}

// NewPhase provides a builder for the Phase struct.
func NewPhaseBuilder() *PhaseBuilder {
	r := PhaseBuilder{
		&Phase{},
	}

	return &r
}

// Build finalize the chain and returns the Phase struct
func (rb *PhaseBuilder) Build() Phase {
	return *rb.v
}

func (rb *PhaseBuilder) Actions(actions *ActionsBuilder) *PhaseBuilder {
	v := actions.Build()
	rb.v.Actions = &v
	return rb
}

func (rb *PhaseBuilder) Configurations(configurations *ConfigurationsBuilder) *PhaseBuilder {
	v := configurations.Build()
	rb.v.Configurations = &v
	return rb
}

func (rb *PhaseBuilder) MinAge(arg int64) *PhaseBuilder {
	rb.v.MinAge = arg
	return rb
}
