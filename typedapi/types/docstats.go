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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DocStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/Stats.ts#L100-L121
type DocStats struct {
	// Count Total number of non-deleted documents across all primary shards assigned to
	// selected nodes.
	// This number is based on documents in Lucene segments and may include
	// documents from nested fields.
	Count int64 `json:"count"`
	// Deleted Total number of deleted documents across all primary shards assigned to
	// selected nodes.
	// This number is based on documents in Lucene segments.
	// Elasticsearch reclaims the disk space of deleted Lucene documents when a
	// segment is merged.
	Deleted *int64 `json:"deleted,omitempty"`
	// TotalSize Human readable total_size_in_bytes
	TotalSize ByteSize `json:"total_size,omitempty"`
	// TotalSizeInBytes Returns the total size in bytes of all documents in this stats.
	// This value may be more reliable than store_stats.size_in_bytes in estimating
	// the index size.
	TotalSizeInBytes int64 `json:"total_size_in_bytes"`
}

func (s *DocStats) UnmarshalJSON(data []byte) error {

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
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int64(v)
				s.Count = f
			}

		case "deleted":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Deleted", err)
				}
				s.Deleted = &value
			case float64:
				f := int64(v)
				s.Deleted = &f
			}

		case "total_size":
			if err := dec.Decode(&s.TotalSize); err != nil {
				return fmt.Errorf("%s | %w", "TotalSize", err)
			}

		case "total_size_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalSizeInBytes", err)
				}
				s.TotalSizeInBytes = value
			case float64:
				f := int64(v)
				s.TotalSizeInBytes = f
			}

		}
	}
	return nil
}

// NewDocStats returns a DocStats.
func NewDocStats() *DocStats {
	r := &DocStats{}

	return r
}
