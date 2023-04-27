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

// Breaker type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/nodes/_types/Stats.ts#L179-L186
type Breaker struct {
	EstimatedSize        *string  `json:"estimated_size,omitempty"`
	EstimatedSizeInBytes *int64   `json:"estimated_size_in_bytes,omitempty"`
	LimitSize            *string  `json:"limit_size,omitempty"`
	LimitSizeInBytes     *int64   `json:"limit_size_in_bytes,omitempty"`
	Overhead             *float32 `json:"overhead,omitempty"`
	Tripped              *float32 `json:"tripped,omitempty"`
}

func (s *Breaker) UnmarshalJSON(data []byte) error {

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

		case "estimated_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.EstimatedSize = &o

		case "estimated_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.EstimatedSizeInBytes = &value
			case float64:
				f := int64(v)
				s.EstimatedSizeInBytes = &f
			}

		case "limit_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.LimitSize = &o

		case "limit_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LimitSizeInBytes = &value
			case float64:
				f := int64(v)
				s.LimitSizeInBytes = &f
			}

		case "overhead":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Overhead = &f
			case float64:
				f := float32(v)
				s.Overhead = &f
			}

		case "tripped":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Tripped = &f
			case float64:
				f := float32(v)
				s.Tripped = &f
			}

		}
	}
	return nil
}

// NewBreaker returns a Breaker.
func NewBreaker() *Breaker {
	r := &Breaker{}

	return r
}
