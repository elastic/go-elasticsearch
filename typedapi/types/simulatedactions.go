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

// SimulatedActions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Action.ts#L93-L97
type SimulatedActions struct {
	Actions []string          `json:"actions"`
	All     *SimulatedActions `json:"all,omitempty"`
	UseAll  bool              `json:"use_all"`
}

// SimulatedActionsBuilder holds SimulatedActions struct and provides a builder API.
type SimulatedActionsBuilder struct {
	v *SimulatedActions
}

// NewSimulatedActions provides a builder for the SimulatedActions struct.
func NewSimulatedActionsBuilder() *SimulatedActionsBuilder {
	r := SimulatedActionsBuilder{
		&SimulatedActions{},
	}

	return &r
}

// Build finalize the chain and returns the SimulatedActions struct
func (rb *SimulatedActionsBuilder) Build() SimulatedActions {
	return *rb.v
}

func (rb *SimulatedActionsBuilder) Actions(actions ...string) *SimulatedActionsBuilder {
	rb.v.Actions = actions
	return rb
}

func (rb *SimulatedActionsBuilder) All(all *SimulatedActionsBuilder) *SimulatedActionsBuilder {
	v := all.Build()
	rb.v.All = &v
	return rb
}

func (rb *SimulatedActionsBuilder) UseAll(useall bool) *SimulatedActionsBuilder {
	rb.v.UseAll = useall
	return rb
}
