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
// https://github.com/elastic/elasticsearch-specification/tree/abf9c2c6bb21328339daa197aae15af2ecbc46f0

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// ExecutionResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/abf9c2c6bb21328339daa197aae15af2ecbc46f0/specification/watcher/_types/Execution.ts#L73-L80
type ExecutionResult struct {
	Actions           []ExecutionResultAction   `json:"actions"`
	Condition         *ExecutionResultCondition `json:"condition,omitempty"`
	ExecutionDuration int64                     `json:"execution_duration"`
	ExecutionTime     DateTime                  `json:"execution_time"`
	Input             *ExecutionResultInput     `json:"input,omitempty"`
	Transform         *ExecutionResultTransform `json:"transform,omitempty"`
}

func (s *ExecutionResult) UnmarshalJSON(data []byte) error {

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

		case "condition":
			if err := dec.Decode(&s.Condition); err != nil {
				return fmt.Errorf("%s | %w", "Condition", err)
			}

		case "execution_duration":
			if err := dec.Decode(&s.ExecutionDuration); err != nil {
				return fmt.Errorf("%s | %w", "ExecutionDuration", err)
			}

		case "execution_time":
			if err := dec.Decode(&s.ExecutionTime); err != nil {
				return fmt.Errorf("%s | %w", "ExecutionTime", err)
			}

		case "input":
			if err := dec.Decode(&s.Input); err != nil {
				return fmt.Errorf("%s | %w", "Input", err)
			}

		case "transform":
			if err := dec.Decode(&s.Transform); err != nil {
				return fmt.Errorf("%s | %w", "Transform", err)
			}

		}
	}
	return nil
}

// NewExecutionResult returns a ExecutionResult.
func NewExecutionResult() *ExecutionResult {
	r := &ExecutionResult{}

	return r
}
