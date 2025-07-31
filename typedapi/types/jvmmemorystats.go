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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// JvmMemoryStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/nodes/_types/Stats.ts#L920-L954
type JvmMemoryStats struct {
	// HeapCommittedInBytes Amount of memory, in bytes, available for use by the heap.
	HeapCommittedInBytes *int64 `json:"heap_committed_in_bytes,omitempty"`
	// HeapMax Maximum amount of memory, available for use by the heap.
	HeapMax ByteSize `json:"heap_max,omitempty"`
	// HeapMaxInBytes Maximum amount of memory, in bytes, available for use by the heap.
	HeapMaxInBytes *int64 `json:"heap_max_in_bytes,omitempty"`
	// HeapUsedInBytes Memory, in bytes, currently in use by the heap.
	HeapUsedInBytes *int64 `json:"heap_used_in_bytes,omitempty"`
	// HeapUsedPercent Percentage of memory currently in use by the heap.
	HeapUsedPercent *int64 `json:"heap_used_percent,omitempty"`
	// NonHeapCommittedInBytes Amount of non-heap memory available, in bytes.
	NonHeapCommittedInBytes *int64 `json:"non_heap_committed_in_bytes,omitempty"`
	// NonHeapUsedInBytes Non-heap memory used, in bytes.
	NonHeapUsedInBytes *int64 `json:"non_heap_used_in_bytes,omitempty"`
	// Pools Contains statistics about heap memory usage for the node.
	Pools map[string]Pool `json:"pools,omitempty"`
}

func (s *JvmMemoryStats) UnmarshalJSON(data []byte) error {

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

		case "heap_committed_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "HeapCommittedInBytes", err)
				}
				s.HeapCommittedInBytes = &value
			case float64:
				f := int64(v)
				s.HeapCommittedInBytes = &f
			}

		case "heap_max":
			if err := dec.Decode(&s.HeapMax); err != nil {
				return fmt.Errorf("%s | %w", "HeapMax", err)
			}

		case "heap_max_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "HeapMaxInBytes", err)
				}
				s.HeapMaxInBytes = &value
			case float64:
				f := int64(v)
				s.HeapMaxInBytes = &f
			}

		case "heap_used_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "HeapUsedInBytes", err)
				}
				s.HeapUsedInBytes = &value
			case float64:
				f := int64(v)
				s.HeapUsedInBytes = &f
			}

		case "heap_used_percent":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "HeapUsedPercent", err)
				}
				s.HeapUsedPercent = &value
			case float64:
				f := int64(v)
				s.HeapUsedPercent = &f
			}

		case "non_heap_committed_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NonHeapCommittedInBytes", err)
				}
				s.NonHeapCommittedInBytes = &value
			case float64:
				f := int64(v)
				s.NonHeapCommittedInBytes = &f
			}

		case "non_heap_used_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NonHeapUsedInBytes", err)
				}
				s.NonHeapUsedInBytes = &value
			case float64:
				f := int64(v)
				s.NonHeapUsedInBytes = &f
			}

		case "pools":
			if s.Pools == nil {
				s.Pools = make(map[string]Pool, 0)
			}
			if err := dec.Decode(&s.Pools); err != nil {
				return fmt.Errorf("%s | %w", "Pools", err)
			}

		}
	}
	return nil
}

// NewJvmMemoryStats returns a JvmMemoryStats.
func NewJvmMemoryStats() *JvmMemoryStats {
	r := &JvmMemoryStats{
		Pools: make(map[string]Pool),
	}

	return r
}
