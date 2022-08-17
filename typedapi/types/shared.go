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

// Shared type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/searchable_snapshots/cache_stats/Response.ts#L34-L43
type Shared struct {
	BytesReadInBytes    ByteSize `json:"bytes_read_in_bytes"`
	BytesWrittenInBytes ByteSize `json:"bytes_written_in_bytes"`
	Evictions           int64    `json:"evictions"`
	NumRegions          int      `json:"num_regions"`
	Reads               int64    `json:"reads"`
	RegionSizeInBytes   ByteSize `json:"region_size_in_bytes"`
	SizeInBytes         ByteSize `json:"size_in_bytes"`
	Writes              int64    `json:"writes"`
}

// SharedBuilder holds Shared struct and provides a builder API.
type SharedBuilder struct {
	v *Shared
}

// NewShared provides a builder for the Shared struct.
func NewSharedBuilder() *SharedBuilder {
	r := SharedBuilder{
		&Shared{},
	}

	return &r
}

// Build finalize the chain and returns the Shared struct
func (rb *SharedBuilder) Build() Shared {
	return *rb.v
}

func (rb *SharedBuilder) BytesReadInBytes(bytesreadinbytes *ByteSizeBuilder) *SharedBuilder {
	v := bytesreadinbytes.Build()
	rb.v.BytesReadInBytes = v
	return rb
}

func (rb *SharedBuilder) BytesWrittenInBytes(byteswritteninbytes *ByteSizeBuilder) *SharedBuilder {
	v := byteswritteninbytes.Build()
	rb.v.BytesWrittenInBytes = v
	return rb
}

func (rb *SharedBuilder) Evictions(evictions int64) *SharedBuilder {
	rb.v.Evictions = evictions
	return rb
}

func (rb *SharedBuilder) NumRegions(numregions int) *SharedBuilder {
	rb.v.NumRegions = numregions
	return rb
}

func (rb *SharedBuilder) Reads(reads int64) *SharedBuilder {
	rb.v.Reads = reads
	return rb
}

func (rb *SharedBuilder) RegionSizeInBytes(regionsizeinbytes *ByteSizeBuilder) *SharedBuilder {
	v := regionsizeinbytes.Build()
	rb.v.RegionSizeInBytes = v
	return rb
}

func (rb *SharedBuilder) SizeInBytes(sizeinbytes *ByteSizeBuilder) *SharedBuilder {
	v := sizeinbytes.Build()
	rb.v.SizeInBytes = v
	return rb
}

func (rb *SharedBuilder) Writes(writes int64) *SharedBuilder {
	rb.v.Writes = writes
	return rb
}
