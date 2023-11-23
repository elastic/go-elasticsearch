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

// NodeBufferPool type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/nodes/_types/Stats.ts#L788-L809
type NodeBufferPool struct {
	// Count Number of buffer pools.
	Count *int64 `json:"count,omitempty"`
	// TotalCapacity Total capacity of buffer pools.
	TotalCapacity *string `json:"total_capacity,omitempty"`
	// TotalCapacityInBytes Total capacity of buffer pools in bytes.
	TotalCapacityInBytes *int64 `json:"total_capacity_in_bytes,omitempty"`
	// Used Size of buffer pools.
	Used *string `json:"used,omitempty"`
	// UsedInBytes Size of buffer pools in bytes.
	UsedInBytes *int64 `json:"used_in_bytes,omitempty"`
}

func (s *NodeBufferPool) UnmarshalJSON(data []byte) error {

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
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Count = &value
			case float64:
				f := int64(v)
				s.Count = &f
			}

		case "total_capacity":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TotalCapacity = &o

		case "total_capacity_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalCapacityInBytes = &value
			case float64:
				f := int64(v)
				s.TotalCapacityInBytes = &f
			}

		case "used":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Used = &o

		case "used_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.UsedInBytes = &value
			case float64:
				f := int64(v)
				s.UsedInBytes = &f
			}

		}
	}
	return nil
}

// NewNodeBufferPool returns a NodeBufferPool.
func NewNodeBufferPool() *NodeBufferPool {
	r := &NodeBufferPool{}

	return r
}
