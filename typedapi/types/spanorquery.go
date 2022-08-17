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

// SpanOrQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/span.ts#L65-L67
type SpanOrQuery struct {
	Boost      *float32    `json:"boost,omitempty"`
	Clauses    []SpanQuery `json:"clauses"`
	QueryName_ *string     `json:"_name,omitempty"`
}

// SpanOrQueryBuilder holds SpanOrQuery struct and provides a builder API.
type SpanOrQueryBuilder struct {
	v *SpanOrQuery
}

// NewSpanOrQuery provides a builder for the SpanOrQuery struct.
func NewSpanOrQueryBuilder() *SpanOrQueryBuilder {
	r := SpanOrQueryBuilder{
		&SpanOrQuery{},
	}

	return &r
}

// Build finalize the chain and returns the SpanOrQuery struct
func (rb *SpanOrQueryBuilder) Build() SpanOrQuery {
	return *rb.v
}

func (rb *SpanOrQueryBuilder) Boost(boost float32) *SpanOrQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *SpanOrQueryBuilder) Clauses(clauses []SpanQueryBuilder) *SpanOrQueryBuilder {
	tmp := make([]SpanQuery, len(clauses))
	for _, value := range clauses {
		tmp = append(tmp, value.Build())
	}
	rb.v.Clauses = tmp
	return rb
}

func (rb *SpanOrQueryBuilder) QueryName_(queryname_ string) *SpanOrQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
