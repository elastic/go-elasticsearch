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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actionstatusoptions"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actiontype"
)

// ExecutionResultAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Execution.ts#L74-L86
type ExecutionResultAction struct {
	Email     *EmailResult                            `json:"email,omitempty"`
	Error     *ErrorCause                             `json:"error,omitempty"`
	Id        Id                                      `json:"id"`
	Index     *IndexResult                            `json:"index,omitempty"`
	Logging   *LoggingResult                          `json:"logging,omitempty"`
	Pagerduty *PagerDutyResult                        `json:"pagerduty,omitempty"`
	Reason    *string                                 `json:"reason,omitempty"`
	Slack     *SlackResult                            `json:"slack,omitempty"`
	Status    actionstatusoptions.ActionStatusOptions `json:"status"`
	Type      actiontype.ActionType                   `json:"type"`
	Webhook   *WebhookResult                          `json:"webhook,omitempty"`
}

// ExecutionResultActionBuilder holds ExecutionResultAction struct and provides a builder API.
type ExecutionResultActionBuilder struct {
	v *ExecutionResultAction
}

// NewExecutionResultAction provides a builder for the ExecutionResultAction struct.
func NewExecutionResultActionBuilder() *ExecutionResultActionBuilder {
	r := ExecutionResultActionBuilder{
		&ExecutionResultAction{},
	}

	return &r
}

// Build finalize the chain and returns the ExecutionResultAction struct
func (rb *ExecutionResultActionBuilder) Build() ExecutionResultAction {
	return *rb.v
}

func (rb *ExecutionResultActionBuilder) Email(email *EmailResultBuilder) *ExecutionResultActionBuilder {
	v := email.Build()
	rb.v.Email = &v
	return rb
}

func (rb *ExecutionResultActionBuilder) Error(error *ErrorCauseBuilder) *ExecutionResultActionBuilder {
	v := error.Build()
	rb.v.Error = &v
	return rb
}

func (rb *ExecutionResultActionBuilder) Id(id Id) *ExecutionResultActionBuilder {
	rb.v.Id = id
	return rb
}

func (rb *ExecutionResultActionBuilder) Index(index *IndexResultBuilder) *ExecutionResultActionBuilder {
	v := index.Build()
	rb.v.Index = &v
	return rb
}

func (rb *ExecutionResultActionBuilder) Logging(logging *LoggingResultBuilder) *ExecutionResultActionBuilder {
	v := logging.Build()
	rb.v.Logging = &v
	return rb
}

func (rb *ExecutionResultActionBuilder) Pagerduty(pagerduty *PagerDutyResultBuilder) *ExecutionResultActionBuilder {
	v := pagerduty.Build()
	rb.v.Pagerduty = &v
	return rb
}

func (rb *ExecutionResultActionBuilder) Reason(reason string) *ExecutionResultActionBuilder {
	rb.v.Reason = &reason
	return rb
}

func (rb *ExecutionResultActionBuilder) Slack(slack *SlackResultBuilder) *ExecutionResultActionBuilder {
	v := slack.Build()
	rb.v.Slack = &v
	return rb
}

func (rb *ExecutionResultActionBuilder) Status(status actionstatusoptions.ActionStatusOptions) *ExecutionResultActionBuilder {
	rb.v.Status = status
	return rb
}

func (rb *ExecutionResultActionBuilder) Type_(type_ actiontype.ActionType) *ExecutionResultActionBuilder {
	rb.v.Type = type_
	return rb
}

func (rb *ExecutionResultActionBuilder) Webhook(webhook *WebhookResultBuilder) *ExecutionResultActionBuilder {
	v := webhook.Build()
	rb.v.Webhook = &v
	return rb
}
