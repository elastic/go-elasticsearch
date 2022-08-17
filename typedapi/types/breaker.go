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

// Breaker type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L173-L180
type Breaker struct {
	EstimatedSize        *string  `json:"estimated_size,omitempty"`
	EstimatedSizeInBytes *int64   `json:"estimated_size_in_bytes,omitempty"`
	LimitSize            *string  `json:"limit_size,omitempty"`
	LimitSizeInBytes     *int64   `json:"limit_size_in_bytes,omitempty"`
	Overhead             *float32 `json:"overhead,omitempty"`
	Tripped              *float32 `json:"tripped,omitempty"`
}

// BreakerBuilder holds Breaker struct and provides a builder API.
type BreakerBuilder struct {
	v *Breaker
}

// NewBreaker provides a builder for the Breaker struct.
func NewBreakerBuilder() *BreakerBuilder {
	r := BreakerBuilder{
		&Breaker{},
	}

	return &r
}

// Build finalize the chain and returns the Breaker struct
func (rb *BreakerBuilder) Build() Breaker {
	return *rb.v
}

func (rb *BreakerBuilder) EstimatedSize(estimatedsize string) *BreakerBuilder {
	rb.v.EstimatedSize = &estimatedsize
	return rb
}

func (rb *BreakerBuilder) EstimatedSizeInBytes(estimatedsizeinbytes int64) *BreakerBuilder {
	rb.v.EstimatedSizeInBytes = &estimatedsizeinbytes
	return rb
}

func (rb *BreakerBuilder) LimitSize(limitsize string) *BreakerBuilder {
	rb.v.LimitSize = &limitsize
	return rb
}

func (rb *BreakerBuilder) LimitSizeInBytes(limitsizeinbytes int64) *BreakerBuilder {
	rb.v.LimitSizeInBytes = &limitsizeinbytes
	return rb
}

func (rb *BreakerBuilder) Overhead(overhead float32) *BreakerBuilder {
	rb.v.Overhead = &overhead
	return rb
}

func (rb *BreakerBuilder) Tripped(tripped float32) *BreakerBuilder {
	rb.v.Tripped = &tripped
	return rb
}
