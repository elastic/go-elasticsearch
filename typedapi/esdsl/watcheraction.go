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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/actiontype"
)

type _watcherAction struct {
	v *types.WatcherAction
}

func NewWatcherAction() *_watcherAction {

	return &_watcherAction{v: types.NewWatcherAction()}

}

func (s *_watcherAction) ActionType(actiontype actiontype.ActionType) *_watcherAction {

	s.v.ActionType = &actiontype
	return s
}

func (s *_watcherAction) Condition(condition types.WatcherConditionVariant) *_watcherAction {

	s.v.Condition = condition.WatcherConditionCaster()

	return s
}

func (s *_watcherAction) Email(email types.EmailActionVariant) *_watcherAction {

	s.v.Email = email.EmailActionCaster()

	return s
}

func (s *_watcherAction) Foreach(foreach string) *_watcherAction {

	s.v.Foreach = &foreach

	return s
}

func (s *_watcherAction) Index(index types.IndexActionVariant) *_watcherAction {

	s.v.Index = index.IndexActionCaster()

	return s
}

func (s *_watcherAction) Logging(logging types.LoggingActionVariant) *_watcherAction {

	s.v.Logging = logging.LoggingActionCaster()

	return s
}

func (s *_watcherAction) MaxIterations(maxiterations int) *_watcherAction {

	s.v.MaxIterations = &maxiterations

	return s
}

func (s *_watcherAction) Name(name string) *_watcherAction {

	s.v.Name = &name

	return s
}

func (s *_watcherAction) Pagerduty(pagerduty types.PagerDutyActionVariant) *_watcherAction {

	s.v.Pagerduty = pagerduty.PagerDutyActionCaster()

	return s
}

func (s *_watcherAction) Slack(slack types.SlackActionVariant) *_watcherAction {

	s.v.Slack = slack.SlackActionCaster()

	return s
}

func (s *_watcherAction) ThrottlePeriod(duration types.DurationVariant) *_watcherAction {

	s.v.ThrottlePeriod = *duration.DurationCaster()

	return s
}

func (s *_watcherAction) ThrottlePeriodInMillis(durationvalueunitmillis int64) *_watcherAction {

	s.v.ThrottlePeriodInMillis = &durationvalueunitmillis

	return s
}

func (s *_watcherAction) Transform(transform types.TransformContainerVariant) *_watcherAction {

	s.v.Transform = transform.TransformContainerCaster()

	return s
}

func (s *_watcherAction) Webhook(webhook types.WebhookActionVariant) *_watcherAction {

	s.v.Webhook = webhook.WebhookActionCaster()

	return s
}

func (s *_watcherAction) WatcherActionCaster() *types.WatcherAction {
	return s.v
}
