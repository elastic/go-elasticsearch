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
// https://github.com/elastic/elasticsearch-specification/tree/9665eef0d78c41f20c4c83e69b7c8155efd58f24

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// MappingStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9665eef0d78c41f20c4c83e69b7c8155efd58f24/specification/indices/stats/types.ts#L188-L195
type MappingStats struct {
	AverageFieldsPerSegment       int64    `json:"average_fields_per_segment"`
	TotalCount                    int64    `json:"total_count"`
	TotalEstimatedOverhead        ByteSize `json:"total_estimated_overhead,omitempty"`
	TotalEstimatedOverheadInBytes int64    `json:"total_estimated_overhead_in_bytes"`
	TotalSegmentFields            int64    `json:"total_segment_fields"`
	TotalSegments                 int64    `json:"total_segments"`
}

func (s *MappingStats) UnmarshalJSON(data []byte) error {

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

		case "average_fields_per_segment":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "AverageFieldsPerSegment", err)
				}
				s.AverageFieldsPerSegment = value
			case float64:
				f := int64(v)
				s.AverageFieldsPerSegment = f
			}

		case "total_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalCount", err)
				}
				s.TotalCount = value
			case float64:
				f := int64(v)
				s.TotalCount = f
			}

		case "total_estimated_overhead":
			if err := dec.Decode(&s.TotalEstimatedOverhead); err != nil {
				return fmt.Errorf("%s | %w", "TotalEstimatedOverhead", err)
			}

		case "total_estimated_overhead_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalEstimatedOverheadInBytes", err)
				}
				s.TotalEstimatedOverheadInBytes = value
			case float64:
				f := int64(v)
				s.TotalEstimatedOverheadInBytes = f
			}

		case "total_segment_fields":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalSegmentFields", err)
				}
				s.TotalSegmentFields = value
			case float64:
				f := int64(v)
				s.TotalSegmentFields = f
			}

		case "total_segments":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalSegments", err)
				}
				s.TotalSegments = value
			case float64:
				f := int64(v)
				s.TotalSegments = f
			}

		}
	}
	return nil
}

// NewMappingStats returns a MappingStats.
func NewMappingStats() *MappingStats {
	r := &MappingStats{}

	return r
}
