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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CircuitBreakerRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6785a6caa1fa3ca5ab3308963d79dce923a3469f/specification/cat/circuit_breaker/types.ts#L22-L68
type CircuitBreakerRecord struct {
	// Breaker Breaker name
	Breaker *string `json:"breaker,omitempty"`
	// Estimated Estimated size
	Estimated *string `json:"estimated,omitempty"`
	// EstimatedBytes Estimated size in bytes
	EstimatedBytes ByteSize `json:"estimated_bytes,omitempty"`
	// Limit Limit size
	Limit *string `json:"limit,omitempty"`
	// LimitBytes Limit size in bytes
	LimitBytes ByteSize `json:"limit_bytes,omitempty"`
	// NodeId Persistent node ID
	NodeId *string `json:"node_id,omitempty"`
	// NodeName Node name
	NodeName *string `json:"node_name,omitempty"`
	// Overhead Overhead
	Overhead *string `json:"overhead,omitempty"`
	// Tripped Tripped count
	Tripped *string `json:"tripped,omitempty"`
}

func (s *CircuitBreakerRecord) UnmarshalJSON(data []byte) error {

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

		case "breaker", "br":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Breaker", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Breaker = &o

		case "estimated", "e":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Estimated", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Estimated = &o

		case "estimated_bytes", "eb":
			if err := dec.Decode(&s.EstimatedBytes); err != nil {
				return fmt.Errorf("%s | %w", "EstimatedBytes", err)
			}

		case "limit", "l":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Limit", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Limit = &o

		case "limit_bytes", "lb":
			if err := dec.Decode(&s.LimitBytes); err != nil {
				return fmt.Errorf("%s | %w", "LimitBytes", err)
			}

		case "node_id", "id":
			if err := dec.Decode(&s.NodeId); err != nil {
				return fmt.Errorf("%s | %w", "NodeId", err)
			}

		case "node_name", "nn":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "NodeName", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NodeName = &o

		case "overhead", "o":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Overhead", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Overhead = &o

		case "tripped", "t":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Tripped", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Tripped = &o

		}
	}
	return nil
}

// NewCircuitBreakerRecord returns a CircuitBreakerRecord.
func NewCircuitBreakerRecord() *CircuitBreakerRecord {
	r := &CircuitBreakerRecord{}

	return r
}
