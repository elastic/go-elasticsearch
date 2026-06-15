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
// https://github.com/elastic/elasticsearch-specification/tree/eb2e22fb2ac404e676d19bcc7bb089647f029026

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Information about a single reindex task, as returned by the reindex
// management APIs.
//
// https://github.com/elastic/elasticsearch-specification/blob/eb2e22fb2ac404e676d19bcc7bb089647f029026/specification/_types/Reindex.ts#L153-L191
type ReindexTaskInfo struct {
	// Cancelled Whether the reindex task has been cancelled.
	Cancelled bool `json:"cancelled"`
	// Description A sanitized description of the reindex operation (source and destination
	// indices, and optionally remote host info).
	Description *string `json:"description,omitempty"`
	// Id The ID of the reindex task, in `nodeId:taskNum` format.
	Id string `json:"id"`
	// RunningTime The elapsed running time of the reindex task, in a human-readable format.
	// Only present when the request includes the `?human=true` query parameter.
	RunningTime Duration `json:"running_time,omitempty"`
	// RunningTimeInNanos The elapsed running time of the reindex task, in nanoseconds.
	RunningTimeInNanos int64 `json:"running_time_in_nanos"`
	// StartTime The time at which the reindex task started, as an ISO 8601 formatted string.
	// Only present when the request includes the `?human=true` query parameter.
	StartTime *string `json:"start_time,omitempty"`
	// StartTimeInMillis The time at which the reindex task started, in milliseconds since the Unix
	// epoch.
	StartTimeInMillis int64 `json:"start_time_in_millis"`
	// Status The current progress of the reindex operation.
	Status *ReindexStatus `json:"status,omitempty"`
}

func (s *ReindexTaskInfo) UnmarshalJSON(data []byte) error {

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

		case "cancelled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Cancelled", err)
				}
				s.Cancelled = value
			case bool:
				s.Cancelled = v
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "running_time":
			if err := dec.Decode(&s.RunningTime); err != nil {
				return fmt.Errorf("%s | %w", "RunningTime", err)
			}

		case "running_time_in_nanos":
			if err := dec.Decode(&s.RunningTimeInNanos); err != nil {
				return fmt.Errorf("%s | %w", "RunningTimeInNanos", err)
			}

		case "start_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "StartTime", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.StartTime = &o

		case "start_time_in_millis":
			if err := dec.Decode(&s.StartTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "StartTimeInMillis", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		}
	}
	return nil
}

// NewReindexTaskInfo returns a ReindexTaskInfo.
func NewReindexTaskInfo() *ReindexTaskInfo {
	r := &ReindexTaskInfo{}

	return r
}
