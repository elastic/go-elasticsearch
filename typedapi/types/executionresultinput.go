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
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/executionresultstatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/inputtype"
)

// ExecutionResultInput type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9665eef0d78c41f20c4c83e69b7c8155efd58f24/specification/watcher/_types/Execution.ts#L134-L151
type ExecutionResultInput struct {
	// Chain The result of each named input, present when the input is a chain input.
	Chain map[string]ExecutionResultInput `json:"chain,omitempty"`
	Error *ErrorCause                     `json:"error,omitempty"`
	// Http The resolved HTTP request, present when the input is an HTTP input.
	Http    *ExecutionResultHttpInput  `json:"http,omitempty"`
	Payload map[string]json.RawMessage `json:"payload,omitempty"`
	// Search The resolved search request, present when the input is a search input.
	Search *ExecutionResultSearchInput                 `json:"search,omitempty"`
	Status executionresultstatus.ExecutionResultStatus `json:"status"`
	Type   inputtype.InputType                         `json:"type"`
}

// NewExecutionResultInput returns a ExecutionResultInput.
func NewExecutionResultInput() *ExecutionResultInput {
	r := &ExecutionResultInput{
		Chain:   make(map[string]ExecutionResultInput),
		Payload: make(map[string]json.RawMessage),
	}

	return r
}
