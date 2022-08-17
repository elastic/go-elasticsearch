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

// Calendar type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/get_calendars/types.ts#L22-L29
type Calendar struct {
	// CalendarId A string that uniquely identifies a calendar.
	CalendarId Id `json:"calendar_id"`
	// Description A description of the calendar.
	Description *string `json:"description,omitempty"`
	// JobIds An array of anomaly detection job identifiers.
	JobIds []Id `json:"job_ids"`
}

// CalendarBuilder holds Calendar struct and provides a builder API.
type CalendarBuilder struct {
	v *Calendar
}

// NewCalendar provides a builder for the Calendar struct.
func NewCalendarBuilder() *CalendarBuilder {
	r := CalendarBuilder{
		&Calendar{},
	}

	return &r
}

// Build finalize the chain and returns the Calendar struct
func (rb *CalendarBuilder) Build() Calendar {
	return *rb.v
}

// CalendarId A string that uniquely identifies a calendar.

func (rb *CalendarBuilder) CalendarId(calendarid Id) *CalendarBuilder {
	rb.v.CalendarId = calendarid
	return rb
}

// Description A description of the calendar.

func (rb *CalendarBuilder) Description(description string) *CalendarBuilder {
	rb.v.Description = &description
	return rb
}

// JobIds An array of anomaly detection job identifiers.

func (rb *CalendarBuilder) JobIds(job_ids ...Id) *CalendarBuilder {
	rb.v.JobIds = job_ids
	return rb
}
