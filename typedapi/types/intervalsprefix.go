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

// IntervalsPrefix type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/fulltext.ts#L110-L114
type IntervalsPrefix struct {
	Analyzer *string `json:"analyzer,omitempty"`
	Prefix   string  `json:"prefix"`
	UseField *Field  `json:"use_field,omitempty"`
}

// IntervalsPrefixBuilder holds IntervalsPrefix struct and provides a builder API.
type IntervalsPrefixBuilder struct {
	v *IntervalsPrefix
}

// NewIntervalsPrefix provides a builder for the IntervalsPrefix struct.
func NewIntervalsPrefixBuilder() *IntervalsPrefixBuilder {
	r := IntervalsPrefixBuilder{
		&IntervalsPrefix{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsPrefix struct
func (rb *IntervalsPrefixBuilder) Build() IntervalsPrefix {
	return *rb.v
}

func (rb *IntervalsPrefixBuilder) Analyzer(analyzer string) *IntervalsPrefixBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *IntervalsPrefixBuilder) Prefix(prefix string) *IntervalsPrefixBuilder {
	rb.v.Prefix = prefix
	return rb
}

func (rb *IntervalsPrefixBuilder) UseField(usefield Field) *IntervalsPrefixBuilder {
	rb.v.UseField = &usefield
	return rb
}
