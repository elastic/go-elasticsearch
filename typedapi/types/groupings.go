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

// Groupings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/rollup/_types/Groupings.ts#L24-L28
type Groupings struct {
	DateHistogram *DateHistogramGrouping `json:"date_histogram,omitempty"`
	Histogram     *HistogramGrouping     `json:"histogram,omitempty"`
	Terms         *TermsGrouping         `json:"terms,omitempty"`
}

// GroupingsBuilder holds Groupings struct and provides a builder API.
type GroupingsBuilder struct {
	v *Groupings
}

// NewGroupings provides a builder for the Groupings struct.
func NewGroupingsBuilder() *GroupingsBuilder {
	r := GroupingsBuilder{
		&Groupings{},
	}

	return &r
}

// Build finalize the chain and returns the Groupings struct
func (rb *GroupingsBuilder) Build() Groupings {
	return *rb.v
}

func (rb *GroupingsBuilder) DateHistogram(datehistogram *DateHistogramGroupingBuilder) *GroupingsBuilder {
	v := datehistogram.Build()
	rb.v.DateHistogram = &v
	return rb
}

func (rb *GroupingsBuilder) Histogram(histogram *HistogramGroupingBuilder) *GroupingsBuilder {
	v := histogram.Build()
	rb.v.Histogram = &v
	return rb
}

func (rb *GroupingsBuilder) Terms(terms *TermsGroupingBuilder) *GroupingsBuilder {
	v := terms.Build()
	rb.v.Terms = &v
	return rb
}
