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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _scheduleContainer struct {
	v *types.ScheduleContainer
}

func NewScheduleContainer() *_scheduleContainer {
	return &_scheduleContainer{v: types.NewScheduleContainer()}
}

// AdditionalScheduleContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_scheduleContainer) AdditionalScheduleContainerProperty(key string, value json.RawMessage) *_scheduleContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalScheduleContainerProperty = tmp
	return s
}

func (s *_scheduleContainer) Cron(cronexpression string) *_scheduleContainer {

	s.v.Cron = &cronexpression

	return s
}

func (s *_scheduleContainer) Daily(daily types.DailyScheduleVariant) *_scheduleContainer {

	s.v.Daily = daily.DailyScheduleCaster()

	return s
}

func (s *_scheduleContainer) Hourly(hourly types.HourlyScheduleVariant) *_scheduleContainer {

	s.v.Hourly = hourly.HourlyScheduleCaster()

	return s
}

func (s *_scheduleContainer) Interval(duration types.DurationVariant) *_scheduleContainer {

	s.v.Interval = *duration.DurationCaster()

	return s
}

func (s *_scheduleContainer) Monthly(monthlies ...types.TimeOfMonthVariant) *_scheduleContainer {

	s.v.Monthly = make([]types.TimeOfMonth, len(monthlies))
	for i, v := range monthlies {
		s.v.Monthly[i] = *v.TimeOfMonthCaster()
	}

	return s
}

func (s *_scheduleContainer) Timezone(timezone string) *_scheduleContainer {

	s.v.Timezone = &timezone

	return s
}

func (s *_scheduleContainer) Weekly(weeklies ...types.TimeOfWeekVariant) *_scheduleContainer {

	s.v.Weekly = make([]types.TimeOfWeek, len(weeklies))
	for i, v := range weeklies {
		s.v.Weekly[i] = *v.TimeOfWeekCaster()
	}

	return s
}

func (s *_scheduleContainer) Yearly(yearlies ...types.TimeOfYearVariant) *_scheduleContainer {

	s.v.Yearly = make([]types.TimeOfYear, len(yearlies))
	for i, v := range yearlies {
		s.v.Yearly[i] = *v.TimeOfYearCaster()
	}

	return s
}

func (s *_scheduleContainer) ScheduleContainerCaster() *types.ScheduleContainer {
	return s.v
}
