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

// CheckpointStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/transform/get_transform_stats/types.ts#L76-L83
type CheckpointStats struct {
	Checkpoint           int64              `json:"checkpoint"`
	CheckpointProgress   *TransformProgress `json:"checkpoint_progress,omitempty"`
	TimeUpperBound       DateTime           `json:"time_upper_bound,omitempty"`
	TimeUpperBoundMillis *int64             `json:"time_upper_bound_millis,omitempty"`
	Timestamp            DateTime           `json:"timestamp,omitempty"`
	TimestampMillis      *int64             `json:"timestamp_millis,omitempty"`
}

func (s *CheckpointStats) UnmarshalJSON(data []byte) error {

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

		case "checkpoint":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Checkpoint = value
			case float64:
				f := int64(v)
				s.Checkpoint = f
			}

		case "checkpoint_progress":
			if err := dec.Decode(&s.CheckpointProgress); err != nil {
				return err
			}

		case "time_upper_bound":
			if err := dec.Decode(&s.TimeUpperBound); err != nil {
				return err
			}

		case "time_upper_bound_millis":
			if err := dec.Decode(&s.TimeUpperBoundMillis); err != nil {
				return err
			}

		case "timestamp":
			if err := dec.Decode(&s.Timestamp); err != nil {
				return err
			}

		case "timestamp_millis":
			if err := dec.Decode(&s.TimestampMillis); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewCheckpointStats returns a CheckpointStats.
func NewCheckpointStats() *CheckpointStats {
	r := &CheckpointStats{}

	return r
}
