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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// SnapshotLifecycle type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/slm/_types/SnapshotLifecycle.ts#L38-L49
type SnapshotLifecycle struct {
	InProgress          *InProgress `json:"in_progress,omitempty"`
	LastFailure         *Invocation `json:"last_failure,omitempty"`
	LastSuccess         *Invocation `json:"last_success,omitempty"`
	ModifiedDate        DateTime    `json:"modified_date,omitempty"`
	ModifiedDateMillis  int64       `json:"modified_date_millis"`
	NextExecution       DateTime    `json:"next_execution,omitempty"`
	NextExecutionMillis int64       `json:"next_execution_millis"`
	Policy              SLMPolicy   `json:"policy"`
	Stats               Statistics  `json:"stats"`
	Version             int64       `json:"version"`
}

func (s *SnapshotLifecycle) UnmarshalJSON(data []byte) error {

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

		case "in_progress":
			if err := dec.Decode(&s.InProgress); err != nil {
				return err
			}

		case "last_failure":
			if err := dec.Decode(&s.LastFailure); err != nil {
				return err
			}

		case "last_success":
			if err := dec.Decode(&s.LastSuccess); err != nil {
				return err
			}

		case "modified_date":
			if err := dec.Decode(&s.ModifiedDate); err != nil {
				return err
			}

		case "modified_date_millis":
			if err := dec.Decode(&s.ModifiedDateMillis); err != nil {
				return err
			}

		case "next_execution":
			if err := dec.Decode(&s.NextExecution); err != nil {
				return err
			}

		case "next_execution_millis":
			if err := dec.Decode(&s.NextExecutionMillis); err != nil {
				return err
			}

		case "policy":
			if err := dec.Decode(&s.Policy); err != nil {
				return err
			}

		case "stats":
			if err := dec.Decode(&s.Stats); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSnapshotLifecycle returns a SnapshotLifecycle.
func NewSnapshotLifecycle() *SnapshotLifecycle {
	r := &SnapshotLifecycle{}

	return r
}
