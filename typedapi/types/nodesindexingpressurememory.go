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

// NodesIndexingPressureMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/nodes/_types/Stats.ts#L59-L64
type NodesIndexingPressureMemory struct {
	Current      *PressureMemory `json:"current,omitempty"`
	Limit        ByteSize        `json:"limit,omitempty"`
	LimitInBytes *int64          `json:"limit_in_bytes,omitempty"`
	Total        *PressureMemory `json:"total,omitempty"`
}

func (s *NodesIndexingPressureMemory) UnmarshalJSON(data []byte) error {

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

		case "current":
			if err := dec.Decode(&s.Current); err != nil {
				return err
			}

		case "limit":
			if err := dec.Decode(&s.Limit); err != nil {
				return err
			}

		case "limit_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LimitInBytes = &value
			case float64:
				f := int64(v)
				s.LimitInBytes = &f
			}

		case "total":
			if err := dec.Decode(&s.Total); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewNodesIndexingPressureMemory returns a NodesIndexingPressureMemory.
func NewNodesIndexingPressureMemory() *NodesIndexingPressureMemory {
	r := &NodesIndexingPressureMemory{}

	return r
}
