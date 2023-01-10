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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// SegmentsStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/Stats.ts#L206-L231
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

// NewSegmentsStats returns a SegmentsStats.
func NewSegmentsStats() *SegmentsStats {
	r := &SegmentsStats{
		FileSizes: make(map[string]ShardFileSizeInfo, 0),
	}

	return r
}
