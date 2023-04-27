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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// Pool type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/nodes/_types/Stats.ts#L345-L350
type Pool struct {
	MaxInBytes      *int64 `json:"max_in_bytes,omitempty"`
	PeakMaxInBytes  *int64 `json:"peak_max_in_bytes,omitempty"`
	PeakUsedInBytes *int64 `json:"peak_used_in_bytes,omitempty"`
	UsedInBytes     *int64 `json:"used_in_bytes,omitempty"`
}

func (s *Pool) UnmarshalJSON(data []byte) error {

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

		case "max_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MaxInBytes = &value
			case float64:
				f := int64(v)
				s.MaxInBytes = &f
			}

		case "peak_max_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PeakMaxInBytes = &value
			case float64:
				f := int64(v)
				s.PeakMaxInBytes = &f
			}

		case "peak_used_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PeakUsedInBytes = &value
			case float64:
				f := int64(v)
				s.PeakUsedInBytes = &f
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
				s.UsedInBytes = &value
			case float64:
				f := int64(v)
				s.UsedInBytes = &f
			}

		}
	}
	return nil
}

// NewPool returns a Pool.
func NewPool() *Pool {
	r := &Pool{}

	return r
}
