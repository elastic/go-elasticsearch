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
// https://github.com/elastic/elasticsearch-specification/tree/7fd0bd13eaf28bd179fc906f57da09e852eb818e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/executionresultstatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/executionresulttransformtype"
)

// ExecutionResultTransform type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7fd0bd13eaf28bd179fc906f57da09e852eb818e/specification/watcher/_types/Execution.ts#L119-L126
type ExecutionResultTransform struct {
	Error   *ErrorCause                                               `json:"error,omitempty"`
	Payload map[string]json.RawMessage                                `json:"payload,omitempty"`
	Reason  *string                                                   `json:"reason,omitempty"`
	Search  *ExecutionResultSearchInput                               `json:"search,omitempty"`
	Status  executionresultstatus.ExecutionResultStatus               `json:"status"`
	Type    executionresulttransformtype.ExecutionResultTransformType `json:"type"`
}

func (s *ExecutionResultTransform) UnmarshalJSON(data []byte) error {

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

		case "error":
			if err := dec.Decode(&s.Error); err != nil {
				return fmt.Errorf("%s | %w", "Error", err)
			}

		case "payload":
			if s.Payload == nil {
				s.Payload = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Payload); err != nil {
				return fmt.Errorf("%s | %w", "Payload", err)
			}

		case "reason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Reason", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = &o

		case "search":
			if err := dec.Decode(&s.Search); err != nil {
				return fmt.Errorf("%s | %w", "Search", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// NewExecutionResultTransform returns a ExecutionResultTransform.
func NewExecutionResultTransform() *ExecutionResultTransform {
	r := &ExecutionResultTransform{
		Payload: make(map[string]json.RawMessage),
	}

	return r
}
