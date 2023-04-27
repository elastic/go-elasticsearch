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

// RunningStateSearchInterval type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/ml/_types/Datafeed.ts#L164-L169
type RunningStateSearchInterval struct {
	End     Duration `json:"end,omitempty"`
	EndMs   int64    `json:"end_ms"`
	Start   Duration `json:"start,omitempty"`
	StartMs int64    `json:"start_ms"`
}

func (s *RunningStateSearchInterval) UnmarshalJSON(data []byte) error {

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

		case "end":
			if err := dec.Decode(&s.End); err != nil {
				return err
			}

		case "end_ms":
			if err := dec.Decode(&s.EndMs); err != nil {
				return err
			}

		case "start":
			if err := dec.Decode(&s.Start); err != nil {
				return err
			}

		case "start_ms":
			if err := dec.Decode(&s.StartMs); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewRunningStateSearchInterval returns a RunningStateSearchInterval.
func NewRunningStateSearchInterval() *RunningStateSearchInterval {
	r := &RunningStateSearchInterval{}

	return r
}
