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

// SpanFirstQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/span.ts#L35-L38
type SpanFirstQuery struct {
	Boost      *float32   `json:"boost,omitempty"`
	End        int        `json:"end"`
	Match      *SpanQuery `json:"match,omitempty"`
	QueryName_ *string    `json:"_name,omitempty"`
}

// SpanFirstQueryBuilder holds SpanFirstQuery struct and provides a builder API.
type SpanFirstQueryBuilder struct {
	v *SpanFirstQuery
}

// NewSpanFirstQuery provides a builder for the SpanFirstQuery struct.
func NewSpanFirstQueryBuilder() *SpanFirstQueryBuilder {
	r := SpanFirstQueryBuilder{
		&SpanFirstQuery{},
	}

	return &r
}

// Build finalize the chain and returns the SpanFirstQuery struct
func (rb *SpanFirstQueryBuilder) Build() SpanFirstQuery {
	return *rb.v
}

func (rb *SpanFirstQueryBuilder) Boost(boost float32) *SpanFirstQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *SpanFirstQueryBuilder) End(end int) *SpanFirstQueryBuilder {
	rb.v.End = end
	return rb
}

func (rb *SpanFirstQueryBuilder) Match(match *SpanQueryBuilder) *SpanFirstQueryBuilder {
	v := match.Build()
	rb.v.Match = &v
	return rb
}

func (rb *SpanFirstQueryBuilder) QueryName_(queryname_ string) *SpanFirstQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
