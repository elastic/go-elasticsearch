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

// FieldMemoryUsage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Stats.ts#L76-L79
type FieldMemoryUsage struct {
	MemorySize        *ByteSize `json:"memory_size,omitempty"`
	MemorySizeInBytes int64     `json:"memory_size_in_bytes"`
}

// FieldMemoryUsageBuilder holds FieldMemoryUsage struct and provides a builder API.
type FieldMemoryUsageBuilder struct {
	v *FieldMemoryUsage
}

// NewFieldMemoryUsage provides a builder for the FieldMemoryUsage struct.
func NewFieldMemoryUsageBuilder() *FieldMemoryUsageBuilder {
	r := FieldMemoryUsageBuilder{
		&FieldMemoryUsage{},
	}

	return &r
}

// Build finalize the chain and returns the FieldMemoryUsage struct
func (rb *FieldMemoryUsageBuilder) Build() FieldMemoryUsage {
	return *rb.v
}

func (rb *FieldMemoryUsageBuilder) MemorySize(memorysize *ByteSizeBuilder) *FieldMemoryUsageBuilder {
	v := memorysize.Build()
	rb.v.MemorySize = &v
	return rb
}

func (rb *FieldMemoryUsageBuilder) MemorySizeInBytes(memorysizeinbytes int64) *FieldMemoryUsageBuilder {
	rb.v.MemorySizeInBytes = memorysizeinbytes
	return rb
}
