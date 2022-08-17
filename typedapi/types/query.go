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

// Query type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L242-L247
type Query struct {
	Count  *int `json:"count,omitempty"`
	Failed *int `json:"failed,omitempty"`
	Paging *int `json:"paging,omitempty"`
	Total  *int `json:"total,omitempty"`
}

// QueryBuilder holds Query struct and provides a builder API.
type QueryBuilder struct {
	v *Query
}

// NewQuery provides a builder for the Query struct.
func NewQueryBuilder() *QueryBuilder {
	r := QueryBuilder{
		&Query{},
	}

	return &r
}

// Build finalize the chain and returns the Query struct
func (rb *QueryBuilder) Build() Query {
	return *rb.v
}

func (rb *QueryBuilder) Count(count int) *QueryBuilder {
	rb.v.Count = &count
	return rb
}

func (rb *QueryBuilder) Failed(failed int) *QueryBuilder {
	rb.v.Failed = &failed
	return rb
}

func (rb *QueryBuilder) Paging(paging int) *QueryBuilder {
	rb.v.Paging = &paging
	return rb
}

func (rb *QueryBuilder) Total(total int) *QueryBuilder {
	rb.v.Total = &total
	return rb
}
