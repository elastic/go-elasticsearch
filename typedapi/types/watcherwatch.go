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

// WatcherWatch type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/xpack/usage/types.ts#L391-L396
type WatcherWatch struct {
	Action    map[string]Counter  `json:"action,omitempty"`
	Condition map[string]Counter  `json:"condition,omitempty"`
	Input     map[string]Counter  `json:"input"`
	Trigger   WatcherWatchTrigger `json:"trigger"`
}

func (s *WatcherWatch) UnmarshalJSON(data []byte) error {

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

		case "action":
			if s.Action == nil {
				s.Action = make(map[string]Counter, 0)
			}
			if err := dec.Decode(&s.Action); err != nil {
				return err
			}

		case "condition":
			if s.Condition == nil {
				s.Condition = make(map[string]Counter, 0)
			}
			if err := dec.Decode(&s.Condition); err != nil {
				return err
			}

		case "input":
			if s.Input == nil {
				s.Input = make(map[string]Counter, 0)
			}
			if err := dec.Decode(&s.Input); err != nil {
				return err
			}

		case "trigger":
			if err := dec.Decode(&s.Trigger); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewWatcherWatch returns a WatcherWatch.
func NewWatcherWatch() *WatcherWatch {
	r := &WatcherWatch{
		Action:    make(map[string]Counter, 0),
		Condition: make(map[string]Counter, 0),
		Input:     make(map[string]Counter, 0),
	}

	return r
}
