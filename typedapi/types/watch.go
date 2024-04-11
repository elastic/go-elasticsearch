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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Watch type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/watcher/_types/Watch.ts#L37-L47
type Watch struct {
	Actions                map[string]WatcherAction `json:"actions"`
	Condition              WatcherCondition         `json:"condition"`
	Input                  WatcherInput             `json:"input"`
	Metadata               Metadata                 `json:"metadata,omitempty"`
	Status                 *WatchStatus             `json:"status,omitempty"`
	ThrottlePeriod         Duration                 `json:"throttle_period,omitempty"`
	ThrottlePeriodInMillis *int64                   `json:"throttle_period_in_millis,omitempty"`
	Transform              *TransformContainer      `json:"transform,omitempty"`
	Trigger                TriggerContainer         `json:"trigger"`
}

func (s *Watch) UnmarshalJSON(data []byte) error {

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

		case "actions":
			if s.Actions == nil {
				s.Actions = make(map[string]WatcherAction, 0)
			}
			if err := dec.Decode(&s.Actions); err != nil {
				return fmt.Errorf("%s | %w", "Actions", err)
			}

		case "condition":
			if err := dec.Decode(&s.Condition); err != nil {
				return fmt.Errorf("%s | %w", "Condition", err)
			}

		case "input":
			if err := dec.Decode(&s.Input); err != nil {
				return fmt.Errorf("%s | %w", "Input", err)
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "throttle_period":
			if err := dec.Decode(&s.ThrottlePeriod); err != nil {
				return fmt.Errorf("%s | %w", "ThrottlePeriod", err)
			}

		case "throttle_period_in_millis":
			if err := dec.Decode(&s.ThrottlePeriodInMillis); err != nil {
				return fmt.Errorf("%s | %w", "ThrottlePeriodInMillis", err)
			}

		case "transform":
			if err := dec.Decode(&s.Transform); err != nil {
				return fmt.Errorf("%s | %w", "Transform", err)
			}

		case "trigger":
			if err := dec.Decode(&s.Trigger); err != nil {
				return fmt.Errorf("%s | %w", "Trigger", err)
			}

		}
	}
	return nil
}

// NewWatch returns a Watch.
func NewWatch() *Watch {
	r := &Watch{
		Actions: make(map[string]WatcherAction, 0),
	}

	return r
}
