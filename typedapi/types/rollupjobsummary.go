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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// RollupJobSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/rollup/get_rollup_index_caps/types.ts#L28-L33
type RollupJobSummary struct {
	Fields       map[string][]RollupJobSummaryField `json:"fields"`
	IndexPattern string                             `json:"index_pattern"`
	JobId        string                             `json:"job_id"`
	RollupIndex  string                             `json:"rollup_index"`
}

func (s *RollupJobSummary) UnmarshalJSON(data []byte) error {

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
				s.Fields = make(map[string][]RollupJobSummaryField, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		case "index_pattern":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexPattern = o

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return err
			}

		case "rollup_index":
			if err := dec.Decode(&s.RollupIndex); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewRollupJobSummary returns a RollupJobSummary.
func NewRollupJobSummary() *RollupJobSummary {
	r := &RollupJobSummary{
		Fields: make(map[string][]RollupJobSummaryField, 0),
	}

	return r
}
