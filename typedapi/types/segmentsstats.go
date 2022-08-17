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

// SegmentsStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Stats.ts#L206-L231
type SegmentsStats struct {
	Count                       int                          `json:"count"`
	DocValuesMemory             *ByteSize                    `json:"doc_values_memory,omitempty"`
	DocValuesMemoryInBytes      int                          `json:"doc_values_memory_in_bytes"`
	FileSizes                   map[string]ShardFileSizeInfo `json:"file_sizes"`
	FixedBitSet                 *ByteSize                    `json:"fixed_bit_set,omitempty"`
	FixedBitSetMemoryInBytes    int                          `json:"fixed_bit_set_memory_in_bytes"`
	IndexWriterMaxMemoryInBytes *int                         `json:"index_writer_max_memory_in_bytes,omitempty"`
	IndexWriterMemory           *ByteSize                    `json:"index_writer_memory,omitempty"`
	IndexWriterMemoryInBytes    int                          `json:"index_writer_memory_in_bytes"`
	MaxUnsafeAutoIdTimestamp    int64                        `json:"max_unsafe_auto_id_timestamp"`
	Memory                      *ByteSize                    `json:"memory,omitempty"`
	MemoryInBytes               int                          `json:"memory_in_bytes"`
	NormsMemory                 *ByteSize                    `json:"norms_memory,omitempty"`
	NormsMemoryInBytes          int                          `json:"norms_memory_in_bytes"`
	PointsMemory                *ByteSize                    `json:"points_memory,omitempty"`
	PointsMemoryInBytes         int                          `json:"points_memory_in_bytes"`
	StoredFieldsMemoryInBytes   int                          `json:"stored_fields_memory_in_bytes"`
	StoredMemory                *ByteSize                    `json:"stored_memory,omitempty"`
	TermVectorsMemoryInBytes    int                          `json:"term_vectors_memory_in_bytes"`
	TermVectoryMemory           *ByteSize                    `json:"term_vectory_memory,omitempty"`
	TermsMemory                 *ByteSize                    `json:"terms_memory,omitempty"`
	TermsMemoryInBytes          int                          `json:"terms_memory_in_bytes"`
	VersionMapMemory            *ByteSize                    `json:"version_map_memory,omitempty"`
	VersionMapMemoryInBytes     int                          `json:"version_map_memory_in_bytes"`
}

// SegmentsStatsBuilder holds SegmentsStats struct and provides a builder API.
type SegmentsStatsBuilder struct {
	v *SegmentsStats
}

// NewSegmentsStats provides a builder for the SegmentsStats struct.
func NewSegmentsStatsBuilder() *SegmentsStatsBuilder {
	r := SegmentsStatsBuilder{
		&SegmentsStats{
			FileSizes: make(map[string]ShardFileSizeInfo, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SegmentsStats struct
func (rb *SegmentsStatsBuilder) Build() SegmentsStats {
	return *rb.v
}

func (rb *SegmentsStatsBuilder) Count(count int) *SegmentsStatsBuilder {
	rb.v.Count = count
	return rb
}

func (rb *SegmentsStatsBuilder) DocValuesMemory(docvaluesmemory *ByteSizeBuilder) *SegmentsStatsBuilder {
	v := docvaluesmemory.Build()
	rb.v.DocValuesMemory = &v
	return rb
}

func (rb *SegmentsStatsBuilder) DocValuesMemoryInBytes(docvaluesmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.DocValuesMemoryInBytes = docvaluesmemoryinbytes
	return rb
}

func (rb *SegmentsStatsBuilder) FileSizes(values map[string]*ShardFileSizeInfoBuilder) *SegmentsStatsBuilder {
	tmp := make(map[string]ShardFileSizeInfo, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.FileSizes = tmp
	return rb
}

func (rb *SegmentsStatsBuilder) FixedBitSet(fixedbitset *ByteSizeBuilder) *SegmentsStatsBuilder {
	v := fixedbitset.Build()
	rb.v.FixedBitSet = &v
	return rb
}

func (rb *SegmentsStatsBuilder) FixedBitSetMemoryInBytes(fixedbitsetmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.FixedBitSetMemoryInBytes = fixedbitsetmemoryinbytes
	return rb
}

func (rb *SegmentsStatsBuilder) IndexWriterMaxMemoryInBytes(indexwritermaxmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.IndexWriterMaxMemoryInBytes = &indexwritermaxmemoryinbytes
	return rb
}

func (rb *SegmentsStatsBuilder) IndexWriterMemory(indexwritermemory *ByteSizeBuilder) *SegmentsStatsBuilder {
	v := indexwritermemory.Build()
	rb.v.IndexWriterMemory = &v
	return rb
}

func (rb *SegmentsStatsBuilder) IndexWriterMemoryInBytes(indexwritermemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.IndexWriterMemoryInBytes = indexwritermemoryinbytes
	return rb
}

func (rb *SegmentsStatsBuilder) MaxUnsafeAutoIdTimestamp(maxunsafeautoidtimestamp int64) *SegmentsStatsBuilder {
	rb.v.MaxUnsafeAutoIdTimestamp = maxunsafeautoidtimestamp
	return rb
}

func (rb *SegmentsStatsBuilder) Memory(memory *ByteSizeBuilder) *SegmentsStatsBuilder {
	v := memory.Build()
	rb.v.Memory = &v
	return rb
}

func (rb *SegmentsStatsBuilder) MemoryInBytes(memoryinbytes int) *SegmentsStatsBuilder {
	rb.v.MemoryInBytes = memoryinbytes
	return rb
}

func (rb *SegmentsStatsBuilder) NormsMemory(normsmemory *ByteSizeBuilder) *SegmentsStatsBuilder {
	v := normsmemory.Build()
	rb.v.NormsMemory = &v
	return rb
}

func (rb *SegmentsStatsBuilder) NormsMemoryInBytes(normsmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.NormsMemoryInBytes = normsmemoryinbytes
	return rb
}

func (rb *SegmentsStatsBuilder) PointsMemory(pointsmemory *ByteSizeBuilder) *SegmentsStatsBuilder {
	v := pointsmemory.Build()
	rb.v.PointsMemory = &v
	return rb
}

func (rb *SegmentsStatsBuilder) PointsMemoryInBytes(pointsmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.PointsMemoryInBytes = pointsmemoryinbytes
	return rb
}

func (rb *SegmentsStatsBuilder) StoredFieldsMemoryInBytes(storedfieldsmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.StoredFieldsMemoryInBytes = storedfieldsmemoryinbytes
	return rb
}

func (rb *SegmentsStatsBuilder) StoredMemory(storedmemory *ByteSizeBuilder) *SegmentsStatsBuilder {
	v := storedmemory.Build()
	rb.v.StoredMemory = &v
	return rb
}

func (rb *SegmentsStatsBuilder) TermVectorsMemoryInBytes(termvectorsmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.TermVectorsMemoryInBytes = termvectorsmemoryinbytes
	return rb
}

func (rb *SegmentsStatsBuilder) TermVectoryMemory(termvectorymemory *ByteSizeBuilder) *SegmentsStatsBuilder {
	v := termvectorymemory.Build()
	rb.v.TermVectoryMemory = &v
	return rb
}

func (rb *SegmentsStatsBuilder) TermsMemory(termsmemory *ByteSizeBuilder) *SegmentsStatsBuilder {
	v := termsmemory.Build()
	rb.v.TermsMemory = &v
	return rb
}

func (rb *SegmentsStatsBuilder) TermsMemoryInBytes(termsmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.TermsMemoryInBytes = termsmemoryinbytes
	return rb
}

func (rb *SegmentsStatsBuilder) VersionMapMemory(versionmapmemory *ByteSizeBuilder) *SegmentsStatsBuilder {
	v := versionmapmemory.Build()
	rb.v.VersionMapMemory = &v
	return rb
}

func (rb *SegmentsStatsBuilder) VersionMapMemoryInBytes(versionmapmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.VersionMapMemoryInBytes = versionmapmemoryinbytes
	return rb
}
