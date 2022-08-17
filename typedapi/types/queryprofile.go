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

// QueryProfile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/profile.ts#L116-L122
type QueryProfile struct {
	Breakdown   QueryBreakdown         `json:"breakdown"`
	Children    []QueryProfile         `json:"children,omitempty"`
	Description string                 `json:"description"`
	TimeInNanos DurationValueUnitNanos `json:"time_in_nanos"`
	Type        string                 `json:"type"`
}

// QueryProfileBuilder holds QueryProfile struct and provides a builder API.
type QueryProfileBuilder struct {
	v *QueryProfile
}

// NewQueryProfile provides a builder for the QueryProfile struct.
func NewQueryProfileBuilder() *QueryProfileBuilder {
	r := QueryProfileBuilder{
		&QueryProfile{},
	}

	return &r
}

// Build finalize the chain and returns the QueryProfile struct
func (rb *QueryProfileBuilder) Build() QueryProfile {
	return *rb.v
}

func (rb *QueryProfileBuilder) Breakdown(breakdown *QueryBreakdownBuilder) *QueryProfileBuilder {
	v := breakdown.Build()
	rb.v.Breakdown = v
	return rb
}

func (rb *QueryProfileBuilder) Children(children []QueryProfileBuilder) *QueryProfileBuilder {
	tmp := make([]QueryProfile, len(children))
	for _, value := range children {
		tmp = append(tmp, value.Build())
	}
	rb.v.Children = tmp
	return rb
}

func (rb *QueryProfileBuilder) Description(description string) *QueryProfileBuilder {
	rb.v.Description = description
	return rb
}

func (rb *QueryProfileBuilder) TimeInNanos(timeinnanos *DurationValueUnitNanosBuilder) *QueryProfileBuilder {
	v := timeinnanos.Build()
	rb.v.TimeInNanos = v
	return rb
}

func (rb *QueryProfileBuilder) Type_(type_ string) *QueryProfileBuilder {
	rb.v.Type = type_
	return rb
}
