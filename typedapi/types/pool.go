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

// Pool type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L339-L344
type Pool struct {
	MaxInBytes      *int64 `json:"max_in_bytes,omitempty"`
	PeakMaxInBytes  *int64 `json:"peak_max_in_bytes,omitempty"`
	PeakUsedInBytes *int64 `json:"peak_used_in_bytes,omitempty"`
	UsedInBytes     *int64 `json:"used_in_bytes,omitempty"`
}

// PoolBuilder holds Pool struct and provides a builder API.
type PoolBuilder struct {
	v *Pool
}

// NewPool provides a builder for the Pool struct.
func NewPoolBuilder() *PoolBuilder {
	r := PoolBuilder{
		&Pool{},
	}

	return &r
}

// Build finalize the chain and returns the Pool struct
func (rb *PoolBuilder) Build() Pool {
	return *rb.v
}

func (rb *PoolBuilder) MaxInBytes(maxinbytes int64) *PoolBuilder {
	rb.v.MaxInBytes = &maxinbytes
	return rb
}

func (rb *PoolBuilder) PeakMaxInBytes(peakmaxinbytes int64) *PoolBuilder {
	rb.v.PeakMaxInBytes = &peakmaxinbytes
	return rb
}

func (rb *PoolBuilder) PeakUsedInBytes(peakusedinbytes int64) *PoolBuilder {
	rb.v.PeakUsedInBytes = &peakusedinbytes
	return rb
}

func (rb *PoolBuilder) UsedInBytes(usedinbytes int64) *PoolBuilder {
	rb.v.UsedInBytes = &usedinbytes
	return rb
}
