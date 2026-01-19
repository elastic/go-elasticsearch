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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scheduleTriggerEvent struct {
	v *types.ScheduleTriggerEvent
}

func NewScheduleTriggerEvent() *_scheduleTriggerEvent {

	return &_scheduleTriggerEvent{v: types.NewScheduleTriggerEvent()}

}

func (s *_scheduleTriggerEvent) ScheduledTime(datetime types.DateTimeVariant) *_scheduleTriggerEvent {

	s.v.ScheduledTime = *datetime.DateTimeCaster()

	return s
}

func (s *_scheduleTriggerEvent) TriggeredTime(datetime types.DateTimeVariant) *_scheduleTriggerEvent {

	s.v.TriggeredTime = *datetime.DateTimeCaster()

	return s
}

func (s *_scheduleTriggerEvent) TriggerEventContainerCaster() *types.TriggerEventContainer {
	container := types.NewTriggerEventContainer()

	container.Schedule = s.v

	return container
}

func (s *_scheduleTriggerEvent) ScheduleTriggerEventCaster() *types.ScheduleTriggerEvent {
	return s.v
}
