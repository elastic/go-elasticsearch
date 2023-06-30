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

// Checkpointing type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/transform/get_transform_stats/types.ts#L82-L89
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
					return err
				}
				s.ChangesLastDetectedAt = &value
			case float64:
				f := int64(v)
				s.ChangesLastDetectedAt = &f
			}

		case "changes_last_detected_at_date_time":
			if err := dec.Decode(&s.ChangesLastDetectedAtDateTime); err != nil {
				return err
			}

		case "last":
			if err := dec.Decode(&s.Last); err != nil {
				return err
			}

		case "last_search_time":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LastSearchTime = &value
			case float64:
				f := int64(v)
				s.LastSearchTime = &f
			}

		case "next":
			if err := dec.Decode(&s.Next); err != nil {
				return err
			}

		case "operations_behind":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
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
