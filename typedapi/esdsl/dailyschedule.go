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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dailySchedule struct {
	v *types.DailySchedule
}

func NewDailySchedule() *_dailySchedule {

	return &_dailySchedule{v: types.NewDailySchedule()}

}

func (s *_dailySchedule) At(ats ...types.ScheduleTimeOfDayVariant) *_dailySchedule {

	for _, v := range ats {

		s.v.At = append(s.v.At, *v.ScheduleTimeOfDayCaster())

	}
	return s
}

func (s *_dailySchedule) ScheduleContainerCaster() *types.ScheduleContainer {
	container := types.NewScheduleContainer()

	container.Daily = s.v

	return container
}

func (s *_dailySchedule) DailyScheduleCaster() *types.DailySchedule {
	return s.v
}
