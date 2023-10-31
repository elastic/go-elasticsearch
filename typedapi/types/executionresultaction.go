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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actionstatusoptions"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actiontype"
)

// ExecutionResultAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/watcher/_types/Execution.ts#L74-L86
type ExecutionResultAction struct {
	Email     *EmailResult                            `json:"email,omitempty"`
	Error     *ErrorCause                             `json:"error,omitempty"`
	Id        string                                  `json:"id"`
	Index     *IndexResult                            `json:"index,omitempty"`
	Logging   *LoggingResult                          `json:"logging,omitempty"`
	Pagerduty *PagerDutyResult                        `json:"pagerduty,omitempty"`
	Reason    *string                                 `json:"reason,omitempty"`
	Slack     *SlackResult                            `json:"slack,omitempty"`
	Status    actionstatusoptions.ActionStatusOptions `json:"status"`
	Type      actiontype.ActionType                   `json:"type"`
	Webhook   *WebhookResult                          `json:"webhook,omitempty"`
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

		case "email":
			if err := dec.Decode(&s.Email); err != nil {
				return err
			}

		case "error":
			if err := dec.Decode(&s.Error); err != nil {
				return err
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "logging":
			if err := dec.Decode(&s.Logging); err != nil {
				return err
			}

		case "pagerduty":
			if err := dec.Decode(&s.Pagerduty); err != nil {
				return err
			}

		case "reason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = &o

		case "slack":
			if err := dec.Decode(&s.Slack); err != nil {
				return err
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "webhook":
			if err := dec.Decode(&s.Webhook); err != nil {
				return err
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
