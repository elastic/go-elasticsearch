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
// https://github.com/elastic/elasticsearch-specification/tree/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Pool type.
//
// https://github.com/elastic/elasticsearch-specification/blob/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c/specification/nodes/_types/Stats.ts#L1166-L1199
type Pool struct {
	// Max Maximum amount of memory available for use by the heap.
	Max *string `json:"max,omitempty"`
	// MaxInBytes Maximum amount of memory, in bytes, available for use by the heap.
	MaxInBytes *int64 `json:"max_in_bytes,omitempty"`
	// PeakMax Largest amount of memory historically used by the heap.
	PeakMax *string `json:"peak_max,omitempty"`
	// PeakMaxInBytes Largest amount of memory, in bytes, historically used by the heap.
	PeakMaxInBytes *int64 `json:"peak_max_in_bytes,omitempty"`
	// PeakUsed Largest amount of memory historically used by the heap.
	PeakUsed *string `json:"peak_used,omitempty"`
	// PeakUsedInBytes Largest amount of memory, in bytes, historically used by the heap.
	PeakUsedInBytes *int64 `json:"peak_used_in_bytes,omitempty"`
	// Used Memory used by the heap.
	Used *string `json:"used,omitempty"`
	// UsedInBytes Memory, in bytes, used by the heap.
	UsedInBytes *int64 `json:"used_in_bytes,omitempty"`
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

		case "max":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Max", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Max = &o

		case "max_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxInBytes", err)
				}
				s.MaxInBytes = &value
			case float64:
				f := int64(v)
				s.MaxInBytes = &f
			}

		case "peak_max":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "PeakMax", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PeakMax = &o

		case "peak_max_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PeakMaxInBytes", err)
				}
				s.PeakMaxInBytes = &value
			case float64:
				f := int64(v)
				s.PeakMaxInBytes = &f
			}

		case "peak_used":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "PeakUsed", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PeakUsed = &o

		case "peak_used_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PeakUsedInBytes", err)
				}
				s.PeakUsedInBytes = &value
			case float64:
				f := int64(v)
				s.PeakUsedInBytes = &f
			}

		case "used":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Used", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Used = &o

		case "used_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "UsedInBytes", err)
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
