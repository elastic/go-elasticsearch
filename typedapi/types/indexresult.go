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

// IndexResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L267-L269
type IndexResult struct {
	Response IndexResultSummary `json:"response"`
}

// IndexResultBuilder holds IndexResult struct and provides a builder API.
type IndexResultBuilder struct {
	v *IndexResult
}

// NewIndexResult provides a builder for the IndexResult struct.
func NewIndexResultBuilder() *IndexResultBuilder {
	r := IndexResultBuilder{
		&IndexResult{},
	}

	return &r
}

// Build finalize the chain and returns the IndexResult struct
func (rb *IndexResultBuilder) Build() IndexResult {
	return *rb.v
}

func (rb *IndexResultBuilder) Response(response *IndexResultSummaryBuilder) *IndexResultBuilder {
	v := response.Build()
	rb.v.Response = v
	return rb
}
