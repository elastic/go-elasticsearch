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
// https://github.com/elastic/elasticsearch-specification/tree/7fd0bd13eaf28bd179fc906f57da09e852eb818e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/actionstatusoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/actiontype"
)

// ExecutionResultAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7fd0bd13eaf28bd179fc906f57da09e852eb818e/specification/watcher/_types/Execution.ts#L108-L117
type ExecutionResultAction struct {
	Condition               *ExecutionResultCondition               `json:"condition,omitempty"`
	Email                   *EmailResult                            `json:"email,omitempty"`
	Error                   *ErrorCause                             `json:"error,omitempty"`
	Foreach                 []ExecutionResultForeachAction          `json:"foreach,omitempty"`
	Id                      string                                  `json:"id"`
	Index                   *IndexResult                            `json:"index,omitempty"`
	Logging                 *LoggingResult                          `json:"logging,omitempty"`
	MaxIterations           *int                                    `json:"max_iterations,omitempty"`
	NumberOfActionsExecuted *int                                    `json:"number_of_actions_executed,omitempty"`
	Pagerduty               *PagerDutyResult                        `json:"pagerduty,omitempty"`
	Reason                  *string                                 `json:"reason,omitempty"`
	Slack                   *SlackResult                            `json:"slack,omitempty"`
	Status                  actionstatusoptions.ActionStatusOptions `json:"status"`
	Transform               *ExecutionResultTransform               `json:"transform,omitempty"`
	Type                    actiontype.ActionType                   `json:"type"`
	Webhook                 *WebhookResult                          `json:"webhook,omitempty"`
}

func (s *ExecutionResultAction) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "condition":
			if err := dec.Decode(&s.Condition); err != nil {
				return fmt.Errorf("%s | %w", "Condition", err)
			}

		case "email":
			if err := dec.Decode(&s.Email); err != nil {
				return fmt.Errorf("%s | %w", "Email", err)
			}

		case "error":
			if err := dec.Decode(&s.Error); err != nil {
				return fmt.Errorf("%s | %w", "Error", err)
			}

		case "foreach":
			if err := dec.Decode(&s.Foreach); err != nil {
				return fmt.Errorf("%s | %w", "Foreach", err)
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "logging":
			if err := dec.Decode(&s.Logging); err != nil {
				return fmt.Errorf("%s | %w", "Logging", err)
			}

		case "max_iterations":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxIterations", err)
				}
				s.MaxIterations = &value
			case float64:
				f := int(v)
				s.MaxIterations = &f
			}

		case "number_of_actions_executed":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfActionsExecuted", err)
				}
				s.NumberOfActionsExecuted = &value
			case float64:
				f := int(v)
				s.NumberOfActionsExecuted = &f
			}

		case "pagerduty":
			if err := dec.Decode(&s.Pagerduty); err != nil {
				return fmt.Errorf("%s | %w", "Pagerduty", err)
			}

		case "reason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Reason", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = &o

		case "slack":
			if err := dec.Decode(&s.Slack); err != nil {
				return fmt.Errorf("%s | %w", "Slack", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "transform":
			if err := dec.Decode(&s.Transform); err != nil {
				return fmt.Errorf("%s | %w", "Transform", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "webhook":
			if err := dec.Decode(&s.Webhook); err != nil {
				return fmt.Errorf("%s | %w", "Webhook", err)
			}

		}
	}
	return nil
}

// NewExecutionResultAction returns a ExecutionResultAction.
func NewExecutionResultAction() *ExecutionResultAction {
	r := &ExecutionResultAction{}

	return r
}
