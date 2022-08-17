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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actiontype"
)

// Action type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Action.ts#L41-L57
type Action struct {
	ActionType             *actiontype.ActionType   `json:"action_type,omitempty"`
	Condition              *ConditionContainer      `json:"condition,omitempty"`
	Email                  *EmailAction             `json:"email,omitempty"`
	Foreach                *string                  `json:"foreach,omitempty"`
	Index                  *IndexAction             `json:"index,omitempty"`
	Logging                *LoggingAction           `json:"logging,omitempty"`
	MaxIterations          *int                     `json:"max_iterations,omitempty"`
	Name                   *Name                    `json:"name,omitempty"`
	Pagerduty              *PagerDutyAction         `json:"pagerduty,omitempty"`
	Slack                  *SlackAction             `json:"slack,omitempty"`
	ThrottlePeriod         *Duration                `json:"throttle_period,omitempty"`
	ThrottlePeriodInMillis *DurationValueUnitMillis `json:"throttle_period_in_millis,omitempty"`
	Transform              *TransformContainer      `json:"transform,omitempty"`
	Webhook                *WebhookAction           `json:"webhook,omitempty"`
}

// ActionBuilder holds Action struct and provides a builder API.
type ActionBuilder struct {
	v *Action
}

// NewAction provides a builder for the Action struct.
func NewActionBuilder() *ActionBuilder {
	r := ActionBuilder{
		&Action{},
	}

	return &r
}

// Build finalize the chain and returns the Action struct
func (rb *ActionBuilder) Build() Action {
	return *rb.v
}

func (rb *ActionBuilder) ActionType(actiontype actiontype.ActionType) *ActionBuilder {
	rb.v.ActionType = &actiontype
	return rb
}

func (rb *ActionBuilder) Condition(condition *ConditionContainerBuilder) *ActionBuilder {
	v := condition.Build()
	rb.v.Condition = &v
	return rb
}

func (rb *ActionBuilder) Email(email *EmailActionBuilder) *ActionBuilder {
	v := email.Build()
	rb.v.Email = &v
	return rb
}

func (rb *ActionBuilder) Foreach(foreach string) *ActionBuilder {
	rb.v.Foreach = &foreach
	return rb
}

func (rb *ActionBuilder) Index(index *IndexActionBuilder) *ActionBuilder {
	v := index.Build()
	rb.v.Index = &v
	return rb
}

func (rb *ActionBuilder) Logging(logging *LoggingActionBuilder) *ActionBuilder {
	v := logging.Build()
	rb.v.Logging = &v
	return rb
}

func (rb *ActionBuilder) MaxIterations(maxiterations int) *ActionBuilder {
	rb.v.MaxIterations = &maxiterations
	return rb
}

func (rb *ActionBuilder) Name(name Name) *ActionBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *ActionBuilder) Pagerduty(pagerduty *PagerDutyActionBuilder) *ActionBuilder {
	v := pagerduty.Build()
	rb.v.Pagerduty = &v
	return rb
}

func (rb *ActionBuilder) Slack(slack *SlackActionBuilder) *ActionBuilder {
	v := slack.Build()
	rb.v.Slack = &v
	return rb
}

func (rb *ActionBuilder) ThrottlePeriod(throttleperiod *DurationBuilder) *ActionBuilder {
	v := throttleperiod.Build()
	rb.v.ThrottlePeriod = &v
	return rb
}

func (rb *ActionBuilder) ThrottlePeriodInMillis(throttleperiodinmillis *DurationValueUnitMillisBuilder) *ActionBuilder {
	v := throttleperiodinmillis.Build()
	rb.v.ThrottlePeriodInMillis = &v
	return rb
}

func (rb *ActionBuilder) Transform(transform *TransformContainerBuilder) *ActionBuilder {
	v := transform.Build()
	rb.v.Transform = &v
	return rb
}

func (rb *ActionBuilder) Webhook(webhook *WebhookActionBuilder) *ActionBuilder {
	v := webhook.Build()
	rb.v.Webhook = &v
	return rb
}
