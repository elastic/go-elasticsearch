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

// JvmMemoryStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L329-L337
type JvmMemoryStats struct {
	HeapCommittedInBytes    *int64          `json:"heap_committed_in_bytes,omitempty"`
	HeapMaxInBytes          *int64          `json:"heap_max_in_bytes,omitempty"`
	HeapUsedInBytes         *int64          `json:"heap_used_in_bytes,omitempty"`
	HeapUsedPercent         *int64          `json:"heap_used_percent,omitempty"`
	NonHeapCommittedInBytes *int64          `json:"non_heap_committed_in_bytes,omitempty"`
	NonHeapUsedInBytes      *int64          `json:"non_heap_used_in_bytes,omitempty"`
	Pools                   map[string]Pool `json:"pools,omitempty"`
}

// JvmMemoryStatsBuilder holds JvmMemoryStats struct and provides a builder API.
type JvmMemoryStatsBuilder struct {
	v *JvmMemoryStats
}

// NewJvmMemoryStats provides a builder for the JvmMemoryStats struct.
func NewJvmMemoryStatsBuilder() *JvmMemoryStatsBuilder {
	r := JvmMemoryStatsBuilder{
		&JvmMemoryStats{
			Pools: make(map[string]Pool, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the JvmMemoryStats struct
func (rb *JvmMemoryStatsBuilder) Build() JvmMemoryStats {
	return *rb.v
}

func (rb *JvmMemoryStatsBuilder) HeapCommittedInBytes(heapcommittedinbytes int64) *JvmMemoryStatsBuilder {
	rb.v.HeapCommittedInBytes = &heapcommittedinbytes
	return rb
}

func (rb *JvmMemoryStatsBuilder) HeapMaxInBytes(heapmaxinbytes int64) *JvmMemoryStatsBuilder {
	rb.v.HeapMaxInBytes = &heapmaxinbytes
	return rb
}

func (rb *JvmMemoryStatsBuilder) HeapUsedInBytes(heapusedinbytes int64) *JvmMemoryStatsBuilder {
	rb.v.HeapUsedInBytes = &heapusedinbytes
	return rb
}

func (rb *JvmMemoryStatsBuilder) HeapUsedPercent(heapusedpercent int64) *JvmMemoryStatsBuilder {
	rb.v.HeapUsedPercent = &heapusedpercent
	return rb
}

func (rb *JvmMemoryStatsBuilder) NonHeapCommittedInBytes(nonheapcommittedinbytes int64) *JvmMemoryStatsBuilder {
	rb.v.NonHeapCommittedInBytes = &nonheapcommittedinbytes
	return rb
}

func (rb *JvmMemoryStatsBuilder) NonHeapUsedInBytes(nonheapusedinbytes int64) *JvmMemoryStatsBuilder {
	rb.v.NonHeapUsedInBytes = &nonheapusedinbytes
	return rb
}

func (rb *JvmMemoryStatsBuilder) Pools(values map[string]*PoolBuilder) *JvmMemoryStatsBuilder {
	tmp := make(map[string]Pool, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Pools = tmp
	return rb
}
