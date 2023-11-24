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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actiontype"
)

// WatcherAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/watcher/_types/Action.ts#L41-L60
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

func (s *WatcherAction) UnmarshalJSON(data []byte) error {

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

		case "action_type":
			if err := dec.Decode(&s.ActionType); err != nil {
				return err
			}

		case "condition":
			if err := dec.Decode(&s.Condition); err != nil {
				return err
			}

		case "email":
			if err := dec.Decode(&s.Email); err != nil {
				return err
			}

		case "foreach":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Foreach = &o

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "logging":
			if err := dec.Decode(&s.Logging); err != nil {
				return err
			}

		case "max_iterations":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxIterations = &value
			case float64:
				f := int(v)
				s.MaxIterations = &f
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "pagerduty":
			if err := dec.Decode(&s.Pagerduty); err != nil {
				return err
			}

		case "slack":
			if err := dec.Decode(&s.Slack); err != nil {
				return err
			}

		case "throttle_period":
			if err := dec.Decode(&s.ThrottlePeriod); err != nil {
				return err
			}

		case "throttle_period_in_millis":
			if err := dec.Decode(&s.ThrottlePeriodInMillis); err != nil {
				return err
			}

		case "transform":
			if err := dec.Decode(&s.Transform); err != nil {
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

// NewWatcherAction returns a WatcherAction.
func NewWatcherAction() *WatcherAction {
	r := &WatcherAction{}

	return r
}
