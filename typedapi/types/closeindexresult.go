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

// CloseIndexResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/close/CloseIndexResponse.ts#L32-L35
type CloseIndexResult struct {
	Closed bool                        `json:"closed"`
	Shards map[string]CloseShardResult `json:"shards,omitempty"`
}

// CloseIndexResultBuilder holds CloseIndexResult struct and provides a builder API.
type CloseIndexResultBuilder struct {
	v *CloseIndexResult
}

// NewCloseIndexResult provides a builder for the CloseIndexResult struct.
func NewCloseIndexResultBuilder() *CloseIndexResultBuilder {
	r := CloseIndexResultBuilder{
		&CloseIndexResult{
			Shards: make(map[string]CloseShardResult, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the CloseIndexResult struct
func (rb *CloseIndexResultBuilder) Build() CloseIndexResult {
	return *rb.v
}

func (rb *CloseIndexResultBuilder) Closed(closed bool) *CloseIndexResultBuilder {
	rb.v.Closed = closed
	return rb
}

func (rb *CloseIndexResultBuilder) Shards(values map[string]*CloseShardResultBuilder) *CloseIndexResultBuilder {
	tmp := make(map[string]CloseShardResult, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Shards = tmp
	return rb
}
