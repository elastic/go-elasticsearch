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

// RangeQuery holds the union for the following types:
//
//	DateRangeQuery
//	NumberRangeQuery
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/term.ts#L92-L94
type RangeQuery interface{}

// RangeQueryBuilder holds RangeQuery struct and provides a builder API.
type RangeQueryBuilder struct {
	v RangeQuery
}

// NewRangeQuery provides a builder for the RangeQuery struct.
func NewRangeQueryBuilder() *RangeQueryBuilder {
	return &RangeQueryBuilder{}
}

// Build finalize the chain and returns the RangeQuery struct
func (u *RangeQueryBuilder) Build() RangeQuery {
	return u.v
}

func (u *RangeQueryBuilder) DateRangeQuery(daterangequery *DateRangeQueryBuilder) *RangeQueryBuilder {
	v := daterangequery.Build()
	u.v = &v
	return u
}

func (u *RangeQueryBuilder) NumberRangeQuery(numberrangequery *NumberRangeQueryBuilder) *RangeQueryBuilder {
	v := numberrangequery.Build()
	u.v = &v
	return u
}
