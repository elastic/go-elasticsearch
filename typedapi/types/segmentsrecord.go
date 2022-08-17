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

// SegmentsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/segments/types.ts#L22-L96
type SegmentsRecord struct {
	// Committed is segment committed
	Committed *string `json:"committed,omitempty"`
	// Compound is segment compound
	Compound *string `json:"compound,omitempty"`
	// DocsCount number of docs in segment
	DocsCount *string `json:"docs.count,omitempty"`
	// DocsDeleted number of deleted docs in segment
	DocsDeleted *string `json:"docs.deleted,omitempty"`
	// Generation segment generation
	Generation *string `json:"generation,omitempty"`
	// Id unique id of node where it lives
	Id *NodeId `json:"id,omitempty"`
	// Index index name
	Index *IndexName `json:"index,omitempty"`
	// Ip ip of node where it lives
	Ip *string `json:"ip,omitempty"`
	// Prirep primary or replica
	Prirep *string `json:"prirep,omitempty"`
	// Searchable is segment searched
	Searchable *string `json:"searchable,omitempty"`
	// Segment segment name
	Segment *string `json:"segment,omitempty"`
	// Shard shard name
	Shard *string `json:"shard,omitempty"`
	// Size segment size in bytes
	Size *ByteSize `json:"size,omitempty"`
	// SizeMemory segment memory in bytes
	SizeMemory *ByteSize `json:"size.memory,omitempty"`
	// Version version
	Version *VersionString `json:"version,omitempty"`
}

// SegmentsRecordBuilder holds SegmentsRecord struct and provides a builder API.
type SegmentsRecordBuilder struct {
	v *SegmentsRecord
}

// NewSegmentsRecord provides a builder for the SegmentsRecord struct.
func NewSegmentsRecordBuilder() *SegmentsRecordBuilder {
	r := SegmentsRecordBuilder{
		&SegmentsRecord{},
	}

	return &r
}

// Build finalize the chain and returns the SegmentsRecord struct
func (rb *SegmentsRecordBuilder) Build() SegmentsRecord {
	return *rb.v
}

// Committed is segment committed

func (rb *SegmentsRecordBuilder) Committed(committed string) *SegmentsRecordBuilder {
	rb.v.Committed = &committed
	return rb
}

// Compound is segment compound

func (rb *SegmentsRecordBuilder) Compound(compound string) *SegmentsRecordBuilder {
	rb.v.Compound = &compound
	return rb
}

// DocsCount number of docs in segment

func (rb *SegmentsRecordBuilder) DocsCount(docscount string) *SegmentsRecordBuilder {
	rb.v.DocsCount = &docscount
	return rb
}

// DocsDeleted number of deleted docs in segment

func (rb *SegmentsRecordBuilder) DocsDeleted(docsdeleted string) *SegmentsRecordBuilder {
	rb.v.DocsDeleted = &docsdeleted
	return rb
}

// Generation segment generation

func (rb *SegmentsRecordBuilder) Generation(generation string) *SegmentsRecordBuilder {
	rb.v.Generation = &generation
	return rb
}

// Id unique id of node where it lives

func (rb *SegmentsRecordBuilder) Id(id NodeId) *SegmentsRecordBuilder {
	rb.v.Id = &id
	return rb
}

// Index index name

func (rb *SegmentsRecordBuilder) Index(index IndexName) *SegmentsRecordBuilder {
	rb.v.Index = &index
	return rb
}

// Ip ip of node where it lives

func (rb *SegmentsRecordBuilder) Ip(ip string) *SegmentsRecordBuilder {
	rb.v.Ip = &ip
	return rb
}

// Prirep primary or replica

func (rb *SegmentsRecordBuilder) Prirep(prirep string) *SegmentsRecordBuilder {
	rb.v.Prirep = &prirep
	return rb
}

// Searchable is segment searched

func (rb *SegmentsRecordBuilder) Searchable(searchable string) *SegmentsRecordBuilder {
	rb.v.Searchable = &searchable
	return rb
}

// Segment segment name

func (rb *SegmentsRecordBuilder) Segment(segment string) *SegmentsRecordBuilder {
	rb.v.Segment = &segment
	return rb
}

// Shard shard name

func (rb *SegmentsRecordBuilder) Shard(shard string) *SegmentsRecordBuilder {
	rb.v.Shard = &shard
	return rb
}

// Size segment size in bytes

func (rb *SegmentsRecordBuilder) Size(size *ByteSizeBuilder) *SegmentsRecordBuilder {
	v := size.Build()
	rb.v.Size = &v
	return rb
}

// SizeMemory segment memory in bytes

func (rb *SegmentsRecordBuilder) SizeMemory(sizememory *ByteSizeBuilder) *SegmentsRecordBuilder {
	v := sizememory.Build()
	rb.v.SizeMemory = &v
	return rb
}

// Version version

func (rb *SegmentsRecordBuilder) Version(version VersionString) *SegmentsRecordBuilder {
	rb.v.Version = &version
	return rb
}
