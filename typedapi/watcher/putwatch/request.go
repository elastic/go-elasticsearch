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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package putwatch

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putwatch
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/watcher/put_watch/WatcherPutWatchRequest.ts#L31-L110
type Request struct {

	// Actions The list of actions that will be run if the condition matches.
	Actions map[string]types.WatcherAction `json:"actions,omitempty"`
	// Condition The condition that defines if the actions should be run.
	Condition *types.WatcherCondition `json:"condition,omitempty"`
	// Input The input that defines the input that loads the data for the watch.
	Input *types.WatcherInput `json:"input,omitempty"`
	// Metadata Metadata JSON that will be copied into the history entries.
	Metadata types.Metadata `json:"metadata,omitempty"`
	// ThrottlePeriod The minimum time between actions being run.
	// The default is 5 seconds.
	// This default can be changed in the config file with the setting
	// `xpack.watcher.throttle.period.default_period`.
	// If both this value and the `throttle_period_in_millis` parameter are
	// specified, Watcher uses the last parameter included in the request.
	ThrottlePeriod types.Duration `json:"throttle_period,omitempty"`
	// ThrottlePeriodInMillis Minimum time in milliseconds between actions being run. Defaults to 5000. If
	// both this value and the throttle_period parameter are specified, Watcher uses
	// the last parameter included in the request.
	ThrottlePeriodInMillis *int64 `json:"throttle_period_in_millis,omitempty"`
	// Transform The transform that processes the watch payload to prepare it for the watch
	// actions.
	Transform *types.TransformContainer `json:"transform,omitempty"`
	// Trigger The trigger that defines when the watch should run.
	Trigger *types.TriggerContainer `json:"trigger,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Actions: make(map[string]types.WatcherAction, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putwatch request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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
				s.Actions = make(map[string]types.WatcherAction, 0)
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
