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

// IntervalsQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/fulltext.ts#L116-L125
type IntervalsQuery struct {
	AllOf      *IntervalsAllOf    `json:"all_of,omitempty"`
	AnyOf      *IntervalsAnyOf    `json:"any_of,omitempty"`
	Boost      *float32           `json:"boost,omitempty"`
	Fuzzy      *IntervalsFuzzy    `json:"fuzzy,omitempty"`
	Match      *IntervalsMatch    `json:"match,omitempty"`
	Prefix     *IntervalsPrefix   `json:"prefix,omitempty"`
	QueryName_ *string            `json:"_name,omitempty"`
	Wildcard   *IntervalsWildcard `json:"wildcard,omitempty"`
}

// IntervalsQueryBuilder holds IntervalsQuery struct and provides a builder API.
type IntervalsQueryBuilder struct {
	v *IntervalsQuery
}

// NewIntervalsQuery provides a builder for the IntervalsQuery struct.
func NewIntervalsQueryBuilder() *IntervalsQueryBuilder {
	r := IntervalsQueryBuilder{
		&IntervalsQuery{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsQuery struct
func (rb *IntervalsQueryBuilder) Build() IntervalsQuery {
	return *rb.v
}

func (rb *IntervalsQueryBuilder) AllOf(allof *IntervalsAllOfBuilder) *IntervalsQueryBuilder {
	v := allof.Build()
	rb.v.AllOf = &v
	return rb
}

func (rb *IntervalsQueryBuilder) AnyOf(anyof *IntervalsAnyOfBuilder) *IntervalsQueryBuilder {
	v := anyof.Build()
	rb.v.AnyOf = &v
	return rb
}

func (rb *IntervalsQueryBuilder) Boost(boost float32) *IntervalsQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *IntervalsQueryBuilder) Fuzzy(fuzzy *IntervalsFuzzyBuilder) *IntervalsQueryBuilder {
	v := fuzzy.Build()
	rb.v.Fuzzy = &v
	return rb
}

func (rb *IntervalsQueryBuilder) Match(match *IntervalsMatchBuilder) *IntervalsQueryBuilder {
	v := match.Build()
	rb.v.Match = &v
	return rb
}

func (rb *IntervalsQueryBuilder) Prefix(prefix *IntervalsPrefixBuilder) *IntervalsQueryBuilder {
	v := prefix.Build()
	rb.v.Prefix = &v
	return rb
}

func (rb *IntervalsQueryBuilder) QueryName_(queryname_ string) *IntervalsQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *IntervalsQueryBuilder) Wildcard(wildcard *IntervalsWildcardBuilder) *IntervalsQueryBuilder {
	v := wildcard.Build()
	rb.v.Wildcard = &v
	return rb
}
