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

// FielddataStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/Stats.ts#L111-L116
type FielddataStats struct {
	Evictions         *int64                      `json:"evictions,omitempty"`
	Fields            map[string]FieldMemoryUsage `json:"fields,omitempty"`
	MemorySize        ByteSize                    `json:"memory_size,omitempty"`
	MemorySizeInBytes int64                       `json:"memory_size_in_bytes"`
}

func (s *FielddataStats) UnmarshalJSON(data []byte) error {

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

		case "evictions":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Evictions = &value
			case float64:
				f := int64(v)
				s.Evictions = &f
			}

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string]FieldMemoryUsage, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		case "memory_size":
			if err := dec.Decode(&s.MemorySize); err != nil {
				return err
			}

		case "memory_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MemorySizeInBytes = value
			case float64:
				f := int64(v)
				s.MemorySizeInBytes = f
			}

		}
	}
	return nil
}

// NewFielddataStats returns a FielddataStats.
func NewFielddataStats() *FielddataStats {
	r := &FielddataStats{
		Fields: make(map[string]FieldMemoryUsage, 0),
	}

	return r
}
