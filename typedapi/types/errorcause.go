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
	"encoding/json"
	"fmt"
)

// ErrorCause type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/Errors.ts#L25-L48
type ErrorCause struct {
	CausedBy *ErrorCause                `json:"caused_by,omitempty"`
	Metadata map[string]json.RawMessage `json:"-"`
	// Reason A human-readable explanation of the error, in english
	Reason    *string      `json:"reason,omitempty"`
	RootCause []ErrorCause `json:"root_cause,omitempty"`
	// StackTrace The server stack trace. Present only if the `error_trace=true` parameter was
	// sent with the request.
	StackTrace *string      `json:"stack_trace,omitempty"`
	Suppressed []ErrorCause `json:"suppressed,omitempty"`
	// Type The type of error
	Type string `json:"type"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s ErrorCause) MarshalJSON() ([]byte, error) {
	type opt ErrorCause
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.Metadata {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "Metadata")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewErrorCause returns a ErrorCause.
func NewErrorCause() *ErrorCause {
	r := &ErrorCause{
		Metadata: make(map[string]json.RawMessage, 0),
	}

	return r
}
