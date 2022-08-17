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

// IntervalsContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/fulltext.ts#L63-L72
type IntervalsContainer struct {
	AllOf    *IntervalsAllOf    `json:"all_of,omitempty"`
	AnyOf    *IntervalsAnyOf    `json:"any_of,omitempty"`
	Fuzzy    *IntervalsFuzzy    `json:"fuzzy,omitempty"`
	Match    *IntervalsMatch    `json:"match,omitempty"`
	Prefix   *IntervalsPrefix   `json:"prefix,omitempty"`
	Wildcard *IntervalsWildcard `json:"wildcard,omitempty"`
}

// IntervalsContainerBuilder holds IntervalsContainer struct and provides a builder API.
type IntervalsContainerBuilder struct {
	v *IntervalsContainer
}

// NewIntervalsContainer provides a builder for the IntervalsContainer struct.
func NewIntervalsContainerBuilder() *IntervalsContainerBuilder {
	r := IntervalsContainerBuilder{
		&IntervalsContainer{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsContainer struct
func (rb *IntervalsContainerBuilder) Build() IntervalsContainer {
	return *rb.v
}

func (rb *IntervalsContainerBuilder) AllOf(allof *IntervalsAllOfBuilder) *IntervalsContainerBuilder {
	v := allof.Build()
	rb.v.AllOf = &v
	return rb
}

func (rb *IntervalsContainerBuilder) AnyOf(anyof *IntervalsAnyOfBuilder) *IntervalsContainerBuilder {
	v := anyof.Build()
	rb.v.AnyOf = &v
	return rb
}

func (rb *IntervalsContainerBuilder) Fuzzy(fuzzy *IntervalsFuzzyBuilder) *IntervalsContainerBuilder {
	v := fuzzy.Build()
	rb.v.Fuzzy = &v
	return rb
}

func (rb *IntervalsContainerBuilder) Match(match *IntervalsMatchBuilder) *IntervalsContainerBuilder {
	v := match.Build()
	rb.v.Match = &v
	return rb
}

func (rb *IntervalsContainerBuilder) Prefix(prefix *IntervalsPrefixBuilder) *IntervalsContainerBuilder {
	v := prefix.Build()
	rb.v.Prefix = &v
	return rb
}

func (rb *IntervalsContainerBuilder) Wildcard(wildcard *IntervalsWildcardBuilder) *IntervalsContainerBuilder {
	v := wildcard.Build()
	rb.v.Wildcard = &v
	return rb
}
