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

// PinnedQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/specialized.ts#L122-L130
type PinnedQuery struct {
	Boost      *float32        `json:"boost,omitempty"`
	Docs       []PinnedDoc     `json:"docs,omitempty"`
	Ids        []Id            `json:"ids,omitempty"`
	Organic    *QueryContainer `json:"organic,omitempty"`
	QueryName_ *string         `json:"_name,omitempty"`
}

// PinnedQueryBuilder holds PinnedQuery struct and provides a builder API.
type PinnedQueryBuilder struct {
	v *PinnedQuery
}

// NewPinnedQuery provides a builder for the PinnedQuery struct.
func NewPinnedQueryBuilder() *PinnedQueryBuilder {
	r := PinnedQueryBuilder{
		&PinnedQuery{},
	}

	return &r
}

// Build finalize the chain and returns the PinnedQuery struct
func (rb *PinnedQueryBuilder) Build() PinnedQuery {
	return *rb.v
}

func (rb *PinnedQueryBuilder) Boost(boost float32) *PinnedQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *PinnedQueryBuilder) Docs(docs []PinnedDocBuilder) *PinnedQueryBuilder {
	tmp := make([]PinnedDoc, len(docs))
	for _, value := range docs {
		tmp = append(tmp, value.Build())
	}
	rb.v.Docs = tmp
	return rb
}

func (rb *PinnedQueryBuilder) Ids(ids ...Id) *PinnedQueryBuilder {
	rb.v.Ids = ids
	return rb
}

func (rb *PinnedQueryBuilder) Organic(organic *QueryContainerBuilder) *PinnedQueryBuilder {
	v := organic.Build()
	rb.v.Organic = &v
	return rb
}

func (rb *PinnedQueryBuilder) QueryName_(queryname_ string) *PinnedQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
