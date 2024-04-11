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

// Checkpointing type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/transform/get_transform_stats/types.ts#L85-L92
type Checkpointing struct {
	ChangesLastDetectedAt         *int64           `json:"changes_last_detected_at,omitempty"`
	ChangesLastDetectedAtDateTime DateTime         `json:"changes_last_detected_at_date_time,omitempty"`
	Last                          CheckpointStats  `json:"last"`
	LastSearchTime                *int64           `json:"last_search_time,omitempty"`
	Next                          *CheckpointStats `json:"next,omitempty"`
	OperationsBehind              *int64           `json:"operations_behind,omitempty"`
}

func (s *Checkpointing) UnmarshalJSON(data []byte) error {

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

		case "changes_last_detected_at":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ChangesLastDetectedAt", err)
				}
				s.ChangesLastDetectedAt = &value
			case float64:
				f := int64(v)
				s.ChangesLastDetectedAt = &f
			}

		case "changes_last_detected_at_date_time":
			if err := dec.Decode(&s.ChangesLastDetectedAtDateTime); err != nil {
				return fmt.Errorf("%s | %w", "ChangesLastDetectedAtDateTime", err)
			}

		case "last":
			if err := dec.Decode(&s.Last); err != nil {
				return fmt.Errorf("%s | %w", "Last", err)
			}

		case "last_search_time":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "LastSearchTime", err)
				}
				s.LastSearchTime = &value
			case float64:
				f := int64(v)
				s.LastSearchTime = &f
			}

		case "next":
			if err := dec.Decode(&s.Next); err != nil {
				return fmt.Errorf("%s | %w", "Next", err)
			}

		case "operations_behind":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "OperationsBehind", err)
				}
				s.OperationsBehind = &value
			case float64:
				f := int64(v)
				s.OperationsBehind = &f
			}

		}
	}
	return nil
}

// NewCheckpointing returns a Checkpointing.
func NewCheckpointing() *Checkpointing {
	r := &Checkpointing{}

	return r
}
