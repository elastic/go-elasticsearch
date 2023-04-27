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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// ScheduleTriggerEvent type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/watcher/_types/Schedule.ts#L98-L101
type ScheduleTriggerEvent struct {
	ScheduledTime DateTime `json:"scheduled_time"`
	TriggeredTime DateTime `json:"triggered_time,omitempty"`
}

func (s *ScheduleTriggerEvent) UnmarshalJSON(data []byte) error {

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

		case "scheduled_time":
			if err := dec.Decode(&s.ScheduledTime); err != nil {
				return err
			}

		case "triggered_time":
			if err := dec.Decode(&s.TriggeredTime); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewScheduleTriggerEvent returns a ScheduleTriggerEvent.
func NewScheduleTriggerEvent() *ScheduleTriggerEvent {
	r := &ScheduleTriggerEvent{}

	return r
}
