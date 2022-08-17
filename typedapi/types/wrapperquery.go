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

// WrapperQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/abstractions.ts#L197-L200
type WrapperQuery struct {
	Boost *float32 `json:"boost,omitempty"`
	// Query A base64 encoded query. The binary data format can be any of JSON, YAML, CBOR
	// or SMILE encodings
	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`
}

// WrapperQueryBuilder holds WrapperQuery struct and provides a builder API.
type WrapperQueryBuilder struct {
	v *WrapperQuery
}

// NewWrapperQuery provides a builder for the WrapperQuery struct.
func NewWrapperQueryBuilder() *WrapperQueryBuilder {
	r := WrapperQueryBuilder{
		&WrapperQuery{},
	}

	return &r
}

// Build finalize the chain and returns the WrapperQuery struct
func (rb *WrapperQueryBuilder) Build() WrapperQuery {
	return *rb.v
}

func (rb *WrapperQueryBuilder) Boost(boost float32) *WrapperQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Query A base64 encoded query. The binary data format can be any of JSON, YAML, CBOR
// or SMILE encodings

func (rb *WrapperQueryBuilder) Query(query string) *WrapperQueryBuilder {
	rb.v.Query = query
	return rb
}

func (rb *WrapperQueryBuilder) QueryName_(queryname_ string) *WrapperQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
