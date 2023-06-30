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

// MemStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/get_memory_stats/types.ts#L65-L88
type MemStats struct {
	// AdjustedTotal If the amount of physical memory has been overridden using the
	// es.total_memory_bytes system property
	// then this reports the overridden value. Otherwise it reports the same value
	// as total.
	AdjustedTotal ByteSize `json:"adjusted_total,omitempty"`
	// AdjustedTotalInBytes If the amount of physical memory has been overridden using the
	// `es.total_memory_bytes` system property
	// then this reports the overridden value in bytes. Otherwise it reports the
	// same value as `total_in_bytes`.
	AdjustedTotalInBytes int `json:"adjusted_total_in_bytes"`
	// Ml Contains statistics about machine learning use of native memory on the node.
	Ml MemMlStats `json:"ml"`
	// Total Total amount of physical memory.
	Total ByteSize `json:"total,omitempty"`
	// TotalInBytes Total amount of physical memory in bytes.
	TotalInBytes int `json:"total_in_bytes"`
}

func (s *MemStats) UnmarshalJSON(data []byte) error {

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

		case "adjusted_total":
			if err := dec.Decode(&s.AdjustedTotal); err != nil {
				return err
			}

		case "adjusted_total_in_bytes":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.AdjustedTotalInBytes = value
			case float64:
				f := int(v)
				s.AdjustedTotalInBytes = f
			}

		case "ml":
			if err := dec.Decode(&s.Ml); err != nil {
				return err
			}

		case "total":
			if err := dec.Decode(&s.Total); err != nil {
				return err
			}

		case "total_in_bytes":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TotalInBytes = value
			case float64:
				f := int(v)
				s.TotalInBytes = f
			}

		}
	}
	return nil
}

// NewMemStats returns a MemStats.
func NewMemStats() *MemStats {
	r := &MemStats{}

	return r
}
