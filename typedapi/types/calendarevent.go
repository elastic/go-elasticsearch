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

	"encoding/json"
)

// CalendarEvent type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/ml/_types/CalendarEvent.ts#L23-L33
type CalendarEvent struct {
	// CalendarId A string that uniquely identifies a calendar.
	CalendarId *string `json:"calendar_id,omitempty"`
	// Description A description of the scheduled event.
	Description string `json:"description"`
	// EndTime The timestamp for the end of the scheduled event in milliseconds since the
	// epoch or ISO 8601 format.
	EndTime DateTime `json:"end_time"`
	EventId *string  `json:"event_id,omitempty"`
	// StartTime The timestamp for the beginning of the scheduled event in milliseconds since
	// the epoch or ISO 8601 format.
	StartTime DateTime `json:"start_time"`
}

func (s *CalendarEvent) UnmarshalJSON(data []byte) error {

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

		case "calendar_id":
			if err := dec.Decode(&s.CalendarId); err != nil {
				return err
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Description = o

		case "end_time":
			if err := dec.Decode(&s.EndTime); err != nil {
				return err
			}

		case "event_id":
			if err := dec.Decode(&s.EventId); err != nil {
				return err
			}

		case "start_time":
			if err := dec.Decode(&s.StartTime); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewCalendarEvent returns a CalendarEvent.
func NewCalendarEvent() *CalendarEvent {
	r := &CalendarEvent{}

	return r
}
