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

// IdsQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/term.ts#L53-L55
type IdsQuery struct {
	Boost      *float32 `json:"boost,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`
	Values     *Ids     `json:"values,omitempty"`
}

// IdsQueryBuilder holds IdsQuery struct and provides a builder API.
type IdsQueryBuilder struct {
	v *IdsQuery
}

// NewIdsQuery provides a builder for the IdsQuery struct.
func NewIdsQueryBuilder() *IdsQueryBuilder {
	r := IdsQueryBuilder{
		&IdsQuery{},
	}

	return &r
}

// Build finalize the chain and returns the IdsQuery struct
func (rb *IdsQueryBuilder) Build() IdsQuery {
	return *rb.v
}

func (rb *IdsQueryBuilder) Boost(boost float32) *IdsQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *IdsQueryBuilder) QueryName_(queryname_ string) *IdsQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *IdsQueryBuilder) Values(values *IdsBuilder) *IdsQueryBuilder {
	v := values.Build()
	rb.v.Values = &v
	return rb
}
