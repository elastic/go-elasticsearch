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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

import (
	"encoding/json"
)

// ErrorCause type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Errors.ts#L25-L48
type ErrorCause struct {
	CausedBy *ErrorCause            `json:"caused_by,omitempty"`
	Metadata map[string]interface{} `json:"-"`
	// Reason A human-readable explanation of the error, in english
	Reason    string       `json:"reason"`
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
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ErrorCauseBuilder holds ErrorCause struct and provides a builder API.
type ErrorCauseBuilder struct {
	v *ErrorCause
}

// NewErrorCause provides a builder for the ErrorCause struct.
func NewErrorCauseBuilder() *ErrorCauseBuilder {
	r := ErrorCauseBuilder{
		&ErrorCause{
			Metadata: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ErrorCause struct
func (rb *ErrorCauseBuilder) Build() ErrorCause {
	return *rb.v
}

func (rb *ErrorCauseBuilder) CausedBy(causedby *ErrorCauseBuilder) *ErrorCauseBuilder {
	v := causedby.Build()
	rb.v.CausedBy = &v
	return rb
}

func (rb *ErrorCauseBuilder) Metadata(value map[string]interface{}) *ErrorCauseBuilder {
	rb.v.Metadata = value
	return rb
}

// Reason A human-readable explanation of the error, in english

func (rb *ErrorCauseBuilder) Reason(reason string) *ErrorCauseBuilder {
	rb.v.Reason = reason
	return rb
}

func (rb *ErrorCauseBuilder) RootCause(root_cause []ErrorCauseBuilder) *ErrorCauseBuilder {
	tmp := make([]ErrorCause, len(root_cause))
	for _, value := range root_cause {
		tmp = append(tmp, value.Build())
	}
	rb.v.RootCause = tmp
	return rb
}

// StackTrace The server stack trace. Present only if the `error_trace=true` parameter was
// sent with the request.

func (rb *ErrorCauseBuilder) StackTrace(stacktrace string) *ErrorCauseBuilder {
	rb.v.StackTrace = &stacktrace
	return rb
}

func (rb *ErrorCauseBuilder) Suppressed(suppressed []ErrorCauseBuilder) *ErrorCauseBuilder {
	tmp := make([]ErrorCause, len(suppressed))
	for _, value := range suppressed {
		tmp = append(tmp, value.Build())
	}
	rb.v.Suppressed = tmp
	return rb
}

// Type The type of error

func (rb *ErrorCauseBuilder) Type_(type_ string) *ErrorCauseBuilder {
	rb.v.Type = type_
	return rb
}
