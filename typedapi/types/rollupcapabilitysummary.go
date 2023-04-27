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

// RollupCapabilitySummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/rollup/get_rollup_caps/types.ts#L29-L34
type RollupCapabilitySummary struct {
	Fields       map[string][]RollupFieldSummary `json:"fields"`
	IndexPattern string                          `json:"index_pattern"`
	JobId        string                          `json:"job_id"`
	RollupIndex  string                          `json:"rollup_index"`
}

func (s *RollupCapabilitySummary) UnmarshalJSON(data []byte) error {

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

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string][]RollupFieldSummary, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		case "index_pattern":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.IndexPattern = o

		case "job_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.JobId = o

		case "rollup_index":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.RollupIndex = o

		}
	}
	return nil
}

// NewRollupCapabilitySummary returns a RollupCapabilitySummary.
func NewRollupCapabilitySummary() *RollupCapabilitySummary {
	r := &RollupCapabilitySummary{
		Fields: make(map[string][]RollupFieldSummary, 0),
	}

	return r
}
