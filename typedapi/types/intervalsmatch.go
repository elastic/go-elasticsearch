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

// IntervalsMatch type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/fulltext.ts#L99-L108
type IntervalsMatch struct {
	Analyzer *string          `json:"analyzer,omitempty"`
	Filter   *IntervalsFilter `json:"filter,omitempty"`
	MaxGaps  *int             `json:"max_gaps,omitempty"`
	Ordered  *bool            `json:"ordered,omitempty"`
	Query    string           `json:"query"`
	UseField *Field           `json:"use_field,omitempty"`
}

// IntervalsMatchBuilder holds IntervalsMatch struct and provides a builder API.
type IntervalsMatchBuilder struct {
	v *IntervalsMatch
}

// NewIntervalsMatch provides a builder for the IntervalsMatch struct.
func NewIntervalsMatchBuilder() *IntervalsMatchBuilder {
	r := IntervalsMatchBuilder{
		&IntervalsMatch{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsMatch struct
func (rb *IntervalsMatchBuilder) Build() IntervalsMatch {
	return *rb.v
}

func (rb *IntervalsMatchBuilder) Analyzer(analyzer string) *IntervalsMatchBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *IntervalsMatchBuilder) Filter(filter *IntervalsFilterBuilder) *IntervalsMatchBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *IntervalsMatchBuilder) MaxGaps(maxgaps int) *IntervalsMatchBuilder {
	rb.v.MaxGaps = &maxgaps
	return rb
}

func (rb *IntervalsMatchBuilder) Ordered(ordered bool) *IntervalsMatchBuilder {
	rb.v.Ordered = &ordered
	return rb
}

func (rb *IntervalsMatchBuilder) Query(query string) *IntervalsMatchBuilder {
	rb.v.Query = query
	return rb
}

func (rb *IntervalsMatchBuilder) UseField(usefield Field) *IntervalsMatchBuilder {
	rb.v.UseField = &usefield
	return rb
}
