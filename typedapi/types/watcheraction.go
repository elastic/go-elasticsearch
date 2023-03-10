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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actiontype"
)

// WatcherAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/watcher/_types/Action.ts#L41-L57
type WatcherAction struct {
	ActionType             *actiontype.ActionType `json:"action_type,omitempty"`
	Condition              *WatcherCondition      `json:"condition,omitempty"`
	Email                  *EmailAction           `json:"email,omitempty"`
	Foreach                *string                `json:"foreach,omitempty"`
	Index                  *IndexAction           `json:"index,omitempty"`
	Logging                *LoggingAction         `json:"logging,omitempty"`
	MaxIterations          *int                   `json:"max_iterations,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	Pagerduty              *PagerDutyAction       `json:"pagerduty,omitempty"`
	Slack                  *SlackAction           `json:"slack,omitempty"`
	ThrottlePeriod         Duration               `json:"throttle_period,omitempty"`
	ThrottlePeriodInMillis *int64                 `json:"throttle_period_in_millis,omitempty"`
	Transform              *TransformContainer    `json:"transform,omitempty"`
	Webhook                *WebhookAction         `json:"webhook,omitempty"`
}

// NewWatcherAction returns a WatcherAction.
func NewWatcherAction() *WatcherAction {
	r := &WatcherAction{}

	return r
}
