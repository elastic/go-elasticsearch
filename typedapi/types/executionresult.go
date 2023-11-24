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
)

// ExecutionResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/watcher/_types/Execution.ts#L60-L66
type ExecutionResult struct {
	Actions           []ExecutionResultAction  `json:"actions"`
	Condition         ExecutionResultCondition `json:"condition"`
	ExecutionDuration int64                    `json:"execution_duration"`
	ExecutionTime     DateTime                 `json:"execution_time"`
	Input             ExecutionResultInput     `json:"input"`
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
				return err
			}

		case "condition":
			if err := dec.Decode(&s.Condition); err != nil {
				return err
			}

		case "execution_duration":
			if err := dec.Decode(&s.ExecutionDuration); err != nil {
				return err
			}

		case "execution_time":
			if err := dec.Decode(&s.ExecutionTime); err != nil {
				return err
			}

		case "input":
			if err := dec.Decode(&s.Input); err != nil {
				return err
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
