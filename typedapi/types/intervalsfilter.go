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

// IntervalsFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/fulltext.ts#L74-L86
type IntervalsFilter struct {
	After          *IntervalsContainer `json:"after,omitempty"`
	Before         *IntervalsContainer `json:"before,omitempty"`
	ContainedBy    *IntervalsContainer `json:"contained_by,omitempty"`
	Containing     *IntervalsContainer `json:"containing,omitempty"`
	NotContainedBy *IntervalsContainer `json:"not_contained_by,omitempty"`
	NotContaining  *IntervalsContainer `json:"not_containing,omitempty"`
	NotOverlapping *IntervalsContainer `json:"not_overlapping,omitempty"`
	Overlapping    *IntervalsContainer `json:"overlapping,omitempty"`
	Script         *Script             `json:"script,omitempty"`
}

// IntervalsFilterBuilder holds IntervalsFilter struct and provides a builder API.
type IntervalsFilterBuilder struct {
	v *IntervalsFilter
}

// NewIntervalsFilter provides a builder for the IntervalsFilter struct.
func NewIntervalsFilterBuilder() *IntervalsFilterBuilder {
	r := IntervalsFilterBuilder{
		&IntervalsFilter{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsFilter struct
func (rb *IntervalsFilterBuilder) Build() IntervalsFilter {
	return *rb.v
}

func (rb *IntervalsFilterBuilder) After(after *IntervalsContainerBuilder) *IntervalsFilterBuilder {
	v := after.Build()
	rb.v.After = &v
	return rb
}

func (rb *IntervalsFilterBuilder) Before(before *IntervalsContainerBuilder) *IntervalsFilterBuilder {
	v := before.Build()
	rb.v.Before = &v
	return rb
}

func (rb *IntervalsFilterBuilder) ContainedBy(containedby *IntervalsContainerBuilder) *IntervalsFilterBuilder {
	v := containedby.Build()
	rb.v.ContainedBy = &v
	return rb
}

func (rb *IntervalsFilterBuilder) Containing(containing *IntervalsContainerBuilder) *IntervalsFilterBuilder {
	v := containing.Build()
	rb.v.Containing = &v
	return rb
}

func (rb *IntervalsFilterBuilder) NotContainedBy(notcontainedby *IntervalsContainerBuilder) *IntervalsFilterBuilder {
	v := notcontainedby.Build()
	rb.v.NotContainedBy = &v
	return rb
}

func (rb *IntervalsFilterBuilder) NotContaining(notcontaining *IntervalsContainerBuilder) *IntervalsFilterBuilder {
	v := notcontaining.Build()
	rb.v.NotContaining = &v
	return rb
}

func (rb *IntervalsFilterBuilder) NotOverlapping(notoverlapping *IntervalsContainerBuilder) *IntervalsFilterBuilder {
	v := notoverlapping.Build()
	rb.v.NotOverlapping = &v
	return rb
}

func (rb *IntervalsFilterBuilder) Overlapping(overlapping *IntervalsContainerBuilder) *IntervalsFilterBuilder {
	v := overlapping.Build()
	rb.v.Overlapping = &v
	return rb
}

func (rb *IntervalsFilterBuilder) Script(script *ScriptBuilder) *IntervalsFilterBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
