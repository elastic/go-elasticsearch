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

// PublishedClusterStates type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L114-L118
type PublishedClusterStates struct {
	CompatibleDiffs   *int64 `json:"compatible_diffs,omitempty"`
	FullStates        *int64 `json:"full_states,omitempty"`
	IncompatibleDiffs *int64 `json:"incompatible_diffs,omitempty"`
}

// PublishedClusterStatesBuilder holds PublishedClusterStates struct and provides a builder API.
type PublishedClusterStatesBuilder struct {
	v *PublishedClusterStates
}

// NewPublishedClusterStates provides a builder for the PublishedClusterStates struct.
func NewPublishedClusterStatesBuilder() *PublishedClusterStatesBuilder {
	r := PublishedClusterStatesBuilder{
		&PublishedClusterStates{},
	}

	return &r
}

// Build finalize the chain and returns the PublishedClusterStates struct
func (rb *PublishedClusterStatesBuilder) Build() PublishedClusterStates {
	return *rb.v
}

func (rb *PublishedClusterStatesBuilder) CompatibleDiffs(compatiblediffs int64) *PublishedClusterStatesBuilder {
	rb.v.CompatibleDiffs = &compatiblediffs
	return rb
}

func (rb *PublishedClusterStatesBuilder) FullStates(fullstates int64) *PublishedClusterStatesBuilder {
	rb.v.FullStates = &fullstates
	return rb
}

func (rb *PublishedClusterStatesBuilder) IncompatibleDiffs(incompatiblediffs int64) *PublishedClusterStatesBuilder {
	rb.v.IncompatibleDiffs = &incompatiblediffs
	return rb
}
