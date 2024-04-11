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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

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
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/indices/stats/types.ts#L186-L190
type MappingStats struct {
	TotalCount                    int64    `json:"total_count"`
	TotalEstimatedOverhead        ByteSize `json:"total_estimated_overhead,omitempty"`
	TotalEstimatedOverheadInBytes int64    `json:"total_estimated_overhead_in_bytes"`
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

		case "total_count":
			var tmp interface{}
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
			var tmp interface{}
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

		}
	}
	return nil
}

// NewMappingStats returns a MappingStats.
func NewMappingStats() *MappingStats {
	r := &MappingStats{}

	return r
}
