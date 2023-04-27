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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexingjobstate"

	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// RollupJobStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/rollup/get_jobs/types.ts#L60-L64
type RollupJobStatus struct {
	CurrentPosition map[string]json.RawMessage        `json:"current_position,omitempty"`
	JobState        indexingjobstate.IndexingJobState `json:"job_state"`
	UpgradedDocId   *bool                             `json:"upgraded_doc_id,omitempty"`
}

func (s *RollupJobStatus) UnmarshalJSON(data []byte) error {

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

		case "current_position":
			if s.CurrentPosition == nil {
				s.CurrentPosition = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.CurrentPosition); err != nil {
				return err
			}

		case "job_state":
			if err := dec.Decode(&s.JobState); err != nil {
				return err
			}

		case "upgraded_doc_id":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.UpgradedDocId = &value
			case bool:
				s.UpgradedDocId = &v
			}

		}
	}
	return nil
}

// NewRollupJobStatus returns a RollupJobStatus.
func NewRollupJobStatus() *RollupJobStatus {
	r := &RollupJobStatus{
		CurrentPosition: make(map[string]json.RawMessage, 0),
	}

	return r
}
