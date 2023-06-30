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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

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
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/Stats.ts#L206-L231
type SegmentsStats struct {
	Count                       int                          `json:"count"`
	DocValuesMemory             ByteSize                     `json:"doc_values_memory,omitempty"`
	DocValuesMemoryInBytes      int64                        `json:"doc_values_memory_in_bytes"`
	FileSizes                   map[string]ShardFileSizeInfo `json:"file_sizes"`
	FixedBitSet                 ByteSize                     `json:"fixed_bit_set,omitempty"`
	FixedBitSetMemoryInBytes    int64                        `json:"fixed_bit_set_memory_in_bytes"`
	IndexWriterMaxMemoryInBytes *int64                       `json:"index_writer_max_memory_in_bytes,omitempty"`
	IndexWriterMemory           ByteSize                     `json:"index_writer_memory,omitempty"`
	IndexWriterMemoryInBytes    int64                        `json:"index_writer_memory_in_bytes"`
	MaxUnsafeAutoIdTimestamp    int64                        `json:"max_unsafe_auto_id_timestamp"`
	Memory                      ByteSize                     `json:"memory,omitempty"`
	MemoryInBytes               int64                        `json:"memory_in_bytes"`
	NormsMemory                 ByteSize                     `json:"norms_memory,omitempty"`
	NormsMemoryInBytes          int64                        `json:"norms_memory_in_bytes"`
	PointsMemory                ByteSize                     `json:"points_memory,omitempty"`
	PointsMemoryInBytes         int64                        `json:"points_memory_in_bytes"`
	StoredFieldsMemoryInBytes   int64                        `json:"stored_fields_memory_in_bytes"`
	StoredMemory                ByteSize                     `json:"stored_memory,omitempty"`
	TermVectorsMemoryInBytes    int64                        `json:"term_vectors_memory_in_bytes"`
	TermVectoryMemory           ByteSize                     `json:"term_vectory_memory,omitempty"`
	TermsMemory                 ByteSize                     `json:"terms_memory,omitempty"`
	TermsMemoryInBytes          int64                        `json:"terms_memory_in_bytes"`
	VersionMapMemory            ByteSize                     `json:"version_map_memory,omitempty"`
	VersionMapMemoryInBytes     int64                        `json:"version_map_memory_in_bytes"`
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
