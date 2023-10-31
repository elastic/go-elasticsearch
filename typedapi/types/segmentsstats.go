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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// SegmentsStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/Stats.ts#L273-L366
type SegmentsStats struct {
	// Count Total number of segments across all shards assigned to selected nodes.
	Count int `json:"count"`
	// DocValuesMemory Total amount of memory used for doc values across all shards assigned to
	// selected nodes.
	DocValuesMemory ByteSize `json:"doc_values_memory,omitempty"`
	// DocValuesMemoryInBytes Total amount, in bytes, of memory used for doc values across all shards
	// assigned to selected nodes.
	DocValuesMemoryInBytes int64 `json:"doc_values_memory_in_bytes"`
	// FileSizes This object is not populated by the cluster stats API.
	// To get information on segment files, use the node stats API.
	FileSizes map[string]ShardFileSizeInfo `json:"file_sizes"`
	// FixedBitSet Total amount of memory used by fixed bit sets across all shards assigned to
	// selected nodes.
	// Fixed bit sets are used for nested object field types and type filters for
	// join fields.
	FixedBitSet ByteSize `json:"fixed_bit_set,omitempty"`
	// FixedBitSetMemoryInBytes Total amount of memory, in bytes, used by fixed bit sets across all shards
	// assigned to selected nodes.
	FixedBitSetMemoryInBytes    int64  `json:"fixed_bit_set_memory_in_bytes"`
	IndexWriterMaxMemoryInBytes *int64 `json:"index_writer_max_memory_in_bytes,omitempty"`
	// IndexWriterMemory Total amount of memory used by all index writers across all shards assigned
	// to selected nodes.
	IndexWriterMemory ByteSize `json:"index_writer_memory,omitempty"`
	// IndexWriterMemoryInBytes Total amount, in bytes, of memory used by all index writers across all shards
	// assigned to selected nodes.
	IndexWriterMemoryInBytes int64 `json:"index_writer_memory_in_bytes"`
	// MaxUnsafeAutoIdTimestamp Unix timestamp, in milliseconds, of the most recently retried indexing
	// request.
	MaxUnsafeAutoIdTimestamp int64 `json:"max_unsafe_auto_id_timestamp"`
	// Memory Total amount of memory used for segments across all shards assigned to
	// selected nodes.
	Memory ByteSize `json:"memory,omitempty"`
	// MemoryInBytes Total amount, in bytes, of memory used for segments across all shards
	// assigned to selected nodes.
	MemoryInBytes int64 `json:"memory_in_bytes"`
	// NormsMemory Total amount of memory used for normalization factors across all shards
	// assigned to selected nodes.
	NormsMemory ByteSize `json:"norms_memory,omitempty"`
	// NormsMemoryInBytes Total amount, in bytes, of memory used for normalization factors across all
	// shards assigned to selected nodes.
	NormsMemoryInBytes int64 `json:"norms_memory_in_bytes"`
	// PointsMemory Total amount of memory used for points across all shards assigned to selected
	// nodes.
	PointsMemory ByteSize `json:"points_memory,omitempty"`
	// PointsMemoryInBytes Total amount, in bytes, of memory used for points across all shards assigned
	// to selected nodes.
	PointsMemoryInBytes int64 `json:"points_memory_in_bytes"`
	// StoredFieldsMemoryInBytes Total amount, in bytes, of memory used for stored fields across all shards
	// assigned to selected nodes.
	StoredFieldsMemoryInBytes int64    `json:"stored_fields_memory_in_bytes"`
	StoredMemory              ByteSize `json:"stored_memory,omitempty"`
	// TermVectorsMemoryInBytes Total amount, in bytes, of memory used for term vectors across all shards
	// assigned to selected nodes.
	TermVectorsMemoryInBytes int64 `json:"term_vectors_memory_in_bytes"`
	// TermVectoryMemory Total amount of memory used for term vectors across all shards assigned to
	// selected nodes.
	TermVectoryMemory ByteSize `json:"term_vectory_memory,omitempty"`
	// TermsMemory Total amount of memory used for terms across all shards assigned to selected
	// nodes.
	TermsMemory ByteSize `json:"terms_memory,omitempty"`
	// TermsMemoryInBytes Total amount, in bytes, of memory used for terms across all shards assigned
	// to selected nodes.
	TermsMemoryInBytes int64 `json:"terms_memory_in_bytes"`
	// VersionMapMemory Total amount of memory used by all version maps across all shards assigned to
	// selected nodes.
	VersionMapMemory ByteSize `json:"version_map_memory,omitempty"`
	// VersionMapMemoryInBytes Total amount, in bytes, of memory used by all version maps across all shards
	// assigned to selected nodes.
	VersionMapMemoryInBytes int64 `json:"version_map_memory_in_bytes"`
}

func (s *SegmentsStats) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "doc_values_memory":
			if err := dec.Decode(&s.DocValuesMemory); err != nil {
				return err
			}

		case "doc_values_memory_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DocValuesMemoryInBytes = value
			case float64:
				f := int64(v)
				s.DocValuesMemoryInBytes = f
			}

		case "file_sizes":
			if s.FileSizes == nil {
				s.FileSizes = make(map[string]ShardFileSizeInfo, 0)
			}
			if err := dec.Decode(&s.FileSizes); err != nil {
				return err
			}

		case "fixed_bit_set":
			if err := dec.Decode(&s.FixedBitSet); err != nil {
				return err
			}

		case "fixed_bit_set_memory_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.FixedBitSetMemoryInBytes = value
			case float64:
				f := int64(v)
				s.FixedBitSetMemoryInBytes = f
			}

		case "index_writer_max_memory_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.IndexWriterMaxMemoryInBytes = &value
			case float64:
				f := int64(v)
				s.IndexWriterMaxMemoryInBytes = &f
			}

		case "index_writer_memory":
			if err := dec.Decode(&s.IndexWriterMemory); err != nil {
				return err
			}

		case "index_writer_memory_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.IndexWriterMemoryInBytes = value
			case float64:
				f := int64(v)
				s.IndexWriterMemoryInBytes = f
			}

		case "max_unsafe_auto_id_timestamp":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MaxUnsafeAutoIdTimestamp = value
			case float64:
				f := int64(v)
				s.MaxUnsafeAutoIdTimestamp = f
			}

		case "memory":
			if err := dec.Decode(&s.Memory); err != nil {
				return err
			}

		case "memory_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MemoryInBytes = value
			case float64:
				f := int64(v)
				s.MemoryInBytes = f
			}

		case "norms_memory":
			if err := dec.Decode(&s.NormsMemory); err != nil {
				return err
			}

		case "norms_memory_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NormsMemoryInBytes = value
			case float64:
				f := int64(v)
				s.NormsMemoryInBytes = f
			}

		case "points_memory":
			if err := dec.Decode(&s.PointsMemory); err != nil {
				return err
			}

		case "points_memory_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PointsMemoryInBytes = value
			case float64:
				f := int64(v)
				s.PointsMemoryInBytes = f
			}

		case "stored_fields_memory_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.StoredFieldsMemoryInBytes = value
			case float64:
				f := int64(v)
				s.StoredFieldsMemoryInBytes = f
			}

		case "stored_memory":
			if err := dec.Decode(&s.StoredMemory); err != nil {
				return err
			}

		case "term_vectors_memory_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TermVectorsMemoryInBytes = value
			case float64:
				f := int64(v)
				s.TermVectorsMemoryInBytes = f
			}

		case "term_vectory_memory":
			if err := dec.Decode(&s.TermVectoryMemory); err != nil {
				return err
			}

		case "terms_memory":
			if err := dec.Decode(&s.TermsMemory); err != nil {
				return err
			}

		case "terms_memory_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TermsMemoryInBytes = value
			case float64:
				f := int64(v)
				s.TermsMemoryInBytes = f
			}

		case "version_map_memory":
			if err := dec.Decode(&s.VersionMapMemory); err != nil {
				return err
			}

		case "version_map_memory_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.VersionMapMemoryInBytes = value
			case float64:
				f := int64(v)
				s.VersionMapMemoryInBytes = f
			}

		}
	}
	return nil
}

// NewSegmentsStats returns a SegmentsStats.
func NewSegmentsStats() *SegmentsStats {
	r := &SegmentsStats{
		FileSizes: make(map[string]ShardFileSizeInfo, 0),
	}

	return r
}
