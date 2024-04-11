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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// TranslogStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/_types/Stats.ts#L397-L405
type TranslogStats struct {
	EarliestLastModifiedAge int64   `json:"earliest_last_modified_age"`
	Operations              int64   `json:"operations"`
	Size                    *string `json:"size,omitempty"`
	SizeInBytes             int64   `json:"size_in_bytes"`
	UncommittedOperations   int     `json:"uncommitted_operations"`
	UncommittedSize         *string `json:"uncommitted_size,omitempty"`
	UncommittedSizeInBytes  int64   `json:"uncommitted_size_in_bytes"`
}

func (s *TranslogStats) UnmarshalJSON(data []byte) error {

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

		case "earliest_last_modified_age":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "EarliestLastModifiedAge", err)
				}
				s.EarliestLastModifiedAge = value
			case float64:
				f := int64(v)
				s.EarliestLastModifiedAge = f
			}

		case "operations":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Operations", err)
				}
				s.Operations = value
			case float64:
				f := int64(v)
				s.Operations = f
			}

		case "size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Size", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Size = &o

		case "size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SizeInBytes", err)
				}
				s.SizeInBytes = value
			case float64:
				f := int64(v)
				s.SizeInBytes = f
			}

		case "uncommitted_operations":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "UncommittedOperations", err)
				}
				s.UncommittedOperations = value
			case float64:
				f := int(v)
				s.UncommittedOperations = f
			}

		case "uncommitted_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "UncommittedSize", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UncommittedSize = &o

		case "uncommitted_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "UncommittedSizeInBytes", err)
				}
				s.UncommittedSizeInBytes = value
			case float64:
				f := int64(v)
				s.UncommittedSizeInBytes = f
			}

		}
	}
	return nil
}

// NewTranslogStats returns a TranslogStats.
func NewTranslogStats() *TranslogStats {
	r := &TranslogStats{}

	return r
}
