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

// FetchProfile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/profile.ts#L137-L144
type FetchProfile struct {
	Breakdown   FetchProfileBreakdown  `json:"breakdown"`
	Children    []FetchProfile         `json:"children,omitempty"`
	Debug       *FetchProfileDebug     `json:"debug,omitempty"`
	Description string                 `json:"description"`
	TimeInNanos DurationValueUnitNanos `json:"time_in_nanos"`
	Type        string                 `json:"type"`
}

// FetchProfileBuilder holds FetchProfile struct and provides a builder API.
type FetchProfileBuilder struct {
	v *FetchProfile
}

// NewFetchProfile provides a builder for the FetchProfile struct.
func NewFetchProfileBuilder() *FetchProfileBuilder {
	r := FetchProfileBuilder{
		&FetchProfile{},
	}

	return &r
}

// Build finalize the chain and returns the FetchProfile struct
func (rb *FetchProfileBuilder) Build() FetchProfile {
	return *rb.v
}

func (rb *FetchProfileBuilder) Breakdown(breakdown *FetchProfileBreakdownBuilder) *FetchProfileBuilder {
	v := breakdown.Build()
	rb.v.Breakdown = v
	return rb
}

func (rb *FetchProfileBuilder) Children(children []FetchProfileBuilder) *FetchProfileBuilder {
	tmp := make([]FetchProfile, len(children))
	for _, value := range children {
		tmp = append(tmp, value.Build())
	}
	rb.v.Children = tmp
	return rb
}

func (rb *FetchProfileBuilder) Debug(debug *FetchProfileDebugBuilder) *FetchProfileBuilder {
	v := debug.Build()
	rb.v.Debug = &v
	return rb
}

func (rb *FetchProfileBuilder) Description(description string) *FetchProfileBuilder {
	rb.v.Description = description
	return rb
}

func (rb *FetchProfileBuilder) TimeInNanos(timeinnanos *DurationValueUnitNanosBuilder) *FetchProfileBuilder {
	v := timeinnanos.Build()
	rb.v.TimeInNanos = v
	return rb
}

func (rb *FetchProfileBuilder) Type_(type_ string) *FetchProfileBuilder {
	rb.v.Type = type_
	return rb
}
