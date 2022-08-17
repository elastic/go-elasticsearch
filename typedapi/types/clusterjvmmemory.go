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

// ClusterJvmMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/types.ts#L156-L159
type ClusterJvmMemory struct {
	HeapMaxInBytes  int64 `json:"heap_max_in_bytes"`
	HeapUsedInBytes int64 `json:"heap_used_in_bytes"`
}

// ClusterJvmMemoryBuilder holds ClusterJvmMemory struct and provides a builder API.
type ClusterJvmMemoryBuilder struct {
	v *ClusterJvmMemory
}

// NewClusterJvmMemory provides a builder for the ClusterJvmMemory struct.
func NewClusterJvmMemoryBuilder() *ClusterJvmMemoryBuilder {
	r := ClusterJvmMemoryBuilder{
		&ClusterJvmMemory{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterJvmMemory struct
func (rb *ClusterJvmMemoryBuilder) Build() ClusterJvmMemory {
	return *rb.v
}

func (rb *ClusterJvmMemoryBuilder) HeapMaxInBytes(heapmaxinbytes int64) *ClusterJvmMemoryBuilder {
	rb.v.HeapMaxInBytes = heapmaxinbytes
	return rb
}

func (rb *ClusterJvmMemoryBuilder) HeapUsedInBytes(heapusedinbytes int64) *ClusterJvmMemoryBuilder {
	rb.v.HeapUsedInBytes = heapusedinbytes
	return rb
}
