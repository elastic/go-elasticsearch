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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _calendarEvent struct {
	v *types.CalendarEvent
}

func NewCalendarEvent(description string) *_calendarEvent {

	tmp := &_calendarEvent{v: types.NewCalendarEvent()}

	tmp.Description(description)

	return tmp

}

func (s *_calendarEvent) CalendarId(id string) *_calendarEvent {

	s.v.CalendarId = &id

	return s
}

func (s *_calendarEvent) Description(description string) *_calendarEvent {

	s.v.Description = description

	return s
}

func (s *_calendarEvent) EndTime(datetime types.DateTimeVariant) *_calendarEvent {

	s.v.EndTime = *datetime.DateTimeCaster()

	return s
}

func (s *_calendarEvent) EventId(id string) *_calendarEvent {

	s.v.EventId = &id

	return s
}

func (s *_calendarEvent) ForceTimeShift(forcetimeshift int) *_calendarEvent {

	s.v.ForceTimeShift = &forcetimeshift

	return s
}

func (s *_calendarEvent) SkipModelUpdate(skipmodelupdate bool) *_calendarEvent {

	s.v.SkipModelUpdate = &skipmodelupdate

	return s
}

func (s *_calendarEvent) SkipResult(skipresult bool) *_calendarEvent {

	s.v.SkipResult = &skipresult

	return s
}

func (s *_calendarEvent) StartTime(datetime types.DateTimeVariant) *_calendarEvent {

	s.v.StartTime = *datetime.DateTimeCaster()

	return s
}

func (s *_calendarEvent) CalendarEventCaster() *types.CalendarEvent {
	return s.v
}
