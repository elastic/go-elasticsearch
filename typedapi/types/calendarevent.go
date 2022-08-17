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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// CalendarEvent type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/CalendarEvent.ts#L23-L33
type CalendarEvent struct {
	// CalendarId A string that uniquely identifies a calendar.
	CalendarId *Id `json:"calendar_id,omitempty"`
	// Description A description of the scheduled event.
	Description string `json:"description"`
	// EndTime The timestamp for the end of the scheduled event in milliseconds since the
	// epoch or ISO 8601 format.
	EndTime DateTime `json:"end_time"`
	EventId *Id      `json:"event_id,omitempty"`
	// StartTime The timestamp for the beginning of the scheduled event in milliseconds since
	// the epoch or ISO 8601 format.
	StartTime DateTime `json:"start_time"`
}

// CalendarEventBuilder holds CalendarEvent struct and provides a builder API.
type CalendarEventBuilder struct {
	v *CalendarEvent
}

// NewCalendarEvent provides a builder for the CalendarEvent struct.
func NewCalendarEventBuilder() *CalendarEventBuilder {
	r := CalendarEventBuilder{
		&CalendarEvent{},
	}

	return &r
}

// Build finalize the chain and returns the CalendarEvent struct
func (rb *CalendarEventBuilder) Build() CalendarEvent {
	return *rb.v
}

// CalendarId A string that uniquely identifies a calendar.

func (rb *CalendarEventBuilder) CalendarId(calendarid Id) *CalendarEventBuilder {
	rb.v.CalendarId = &calendarid
	return rb
}

// Description A description of the scheduled event.

func (rb *CalendarEventBuilder) Description(description string) *CalendarEventBuilder {
	rb.v.Description = description
	return rb
}

// EndTime The timestamp for the end of the scheduled event in milliseconds since the
// epoch or ISO 8601 format.

func (rb *CalendarEventBuilder) EndTime(endtime *DateTimeBuilder) *CalendarEventBuilder {
	v := endtime.Build()
	rb.v.EndTime = v
	return rb
}

func (rb *CalendarEventBuilder) EventId(eventid Id) *CalendarEventBuilder {
	rb.v.EventId = &eventid
	return rb
}

// StartTime The timestamp for the beginning of the scheduled event in milliseconds since
// the epoch or ISO 8601 format.

func (rb *CalendarEventBuilder) StartTime(starttime *DateTimeBuilder) *CalendarEventBuilder {
	v := starttime.Build()
	rb.v.StartTime = v
	return rb
}
