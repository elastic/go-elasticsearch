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
	"strconv"
)

// WatchStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/watcher/_types/Watch.ts#L49-L56
type WatchStatus struct {
	Actions          WatcherStatusActions `json:"actions"`
	ExecutionState   *string              `json:"execution_state,omitempty"`
	LastChecked      DateTime             `json:"last_checked,omitempty"`
	LastMetCondition DateTime             `json:"last_met_condition,omitempty"`
	State            ActivationState      `json:"state"`
	Version          int64                `json:"version"`
}

func (s *WatchStatus) UnmarshalJSON(data []byte) error {

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
			if err := dec.Decode(&s.Actions); err != nil {
				return fmt.Errorf("%s | %w", "Actions", err)
			}

		case "execution_state":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ExecutionState", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ExecutionState = &o

		case "last_checked":
			if err := dec.Decode(&s.LastChecked); err != nil {
				return fmt.Errorf("%s | %w", "LastChecked", err)
			}

		case "last_met_condition":
			if err := dec.Decode(&s.LastMetCondition); err != nil {
				return fmt.Errorf("%s | %w", "LastMetCondition", err)
			}

		case "state":
			if err := dec.Decode(&s.State); err != nil {
				return fmt.Errorf("%s | %w", "State", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewWatchStatus returns a WatchStatus.
func NewWatchStatus() *WatchStatus {
	r := &WatchStatus{}

	return r
}
