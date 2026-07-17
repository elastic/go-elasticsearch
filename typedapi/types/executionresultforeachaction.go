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
// https://github.com/elastic/elasticsearch-specification/tree/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// The result of a single action execution.
//
// https://github.com/elastic/elasticsearch-specification/blob/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c/specification/watcher/_types/Execution.ts#L94-L106
type ExecutionResultForeachAction struct {
	Email     *EmailResult     `json:"email,omitempty"`
	Error     *ErrorCause      `json:"error,omitempty"`
	Index     *IndexResult     `json:"index,omitempty"`
	Logging   *LoggingResult   `json:"logging,omitempty"`
	Pagerduty *PagerDutyResult `json:"pagerduty,omitempty"`
	Reason    *string          `json:"reason,omitempty"`
	Slack     *SlackResult     `json:"slack,omitempty"`
	Webhook   *WebhookResult   `json:"webhook,omitempty"`
}

func (s *ExecutionResultForeachAction) UnmarshalJSON(data []byte) error {

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

		case "email":
			if err := dec.Decode(&s.Email); err != nil {
				return fmt.Errorf("%s | %w", "Email", err)
			}

		case "error":
			if err := dec.Decode(&s.Error); err != nil {
				return fmt.Errorf("%s | %w", "Error", err)
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "logging":
			if err := dec.Decode(&s.Logging); err != nil {
				return fmt.Errorf("%s | %w", "Logging", err)
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

		case "webhook":
			if err := dec.Decode(&s.Webhook); err != nil {
				return fmt.Errorf("%s | %w", "Webhook", err)
			}

		}
	}
	return nil
}

// NewExecutionResultForeachAction returns a ExecutionResultForeachAction.
func NewExecutionResultForeachAction() *ExecutionResultForeachAction {
	r := &ExecutionResultForeachAction{}

	return r
}
