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

// ClusterJvmMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cluster/stats/types.ts#L294-L303
type ClusterJvmMemory struct {
	// HeapMaxInBytes Maximum amount of memory, in bytes, available for use by the heap across all
	// selected nodes.
	HeapMaxInBytes int64 `json:"heap_max_in_bytes"`
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

		case "heap_max_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.HeapMaxInBytes = value
			case float64:
				f := int64(v)
				s.HeapMaxInBytes = f
			}

		case "heap_used_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
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
