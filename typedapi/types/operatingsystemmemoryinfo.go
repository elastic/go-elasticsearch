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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// OperatingSystemMemoryInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cluster/stats/types.ts#L541-L568
type OperatingSystemMemoryInfo struct {
	// AdjustedTotalInBytes Total amount, in bytes, of memory across all selected nodes, but using the
	// value specified using the `es.total_memory_bytes` system property instead of
	// measured total memory for those nodes where that system property was set.
	AdjustedTotalInBytes *int64 `json:"adjusted_total_in_bytes,omitempty"`
	// FreeInBytes Amount, in bytes, of free physical memory across all selected nodes.
	FreeInBytes int64 `json:"free_in_bytes"`
	// FreePercent Percentage of free physical memory across all selected nodes.
	FreePercent int `json:"free_percent"`
	// TotalInBytes Total amount, in bytes, of physical memory across all selected nodes.
	TotalInBytes int64 `json:"total_in_bytes"`
	// UsedInBytes Amount, in bytes, of physical memory in use across all selected nodes.
	UsedInBytes int64 `json:"used_in_bytes"`
	// UsedPercent Percentage of physical memory in use across all selected nodes.
	UsedPercent int `json:"used_percent"`
}

func (s *OperatingSystemMemoryInfo) UnmarshalJSON(data []byte) error {

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

		case "adjusted_total_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.AdjustedTotalInBytes = &value
			case float64:
				f := int64(v)
				s.AdjustedTotalInBytes = &f
			}

		case "free_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.FreeInBytes = value
			case float64:
				f := int64(v)
				s.FreeInBytes = f
			}

		case "free_percent":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.FreePercent = value
			case float64:
				f := int(v)
				s.FreePercent = f
			}

		case "total_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalInBytes = value
			case float64:
				f := int64(v)
				s.TotalInBytes = f
			}

		case "used_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.UsedInBytes = value
			case float64:
				f := int64(v)
				s.UsedInBytes = f
			}

		case "used_percent":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.UsedPercent = value
			case float64:
				f := int(v)
				s.UsedPercent = f
			}

		}
	}
	return nil
}

// NewOperatingSystemMemoryInfo returns a OperatingSystemMemoryInfo.
func NewOperatingSystemMemoryInfo() *OperatingSystemMemoryInfo {
	r := &OperatingSystemMemoryInfo{}

	return r
}
