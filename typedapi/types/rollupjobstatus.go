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
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexingjobstate"
)

// RollupJobStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/rollup/get_jobs/types.ts#L60-L64
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
