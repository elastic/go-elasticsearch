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

// SpanMultiTermQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/span.ts#L44-L47
type SpanMultiTermQuery struct {
	Boost *float32 `json:"boost,omitempty"`
	// Match Should be a multi term query (one of wildcard, fuzzy, prefix, range or regexp
	// query)
	Match      *QueryContainer `json:"match,omitempty"`
	QueryName_ *string         `json:"_name,omitempty"`
}

// SpanMultiTermQueryBuilder holds SpanMultiTermQuery struct and provides a builder API.
type SpanMultiTermQueryBuilder struct {
	v *SpanMultiTermQuery
}

// NewSpanMultiTermQuery provides a builder for the SpanMultiTermQuery struct.
func NewSpanMultiTermQueryBuilder() *SpanMultiTermQueryBuilder {
	r := SpanMultiTermQueryBuilder{
		&SpanMultiTermQuery{},
	}

	return &r
}

// Build finalize the chain and returns the SpanMultiTermQuery struct
func (rb *SpanMultiTermQueryBuilder) Build() SpanMultiTermQuery {
	return *rb.v
}

func (rb *SpanMultiTermQueryBuilder) Boost(boost float32) *SpanMultiTermQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Match Should be a multi term query (one of wildcard, fuzzy, prefix, range or regexp
// query)

func (rb *SpanMultiTermQueryBuilder) Match(match *QueryContainerBuilder) *SpanMultiTermQueryBuilder {
	v := match.Build()
	rb.v.Match = &v
	return rb
}

func (rb *SpanMultiTermQueryBuilder) QueryName_(queryname_ string) *SpanMultiTermQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
