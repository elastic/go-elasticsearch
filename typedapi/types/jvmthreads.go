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
// https://github.com/elastic/elasticsearch-specification/tree/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// JvmThreads type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757/specification/nodes/_types/Stats.ts#L897-L906
type JvmThreads struct {
	// Count Number of active threads in use by JVM.
	Count *int64 `json:"count,omitempty"`
	// PeakCount Highest number of threads used by JVM.
	PeakCount *int64 `json:"peak_count,omitempty"`
}

func (s *JvmThreads) UnmarshalJSON(data []byte) error {

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
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = &value
			case float64:
				f := int64(v)
				s.Count = &f
			}

		case "peak_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PeakCount", err)
				}
				s.PeakCount = &value
			case float64:
				f := int64(v)
				s.PeakCount = &f
			}

		}
	}
	return nil
}

// NewJvmThreads returns a JvmThreads.
func NewJvmThreads() *JvmThreads {
	r := &JvmThreads{}

	return r
}
