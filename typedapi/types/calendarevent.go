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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CalendarEvent type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/CalendarEvent.ts#L24-L44
type CalendarEvent struct {
	// CalendarId A string that uniquely identifies a calendar.
	CalendarId *string `json:"calendar_id,omitempty"`
	// Description A description of the scheduled event.
	Description string `json:"description"`
	// EndTime The timestamp for the end of the scheduled event in milliseconds since the
	// epoch or ISO 8601 format.
	EndTime DateTime `json:"end_time"`
	EventId *string  `json:"event_id,omitempty"`
	// ForceTimeShift Shift time by this many seconds. For example adjust time for daylight savings
	// changes
	ForceTimeShift *int `json:"force_time_shift,omitempty"`
	// SkipModelUpdate When true the model will not be updated for this calendar period.
	SkipModelUpdate *bool `json:"skip_model_update,omitempty"`
	// SkipResult When true the model will not create results for this calendar period.
	SkipResult *bool `json:"skip_result,omitempty"`
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
				return fmt.Errorf("%s | %w", "CalendarId", err)
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
			s.Description = o

		case "end_time":
			if err := dec.Decode(&s.EndTime); err != nil {
				return fmt.Errorf("%s | %w", "EndTime", err)
			}

		case "event_id":
			if err := dec.Decode(&s.EventId); err != nil {
				return fmt.Errorf("%s | %w", "EventId", err)
			}

		case "force_time_shift":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ForceTimeShift", err)
				}
				s.ForceTimeShift = &value
			case float64:
				f := int(v)
				s.ForceTimeShift = &f
			}

		case "skip_model_update":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SkipModelUpdate", err)
				}
				s.SkipModelUpdate = &value
			case bool:
				s.SkipModelUpdate = &v
			}

		case "skip_result":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SkipResult", err)
				}
				s.SkipResult = &value
			case bool:
				s.SkipResult = &v
			}

		case "start_time":
			if err := dec.Decode(&s.StartTime); err != nil {
				return fmt.Errorf("%s | %w", "StartTime", err)
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

type CalendarEventVariant interface {
	CalendarEventCaster() *CalendarEvent
}

func (s *CalendarEvent) CalendarEventCaster() *CalendarEvent {
	return s
}
