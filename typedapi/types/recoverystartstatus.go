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
	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// RecoveryStartStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/indices/recovery/types.ts#L91-L96
type RecoveryStartStatus struct {
	CheckIndexTime         Duration `json:"check_index_time,omitempty"`
	CheckIndexTimeInMillis int64    `json:"check_index_time_in_millis"`
	TotalTime              Duration `json:"total_time,omitempty"`
	TotalTimeInMillis      int64    `json:"total_time_in_millis"`
}

func (s *RecoveryStartStatus) UnmarshalJSON(data []byte) error {

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

		case "check_index_time":
			if err := dec.Decode(&s.CheckIndexTime); err != nil {
				return err
			}

		case "check_index_time_in_millis":
			if err := dec.Decode(&s.CheckIndexTimeInMillis); err != nil {
				return err
			}

		case "total_time":
			if err := dec.Decode(&s.TotalTime); err != nil {
				return err
			}

		case "total_time_in_millis":
			if err := dec.Decode(&s.TotalTimeInMillis); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewRecoveryStartStatus returns a RecoveryStartStatus.
func NewRecoveryStartStatus() *RecoveryStartStatus {
	r := &RecoveryStartStatus{}

	return r
}
