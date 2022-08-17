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

// FieldSizeUsage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Stats.ts#L59-L62
type FieldSizeUsage struct {
	Size        *ByteSize `json:"size,omitempty"`
	SizeInBytes int64     `json:"size_in_bytes"`
}

// FieldSizeUsageBuilder holds FieldSizeUsage struct and provides a builder API.
type FieldSizeUsageBuilder struct {
	v *FieldSizeUsage
}

// NewFieldSizeUsage provides a builder for the FieldSizeUsage struct.
func NewFieldSizeUsageBuilder() *FieldSizeUsageBuilder {
	r := FieldSizeUsageBuilder{
		&FieldSizeUsage{},
	}

	return &r
}

// Build finalize the chain and returns the FieldSizeUsage struct
func (rb *FieldSizeUsageBuilder) Build() FieldSizeUsage {
	return *rb.v
}

func (rb *FieldSizeUsageBuilder) Size(size *ByteSizeBuilder) *FieldSizeUsageBuilder {
	v := size.Build()
	rb.v.Size = &v
	return rb
}

func (rb *FieldSizeUsageBuilder) SizeInBytes(sizeinbytes int64) *FieldSizeUsageBuilder {
	rb.v.SizeInBytes = sizeinbytes
	return rb
}
