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

// ClusterJvmMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/stats/types.ts#L395-L412
type ClusterJvmMemory struct {
	// HeapMax Maximum amount of memory available for use by the heap across all selected
	// nodes.
	HeapMax ByteSize `json:"heap_max,omitempty"`
	// HeapMaxInBytes Maximum amount of memory, in bytes, available for use by the heap across all
	// selected nodes.
	HeapMaxInBytes int64 `json:"heap_max_in_bytes"`
	// HeapUsed Memory currently in use by the heap across all selected nodes.
	HeapUsed ByteSize `json:"heap_used,omitempty"`
	// HeapUsedInBytes Memory, in bytes, currently in use by the heap across all selected nodes.
	HeapUsedInBytes int64 `json:"heap_used_in_bytes"`
}

func (s *ClusterJvmMemory) UnmarshalJSON(data []byte) error {

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
				s.HeapMaxInBytes = value
			case float64:
				f := int64(v)
				s.HeapMaxInBytes = f
			}

		case "heap_used":
			if err := dec.Decode(&s.HeapUsed); err != nil {
				return fmt.Errorf("%s | %w", "HeapUsed", err)
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
				s.HeapUsedInBytes = value
			case float64:
				f := int64(v)
				s.HeapUsedInBytes = f
			}

		}
	}
	return nil
}

// NewClusterJvmMemory returns a ClusterJvmMemory.
func NewClusterJvmMemory() *ClusterJvmMemory {
	r := &ClusterJvmMemory{}

	return r
}
