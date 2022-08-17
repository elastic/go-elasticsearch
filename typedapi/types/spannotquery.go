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

// SpanNotQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/span.ts#L55-L63
type SpanNotQuery struct {
	Boost      *float32   `json:"boost,omitempty"`
	Dist       *int       `json:"dist,omitempty"`
	Exclude    *SpanQuery `json:"exclude,omitempty"`
	Include    *SpanQuery `json:"include,omitempty"`
	Post       *int       `json:"post,omitempty"`
	Pre        *int       `json:"pre,omitempty"`
	QueryName_ *string    `json:"_name,omitempty"`
}

// SpanNotQueryBuilder holds SpanNotQuery struct and provides a builder API.
type SpanNotQueryBuilder struct {
	v *SpanNotQuery
}

// NewSpanNotQuery provides a builder for the SpanNotQuery struct.
func NewSpanNotQueryBuilder() *SpanNotQueryBuilder {
	r := SpanNotQueryBuilder{
		&SpanNotQuery{},
	}

	return &r
}

// Build finalize the chain and returns the SpanNotQuery struct
func (rb *SpanNotQueryBuilder) Build() SpanNotQuery {
	return *rb.v
}

func (rb *SpanNotQueryBuilder) Boost(boost float32) *SpanNotQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *SpanNotQueryBuilder) Dist(dist int) *SpanNotQueryBuilder {
	rb.v.Dist = &dist
	return rb
}

func (rb *SpanNotQueryBuilder) Exclude(exclude *SpanQueryBuilder) *SpanNotQueryBuilder {
	v := exclude.Build()
	rb.v.Exclude = &v
	return rb
}

func (rb *SpanNotQueryBuilder) Include(include *SpanQueryBuilder) *SpanNotQueryBuilder {
	v := include.Build()
	rb.v.Include = &v
	return rb
}

func (rb *SpanNotQueryBuilder) Post(post int) *SpanNotQueryBuilder {
	rb.v.Post = &post
	return rb
}

func (rb *SpanNotQueryBuilder) Pre(pre int) *SpanNotQueryBuilder {
	rb.v.Pre = &pre
	return rb
}

func (rb *SpanNotQueryBuilder) QueryName_(queryname_ string) *SpanNotQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
