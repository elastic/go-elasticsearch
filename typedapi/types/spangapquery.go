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

// SpanGapQuery type alias.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/span.ts#L40-L42
type SpanGapQuery map[Field]int

// SpanGapQueryBuilder holds SpanGapQuery struct and provides a builder API.
type SpanGapQueryBuilder struct {
	v SpanGapQuery
}

// NewSpanGapQuery provides a builder for the SpanGapQuery struct.
func NewSpanGapQueryBuilder() *SpanGapQueryBuilder {
	return &SpanGapQueryBuilder{}
}

// Build finalize the chain and returns the SpanGapQuery struct
func (b *SpanGapQueryBuilder) Build() SpanGapQuery {
	return b.v
}

func (b *SpanGapQueryBuilder) SpanGapQuery(value SpanGapQuery) *SpanGapQueryBuilder {
	b.v = value
	return b
}
