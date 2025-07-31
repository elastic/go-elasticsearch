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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

// An ElasticsearchError represent the exception raised
// by the server and sent as json payloads.
type ElasticsearchError struct {
	ErrorCause ErrorCause `json:"error"`
	Status     int        `json:"status"`
}

// Error implements error string serialization of the ElasticsearchError.
func (e ElasticsearchError) Error() string {
	var reason string
	if e.ErrorCause.Reason != nil {
		reason = *e.ErrorCause.Reason
	}
	return fmt.Sprintf("status: %d, failed: [%s], reason: %s", e.Status, e.ErrorCause.Type, reason)
}

// Is implements errors.Is interface to allow value comparison within ElasticsearchError.
// It checks for always present values only: Status & ErrorCause.Type.
func (e ElasticsearchError) Is(err error) bool {
	prefix := fmt.Sprintf("status: %d, failed: [%s]", e.Status, e.ErrorCause.Type)
	return strings.HasPrefix(err.Error(), prefix)
}

// As implements errors.As interface to allow type matching of ElasticsearchError.
func (e ElasticsearchError) As(err interface{}) bool {
	if _, ok := err.(*ElasticsearchError); ok {
		return true
	}
	return false
}

// NewElasticsearchError returns a ElasticsearchError.
func NewElasticsearchError() *ElasticsearchError {
	r := &ElasticsearchError{}

	return r
}

func (e *ElasticsearchError) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))

	// Expect start object
	t, err := dec.Token()
	if err != nil {
		return err
	}
	if delim, ok := t.(json.Delim); !ok || delim != '{' {
		return fmt.Errorf("expected start object")
	}

	for dec.More() {
		t, err := dec.Token()
		if err != nil {
			return err
		}
		key, ok := t.(string)
		if !ok {
			return fmt.Errorf("expected string key")
		}

		switch key {
		case "error":
			if err := dec.Decode(&e.ErrorCause); err != nil {
				return fmt.Errorf("error decoding error: %w", err)
			}
		case "status":
			if err := dec.Decode(&e.Status); err != nil {
				return fmt.Errorf("error decoding status: %w", err)
			}
		default:
			// Unknown field: store in ErrorCause.Metadata
			if e.ErrorCause.Metadata == nil {
				e.ErrorCause.Metadata = make(map[string]json.RawMessage)
			}
			var raw json.RawMessage
			if err := dec.Decode(&raw); err != nil {
				return fmt.Errorf("error decoding unknown field %s: %w", key, err)
			}
			e.ErrorCause.Metadata[key] = raw
		}
	}

	// Expect end object
	_, err = dec.Token()
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}
