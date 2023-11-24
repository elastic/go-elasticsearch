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
)

// SnapshotStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/snapshot/_types/SnapshotStats.ts#L23-L29
type SnapshotStats struct {
	Incremental       FileCountSnapshotStats `json:"incremental"`
	StartTimeInMillis int64                  `json:"start_time_in_millis"`
	Time              Duration               `json:"time,omitempty"`
	TimeInMillis      int64                  `json:"time_in_millis"`
	Total             FileCountSnapshotStats `json:"total"`
}

func (s *SnapshotStats) UnmarshalJSON(data []byte) error {

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

		case "incremental":
			if err := dec.Decode(&s.Incremental); err != nil {
				return err
			}

		case "start_time_in_millis":
			if err := dec.Decode(&s.StartTimeInMillis); err != nil {
				return err
			}

		case "time":
			if err := dec.Decode(&s.Time); err != nil {
				return err
			}

		case "time_in_millis":
			if err := dec.Decode(&s.TimeInMillis); err != nil {
				return err
			}

		case "total":
			if err := dec.Decode(&s.Total); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSnapshotStats returns a SnapshotStats.
func NewSnapshotStats() *SnapshotStats {
	r := &SnapshotStats{}

	return r
}
