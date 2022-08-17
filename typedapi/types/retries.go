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

// Retries type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Retries.ts#L22-L25
type Retries struct {
	Bulk   int64 `json:"bulk"`
	Search int64 `json:"search"`
}

// RetriesBuilder holds Retries struct and provides a builder API.
type RetriesBuilder struct {
	v *Retries
}

// NewRetries provides a builder for the Retries struct.
func NewRetriesBuilder() *RetriesBuilder {
	r := RetriesBuilder{
		&Retries{},
	}

	return &r
}

// Build finalize the chain and returns the Retries struct
func (rb *RetriesBuilder) Build() Retries {
	return *rb.v
}

func (rb *RetriesBuilder) Bulk(bulk int64) *RetriesBuilder {
	rb.v.Bulk = bulk
	return rb
}

func (rb *RetriesBuilder) Search(search int64) *RetriesBuilder {
	rb.v.Search = search
	return rb
}
