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
// https://github.com/elastic/elasticsearch-specification/tree/9665eef0d78c41f20c4c83e69b7c8155efd58f24

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ExecutionResultHttpInput type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9665eef0d78c41f20c4c83e69b7c8155efd58f24/specification/watcher/_types/Execution.ts#L157-L164
type ExecutionResultHttpInput struct {
	Request HttpInputRequestResult `json:"request"`
	// StatusCode The HTTP status code returned by the request. It is only present when the
	// request was executed.
	StatusCode *int `json:"status_code,omitempty"`
}

func (s *ExecutionResultHttpInput) UnmarshalJSON(data []byte) error {

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

		case "request":
			if err := dec.Decode(&s.Request); err != nil {
				return fmt.Errorf("%s | %w", "Request", err)
			}

		case "status_code":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "StatusCode", err)
				}
				s.StatusCode = &value
			case float64:
				f := int(v)
				s.StatusCode = &f
			}

		}
	}
	return nil
}

// NewExecutionResultHttpInput returns a ExecutionResultHttpInput.
func NewExecutionResultHttpInput() *ExecutionResultHttpInput {
	r := &ExecutionResultHttpInput{}

	return r
}
