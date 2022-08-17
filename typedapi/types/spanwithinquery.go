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

// SpanWithinQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/span.ts#L74-L77
type SpanWithinQuery struct {
	Big        *SpanQuery `json:"big,omitempty"`
	Boost      *float32   `json:"boost,omitempty"`
	Little     *SpanQuery `json:"little,omitempty"`
	QueryName_ *string    `json:"_name,omitempty"`
}

// SpanWithinQueryBuilder holds SpanWithinQuery struct and provides a builder API.
type SpanWithinQueryBuilder struct {
	v *SpanWithinQuery
}

// NewSpanWithinQuery provides a builder for the SpanWithinQuery struct.
func NewSpanWithinQueryBuilder() *SpanWithinQueryBuilder {
	r := SpanWithinQueryBuilder{
		&SpanWithinQuery{},
	}

	return &r
}

// Build finalize the chain and returns the SpanWithinQuery struct
func (rb *SpanWithinQueryBuilder) Build() SpanWithinQuery {
	return *rb.v
}

func (rb *SpanWithinQueryBuilder) Big(big *SpanQueryBuilder) *SpanWithinQueryBuilder {
	v := big.Build()
	rb.v.Big = &v
	return rb
}

func (rb *SpanWithinQueryBuilder) Boost(boost float32) *SpanWithinQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *SpanWithinQueryBuilder) Little(little *SpanQueryBuilder) *SpanWithinQueryBuilder {
	v := little.Build()
	rb.v.Little = &v
	return rb
}

func (rb *SpanWithinQueryBuilder) QueryName_(queryname_ string) *SpanWithinQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
