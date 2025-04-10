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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _scheduleTimeOfDay struct {
	v types.ScheduleTimeOfDay
}

func NewScheduleTimeOfDay() *_scheduleTimeOfDay {
	return &_scheduleTimeOfDay{v: nil}
}

func (u *_scheduleTimeOfDay) String(string string) *_scheduleTimeOfDay {

	u.v = &string

	return u
}

func (u *_scheduleTimeOfDay) HourAndMinute(hourandminute types.HourAndMinuteVariant) *_scheduleTimeOfDay {

	u.v = &hourandminute

	return u
}

// Interface implementation for HourAndMinute in ScheduleTimeOfDay union
func (u *_hourAndMinute) ScheduleTimeOfDayCaster() *types.ScheduleTimeOfDay {
	t := types.ScheduleTimeOfDay(u.v)
	return &t
}

func (u *_scheduleTimeOfDay) ScheduleTimeOfDayCaster() *types.ScheduleTimeOfDay {
	return &u.v
}
