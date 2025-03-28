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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _watch struct {
	v *types.Watch
}

func NewWatch(condition types.WatcherConditionVariant, input types.WatcherInputVariant, trigger types.TriggerContainerVariant) *_watch {

	tmp := &_watch{v: types.NewWatch()}

	tmp.Condition(condition)

	tmp.Input(input)

	tmp.Trigger(trigger)

	return tmp

}

func (s *_watch) Actions(actions map[string]types.WatcherAction) *_watch {

	s.v.Actions = actions
	return s
}

func (s *_watch) AddAction(key string, value types.WatcherActionVariant) *_watch {

	var tmp map[string]types.WatcherAction
	if s.v.Actions == nil {
		s.v.Actions = make(map[string]types.WatcherAction)
	} else {
		tmp = s.v.Actions
	}

	tmp[key] = *value.WatcherActionCaster()

	s.v.Actions = tmp
	return s
}

func (s *_watch) Condition(condition types.WatcherConditionVariant) *_watch {

	s.v.Condition = *condition.WatcherConditionCaster()

	return s
}

func (s *_watch) Input(input types.WatcherInputVariant) *_watch {

	s.v.Input = *input.WatcherInputCaster()

	return s
}

func (s *_watch) Metadata(metadata types.MetadataVariant) *_watch {

	s.v.Metadata = *metadata.MetadataCaster()

	return s
}

func (s *_watch) Status(status types.WatchStatusVariant) *_watch {

	s.v.Status = status.WatchStatusCaster()

	return s
}

func (s *_watch) ThrottlePeriod(duration types.DurationVariant) *_watch {

	s.v.ThrottlePeriod = *duration.DurationCaster()

	return s
}

func (s *_watch) ThrottlePeriodInMillis(durationvalueunitmillis int64) *_watch {

	s.v.ThrottlePeriodInMillis = &durationvalueunitmillis

	return s
}

func (s *_watch) Transform(transform types.TransformContainerVariant) *_watch {

	s.v.Transform = transform.TransformContainerCaster()

	return s
}

func (s *_watch) Trigger(trigger types.TriggerContainerVariant) *_watch {

	s.v.Trigger = *trigger.TriggerContainerCaster()

	return s
}

func (s *_watch) WatchCaster() *types.Watch {
	return s.v
}
