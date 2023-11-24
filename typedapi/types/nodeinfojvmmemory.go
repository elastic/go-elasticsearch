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

// NodeInfoJvmMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/nodes/info/types.ts#L305-L316
type NodeInfoJvmMemory struct {
	DirectMax          ByteSize `json:"direct_max,omitempty"`
	DirectMaxInBytes   int64    `json:"direct_max_in_bytes"`
	HeapInit           ByteSize `json:"heap_init,omitempty"`
	HeapInitInBytes    int64    `json:"heap_init_in_bytes"`
	HeapMax            ByteSize `json:"heap_max,omitempty"`
	HeapMaxInBytes     int64    `json:"heap_max_in_bytes"`
	NonHeapInit        ByteSize `json:"non_heap_init,omitempty"`
	NonHeapInitInBytes int64    `json:"non_heap_init_in_bytes"`
	NonHeapMax         ByteSize `json:"non_heap_max,omitempty"`
	NonHeapMaxInBytes  int64    `json:"non_heap_max_in_bytes"`
}

func (s *NodeInfoJvmMemory) UnmarshalJSON(data []byte) error {

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

		case "direct_max":
			if err := dec.Decode(&s.DirectMax); err != nil {
				return err
			}

		case "direct_max_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DirectMaxInBytes = value
			case float64:
				f := int64(v)
				s.DirectMaxInBytes = f
			}

		case "heap_init":
			if err := dec.Decode(&s.HeapInit); err != nil {
				return err
			}

		case "heap_init_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.HeapInitInBytes = value
			case float64:
				f := int64(v)
				s.HeapInitInBytes = f
			}

		case "heap_max":
			if err := dec.Decode(&s.HeapMax); err != nil {
				return err
			}

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

		case "non_heap_init":
			if err := dec.Decode(&s.NonHeapInit); err != nil {
				return err
			}

		case "non_heap_init_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NonHeapInitInBytes = value
			case float64:
				f := int64(v)
				s.NonHeapInitInBytes = f
			}

		case "non_heap_max":
			if err := dec.Decode(&s.NonHeapMax); err != nil {
				return err
			}

		case "non_heap_max_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NonHeapMaxInBytes = value
			case float64:
				f := int64(v)
				s.NonHeapMaxInBytes = f
			}

		}
	}
	return nil
}

// NewNodeInfoJvmMemory returns a NodeInfoJvmMemory.
func NewNodeInfoJvmMemory() *NodeInfoJvmMemory {
	r := &NodeInfoJvmMemory{}

	return r
}
