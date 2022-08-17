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

// QueryBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/abstractions.ts#L175-L179
type QueryBase struct {
	Boost      *float32 `json:"boost,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`
}

// QueryBaseBuilder holds QueryBase struct and provides a builder API.
type QueryBaseBuilder struct {
	v *QueryBase
}

// NewQueryBase provides a builder for the QueryBase struct.
func NewQueryBaseBuilder() *QueryBaseBuilder {
	r := QueryBaseBuilder{
		&QueryBase{},
	}

	return &r
}

// Build finalize the chain and returns the QueryBase struct
func (rb *QueryBaseBuilder) Build() QueryBase {
	return *rb.v
}

func (rb *QueryBaseBuilder) Boost(boost float32) *QueryBaseBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *QueryBaseBuilder) QueryName_(queryname_ string) *QueryBaseBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
